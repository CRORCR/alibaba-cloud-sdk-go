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

// ModifyDedicatedHostGroupAttribute invokes the rds.ModifyDedicatedHostGroupAttribute API synchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostgroupattribute.html
func (client *Client) ModifyDedicatedHostGroupAttribute(request *ModifyDedicatedHostGroupAttributeRequest) (response *ModifyDedicatedHostGroupAttributeResponse, err error) {
	response = CreateModifyDedicatedHostGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDedicatedHostGroupAttributeWithChan invokes the rds.ModifyDedicatedHostGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostgroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedHostGroupAttributeWithChan(request *ModifyDedicatedHostGroupAttributeRequest) (<-chan *ModifyDedicatedHostGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyDedicatedHostGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDedicatedHostGroupAttribute(request)
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

// ModifyDedicatedHostGroupAttributeWithCallback invokes the rds.ModifyDedicatedHostGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostgroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedHostGroupAttributeWithCallback(request *ModifyDedicatedHostGroupAttributeRequest, callback func(response *ModifyDedicatedHostGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDedicatedHostGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDedicatedHostGroupAttribute(request)
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

// ModifyDedicatedHostGroupAttributeRequest is the request struct for api ModifyDedicatedHostGroupAttribute
type ModifyDedicatedHostGroupAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CpuAllocationRatio     requests.Integer `position:"Query" name:"CpuAllocationRatio"`
	DedicatedHostGroupId   string           `position:"Query" name:"DedicatedHostGroupId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	DiskAllocationRatio    requests.Integer `position:"Query" name:"DiskAllocationRatio"`
	MemAllocationRatio     requests.Integer `position:"Query" name:"MemAllocationRatio"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	HostReplacePolicy      string           `position:"Query" name:"HostReplacePolicy"`
	DedicatedHostGroupDesc string           `position:"Query" name:"DedicatedHostGroupDesc"`
	AllocationPolicy       string           `position:"Query" name:"AllocationPolicy"`
}

// ModifyDedicatedHostGroupAttributeResponse is the response struct for api ModifyDedicatedHostGroupAttribute
type ModifyDedicatedHostGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDedicatedHostGroupAttributeRequest creates a request to invoke ModifyDedicatedHostGroupAttribute API
func CreateModifyDedicatedHostGroupAttributeRequest() (request *ModifyDedicatedHostGroupAttributeRequest) {
	request = &ModifyDedicatedHostGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDedicatedHostGroupAttribute", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDedicatedHostGroupAttributeResponse creates a response to parse from ModifyDedicatedHostGroupAttribute response
func CreateModifyDedicatedHostGroupAttributeResponse() (response *ModifyDedicatedHostGroupAttributeResponse) {
	response = &ModifyDedicatedHostGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
