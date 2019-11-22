/* Ecology Overkill


Step 1. Build webcrawler to GET-scrape wikipedia pages {from crawllist.txt}.
Step 2. Parse source code to remove all <tags></tags>
*/

package main

import "EcologyOverkill/engine"

func main() {
	engine.ScrapeFromFile("./files/crawllist.txt", "./files/filtereddata.txt")
}
