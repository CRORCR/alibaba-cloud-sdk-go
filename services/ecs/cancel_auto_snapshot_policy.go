package ecs

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

// CancelAutoSnapshotPolicy invokes the ecs.CancelAutoSnapshotPolicy API synchronously
// api document: https://help.aliyun.com/api/ecs/cancelautosnapshotpolicy.html
func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (response *CancelAutoSnapshotPolicyResponse, err error) {
	response = CreateCancelAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CancelAutoSnapshotPolicyWithChan invokes the ecs.CancelAutoSnapshotPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/cancelautosnapshotpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelAutoSnapshotPolicyWithChan(request *CancelAutoSnapshotPolicyRequest) (<-chan *CancelAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *CancelAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelAutoSnapshotPolicy(request)
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

// CancelAutoSnapshotPolicyWithCallback invokes the ecs.CancelAutoSnapshotPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/cancelautosnapshotpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelAutoSnapshotPolicyWithCallback(request *CancelAutoSnapshotPolicyRequest, callback func(response *CancelAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.CancelAutoSnapshotPolicy(request)
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

// CancelAutoSnapshotPolicyRequest is the request struct for api CancelAutoSnapshotPolicy
type CancelAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DiskIds              string           `position:"Query" name:"diskIds"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CancelAutoSnapshotPolicyResponse is the response struct for api CancelAutoSnapshotPolicy
type CancelAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelAutoSnapshotPolicyRequest creates a request to invoke CancelAutoSnapshotPolicy API
func CreateCancelAutoSnapshotPolicyRequest() (request *CancelAutoSnapshotPolicyRequest) {
	request = &CancelAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CancelAutoSnapshotPolicy", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelAutoSnapshotPolicyResponse creates a response to parse from CancelAutoSnapshotPolicy response
func CreateCancelAutoSnapshotPolicyResponse() (response *CancelAutoSnapshotPolicyResponse) {
	response = &CancelAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
