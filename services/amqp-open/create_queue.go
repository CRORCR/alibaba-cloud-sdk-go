package amqp_open

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

// CreateQueue invokes the amqp_open.CreateQueue API synchronously
// api document: https://help.aliyun.com/api/amqp-open/createqueue.html
func (client *Client) CreateQueue(request *CreateQueueRequest) (response *CreateQueueResponse, err error) {
	response = CreateCreateQueueResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQueueWithChan invokes the amqp_open.CreateQueue API asynchronously
// api document: https://help.aliyun.com/api/amqp-open/createqueue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateQueueWithChan(request *CreateQueueRequest) (<-chan *CreateQueueResponse, <-chan error) {
	responseChan := make(chan *CreateQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQueue(request)
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

// CreateQueueWithCallback invokes the amqp_open.CreateQueue API asynchronously
// api document: https://help.aliyun.com/api/amqp-open/createqueue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateQueueWithCallback(request *CreateQueueRequest, callback func(response *CreateQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQueueResponse
		var err error
		defer close(result)
		response, err = client.CreateQueue(request)
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

// CreateQueueRequest is the request struct for api CreateQueue
type CreateQueueRequest struct {
	*requests.RpcRequest
	QueueName            string           `position:"Body" name:"QueueName"`
	DeadLetterRoutingKey string           `position:"Body" name:"DeadLetterRoutingKey"`
	MaxLength            requests.Integer `position:"Body" name:"MaxLength"`
	AutoExpireState      requests.Integer `position:"Body" name:"AutoExpireState"`
	DeadLetterExchange   string           `position:"Body" name:"DeadLetterExchange"`
	InstanceId           string           `position:"Body" name:"InstanceId"`
	ExclusiveState       requests.Boolean `position:"Body" name:"ExclusiveState"`
	AutoDeleteState      requests.Boolean `position:"Body" name:"AutoDeleteState"`
	MessageTTL           requests.Integer `position:"Body" name:"MessageTTL"`
	VirtualHost          string           `position:"Body" name:"VirtualHost"`
	MaximumPriority      requests.Integer `position:"Body" name:"MaximumPriority"`
}

// CreateQueueResponse is the response struct for api CreateQueue
type CreateQueueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateQueueRequest creates a request to invoke CreateQueue API
func CreateCreateQueueRequest() (request *CreateQueueRequest) {
	request = &CreateQueueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("amqp-open", "2019-12-12", "CreateQueue", "onsproxy", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateQueueResponse creates a response to parse from CreateQueue response
func CreateCreateQueueResponse() (response *CreateQueueResponse) {
	response = &CreateQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
