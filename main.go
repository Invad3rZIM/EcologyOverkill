/* Ecology Overkill


Step 1. Build webcrawler to GET-scrape & parse wikipedia pages {from crawllist.txt}.
Step 2. Build a parse engine that takes the webscrape-data, filters it, then autoformats

Both functions take a "source" and a "target". The source will be read and the target will be written to.
*/

package main

import "EcologyOverkill/engine"

func main() {
	//engine.ScrapeFromFile("./files/crawllist.txt", "./files/filtereddata.txt")
	engine.RenderTreeFromFile("./files/modifieddata.txt", "./files/output.txt")
}
