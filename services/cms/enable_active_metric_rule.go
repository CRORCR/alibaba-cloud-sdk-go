package cms

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

// EnableActiveMetricRule invokes the cms.EnableActiveMetricRule API synchronously
// api document: https://help.aliyun.com/api/cms/enableactivemetricrule.html
func (client *Client) EnableActiveMetricRule(request *EnableActiveMetricRuleRequest) (response *EnableActiveMetricRuleResponse, err error) {
	response = CreateEnableActiveMetricRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableActiveMetricRuleWithChan invokes the cms.EnableActiveMetricRule API asynchronously
// api document: https://help.aliyun.com/api/cms/enableactivemetricrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableActiveMetricRuleWithChan(request *EnableActiveMetricRuleRequest) (<-chan *EnableActiveMetricRuleResponse, <-chan error) {
	responseChan := make(chan *EnableActiveMetricRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableActiveMetricRule(request)
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

// EnableActiveMetricRuleWithCallback invokes the cms.EnableActiveMetricRule API asynchronously
// api document: https://help.aliyun.com/api/cms/enableactivemetricrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableActiveMetricRuleWithCallback(request *EnableActiveMetricRuleRequest, callback func(response *EnableActiveMetricRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableActiveMetricRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableActiveMetricRule(request)
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

// EnableActiveMetricRuleRequest is the request struct for api EnableActiveMetricRule
type EnableActiveMetricRuleRequest struct {
	*requests.RpcRequest
	Product string `position:"Query" name:"Product"`
}

// EnableActiveMetricRuleResponse is the response struct for api EnableActiveMetricRule
type EnableActiveMetricRuleResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableActiveMetricRuleRequest creates a request to invoke EnableActiveMetricRule API
func CreateEnableActiveMetricRuleRequest() (request *EnableActiveMetricRuleRequest) {
	request = &EnableActiveMetricRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "EnableActiveMetricRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableActiveMetricRuleResponse creates a response to parse from EnableActiveMetricRule response
func CreateEnableActiveMetricRuleResponse() (response *EnableActiveMetricRuleResponse) {
	response = &EnableActiveMetricRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
