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

// DisableMetricRules invokes the cms.DisableMetricRules API synchronously
// api document: https://help.aliyun.com/api/cms/disablemetricrules.html
func (client *Client) DisableMetricRules(request *DisableMetricRulesRequest) (response *DisableMetricRulesResponse, err error) {
	response = CreateDisableMetricRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DisableMetricRulesWithChan invokes the cms.DisableMetricRules API asynchronously
// api document: https://help.aliyun.com/api/cms/disablemetricrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableMetricRulesWithChan(request *DisableMetricRulesRequest) (<-chan *DisableMetricRulesResponse, <-chan error) {
	responseChan := make(chan *DisableMetricRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableMetricRules(request)
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

// DisableMetricRulesWithCallback invokes the cms.DisableMetricRules API asynchronously
// api document: https://help.aliyun.com/api/cms/disablemetricrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableMetricRulesWithCallback(request *DisableMetricRulesRequest, callback func(response *DisableMetricRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableMetricRulesResponse
		var err error
		defer close(result)
		response, err = client.DisableMetricRules(request)
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

// DisableMetricRulesRequest is the request struct for api DisableMetricRules
type DisableMetricRulesRequest struct {
	*requests.RpcRequest
	RuleId *[]string `position:"Query" name:"RuleId"  type:"Repeated"`
}

// DisableMetricRulesResponse is the response struct for api DisableMetricRules
type DisableMetricRulesResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableMetricRulesRequest creates a request to invoke DisableMetricRules API
func CreateDisableMetricRulesRequest() (request *DisableMetricRulesRequest) {
	request = &DisableMetricRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DisableMetricRules", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableMetricRulesResponse creates a response to parse from DisableMetricRules response
func CreateDisableMetricRulesResponse() (response *DisableMetricRulesResponse) {
	response = &DisableMetricRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
