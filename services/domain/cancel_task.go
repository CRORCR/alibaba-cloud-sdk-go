package domain

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

// CancelTask invokes the domain.CancelTask API synchronously
// api document: https://help.aliyun.com/api/domain/canceltask.html
func (client *Client) CancelTask(request *CancelTaskRequest) (response *CancelTaskResponse, err error) {
	response = CreateCancelTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CancelTaskWithChan invokes the domain.CancelTask API asynchronously
// api document: https://help.aliyun.com/api/domain/canceltask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelTaskWithChan(request *CancelTaskRequest) (<-chan *CancelTaskResponse, <-chan error) {
	responseChan := make(chan *CancelTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelTask(request)
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

// CancelTaskWithCallback invokes the domain.CancelTask API asynchronously
// api document: https://help.aliyun.com/api/domain/canceltask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelTaskWithCallback(request *CancelTaskRequest, callback func(response *CancelTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelTaskResponse
		var err error
		defer close(result)
		response, err = client.CancelTask(request)
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

// CancelTaskRequest is the request struct for api CancelTask
type CancelTaskRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TaskNo       string `position:"Query" name:"TaskNo"`
	Lang         string `position:"Query" name:"Lang"`
}

// CancelTaskResponse is the response struct for api CancelTask
type CancelTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelTaskRequest creates a request to invoke CancelTask API
func CreateCancelTaskRequest() (request *CancelTaskRequest) {
	request = &CancelTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "CancelTask", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelTaskResponse creates a response to parse from CancelTask response
func CreateCancelTaskResponse() (response *CancelTaskResponse) {
	response = &CancelTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
