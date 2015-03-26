package client

import (
	integration "github.com/opsgenie/opsgenie-go-sdk/integration"
	goreq "github.com/franela/goreq"
	"errors"
	"fmt"
)

const (
	ENABLE_INTEGRATION_URL 		= ENDPOINT_URL + "/v1/json/integration/enable"
	DISABLE_INTEGRATION_URL 	= ENDPOINT_URL + "/v1/json/integration/disable"
)

type OpsGenieIntegrationClient struct {
	apiKey string
}

func (cli *OpsGenieIntegrationClient) Enable(req integration.EnableIntegrationRequest) (*integration.EnableIntegrationResponse, error){
	req.ApiKey = cli.apiKey
	// validate mandatory fields: id/name, apiKey
	if req.ApiKey == "" && req.Id == ""{
		return nil, errors.New("Api Key or Id should be provided")	
	}
	if req.ApiKey != "" && req.Id != "" {
		return nil, errors.New("Either Api Key or Id should be provided, not both")	
	}
	// send the request
	resp, err := goreq.Request{ Method: "POST", Uri: ENABLE_INTEGRATION_URL, Body: req, }.Do()	
	if err != nil {
		return nil, errors.New("Can not enable the integration, unable to send the request")
	}
	// check for the returning http status, 4xx: client errors, 5xx: server errors
	statusCode := resp.StatusCode
	if statusCode >= 400 && statusCode < 500 {
		return nil, errors.New( fmt.Sprintf("Client error %d occured", statusCode) )
	}
	if statusCode >= 500  {
		return nil, errors.New( fmt.Sprintf("Server error %d occured", statusCode) )
	}
	// try to parse the returning JSON into the response
	var enableIntegrationResp integration.EnableIntegrationResponse
	if err = resp.Body.FromJsonTo(&enableIntegrationResp); err != nil {
		return nil, errors.New("Server response can not be parsed")
	}
	// parsed successfuly with no errors
	return &enableIntegrationResp, nil	
}

func (cli *OpsGenieIntegrationClient) Disable(req integration.DisableIntegrationRequest) (*integration.DisableIntegrationResponse, error){
	req.ApiKey = cli.apiKey
	// validate mandatory fields: id/name, apiKey
	if req.ApiKey == "" && req.Id == ""{
		return nil, errors.New("Api Key or Id should be provided")	
	}
	if req.ApiKey != "" && req.Id != "" {
		return nil, errors.New("Either Api Key or Id should be provided, not both")	
	}
	// send the request
	resp, err := goreq.Request{ Method: "POST", Uri: DISABLE_INTEGRATION_URL, Body: req, }.Do()	
	if err != nil {
		return nil, errors.New("Can not disable the integration, unable to send the request")
	}
	// check for the returning http status, 4xx: client errors, 5xx: server errors
	statusCode := resp.StatusCode
	if statusCode >= 400 && statusCode < 500 {
		return nil, errors.New( fmt.Sprintf("Client error %d occured", statusCode) )
	}
	if statusCode >= 500  {
		return nil, errors.New( fmt.Sprintf("Server error %d occured", statusCode) )
	}
	// try to parse the returning JSON into the response
	var disableIntegrationResp integration.DisableIntegrationResponse
	if err = resp.Body.FromJsonTo(&disableIntegrationResp); err != nil {
		return nil, errors.New("Server response can not be parsed")
	}
	// parsed successfuly with no errors
	return &disableIntegrationResp, nil	
}
