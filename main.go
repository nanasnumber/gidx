package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIndexURL(s string) {
	res, err := http.Get(s)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	fmt.Println(string(resBody))
}

func main() {
	getIndexURL("https://www.google.com/search?q=site%3Asoumissionexcavation.ca&oq=site%3Asoumissionexcavation.ca&aqs=chrome..69i57j69i58.11371j0j4&sourceid=chrome&ie=UTF-8")
}
