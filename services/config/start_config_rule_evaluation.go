package config

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

// StartConfigRuleEvaluation invokes the config.StartConfigRuleEvaluation API synchronously
// api document: https://help.aliyun.com/api/config/startconfigruleevaluation.html
func (client *Client) StartConfigRuleEvaluation(request *StartConfigRuleEvaluationRequest) (response *StartConfigRuleEvaluationResponse, err error) {
	response = CreateStartConfigRuleEvaluationResponse()
	err = client.DoAction(request, response)
	return
}

// StartConfigRuleEvaluationWithChan invokes the config.StartConfigRuleEvaluation API asynchronously
// api document: https://help.aliyun.com/api/config/startconfigruleevaluation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartConfigRuleEvaluationWithChan(request *StartConfigRuleEvaluationRequest) (<-chan *StartConfigRuleEvaluationResponse, <-chan error) {
	responseChan := make(chan *StartConfigRuleEvaluationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartConfigRuleEvaluation(request)
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

// StartConfigRuleEvaluationWithCallback invokes the config.StartConfigRuleEvaluation API asynchronously
// api document: https://help.aliyun.com/api/config/startconfigruleevaluation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartConfigRuleEvaluationWithCallback(request *StartConfigRuleEvaluationRequest, callback func(response *StartConfigRuleEvaluationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartConfigRuleEvaluationResponse
		var err error
		defer close(result)
		response, err = client.StartConfigRuleEvaluation(request)
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

// StartConfigRuleEvaluationRequest is the request struct for api StartConfigRuleEvaluation
type StartConfigRuleEvaluationRequest struct {
	*requests.RpcRequest
	ConfigRuleId string           `position:"Query" name:"ConfigRuleId"`
	MultiAccount requests.Boolean `position:"Query" name:"MultiAccount"`
	MemberId     requests.Integer `position:"Query" name:"MemberId"`
}

// StartConfigRuleEvaluationResponse is the response struct for api StartConfigRuleEvaluation
type StartConfigRuleEvaluationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateStartConfigRuleEvaluationRequest creates a request to invoke StartConfigRuleEvaluation API
func CreateStartConfigRuleEvaluationRequest() (request *StartConfigRuleEvaluationRequest) {
	request = &StartConfigRuleEvaluationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "StartConfigRuleEvaluation", "config", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartConfigRuleEvaluationResponse creates a response to parse from StartConfigRuleEvaluation response
func CreateStartConfigRuleEvaluationResponse() (response *StartConfigRuleEvaluationResponse) {
	response = &StartConfigRuleEvaluationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
