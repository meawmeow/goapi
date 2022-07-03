package security

import (
	"encoding/json"
	"fiberapiv1/configs"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func LineVerifyToken(token string) map[string]interface{} {
	var uri = "https://api.line.me/oauth2/v2.1/verify"
	data := url.Values{}
	data.Set("id_token", token)
	data.Set("client_id", configs.GetLineClientID())

	encodedData := data.Encode()
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(encodedData))
	//req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBufferString(encodedData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	/// http client end ///
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return res
}
