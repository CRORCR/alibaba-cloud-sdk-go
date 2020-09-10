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

// ReleaseService invokes the eas.ReleaseService API synchronously
// api document: https://help.aliyun.com/api/eas/releaseservice.html
func (client *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
	response = CreateReleaseServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseServiceWithChan invokes the eas.ReleaseService API asynchronously
// api document: https://help.aliyun.com/api/eas/releaseservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseServiceWithChan(request *ReleaseServiceRequest) (<-chan *ReleaseServiceResponse, <-chan error) {
	responseChan := make(chan *ReleaseServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseService(request)
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

// ReleaseServiceWithCallback invokes the eas.ReleaseService API asynchronously
// api document: https://help.aliyun.com/api/eas/releaseservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseServiceWithCallback(request *ReleaseServiceRequest, callback func(response *ReleaseServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseServiceResponse
		var err error
		defer close(result)
		response, err = client.ReleaseService(request)
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

// ReleaseServiceRequest is the request struct for api ReleaseService
type ReleaseServiceRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"service_name"`
	Region      string `position:"Path" name:"region"`
}

// ReleaseServiceResponse is the response struct for api ReleaseService
type ReleaseServiceResponse struct {
	*responses.BaseResponse
}

// CreateReleaseServiceRequest creates a request to invoke ReleaseService API
func CreateReleaseServiceRequest() (request *ReleaseServiceRequest) {
	request = &ReleaseServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "ReleaseService", "/api/services/[region]/[service_name]/release", "", "")
	request.Method = requests.PUT
	return
}

// CreateReleaseServiceResponse creates a response to parse from ReleaseService response
func CreateReleaseServiceResponse() (response *ReleaseServiceResponse) {
	response = &ReleaseServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
