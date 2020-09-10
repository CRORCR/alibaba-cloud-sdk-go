package iot

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

// GetNodesAddingTask invokes the iot.GetNodesAddingTask API synchronously
// api document: https://help.aliyun.com/api/iot/getnodesaddingtask.html
func (client *Client) GetNodesAddingTask(request *GetNodesAddingTaskRequest) (response *GetNodesAddingTaskResponse, err error) {
	response = CreateGetNodesAddingTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetNodesAddingTaskWithChan invokes the iot.GetNodesAddingTask API asynchronously
// api document: https://help.aliyun.com/api/iot/getnodesaddingtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNodesAddingTaskWithChan(request *GetNodesAddingTaskRequest) (<-chan *GetNodesAddingTaskResponse, <-chan error) {
	responseChan := make(chan *GetNodesAddingTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNodesAddingTask(request)
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

// GetNodesAddingTaskWithCallback invokes the iot.GetNodesAddingTask API asynchronously
// api document: https://help.aliyun.com/api/iot/getnodesaddingtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNodesAddingTaskWithCallback(request *GetNodesAddingTaskRequest, callback func(response *GetNodesAddingTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNodesAddingTaskResponse
		var err error
		defer close(result)
		response, err = client.GetNodesAddingTask(request)
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

// GetNodesAddingTaskRequest is the request struct for api GetNodesAddingTask
type GetNodesAddingTaskRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	TaskId        string `position:"Query" name:"TaskId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GetNodesAddingTaskResponse is the response struct for api GetNodesAddingTask
type GetNodesAddingTaskResponse struct {
	*responses.BaseResponse
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	Success        bool                               `json:"Success" xml:"Success"`
	Code           string                             `json:"Code" xml:"Code"`
	ErrorMessage   string                             `json:"ErrorMessage" xml:"ErrorMessage"`
	TaskId         string                             `json:"TaskId" xml:"TaskId"`
	TaskState      string                             `json:"TaskState" xml:"TaskState"`
	TotalCount     int64                              `json:"TotalCount" xml:"TotalCount"`
	SuccessCount   int64                              `json:"SuccessCount" xml:"SuccessCount"`
	SuccessDevEuis SuccessDevEuisInGetNodesAddingTask `json:"SuccessDevEuis" xml:"SuccessDevEuis"`
}

// CreateGetNodesAddingTaskRequest creates a request to invoke GetNodesAddingTask API
func CreateGetNodesAddingTaskRequest() (request *GetNodesAddingTaskRequest) {
	request = &GetNodesAddingTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetNodesAddingTask", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNodesAddingTaskResponse creates a response to parse from GetNodesAddingTask response
func CreateGetNodesAddingTaskResponse() (response *GetNodesAddingTaskResponse) {
	response = &GetNodesAddingTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
