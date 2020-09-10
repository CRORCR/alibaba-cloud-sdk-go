package yundun_ds

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

// ModifyEventTypeStatus invokes the yundun_ds.ModifyEventTypeStatus API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/modifyeventtypestatus.html
func (client *Client) ModifyEventTypeStatus(request *ModifyEventTypeStatusRequest) (response *ModifyEventTypeStatusResponse, err error) {
	response = CreateModifyEventTypeStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyEventTypeStatusWithChan invokes the yundun_ds.ModifyEventTypeStatus API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/modifyeventtypestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyEventTypeStatusWithChan(request *ModifyEventTypeStatusRequest) (<-chan *ModifyEventTypeStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyEventTypeStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyEventTypeStatus(request)
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

// ModifyEventTypeStatusWithCallback invokes the yundun_ds.ModifyEventTypeStatus API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/modifyeventtypestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyEventTypeStatusWithCallback(request *ModifyEventTypeStatusRequest, callback func(response *ModifyEventTypeStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyEventTypeStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyEventTypeStatus(request)
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

// ModifyEventTypeStatusRequest is the request struct for api ModifyEventTypeStatus
type ModifyEventTypeStatusRequest struct {
	*requests.RpcRequest
	SubTypeIds string `position:"Query" name:"SubTypeIds"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// ModifyEventTypeStatusResponse is the response struct for api ModifyEventTypeStatus
type ModifyEventTypeStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyEventTypeStatusRequest creates a request to invoke ModifyEventTypeStatus API
func CreateModifyEventTypeStatusRequest() (request *ModifyEventTypeStatusRequest) {
	request = &ModifyEventTypeStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "ModifyEventTypeStatus", "sddp", "openAPI")
	return
}

// CreateModifyEventTypeStatusResponse creates a response to parse from ModifyEventTypeStatus response
func CreateModifyEventTypeStatusResponse() (response *ModifyEventTypeStatusResponse) {
	response = &ModifyEventTypeStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
