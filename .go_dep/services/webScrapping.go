package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/anaskhan96/soup"
)

func GetAmazon(query string) map[string]string {
	var item = make(map[string]string)
	query = strings.Replace(query, " ", "+", -1)
	// add headers to the request
	soup.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	res, err := soup.Get("https://www.amazon.in/s?k=" + query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	doc := soup.HTMLParse(res)
	soup.SetDebug(true)
	//fmt.Println(doc.FullText())
	block := doc.Find("div", "data-component-type", "s-search-result")
	item["title"] = block.Find("span", "class", "a-size-medium").Text()
	item["price"] = block.Find("span", "class", "a-price-whole").Text()

	// if strings.Contains(item["title"], query) {
	// 	return item
	// } else {
	// 	b := map[string]string{
	// 		"message": "Item not found in amazon",
	// 	}
	// 	return b
	// }
	return item

}
func GetFlipkart(query string) map[string]string {
	var item = make(map[string]string)
	query = strings.Replace(query, " ", "+", -1)
	soup.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	res, err := soup.Get("https://www.flipkart.com/search?q=" + query)
	println("")
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(res)
	soup.SetDebug(true)
	block := doc.Find("a", "class", "_1fQZEK")
	item["title"] = block.Find("div", "class", "_4rR01T").Text()
	item["price"] = block.Find("div", "class", "_30jeq3").Text()
	return item
	// if strings.Contains(item["title"], query) {

	// } else {
	// 	b := map[string]string{
	// 		"query":   query,
	// 		"message": "Item not found in flipkart",
	// 	}
	// 	return b
	// }
}

// func main() {
// 	// fmt.Print("Hello World")
// 	fmt.Print(getAmazon("iphone 11"))
// 	fmt.Print("\n")
// 	fmt.Print(getFlipkart("iphone 11"))
// }
