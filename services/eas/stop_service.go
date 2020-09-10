package eas

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

// StopService invokes the eas.StopService API synchronously
// api document: https://help.aliyun.com/api/eas/stopservice.html
func (client *Client) StopService(request *StopServiceRequest) (response *StopServiceResponse, err error) {
	response = CreateStopServiceResponse()
	err = client.DoAction(request, response)
	return
}

// StopServiceWithChan invokes the eas.StopService API asynchronously
// api document: https://help.aliyun.com/api/eas/stopservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopServiceWithChan(request *StopServiceRequest) (<-chan *StopServiceResponse, <-chan error) {
	responseChan := make(chan *StopServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopService(request)
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

// StopServiceWithCallback invokes the eas.StopService API asynchronously
// api document: https://help.aliyun.com/api/eas/stopservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopServiceWithCallback(request *StopServiceRequest, callback func(response *StopServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopServiceResponse
		var err error
		defer close(result)
		response, err = client.StopService(request)
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

// StopServiceRequest is the request struct for api StopService
type StopServiceRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"service_name"`
	Region      string `position:"Path" name:"region"`
}

// StopServiceResponse is the response struct for api StopService
type StopServiceResponse struct {
	*responses.BaseResponse
}

// CreateStopServiceRequest creates a request to invoke StopService API
func CreateStopServiceRequest() (request *StopServiceRequest) {
	request = &StopServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "StopService", "/api/services/[region]/[service_name]/stop", "", "")
	request.Method = requests.PUT
	return
}

// CreateStopServiceResponse creates a response to parse from StopService response
func CreateStopServiceResponse() (response *StopServiceResponse) {
	response = &StopServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
