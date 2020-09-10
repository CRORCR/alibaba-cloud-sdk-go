package ddoscoo

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

// CreateSchedulerRule invokes the ddoscoo.CreateSchedulerRule API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/createschedulerrule.html
func (client *Client) CreateSchedulerRule(request *CreateSchedulerRuleRequest) (response *CreateSchedulerRuleResponse, err error) {
	response = CreateCreateSchedulerRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSchedulerRuleWithChan invokes the ddoscoo.CreateSchedulerRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/createschedulerrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSchedulerRuleWithChan(request *CreateSchedulerRuleRequest) (<-chan *CreateSchedulerRuleResponse, <-chan error) {
	responseChan := make(chan *CreateSchedulerRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSchedulerRule(request)
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

// CreateSchedulerRuleWithCallback invokes the ddoscoo.CreateSchedulerRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/createschedulerrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSchedulerRuleWithCallback(request *CreateSchedulerRuleRequest, callback func(response *CreateSchedulerRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSchedulerRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateSchedulerRule(request)
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

// CreateSchedulerRuleRequest is the request struct for api CreateSchedulerRule
type CreateSchedulerRuleRequest struct {
	*requests.RpcRequest
	Rules           string           `position:"Query" name:"Rules"`
	RuleName        string           `position:"Query" name:"RuleName"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Param           string           `position:"Query" name:"Param"`
	RuleType        requests.Integer `position:"Query" name:"RuleType"`
}

// CreateSchedulerRuleResponse is the response struct for api CreateSchedulerRule
type CreateSchedulerRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Cname     string `json:"Cname" xml:"Cname"`
	RuleName  string `json:"RuleName" xml:"RuleName"`
}

// CreateCreateSchedulerRuleRequest creates a request to invoke CreateSchedulerRule API
func CreateCreateSchedulerRuleRequest() (request *CreateSchedulerRuleRequest) {
	request = &CreateSchedulerRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "CreateSchedulerRule", "ddoscoo", "openAPI")
	return
}

// CreateCreateSchedulerRuleResponse creates a response to parse from CreateSchedulerRule response
func CreateCreateSchedulerRuleResponse() (response *CreateSchedulerRuleResponse) {
	response = &CreateSchedulerRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
