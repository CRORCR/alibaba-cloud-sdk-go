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

// GetServiceWorker invokes the eas.GetServiceWorker API synchronously
// api document: https://help.aliyun.com/api/eas/getserviceworker.html
func (client *Client) GetServiceWorker(request *GetServiceWorkerRequest) (response *GetServiceWorkerResponse, err error) {
	response = CreateGetServiceWorkerResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceWorkerWithChan invokes the eas.GetServiceWorker API asynchronously
// api document: https://help.aliyun.com/api/eas/getserviceworker.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceWorkerWithChan(request *GetServiceWorkerRequest) (<-chan *GetServiceWorkerResponse, <-chan error) {
	responseChan := make(chan *GetServiceWorkerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceWorker(request)
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

// GetServiceWorkerWithCallback invokes the eas.GetServiceWorker API asynchronously
// api document: https://help.aliyun.com/api/eas/getserviceworker.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceWorkerWithCallback(request *GetServiceWorkerRequest, callback func(response *GetServiceWorkerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceWorkerResponse
		var err error
		defer close(result)
		response, err = client.GetServiceWorker(request)
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

// GetServiceWorkerRequest is the request struct for api GetServiceWorker
type GetServiceWorkerRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"service_name"`
	Region      string `position:"Path" name:"region"`
}

// GetServiceWorkerResponse is the response struct for api GetServiceWorker
type GetServiceWorkerResponse struct {
	*responses.BaseResponse
}

// CreateGetServiceWorkerRequest creates a request to invoke GetServiceWorker API
func CreateGetServiceWorkerRequest() (request *GetServiceWorkerRequest) {
	request = &GetServiceWorkerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "GetServiceWorker", "/api/services/[region]/[service_name]/worker", "", "")
	request.Method = requests.GET
	return
}

// CreateGetServiceWorkerResponse creates a response to parse from GetServiceWorker response
func CreateGetServiceWorkerResponse() (response *GetServiceWorkerResponse) {
	response = &GetServiceWorkerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
