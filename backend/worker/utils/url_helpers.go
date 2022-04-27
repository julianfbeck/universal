package utils

import (
	"fmt"
	"regexp"
)

//extract the asin from amazon url
func ExtractAsinFromURL(url string) (string, error) {
	reg := regexp.MustCompile(`dp/\w{10}`)
	result := reg.Find([]byte(url))
	//check if result is empty
	if len(result) == 0 {
		return "", fmt.Errorf("no asin found in url")
	}
	return string(result[3:13]), nil

}
