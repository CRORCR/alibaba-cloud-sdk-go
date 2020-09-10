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

// ModifyLogServiceStatus invokes the waf_openapi.ModifyLogServiceStatus API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifylogservicestatus.html
func (client *Client) ModifyLogServiceStatus(request *ModifyLogServiceStatusRequest) (response *ModifyLogServiceStatusResponse, err error) {
	response = CreateModifyLogServiceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLogServiceStatusWithChan invokes the waf_openapi.ModifyLogServiceStatus API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifylogservicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogServiceStatusWithChan(request *ModifyLogServiceStatusRequest) (<-chan *ModifyLogServiceStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyLogServiceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLogServiceStatus(request)
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

// ModifyLogServiceStatusWithCallback invokes the waf_openapi.ModifyLogServiceStatus API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifylogservicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogServiceStatusWithCallback(request *ModifyLogServiceStatusRequest, callback func(response *ModifyLogServiceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLogServiceStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyLogServiceStatus(request)
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

// ModifyLogServiceStatusRequest is the request struct for api ModifyLogServiceStatus
type ModifyLogServiceStatusRequest struct {
	*requests.RpcRequest
	Enabled    requests.Integer `position:"Query" name:"Enabled"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	Domain     string           `position:"Query" name:"Domain"`
	Lang       string           `position:"Query" name:"Lang"`
}

// ModifyLogServiceStatusResponse is the response struct for api ModifyLogServiceStatus
type ModifyLogServiceStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLogServiceStatusRequest creates a request to invoke ModifyLogServiceStatus API
func CreateModifyLogServiceStatusRequest() (request *ModifyLogServiceStatusRequest) {
	request = &ModifyLogServiceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "ModifyLogServiceStatus", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLogServiceStatusResponse creates a response to parse from ModifyLogServiceStatus response
func CreateModifyLogServiceStatusResponse() (response *ModifyLogServiceStatusResponse) {
	response = &ModifyLogServiceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
