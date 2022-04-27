package main

import (
	"log"

	"github.com/julianfbeck/universal/backend/services/stock/worker/services"
)

//main function
func main() {
	log.Println("Starting webscraper...")
	var scrapedProduct services.AmazonProductDe
	scrapedProduct.Asin = "B01LZ3KHAL"
	scrapedProduct.Asin = "B089NKGWQQ"
	_, err := scrapedProduct.GetProductInfoByASIN()
	//log error
	if err != nil {
		log.Println(err)
	}
	//log result
	log.Println(scrapedProduct)
}
