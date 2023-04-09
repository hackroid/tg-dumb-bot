package static

import (
	"github.com/gocolly/colly/v2"
	"log"
	"time"
)

func test_pppixiv() {
	log.Printf("test_pppixiv()......")
	// Create a new Colly collector instance
	c := colly.NewCollector()

	c.SetRequestTimeout(60 * time.Second)

	// Set the URL to the pppixiv service endpoint that you want to send a POST request to
	postURL := "http://pppixiv:5000/getIllustListByUid"

	// Create a new FormData object and set the POST parameters
	var data map[string]string
	data = make(map[string]string)
	data["uid"] = "4837211"

	// Define the callback function to be executed after the request is completed
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Response: -> %s\n", string(r.Body))
	})

	// Send the POST request using Colly's PostForm method
	err := c.Post(postURL, data)
	if err != nil {
		log.Printf("err: -> %s\n", err)
	}
}
