package RadwareGoSDK

func (new_client *New_Client) CreateItem(API string, Data map[string]interface{}, additional_header map[string]string) (int, string, error) {
	return(new_client.Request("POST", API, Data, additional_header))
}

func (new_client *New_Client) GetItem(API string, Data map[string]interface{}, additional_header map[string]string) (int, string, error) {
	return(new_client.Request("GET", API, Data, additional_header))
}

func (new_client *New_Client) UpdateItem(API string, Data map[string]interface{}, additional_header map[string]string) (int, string, error) {
	return(new_client.Request("PUT", API, Data, additional_header))
}

func (new_client *New_Client) DeleteItem(API string, Data map[string]interface{}, additional_header map[string]string) (int, string, error) {
	return(new_client.Request("DELETE", API, Data, additional_header))
}

func (new_client *New_Client) CreateItemString(API string, Data string, additional_header map[string]string) (int, string, error) {
	return(new_client.Request("POST", API, Data, additional_header))
}
