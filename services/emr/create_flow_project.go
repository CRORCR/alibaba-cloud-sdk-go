package emr

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

// CreateFlowProject invokes the emr.CreateFlowProject API synchronously
// api document: https://help.aliyun.com/api/emr/createflowproject.html
func (client *Client) CreateFlowProject(request *CreateFlowProjectRequest) (response *CreateFlowProjectResponse, err error) {
	response = CreateCreateFlowProjectResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFlowProjectWithChan invokes the emr.CreateFlowProject API asynchronously
// api document: https://help.aliyun.com/api/emr/createflowproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFlowProjectWithChan(request *CreateFlowProjectRequest) (<-chan *CreateFlowProjectResponse, <-chan error) {
	responseChan := make(chan *CreateFlowProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFlowProject(request)
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

// CreateFlowProjectWithCallback invokes the emr.CreateFlowProject API asynchronously
// api document: https://help.aliyun.com/api/emr/createflowproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFlowProjectWithCallback(request *CreateFlowProjectRequest, callback func(response *CreateFlowProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFlowProjectResponse
		var err error
		defer close(result)
		response, err = client.CreateFlowProject(request)
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

// CreateFlowProjectRequest is the request struct for api CreateFlowProject
type CreateFlowProjectRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	Name        string `position:"Query" name:"Name"`
}

// CreateFlowProjectResponse is the response struct for api CreateFlowProject
type CreateFlowProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateCreateFlowProjectRequest creates a request to invoke CreateFlowProject API
func CreateCreateFlowProjectRequest() (request *CreateFlowProjectRequest) {
	request = &CreateFlowProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowProject", "emr", "openAPI")
	return
}

// CreateCreateFlowProjectResponse creates a response to parse from CreateFlowProject response
func CreateCreateFlowProjectResponse() (response *CreateFlowProjectResponse) {
	response = &CreateFlowProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
