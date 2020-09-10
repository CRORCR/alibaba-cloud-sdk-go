package waf_openapi

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

// CreateProtectionModuleRule invokes the waf_openapi.CreateProtectionModuleRule API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/createprotectionmodulerule.html
func (client *Client) CreateProtectionModuleRule(request *CreateProtectionModuleRuleRequest) (response *CreateProtectionModuleRuleResponse, err error) {
	response = CreateCreateProtectionModuleRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProtectionModuleRuleWithChan invokes the waf_openapi.CreateProtectionModuleRule API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/createprotectionmodulerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProtectionModuleRuleWithChan(request *CreateProtectionModuleRuleRequest) (<-chan *CreateProtectionModuleRuleResponse, <-chan error) {
	responseChan := make(chan *CreateProtectionModuleRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProtectionModuleRule(request)
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

// CreateProtectionModuleRuleWithCallback invokes the waf_openapi.CreateProtectionModuleRule API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/createprotectionmodulerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProtectionModuleRuleWithCallback(request *CreateProtectionModuleRuleRequest, callback func(response *CreateProtectionModuleRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProtectionModuleRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateProtectionModuleRule(request)
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

// CreateProtectionModuleRuleRequest is the request struct for api CreateProtectionModuleRule
type CreateProtectionModuleRuleRequest struct {
	*requests.RpcRequest
	Rule        string `position:"Query" name:"Rule"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	Lang        string `position:"Query" name:"Lang"`
	DefenseType string `position:"Query" name:"DefenseType"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Domain      string `position:"Query" name:"Domain"`
	Region      string `position:"Query" name:"Region"`
}

// CreateProtectionModuleRuleResponse is the response struct for api CreateProtectionModuleRule
type CreateProtectionModuleRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateProtectionModuleRuleRequest creates a request to invoke CreateProtectionModuleRule API
func CreateCreateProtectionModuleRuleRequest() (request *CreateProtectionModuleRuleRequest) {
	request = &CreateProtectionModuleRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "CreateProtectionModuleRule", "waf", "openAPI")
	return
}

// CreateCreateProtectionModuleRuleResponse creates a response to parse from CreateProtectionModuleRule response
func CreateCreateProtectionModuleRuleResponse() (response *CreateProtectionModuleRuleResponse) {
	response = &CreateProtectionModuleRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
