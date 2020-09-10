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

// DescribeConfigRule invokes the config.DescribeConfigRule API synchronously
// api document: https://help.aliyun.com/api/config/describeconfigrule.html
func (client *Client) DescribeConfigRule(request *DescribeConfigRuleRequest) (response *DescribeConfigRuleResponse, err error) {
	response = CreateDescribeConfigRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigRuleWithChan invokes the config.DescribeConfigRule API asynchronously
// api document: https://help.aliyun.com/api/config/describeconfigrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigRuleWithChan(request *DescribeConfigRuleRequest) (<-chan *DescribeConfigRuleResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigRule(request)
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

// DescribeConfigRuleWithCallback invokes the config.DescribeConfigRule API asynchronously
// api document: https://help.aliyun.com/api/config/describeconfigrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigRuleWithCallback(request *DescribeConfigRuleRequest, callback func(response *DescribeConfigRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigRuleResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigRule(request)
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

// DescribeConfigRuleRequest is the request struct for api DescribeConfigRule
type DescribeConfigRuleRequest struct {
	*requests.RpcRequest
	ConfigRuleId string           `position:"Query" name:"ConfigRuleId"`
	MultiAccount requests.Boolean `position:"Query" name:"MultiAccount"`
	MemberId     requests.Integer `position:"Query" name:"MemberId"`
}

// DescribeConfigRuleResponse is the response struct for api DescribeConfigRule
type DescribeConfigRuleResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	ConfigRule ConfigRule `json:"ConfigRule" xml:"ConfigRule"`
}

// CreateDescribeConfigRuleRequest creates a request to invoke DescribeConfigRule API
func CreateDescribeConfigRuleRequest() (request *DescribeConfigRuleRequest) {
	request = &DescribeConfigRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "DescribeConfigRule", "config", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeConfigRuleResponse creates a response to parse from DescribeConfigRule response
func CreateDescribeConfigRuleResponse() (response *DescribeConfigRuleResponse) {
	response = &DescribeConfigRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
