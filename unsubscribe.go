package mailchimp

import (
	"encoding/json"
	"fmt"
	"github.com/projectleo/go-mailchimp/v3/status"
	"io/ioutil"
)

type Member struct {
	EmailAddress string `json:"email_address"`
	Status string `json:"status"`
	MergeFields map[string]interface{} `json:"merge_fields"`
}
// Unsubscribe ...
func (c *Client) UnSubscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	// Make request
	body := struct{
		UpdateExisting bool `json:"update_existing"`
		Members []Member `json:"members"`
	}{
		UpdateExisting: true,
		Members: []Member{
			{
				EmailAddress: email,
				Status: status.Unsubscribed,
				MergeFields: mergeFields,
			},
		},
	}

	resp, err := c.do(
		"POST",
		fmt.Sprintf("/lists/%s", listID),
		&body,
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
