package services_test

import (
	"testing"

	"github.com/julianfbeck/universal/backend/services/stock/worker/services"
)

//array of amazon amazon string array
var amazonStrings = []string{
	"B01LZ3KHAL",
	"B089NKGWQQ",
	"B077JBQZPX",
}

func TestHelloName(t *testing.T) {
	//test for each amazon string
	for _, asin := range amazonStrings {
		//instantiate amazon product
		var product = services.AmazonProductDe{
			Asin: asin,
		}
		//get product info
		_, err := product.GetProductInfoByASIN()
		//log error
		if err != nil {
			t.Error(err)
		}
		//log result
		t.Log(product)
	}
}

//test asin
