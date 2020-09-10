package webplus

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// DryRunTerminateAppEnv invokes the webplus.DryRunTerminateAppEnv API synchronously
// api document: https://help.aliyun.com/api/webplus/dryrunterminateappenv.html
func (client *Client) DryRunTerminateAppEnv(request *DryRunTerminateAppEnvRequest) (response *DryRunTerminateAppEnvResponse, err error) {
	response = CreateDryRunTerminateAppEnvResponse()
	err = client.DoAction(request, response)
	return
}

// DryRunTerminateAppEnvWithChan invokes the webplus.DryRunTerminateAppEnv API asynchronously
// api document: https://help.aliyun.com/api/webplus/dryrunterminateappenv.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DryRunTerminateAppEnvWithChan(request *DryRunTerminateAppEnvRequest) (<-chan *DryRunTerminateAppEnvResponse, <-chan error) {
	responseChan := make(chan *DryRunTerminateAppEnvResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DryRunTerminateAppEnv(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DryRunTerminateAppEnvWithCallback invokes the webplus.DryRunTerminateAppEnv API asynchronously
// api document: https://help.aliyun.com/api/webplus/dryrunterminateappenv.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DryRunTerminateAppEnvWithCallback(request *DryRunTerminateAppEnvRequest, callback func(response *DryRunTerminateAppEnvResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DryRunTerminateAppEnvResponse
		var err error
		defer close(result)
		response, err = client.DryRunTerminateAppEnv(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DryRunTerminateAppEnvRequest is the request struct for api DryRunTerminateAppEnv
type DryRunTerminateAppEnvRequest struct {
	*requests.RoaRequest
	EnvId string `position:"Body" name:"EnvId"`
}

// DryRunTerminateAppEnvResponse is the response struct for api DryRunTerminateAppEnv
type DryRunTerminateAppEnvResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Code       string     `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	DryRunInfo DryRunInfo `json:"DryRunInfo" xml:"DryRunInfo"`
}

// CreateDryRunTerminateAppEnvRequest creates a request to invoke DryRunTerminateAppEnv API
func CreateDryRunTerminateAppEnvRequest() (request *DryRunTerminateAppEnvRequest) {
	request = &DryRunTerminateAppEnvRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DryRunTerminateAppEnv", "/pop/v1/wam/appEnv/dryRunTerminate", "", "")
	request.Method = requests.POST
	return
}

// CreateDryRunTerminateAppEnvResponse creates a response to parse from DryRunTerminateAppEnv response
func CreateDryRunTerminateAppEnvResponse() (response *DryRunTerminateAppEnvResponse) {
	response = &DryRunTerminateAppEnvResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
