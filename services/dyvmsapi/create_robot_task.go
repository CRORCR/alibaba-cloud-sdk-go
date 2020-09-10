package dyvmsapi

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

// CreateRobotTask invokes the dyvmsapi.CreateRobotTask API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/createrobottask.html
func (client *Client) CreateRobotTask(request *CreateRobotTaskRequest) (response *CreateRobotTaskResponse, err error) {
	response = CreateCreateRobotTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRobotTaskWithChan invokes the dyvmsapi.CreateRobotTask API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/createrobottask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRobotTaskWithChan(request *CreateRobotTaskRequest) (<-chan *CreateRobotTaskResponse, <-chan error) {
	responseChan := make(chan *CreateRobotTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRobotTask(request)
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

// CreateRobotTaskWithCallback invokes the dyvmsapi.CreateRobotTask API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/createrobottask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRobotTaskWithCallback(request *CreateRobotTaskRequest, callback func(response *CreateRobotTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRobotTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateRobotTask(request)
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

// CreateRobotTaskRequest is the request struct for api CreateRobotTask
type CreateRobotTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RecallStateCodes     string           `position:"Query" name:"RecallStateCodes"`
	TaskName             string           `position:"Query" name:"TaskName"`
	RecallTimes          requests.Integer `position:"Query" name:"RecallTimes"`
	IsSelfLine           requests.Boolean `position:"Query" name:"IsSelfLine"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RetryType            requests.Integer `position:"Query" name:"RetryType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DialogId             requests.Integer `position:"Query" name:"DialogId"`
	Caller               string           `position:"Query" name:"Caller"`
	NumberStatusIdent    requests.Boolean `position:"Query" name:"NumberStatusIdent"`
	CorpName             string           `position:"Query" name:"CorpName"`
	RecallInterval       requests.Integer `position:"Query" name:"RecallInterval"`
}

// CreateRobotTaskResponse is the response struct for api CreateRobotTask
type CreateRobotTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateRobotTaskRequest creates a request to invoke CreateRobotTask API
func CreateCreateRobotTaskRequest() (request *CreateRobotTaskRequest) {
	request = &CreateRobotTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "CreateRobotTask", "dyvms", "openAPI")
	return
}

// CreateCreateRobotTaskResponse creates a response to parse from CreateRobotTask response
func CreateCreateRobotTaskResponse() (response *CreateRobotTaskResponse) {
	response = &CreateRobotTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
