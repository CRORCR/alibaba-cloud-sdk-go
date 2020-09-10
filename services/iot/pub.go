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

// Pub invokes the iot.Pub API synchronously
// api document: https://help.aliyun.com/api/iot/pub.html
func (client *Client) Pub(request *PubRequest) (response *PubResponse, err error) {
	response = CreatePubResponse()
	err = client.DoAction(request, response)
	return
}

// PubWithChan invokes the iot.Pub API asynchronously
// api document: https://help.aliyun.com/api/iot/pub.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PubWithChan(request *PubRequest) (<-chan *PubResponse, <-chan error) {
	responseChan := make(chan *PubResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Pub(request)
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

// PubWithCallback invokes the iot.Pub API asynchronously
// api document: https://help.aliyun.com/api/iot/pub.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PubWithCallback(request *PubRequest, callback func(response *PubResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PubResponse
		var err error
		defer close(result)
		response, err = client.Pub(request)
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

// PubRequest is the request struct for api Pub
type PubRequest struct {
	*requests.RpcRequest
	MessageContent string           `position:"Query" name:"MessageContent"`
	Qos            requests.Integer `position:"Query" name:"Qos"`
	IotInstanceId  string           `position:"Query" name:"IotInstanceId"`
	TopicFullName  string           `position:"Query" name:"TopicFullName"`
	ProductKey     string           `position:"Query" name:"ProductKey"`
	ApiProduct     string           `position:"Body" name:"ApiProduct"`
	ApiRevision    string           `position:"Body" name:"ApiRevision"`
}

// PubResponse is the response struct for api Pub
type PubResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	MessageId    string `json:"MessageId" xml:"MessageId"`
}

// CreatePubRequest creates a request to invoke Pub API
func CreatePubRequest() (request *PubRequest) {
	request = &PubRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "Pub", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePubResponse creates a response to parse from Pub response
func CreatePubResponse() (response *PubResponse) {
	response = &PubResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
