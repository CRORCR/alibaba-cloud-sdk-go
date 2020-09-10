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

// DeletePhysicalConnection invokes the ecs.DeletePhysicalConnection API synchronously
// api document: https://help.aliyun.com/api/ecs/deletephysicalconnection.html
func (client *Client) DeletePhysicalConnection(request *DeletePhysicalConnectionRequest) (response *DeletePhysicalConnectionResponse, err error) {
	response = CreateDeletePhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePhysicalConnectionWithChan invokes the ecs.DeletePhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletephysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePhysicalConnectionWithChan(request *DeletePhysicalConnectionRequest) (<-chan *DeletePhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *DeletePhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePhysicalConnection(request)
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

// DeletePhysicalConnectionWithCallback invokes the ecs.DeletePhysicalConnection API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletephysicalconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePhysicalConnectionWithCallback(request *DeletePhysicalConnectionRequest, callback func(response *DeletePhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.DeletePhysicalConnection(request)
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

// DeletePhysicalConnectionRequest is the request struct for api DeletePhysicalConnection
type DeletePhysicalConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
}

// DeletePhysicalConnectionResponse is the response struct for api DeletePhysicalConnection
type DeletePhysicalConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePhysicalConnectionRequest creates a request to invoke DeletePhysicalConnection API
func CreateDeletePhysicalConnectionRequest() (request *DeletePhysicalConnectionRequest) {
	request = &DeletePhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeletePhysicalConnection", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePhysicalConnectionResponse creates a response to parse from DeletePhysicalConnection response
func CreateDeletePhysicalConnectionResponse() (response *DeletePhysicalConnectionResponse) {
	response = &DeletePhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
