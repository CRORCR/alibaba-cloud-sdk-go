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

// DeleteDemand invokes the ecs.DeleteDemand API synchronously
// api document: https://help.aliyun.com/api/ecs/deletedemand.html
func (client *Client) DeleteDemand(request *DeleteDemandRequest) (response *DeleteDemandResponse, err error) {
	response = CreateDeleteDemandResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDemandWithChan invokes the ecs.DeleteDemand API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletedemand.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDemandWithChan(request *DeleteDemandRequest) (<-chan *DeleteDemandResponse, <-chan error) {
	responseChan := make(chan *DeleteDemandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDemand(request)
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

// DeleteDemandWithCallback invokes the ecs.DeleteDemand API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletedemand.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDemandWithCallback(request *DeleteDemandRequest, callback func(response *DeleteDemandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDemandResponse
		var err error
		defer close(result)
		response, err = client.DeleteDemand(request)
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

// DeleteDemandRequest is the request struct for api DeleteDemand
type DeleteDemandRequest struct {
	*requests.RpcRequest
	Reason               string           `position:"Query" name:"Reason"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DemandId             string           `position:"Query" name:"DemandId"`
}

// DeleteDemandResponse is the response struct for api DeleteDemand
type DeleteDemandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDemandRequest creates a request to invoke DeleteDemand API
func CreateDeleteDemandRequest() (request *DeleteDemandRequest) {
	request = &DeleteDemandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteDemand", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDemandResponse creates a response to parse from DeleteDemand response
func CreateDeleteDemandResponse() (response *DeleteDemandResponse) {
	response = &DeleteDemandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
