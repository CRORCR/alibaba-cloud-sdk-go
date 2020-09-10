package sas

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

// ModifyAntiBruteForceRule invokes the sas.ModifyAntiBruteForceRule API synchronously
// api document: https://help.aliyun.com/api/sas/modifyantibruteforcerule.html
func (client *Client) ModifyAntiBruteForceRule(request *ModifyAntiBruteForceRuleRequest) (response *ModifyAntiBruteForceRuleResponse, err error) {
	response = CreateModifyAntiBruteForceRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAntiBruteForceRuleWithChan invokes the sas.ModifyAntiBruteForceRule API asynchronously
// api document: https://help.aliyun.com/api/sas/modifyantibruteforcerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAntiBruteForceRuleWithChan(request *ModifyAntiBruteForceRuleRequest) (<-chan *ModifyAntiBruteForceRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyAntiBruteForceRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAntiBruteForceRule(request)
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

// ModifyAntiBruteForceRuleWithCallback invokes the sas.ModifyAntiBruteForceRule API asynchronously
// api document: https://help.aliyun.com/api/sas/modifyantibruteforcerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAntiBruteForceRuleWithCallback(request *ModifyAntiBruteForceRuleRequest, callback func(response *ModifyAntiBruteForceRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAntiBruteForceRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyAntiBruteForceRule(request)
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

// ModifyAntiBruteForceRuleRequest is the request struct for api ModifyAntiBruteForceRule
type ModifyAntiBruteForceRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ForbiddenTime   requests.Integer `position:"Query" name:"ForbiddenTime"`
	FailCount       requests.Integer `position:"Query" name:"FailCount"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	EnableSmartRule requests.Boolean `position:"Query" name:"EnableSmartRule"`
	UuidList        *[]string        `position:"Query" name:"UuidList"  type:"Repeated"`
	Id              requests.Integer `position:"Query" name:"Id"`
	Name            string           `position:"Query" name:"Name"`
	Span            requests.Integer `position:"Query" name:"Span"`
	DefaultRule     requests.Boolean `position:"Query" name:"DefaultRule"`
}

// ModifyAntiBruteForceRuleResponse is the response struct for api ModifyAntiBruteForceRule
type ModifyAntiBruteForceRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAntiBruteForceRuleRequest creates a request to invoke ModifyAntiBruteForceRule API
func CreateModifyAntiBruteForceRuleRequest() (request *ModifyAntiBruteForceRuleRequest) {
	request = &ModifyAntiBruteForceRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyAntiBruteForceRule", "sas", "openAPI")
	return
}

// CreateModifyAntiBruteForceRuleResponse creates a response to parse from ModifyAntiBruteForceRule response
func CreateModifyAntiBruteForceRuleResponse() (response *ModifyAntiBruteForceRuleResponse) {
	response = &ModifyAntiBruteForceRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
