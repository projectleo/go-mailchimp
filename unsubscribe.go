package mailchimp

import (
	"encoding/json"
	"fmt"
	"github.com/projectleo/go-mailchimp/v3/status"
	"io/ioutil"
)

// Unsubscribe ...
func (c *Client) UnSubscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	// Make request
	members := map[string]interface{}{
		"email_address": email,
		"status":        status.Unsubscribed,
		"merge_fields":  mergeFields,
	}
	resp, err := c.do(
		"POST",
		fmt.Sprintf("/lists/%s", listID),
		&members,
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
		if err := json.Unmarshal(data, memberResponse); err != nil {
			return nil, err
		}
		return memberResponse, nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
