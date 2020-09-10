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

// StartRestoreTask invokes the dbs.StartRestoreTask API synchronously
func (client *Client) StartRestoreTask(request *StartRestoreTaskRequest) (response *StartRestoreTaskResponse, err error) {
	response = CreateStartRestoreTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StartRestoreTaskWithChan invokes the dbs.StartRestoreTask API asynchronously
func (client *Client) StartRestoreTaskWithChan(request *StartRestoreTaskRequest) (<-chan *StartRestoreTaskResponse, <-chan error) {
	responseChan := make(chan *StartRestoreTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartRestoreTask(request)
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

// StartRestoreTaskWithCallback invokes the dbs.StartRestoreTask API asynchronously
func (client *Client) StartRestoreTaskWithCallback(request *StartRestoreTaskRequest, callback func(response *StartRestoreTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartRestoreTaskResponse
		var err error
		defer close(result)
		response, err = client.StartRestoreTask(request)
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

// StartRestoreTaskRequest is the request struct for api StartRestoreTask
type StartRestoreTaskRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	RestoreTaskId string `position:"Query" name:"RestoreTaskId"`
}

// StartRestoreTaskResponse is the response struct for api StartRestoreTask
type StartRestoreTaskResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RestoreTaskId  string `json:"RestoreTaskId" xml:"RestoreTaskId"`
}

// CreateStartRestoreTaskRequest creates a request to invoke StartRestoreTask API
func CreateStartRestoreTaskRequest() (request *StartRestoreTaskRequest) {
	request = &StartRestoreTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "StartRestoreTask", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartRestoreTaskResponse creates a response to parse from StartRestoreTask response
func CreateStartRestoreTaskResponse() (response *StartRestoreTaskResponse) {
	response = &StartRestoreTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
