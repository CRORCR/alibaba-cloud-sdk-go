package iot

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

// CreateConsumerGroupSubscribeRelation invokes the iot.CreateConsumerGroupSubscribeRelation API synchronously
// api document: https://help.aliyun.com/api/iot/createconsumergroupsubscriberelation.html
func (client *Client) CreateConsumerGroupSubscribeRelation(request *CreateConsumerGroupSubscribeRelationRequest) (response *CreateConsumerGroupSubscribeRelationResponse, err error) {
	response = CreateCreateConsumerGroupSubscribeRelationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateConsumerGroupSubscribeRelationWithChan invokes the iot.CreateConsumerGroupSubscribeRelation API asynchronously
// api document: https://help.aliyun.com/api/iot/createconsumergroupsubscriberelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateConsumerGroupSubscribeRelationWithChan(request *CreateConsumerGroupSubscribeRelationRequest) (<-chan *CreateConsumerGroupSubscribeRelationResponse, <-chan error) {
	responseChan := make(chan *CreateConsumerGroupSubscribeRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateConsumerGroupSubscribeRelation(request)
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

// CreateConsumerGroupSubscribeRelationWithCallback invokes the iot.CreateConsumerGroupSubscribeRelation API asynchronously
// api document: https://help.aliyun.com/api/iot/createconsumergroupsubscriberelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateConsumerGroupSubscribeRelationWithCallback(request *CreateConsumerGroupSubscribeRelationRequest, callback func(response *CreateConsumerGroupSubscribeRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateConsumerGroupSubscribeRelationResponse
		var err error
		defer close(result)
		response, err = client.CreateConsumerGroupSubscribeRelation(request)
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

// CreateConsumerGroupSubscribeRelationRequest is the request struct for api CreateConsumerGroupSubscribeRelation
type CreateConsumerGroupSubscribeRelationRequest struct {
	*requests.RpcRequest
	ConsumerGroupId string `position:"Query" name:"ConsumerGroupId"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	ProductKey      string `position:"Query" name:"ProductKey"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
}

// CreateConsumerGroupSubscribeRelationResponse is the response struct for api CreateConsumerGroupSubscribeRelation
type CreateConsumerGroupSubscribeRelationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateCreateConsumerGroupSubscribeRelationRequest creates a request to invoke CreateConsumerGroupSubscribeRelation API
func CreateCreateConsumerGroupSubscribeRelationRequest() (request *CreateConsumerGroupSubscribeRelationRequest) {
	request = &CreateConsumerGroupSubscribeRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateConsumerGroupSubscribeRelation", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateConsumerGroupSubscribeRelationResponse creates a response to parse from CreateConsumerGroupSubscribeRelation response
func CreateCreateConsumerGroupSubscribeRelationResponse() (response *CreateConsumerGroupSubscribeRelationResponse) {
	response = &CreateConsumerGroupSubscribeRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
