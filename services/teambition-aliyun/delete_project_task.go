package teambition_aliyun

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

// DeleteProjectTask invokes the teambition_aliyun.DeleteProjectTask API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/deleteprojecttask.html
func (client *Client) DeleteProjectTask(request *DeleteProjectTaskRequest) (response *DeleteProjectTaskResponse, err error) {
	response = CreateDeleteProjectTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteProjectTaskWithChan invokes the teambition_aliyun.DeleteProjectTask API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/deleteprojecttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProjectTaskWithChan(request *DeleteProjectTaskRequest) (<-chan *DeleteProjectTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteProjectTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteProjectTask(request)
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

// DeleteProjectTaskWithCallback invokes the teambition_aliyun.DeleteProjectTask API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/deleteprojecttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProjectTaskWithCallback(request *DeleteProjectTaskRequest, callback func(response *DeleteProjectTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteProjectTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteProjectTask(request)
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

// DeleteProjectTaskRequest is the request struct for api DeleteProjectTask
type DeleteProjectTaskRequest struct {
	*requests.RpcRequest
	OrgId  string `position:"Body" name:"OrgId"`
	TaskId string `position:"Body" name:"TaskId"`
}

// DeleteProjectTaskResponse is the response struct for api DeleteProjectTask
type DeleteProjectTaskResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     bool   `json:"Object" xml:"Object"`
}

// CreateDeleteProjectTaskRequest creates a request to invoke DeleteProjectTask API
func CreateDeleteProjectTaskRequest() (request *DeleteProjectTaskRequest) {
	request = &DeleteProjectTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "DeleteProjectTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteProjectTaskResponse creates a response to parse from DeleteProjectTask response
func CreateDeleteProjectTaskResponse() (response *DeleteProjectTaskResponse) {
	response = &DeleteProjectTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
