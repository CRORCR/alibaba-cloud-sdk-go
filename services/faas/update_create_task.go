package faas

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

// UpdateCreateTask invokes the faas.UpdateCreateTask API synchronously
// api document: https://help.aliyun.com/api/faas/updatecreatetask.html
func (client *Client) UpdateCreateTask(request *UpdateCreateTaskRequest) (response *UpdateCreateTaskResponse, err error) {
	response = CreateUpdateCreateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCreateTaskWithChan invokes the faas.UpdateCreateTask API asynchronously
// api document: https://help.aliyun.com/api/faas/updatecreatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCreateTaskWithChan(request *UpdateCreateTaskRequest) (<-chan *UpdateCreateTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateCreateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCreateTask(request)
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

// UpdateCreateTaskWithCallback invokes the faas.UpdateCreateTask API asynchronously
// api document: https://help.aliyun.com/api/faas/updatecreatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCreateTaskWithCallback(request *UpdateCreateTaskRequest, callback func(response *UpdateCreateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCreateTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateCreateTask(request)
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

// UpdateCreateTaskRequest is the request struct for api UpdateCreateTask
type UpdateCreateTaskRequest struct {
	*requests.RpcRequest
	State               string           `position:"Query" name:"State"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	FpgaImageObjectName string           `position:"Query" name:"FpgaImageObjectName"`
	FpgaImageUUID       string           `position:"Query" name:"FpgaImageUUID"`
	CallerUid           requests.Integer `position:"Query" name:"callerUid"`
}

// UpdateCreateTaskResponse is the response struct for api UpdateCreateTask
type UpdateCreateTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateCreateTaskRequest creates a request to invoke UpdateCreateTask API
func CreateUpdateCreateTaskRequest() (request *UpdateCreateTaskRequest) {
	request = &UpdateCreateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2017-08-24", "UpdateCreateTask", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateCreateTaskResponse creates a response to parse from UpdateCreateTask response
func CreateUpdateCreateTaskResponse() (response *UpdateCreateTaskResponse) {
	response = &UpdateCreateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
