package engine

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/* ScrapeFromFile performs a series of Get requests to capture plant classification data
from wikipedia links.

*/
func ScrapeFromFile(source string, target string) {
	//Read Source
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//create a new target to load data into
	f, err := os.Create(target)
	defer f.Close()

	index := 0

	for scanner.Scan() {
		//Scrapes the URL's and returns a sanitized string with the relevant data...
		sanitizedBody, err := scrapeURL(scanner.Text())

		if err != nil {
			fmt.Println(err.Error())
		}

		filtered := filterData(sanitizedBody)
		saveScraperData(f, filtered, index) //saves the data to target file

		index = index + 1
	}
}

/* filterData takes an unfiltered array and returned a collapsed version of grouped data.

data in the unfiltered resembles [ a,:,b,c,:,d,e,:,f]
filtered version will parse into [a b,c d,e f]
this is easier to validate...
*/
func filterData(body []string) []string {
	index := 0

	filtered := []string{}
	line := ""

	for index < len(body) {
		line = body[index]

		if body[index] == ":" {
			line = body[index-1] + " " + body[index+1]

			if line[len(line)-1] == []byte(".")[0] {
				line = line + " " + body[index+2]
			}

			filtered = append(filtered, line)
		}

		index = index + 1
	}

	return filtered
}

//scrapeURL takes in a url and returns a sanitized []string list containing all the data scraped
func scrapeURL(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		//fmt.Printf("%s", err)
		return nil, err
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			//	fmt.Printf("%s", err)
			os.Exit(1)
		}

		san := sanitizeBody(string(body))
		strs := regexp.MustCompile("[^\\s]+").FindAllString(san, -1) //regex out the whitespaces
		return strs, nil
	}
}

/* SanitizateBody does 4 things really...

1. Isolate the relevant data (beginning with kingdom, ending with "/b /i /div")
Conveniently, every wikipedia page formats this the same...

2.	Remove all tags
3.Remove all special characters ; : . , etc

4. Remove Misc Leftovers

*/
func sanitizeBody(generalBody string) string {
	body := func(body string) string {
		//Strip all newlines...
		body = strings.Replace(generalBody, "\n", "", -1)

		//strip everything not relevant to the classification...

		body = body[strings.Index(body, "<td>Kingdom:"):]
		endIndex := strings.Index(body, "</b></i></div>")

		if endIndex == -1 { //account for the 1 edgecase
			endIndex = strings.Index(body, "</a></small></td></tr>")
		}

		if endIndex != -1 {
			body = body[:endIndex+4]
		}

		return body
	}(generalBody)

	//remove all tags and misc garbage from each body...
	body = stripTags(body)                         //removes all <tags></tags>
	body = stripMisc(body)                         //removes stuff like  P.&#160; C.&#160; P.&#160;
	body = strings.Replace(body, "Adans.", "", -1) //random junk that's sometimes at the end...
	body = strings.Replace(body, ":", " : ", -1)   //" : " will be used as a parsing regex...
	body = strings.Replace(body, "(unranked)", " ", -1)
	return body
}

/* stripTags  removes all html tags and replaces them with " " */
func stripTags(body string) string {
	for strings.Index(body, "<") != -1 {
		leftCut := strings.Index(body, "<")
		rightCut := strings.Index(body, ">")

		if leftCut != -1 && rightCut != -1 && leftCut < rightCut {
			body = " " + body[0:leftCut] + " " + body[rightCut+1:]
		}
	}

	return body
}

//for some reason, every string has some weird expression that looks like...
//     P.&#160; C.&#160; P.&#160; => "{someletter}.{...};" and so I parse this out...
func stripMisc(body string) string {
	if strings.Index(body, ".") != -1 {
		leftCut := strings.Index(body, ".") + 1
		rightCut := strings.Index(body, ";")

		if leftCut != -1 && rightCut != -1 && leftCut < rightCut {
			body = " " + body[0:leftCut] + " " + body[rightCut+1:]
		}
	}

	if strings.Index(body, "&#") != -1 {
		leftCut := strings.Index(body, "&#")
		rightCut := strings.LastIndex(body, ";")

		if leftCut != -1 && rightCut != -1 && leftCut < rightCut {
			body = " " + body[0:leftCut] + " " + body[rightCut+1:]
		}
	}

	return body
}
