package qualitycheck

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

// CreateTaskAssignRule invokes the qualitycheck.CreateTaskAssignRule API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/createtaskassignrule.html
func (client *Client) CreateTaskAssignRule(request *CreateTaskAssignRuleRequest) (response *CreateTaskAssignRuleResponse, err error) {
	response = CreateCreateTaskAssignRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTaskAssignRuleWithChan invokes the qualitycheck.CreateTaskAssignRule API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/createtaskassignrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTaskAssignRuleWithChan(request *CreateTaskAssignRuleRequest) (<-chan *CreateTaskAssignRuleResponse, <-chan error) {
	responseChan := make(chan *CreateTaskAssignRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTaskAssignRule(request)
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

// CreateTaskAssignRuleWithCallback invokes the qualitycheck.CreateTaskAssignRule API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/createtaskassignrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTaskAssignRuleWithCallback(request *CreateTaskAssignRuleRequest, callback func(response *CreateTaskAssignRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTaskAssignRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateTaskAssignRule(request)
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

// CreateTaskAssignRuleRequest is the request struct for api CreateTaskAssignRule
type CreateTaskAssignRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// CreateTaskAssignRuleResponse is the response struct for api CreateTaskAssignRule
type CreateTaskAssignRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCreateTaskAssignRuleRequest creates a request to invoke CreateTaskAssignRule API
func CreateCreateTaskAssignRuleRequest() (request *CreateTaskAssignRuleRequest) {
	request = &CreateTaskAssignRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "CreateTaskAssignRule", "", "")
	return
}

// CreateCreateTaskAssignRuleResponse creates a response to parse from CreateTaskAssignRule response
func CreateCreateTaskAssignRuleResponse() (response *CreateTaskAssignRuleResponse) {
	response = &CreateTaskAssignRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
