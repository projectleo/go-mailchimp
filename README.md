# go-mailchimp

A Golang SDK for Mailchimp API v3.

[![Travis Status for projectboard/go-mailchimp](https://travis-ci.org/RichardKnop/go-mailchimp.svg?branch=master)](https://travis-ci.org/RichardKnop/go-mailchimp)

## Usage

```go
package main

import (
	"log"

	"github.com/projectboard/go-mailchimp/v3"
)

func main() {
	client, err := mailchimp.NewClient("the_api_key-us13", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the email is already subscribed
	memberResponse, err := client.CheckSubscription(
		"listID",
		"john@reese.com",
	)

	// User is already subscribed, update the subscription
	if err == nil {
		memberResponse, err = client.UpdateSubscription(
			"listID",
			"john@reese.com",
			map[string]interface{}{},
		)

		if err != nil {
			// Check the error response
			errResponse, ok := err.(*mailchimp.ErrorResponse)

			// Could not type assert error response
			if !ok {
				log.Fatal(err)
			}

			log.Fatal(errResponse)
		}

		log.Printf(
			"%s's subscription has been updated. Status: %s",
			memberResponse.EmailAddress,
			memberResponse.Status,
		)
		return
	}

	if err != nil {
		// Check the error response
		errResponse, ok := err.(*mailchimp.ErrorResponse)

		// Could not type assert error response
		if !ok {
			log.Fatal(err)
		}

		// 404 means we can process and subscribe user,
		// error other than 404 means we return error
		if errResponse.Status != http.StatusNotFound {
			log.Fatal(errResponse)
		}
	}

	// Subscribe the email
	memberResponse, err = client.Subscribe(
		"listID",
		"john@reese.com",
		map[string]interface{}{},
	)

	if err != nil {
		// Check the error response
		errResponse, ok := err.(*mailchimp.ErrorResponse)

		// Could not type assert error response
		if !ok {
			log.Fatal(err)
		}

		log.Fatal(errResponse)
	}

	log.Printf(
		"%s has been subscribed successfully. Status: %s",
		memberResponse.EmailAddress,
		memberResponse.Status,
	)
}
```
