package wns

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
)

const unknownErrorText = "Unknown error."

var responseCodes = map[int]string{
	400: "One or more headers were specified incorrectly or conflict with another header.",
	401: "The cloud service did not present a valid authentication ticket. The OAuth ticket may be invalid.",
	403: "The cloud service is not authorized to send a notification to this URI even though they are authenticated.",
	404: "The channel URI is not valid or is not recognized by WNS.",
	405: "Invalid method (GET, DELETE, CREATE); only POST is allowed.",
	406: "The cloud service exceeded its throttle limit.",
	410: "The channel expired.",
	413: "The notification payload exceeds the 5000 byte size limit.",
	500: "An internal failure caused notification delivery to fail.",
	503: "The server is currently unavailable.",
}

type Client struct {
	clientId     string
	clientSecret string
	accessToken  string
}

type ClientError struct {
	Code    int
	Message string
}

func (e *ClientError) Error() string {
	return e.Message
}

func (c *Client) Init(clientId string, clientSecret string) {
	c.clientId = clientId
	c.clientSecret = clientSecret
	c.accessToken = c.getAccessToken(clientId, clientSecret)
}

func (c *Client) getAccessToken(clientId string, clientSecret string) string {
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", clientId)
	data.Add("client_secret", clientSecret)
	data.Add("scope", "notify.windows.com")

	resp, _ := http.PostForm("https://login.live.com/accesstoken.srf", data)

	if resp.StatusCode == 200 {
		var body struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
		}
		json.NewDecoder(resp.Body).Decode(&body)
		return body.AccessToken
	}

	return ""
}

func (c *Client) refreshAccessToken() {
	c.accessToken = c.getAccessToken(c.clientId, c.clientSecret)
}

func (c *Client) Send(channelUri string, notification NotificationInterface) (success bool, error *ClientError) {
	cl := http.Client{}

	xml, _ := notification.GetXml()
	req, _ := http.NewRequest("POST", channelUri, bytes.NewBuffer([]byte(xml)))
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("X-WNS-Type", notification.GetWnsType())
	req.Header.Add("Authorization", "Bearer "+c.accessToken)

	resp, err := cl.Do(req)
	if err != nil {
		return false, &ClientError{resp.StatusCode, unknownErrorText}
	}

	switch resp.StatusCode {
	case 200:
		return true, nil
	case 401:
		c.refreshAccessToken()
		return c.Send(channelUri, notification)
	default:
		message, ok := responseCodes[resp.StatusCode]
		if !ok {
			message = unknownErrorText
		}
		return false, &ClientError{resp.StatusCode, message}
	}

	return false, nil
}
