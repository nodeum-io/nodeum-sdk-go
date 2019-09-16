package legacy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ClientV1 struct {
	BaseUrl  string
	Username string
	Password string
}

type StandardResponse struct {
	Code            int           `json:"code"`
	Error           error         `json:"error"`
	ErrorParameters []interface{} `json:"error_parameters"`
	Status          string        `json:"status"`
	Message         string        `json:"message"`
}

func NewBasicAuthClientV1(nodeum, username, password string) *ClientV1 {
	baseUrl := fmt.Sprintf("http://%s/api/", nodeum)
	return &ClientV1{
		BaseUrl:  baseUrl,
		Username: username,
		Password: password,
	}
}

func (api *ClientV1) Req(req *http.Request, data interface{}) (err error) {
	req.SetBasicAuth(api.Username, api.Password)
	if req.URL.Host == "" {
		req.URL, err = url.Parse(api.BaseUrl + req.URL.String())

		if err != nil {
			return
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if 200 != resp.StatusCode {
		err = fmt.Errorf("%s", body)
	}

	var std StandardResponse
	json.Unmarshal(body, &std)
	err = json.Unmarshal(body, &data)

	if std.Code >= 400 && std.Code < 600 || std.Status == "ERROR" {
		err = fmt.Errorf("%s", std.Message)
	}

	return
}
