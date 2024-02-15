package RadwareGoSDK

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var temp_request *http.Request
var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func Login(product string, host string, username string, password string) {
	if "ALTEON" == strings.ToUpper(product) {
		url := "https://"+host
		authHeader := basicAuthHeader(username, password)
		login_request, err := http.NewRequest("POST", url, nil)
		if err != nil {
			fmt.Println("Error creating login request:", err)
			return 
		}
		login_request.Header.Set("Authorization", authHeader)
		login_request.Header.Set("Content-Type", "application/json")
		temp_request = login_request
		return
	} else if "CYBERCONTROLLER" == strings.ToUpper(product) {
		url := "https://" + host + "/mgmt/system/user/login"
		Data := map[string]string{"username": "radware", "password": "radware"}
		Bytes, err := json.Marshal(Data)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return 
		}
		login_request, err := http.NewRequest("POST", url, bytes.NewBuffer(Bytes))
		if err != nil {
			fmt.Println("Error creating login request:", err)
			return 
		}
		login_request.Header.Set("Content-Type", "application/json")
		login_response, err := client.Do(login_request)

		if err != nil {
			fmt.Println("Error making login request:", err)
			return 
		}
		defer login_response.Body.Close()
		JsessionID, err := extractJSessionID(login_response)
		if err != nil {
			fmt.Println("Error extracting JSESSIONID:", err)
			return 
		}
		temp_request = login_request
		temp_request.Header.Set("Content-Type", "application/json")
		temp_request.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s", JsessionID))
		fmt.Println("Login Successful")
		return 
	}

}

func Request(host string, method string, API string, Data map[string]interface{}) (int, string, error) {
	URL := "https://" + host + API
	APIBytes, err := json.Marshal(Data)
	if err != nil {
		return 0, "Error encoding JSON", err
	}
	New_Request, err := http.NewRequest(strings.ToUpper(method), URL, bytes.NewBuffer(APIBytes))
	if err != nil {
		return 0, "Error creating API request", err
	}

	New_Request.Header = temp_request.Header.Clone()
	APIResponse, err := client.Do(New_Request)
	if err != nil {
		return 0, "Error making API request", err
	}
	defer APIResponse.Body.Close()

	// Read the response body of the API call
	APIResponseBody, err := ioutil.ReadAll(APIResponse.Body)
	if err != nil {
		return APIResponse.StatusCode, "Error reading API response body", err
	}
	
	return APIResponse.StatusCode, string(APIResponseBody), nil
}

func extractJSessionID(response *http.Response) (string, error) {
	for _, cookie := range response.Cookies() {
		if cookie.Name == "JSESSIONID" {
			return cookie.Value, nil
		}
	}
	return "", fmt.Errorf("JSESSIONID not found in cookies")
}

func basicAuthHeader(username, password string) string {
	auth := username + ":" + password
	b64 := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + b64
}
