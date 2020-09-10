package ccc

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

// ListAgentEvents invokes the ccc.ListAgentEvents API synchronously
// api document: https://help.aliyun.com/api/ccc/listagentevents.html
func (client *Client) ListAgentEvents(request *ListAgentEventsRequest) (response *ListAgentEventsResponse, err error) {
	response = CreateListAgentEventsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAgentEventsWithChan invokes the ccc.ListAgentEvents API asynchronously
// api document: https://help.aliyun.com/api/ccc/listagentevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAgentEventsWithChan(request *ListAgentEventsRequest) (<-chan *ListAgentEventsResponse, <-chan error) {
	responseChan := make(chan *ListAgentEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAgentEvents(request)
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

// ListAgentEventsWithCallback invokes the ccc.ListAgentEvents API asynchronously
// api document: https://help.aliyun.com/api/ccc/listagentevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAgentEventsWithCallback(request *ListAgentEventsRequest, callback func(response *ListAgentEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAgentEventsResponse
		var err error
		defer close(result)
		response, err = client.ListAgentEvents(request)
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

// ListAgentEventsRequest is the request struct for api ListAgentEvents
type ListAgentEventsRequest struct {
	*requests.RpcRequest
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	StopTime   requests.Integer `position:"Query" name:"StopTime"`
	RamId      *[]string        `position:"Query" name:"RamId"  type:"Repeated"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Event      *[]string        `position:"Query" name:"Event"  type:"Repeated"`
}

// ListAgentEventsResponse is the response struct for api ListAgentEvents
type ListAgentEventsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	AgentEventList AgentEventList `json:"AgentEventList" xml:"AgentEventList"`
}

// CreateListAgentEventsRequest creates a request to invoke ListAgentEvents API
func CreateListAgentEventsRequest() (request *ListAgentEventsRequest) {
	request = &ListAgentEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListAgentEvents", "", "")
	return
}

// CreateListAgentEventsResponse creates a response to parse from ListAgentEvents response
func CreateListAgentEventsResponse() (response *ListAgentEventsResponse) {
	response = &ListAgentEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
