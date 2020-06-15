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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDomainIpv6Status invokes the waf_openapi.ModifyDomainIpv6Status API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifydomainipv6status.html
func (client *Client) ModifyDomainIpv6Status(request *ModifyDomainIpv6StatusRequest) (response *ModifyDomainIpv6StatusResponse, err error) {
	response = CreateModifyDomainIpv6StatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDomainIpv6StatusWithChan invokes the waf_openapi.ModifyDomainIpv6Status API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifydomainipv6status.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDomainIpv6StatusWithChan(request *ModifyDomainIpv6StatusRequest) (<-chan *ModifyDomainIpv6StatusResponse, <-chan error) {
	responseChan := make(chan *ModifyDomainIpv6StatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDomainIpv6Status(request)
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

// ModifyDomainIpv6StatusWithCallback invokes the waf_openapi.ModifyDomainIpv6Status API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifydomainipv6status.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDomainIpv6StatusWithCallback(request *ModifyDomainIpv6StatusRequest, callback func(response *ModifyDomainIpv6StatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDomainIpv6StatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyDomainIpv6Status(request)
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

// ModifyDomainIpv6StatusRequest is the request struct for api ModifyDomainIpv6Status
type ModifyDomainIpv6StatusRequest struct {
	*requests.RpcRequest
	WafVersion string `position:"Query" name:"WafVersion"`
	Enabled    string `position:"Query" name:"Enabled"`
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Domain     string `position:"Query" name:"Domain"`
	Lang       string `position:"Query" name:"Lang"`
}

// ModifyDomainIpv6StatusResponse is the response struct for api ModifyDomainIpv6Status
type ModifyDomainIpv6StatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDomainIpv6StatusRequest creates a request to invoke ModifyDomainIpv6Status API
func CreateModifyDomainIpv6StatusRequest() (request *ModifyDomainIpv6StatusRequest) {
	request = &ModifyDomainIpv6StatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "ModifyDomainIpv6Status", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDomainIpv6StatusResponse creates a response to parse from ModifyDomainIpv6Status response
func CreateModifyDomainIpv6StatusResponse() (response *ModifyDomainIpv6StatusResponse) {
	response = &ModifyDomainIpv6StatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
