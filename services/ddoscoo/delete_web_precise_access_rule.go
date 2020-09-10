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

// DeleteWebPreciseAccessRule invokes the ddoscoo.DeleteWebPreciseAccessRule API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/deletewebpreciseaccessrule.html
func (client *Client) DeleteWebPreciseAccessRule(request *DeleteWebPreciseAccessRuleRequest) (response *DeleteWebPreciseAccessRuleResponse, err error) {
	response = CreateDeleteWebPreciseAccessRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWebPreciseAccessRuleWithChan invokes the ddoscoo.DeleteWebPreciseAccessRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/deletewebpreciseaccessrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteWebPreciseAccessRuleWithChan(request *DeleteWebPreciseAccessRuleRequest) (<-chan *DeleteWebPreciseAccessRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteWebPreciseAccessRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWebPreciseAccessRule(request)
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

// DeleteWebPreciseAccessRuleWithCallback invokes the ddoscoo.DeleteWebPreciseAccessRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/deletewebpreciseaccessrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteWebPreciseAccessRuleWithCallback(request *DeleteWebPreciseAccessRuleRequest, callback func(response *DeleteWebPreciseAccessRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWebPreciseAccessRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteWebPreciseAccessRule(request)
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

// DeleteWebPreciseAccessRuleRequest is the request struct for api DeleteWebPreciseAccessRule
type DeleteWebPreciseAccessRuleRequest struct {
	*requests.RpcRequest
	ResourceGroupId string    `position:"Query" name:"ResourceGroupId"`
	SourceIp        string    `position:"Query" name:"SourceIp"`
	RuleNames       *[]string `position:"Query" name:"RuleNames"  type:"Repeated"`
	Domain          string    `position:"Query" name:"Domain"`
}

// DeleteWebPreciseAccessRuleResponse is the response struct for api DeleteWebPreciseAccessRule
type DeleteWebPreciseAccessRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteWebPreciseAccessRuleRequest creates a request to invoke DeleteWebPreciseAccessRule API
func CreateDeleteWebPreciseAccessRuleRequest() (request *DeleteWebPreciseAccessRuleRequest) {
	request = &DeleteWebPreciseAccessRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DeleteWebPreciseAccessRule", "ddoscoo", "openAPI")
	return
}

// CreateDeleteWebPreciseAccessRuleResponse creates a response to parse from DeleteWebPreciseAccessRule response
func CreateDeleteWebPreciseAccessRuleResponse() (response *DeleteWebPreciseAccessRuleResponse) {
	response = &DeleteWebPreciseAccessRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
