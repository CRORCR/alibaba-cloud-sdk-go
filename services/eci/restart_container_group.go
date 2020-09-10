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

package eci

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// RestartContainerGroup invokes the eci.RestartContainerGroup API synchronously
// api document: https://help.aliyun.com/api/eci/restartcontainergroup.html
func (client *Client) RestartContainerGroup(request *RestartContainerGroupRequest) (response *RestartContainerGroupResponse, err error) {
	response = CreateRestartContainerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RestartContainerGroupWithChan invokes the eci.RestartContainerGroup API asynchronously
// api document: https://help.aliyun.com/api/eci/restartcontainergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartContainerGroupWithChan(request *RestartContainerGroupRequest) (<-chan *RestartContainerGroupResponse, <-chan error) {
	responseChan := make(chan *RestartContainerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartContainerGroup(request)
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

// RestartContainerGroupWithCallback invokes the eci.RestartContainerGroup API asynchronously
// api document: https://help.aliyun.com/api/eci/restartcontainergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartContainerGroupWithCallback(request *RestartContainerGroupRequest, callback func(response *RestartContainerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartContainerGroupResponse
		var err error
		defer close(result)
		response, err = client.RestartContainerGroup(request)
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

// RestartContainerGroupRequest is the request struct for api RestartContainerGroup
type RestartContainerGroupRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RegionId             string           `position:"Query" name:"RegionId"`
	ContainerGroupId     string           `position:"Query" name:"ContainerGroupId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	VkClientVersion      string           `position:"Query" name:"VkClientVersion"`
}

// RestartContainerGroupResponse is the response struct for api RestartContainerGroup
type RestartContainerGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestartContainerGroupRequest creates a request to invoke RestartContainerGroup API
func CreateRestartContainerGroupRequest() (request *RestartContainerGroupRequest) {
	request = &RestartContainerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "RestartContainerGroup", "eci", "openAPI")
	return
}

// CreateRestartContainerGroupResponse creates a response to parse from RestartContainerGroup response
func CreateRestartContainerGroupResponse() (response *RestartContainerGroupResponse) {
	response = &RestartContainerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
