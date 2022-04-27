package utils_test

import (
	"fmt"
	"testing"

	"github.com/julianfbeck/universal/backend/services/stock/worker/utils"
)

func TestExtractAsinFromURL(t *testing.T) {
	testString := "	//https://www.amazon.de/Fellow-EKG-Elektrischer-Wasserkocher-Temperaturregelung/dp/B077JBQZPX/?_encoding=UTF8&pd_rd_w=FfWes&pf_rd_p=e0dea9f4-4365-4df5-bbed-53ced537bf12&pf_rd_r=9RV4JMTSQV8CK7MJFCGR&pd_rd_r=aaae9cd8-76d1-4ea5-afbd-f6c58595d1bc&pd_rd_wg=Wmkqo&ref_=pd_gw_bmx_gp_pe3ez0mo"

	//extract asin
	asin, err := utils.ExtractAsinFromURL(testString)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(asin)
}
