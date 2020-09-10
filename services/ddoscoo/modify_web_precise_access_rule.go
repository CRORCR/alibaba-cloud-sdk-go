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

// ModifyWebPreciseAccessRule invokes the ddoscoo.ModifyWebPreciseAccessRule API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifywebpreciseaccessrule.html
func (client *Client) ModifyWebPreciseAccessRule(request *ModifyWebPreciseAccessRuleRequest) (response *ModifyWebPreciseAccessRuleResponse, err error) {
	response = CreateModifyWebPreciseAccessRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebPreciseAccessRuleWithChan invokes the ddoscoo.ModifyWebPreciseAccessRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifywebpreciseaccessrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebPreciseAccessRuleWithChan(request *ModifyWebPreciseAccessRuleRequest) (<-chan *ModifyWebPreciseAccessRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyWebPreciseAccessRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebPreciseAccessRule(request)
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

// ModifyWebPreciseAccessRuleWithCallback invokes the ddoscoo.ModifyWebPreciseAccessRule API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/modifywebpreciseaccessrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebPreciseAccessRuleWithCallback(request *ModifyWebPreciseAccessRuleRequest, callback func(response *ModifyWebPreciseAccessRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebPreciseAccessRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebPreciseAccessRule(request)
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

// ModifyWebPreciseAccessRuleRequest is the request struct for api ModifyWebPreciseAccessRule
type ModifyWebPreciseAccessRuleRequest struct {
	*requests.RpcRequest
	Expires         requests.Integer `position:"Query" name:"Expires"`
	Rules           string           `position:"Query" name:"Rules"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Domain          string           `position:"Query" name:"Domain"`
}

// ModifyWebPreciseAccessRuleResponse is the response struct for api ModifyWebPreciseAccessRule
type ModifyWebPreciseAccessRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebPreciseAccessRuleRequest creates a request to invoke ModifyWebPreciseAccessRule API
func CreateModifyWebPreciseAccessRuleRequest() (request *ModifyWebPreciseAccessRuleRequest) {
	request = &ModifyWebPreciseAccessRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyWebPreciseAccessRule", "ddoscoo", "openAPI")
	return
}

// CreateModifyWebPreciseAccessRuleResponse creates a response to parse from ModifyWebPreciseAccessRule response
func CreateModifyWebPreciseAccessRuleResponse() (response *ModifyWebPreciseAccessRuleResponse) {
	response = &ModifyWebPreciseAccessRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
