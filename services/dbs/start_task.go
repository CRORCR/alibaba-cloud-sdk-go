package dbs

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

// StartTask invokes the dbs.StartTask API synchronously
func (client *Client) StartTask(request *StartTaskRequest) (response *StartTaskResponse, err error) {
	response = CreateStartTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StartTaskWithChan invokes the dbs.StartTask API asynchronously
func (client *Client) StartTaskWithChan(request *StartTaskRequest) (<-chan *StartTaskResponse, <-chan error) {
	responseChan := make(chan *StartTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartTask(request)
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

// StartTaskWithCallback invokes the dbs.StartTask API asynchronously
func (client *Client) StartTaskWithCallback(request *StartTaskRequest, callback func(response *StartTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartTaskResponse
		var err error
		defer close(result)
		response, err = client.StartTask(request)
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

// StartTaskRequest is the request struct for api StartTask
type StartTaskRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     string `position:"Query" name:"OwnerId"`
	TaskId      string `position:"Query" name:"TaskId"`
}

// StartTaskResponse is the response struct for api StartTask
type StartTaskResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	JobTypeName    string `json:"JobTypeName" xml:"JobTypeName"`
}

// CreateStartTaskRequest creates a request to invoke StartTask API
func CreateStartTaskRequest() (request *StartTaskRequest) {
	request = &StartTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "StartTask", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartTaskResponse creates a response to parse from StartTask response
func CreateStartTaskResponse() (response *StartTaskResponse) {
	response = &StartTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
