package ess

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

// DeleteLifecycleHook invokes the ess.DeleteLifecycleHook API synchronously
// api document: https://help.aliyun.com/api/ess/deletelifecyclehook.html
func (client *Client) DeleteLifecycleHook(request *DeleteLifecycleHookRequest) (response *DeleteLifecycleHookResponse, err error) {
	response = CreateDeleteLifecycleHookResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLifecycleHookWithChan invokes the ess.DeleteLifecycleHook API asynchronously
// api document: https://help.aliyun.com/api/ess/deletelifecyclehook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLifecycleHookWithChan(request *DeleteLifecycleHookRequest) (<-chan *DeleteLifecycleHookResponse, <-chan error) {
	responseChan := make(chan *DeleteLifecycleHookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLifecycleHook(request)
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

// DeleteLifecycleHookWithCallback invokes the ess.DeleteLifecycleHook API asynchronously
// api document: https://help.aliyun.com/api/ess/deletelifecyclehook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLifecycleHookWithCallback(request *DeleteLifecycleHookRequest, callback func(response *DeleteLifecycleHookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLifecycleHookResponse
		var err error
		defer close(result)
		response, err = client.DeleteLifecycleHook(request)
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

// DeleteLifecycleHookRequest is the request struct for api DeleteLifecycleHook
type DeleteLifecycleHookRequest struct {
	*requests.RpcRequest
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	LifecycleHookName    string           `position:"Query" name:"LifecycleHookName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	LifecycleHookId      string           `position:"Query" name:"LifecycleHookId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLifecycleHookResponse is the response struct for api DeleteLifecycleHook
type DeleteLifecycleHookResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLifecycleHookRequest creates a request to invoke DeleteLifecycleHook API
func CreateDeleteLifecycleHookRequest() (request *DeleteLifecycleHookRequest) {
	request = &DeleteLifecycleHookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DeleteLifecycleHook", "ess", "openAPI")
	return
}

// CreateDeleteLifecycleHookResponse creates a response to parse from DeleteLifecycleHook response
func CreateDeleteLifecycleHookResponse() (response *DeleteLifecycleHookResponse) {
	response = &DeleteLifecycleHookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
