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

// DeleteAppEnv invokes the webplus.DeleteAppEnv API synchronously
// api document: https://help.aliyun.com/api/webplus/deleteappenv.html
func (client *Client) DeleteAppEnv(request *DeleteAppEnvRequest) (response *DeleteAppEnvResponse, err error) {
	response = CreateDeleteAppEnvResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAppEnvWithChan invokes the webplus.DeleteAppEnv API asynchronously
// api document: https://help.aliyun.com/api/webplus/deleteappenv.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAppEnvWithChan(request *DeleteAppEnvRequest) (<-chan *DeleteAppEnvResponse, <-chan error) {
	responseChan := make(chan *DeleteAppEnvResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAppEnv(request)
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

// DeleteAppEnvWithCallback invokes the webplus.DeleteAppEnv API asynchronously
// api document: https://help.aliyun.com/api/webplus/deleteappenv.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAppEnvWithCallback(request *DeleteAppEnvRequest, callback func(response *DeleteAppEnvResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAppEnvResponse
		var err error
		defer close(result)
		response, err = client.DeleteAppEnv(request)
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

// DeleteAppEnvRequest is the request struct for api DeleteAppEnv
type DeleteAppEnvRequest struct {
	*requests.RoaRequest
	EnvId string `position:"Query" name:"EnvId"`
}

// DeleteAppEnvResponse is the response struct for api DeleteAppEnv
type DeleteAppEnvResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteAppEnvRequest creates a request to invoke DeleteAppEnv API
func CreateDeleteAppEnvRequest() (request *DeleteAppEnvRequest) {
	request = &DeleteAppEnvRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DeleteAppEnv", "/pop/v1/wam/appEnv", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteAppEnvResponse creates a response to parse from DeleteAppEnv response
func CreateDeleteAppEnvResponse() (response *DeleteAppEnvResponse) {
	response = &DeleteAppEnvResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
