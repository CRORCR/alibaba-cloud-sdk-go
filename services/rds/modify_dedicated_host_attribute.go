package rds

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

// ModifyDedicatedHostAttribute invokes the rds.ModifyDedicatedHostAttribute API synchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostattribute.html
func (client *Client) ModifyDedicatedHostAttribute(request *ModifyDedicatedHostAttributeRequest) (response *ModifyDedicatedHostAttributeResponse, err error) {
	response = CreateModifyDedicatedHostAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDedicatedHostAttributeWithChan invokes the rds.ModifyDedicatedHostAttribute API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedHostAttributeWithChan(request *ModifyDedicatedHostAttributeRequest) (<-chan *ModifyDedicatedHostAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyDedicatedHostAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDedicatedHostAttribute(request)
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

// ModifyDedicatedHostAttributeWithCallback invokes the rds.ModifyDedicatedHostAttribute API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedHostAttributeWithCallback(request *ModifyDedicatedHostAttributeRequest, callback func(response *ModifyDedicatedHostAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDedicatedHostAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDedicatedHostAttribute(request)
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

// ModifyDedicatedHostAttributeRequest is the request struct for api ModifyDedicatedHostAttribute
type ModifyDedicatedHostAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HostName             string           `position:"Query" name:"HostName"`
	AllocationStatus     string           `position:"Query" name:"AllocationStatus"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDedicatedHostAttributeResponse is the response struct for api ModifyDedicatedHostAttribute
type ModifyDedicatedHostAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDedicatedHostAttributeRequest creates a request to invoke ModifyDedicatedHostAttribute API
func CreateModifyDedicatedHostAttributeRequest() (request *ModifyDedicatedHostAttributeRequest) {
	request = &ModifyDedicatedHostAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDedicatedHostAttribute", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDedicatedHostAttributeResponse creates a response to parse from ModifyDedicatedHostAttribute response
func CreateModifyDedicatedHostAttributeResponse() (response *ModifyDedicatedHostAttributeResponse) {
	response = &ModifyDedicatedHostAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
