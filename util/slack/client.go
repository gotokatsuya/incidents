package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel,omitempty"`
}

func Post(r Request, incommingWebhookURL string) error {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return err
	}

	v := url.Values{}
	v.Add("payload", string(jsonBody))

	req, err := http.NewRequest("POST", incommingWebhookURL, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
