package mailchimp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// Unsubscribe ...
func (c *Client) UnSubscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	// Make request
	hash := md5.New()
	hash.Write([]byte(email))
	subscriberHash := hex.EncodeToString(hash.Sum(nil))

	resp, err := c.do(
		"POST",
		fmt.Sprintf("/lists/%s/members/%s/actions/delete-permanent", listID, string(subscriberHash)),
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		// Unmarshal response into MemberResponse struct
		memberResponse := new(MemberResponse)
		return memberResponse, nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
