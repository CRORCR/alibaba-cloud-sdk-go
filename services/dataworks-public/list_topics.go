package dataworks_public

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

// ListTopics invokes the dataworks_public.ListTopics API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/listtopics.html
func (client *Client) ListTopics(request *ListTopicsRequest) (response *ListTopicsResponse, err error) {
	response = CreateListTopicsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTopicsWithChan invokes the dataworks_public.ListTopics API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listtopics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTopicsWithChan(request *ListTopicsRequest) (<-chan *ListTopicsResponse, <-chan error) {
	responseChan := make(chan *ListTopicsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTopics(request)
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

// ListTopicsWithCallback invokes the dataworks_public.ListTopics API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listtopics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTopicsWithCallback(request *ListTopicsRequest, callback func(response *ListTopicsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTopicsResponse
		var err error
		defer close(result)
		response, err = client.ListTopics(request)
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

// ListTopicsRequest is the request struct for api ListTopics
type ListTopicsRequest struct {
	*requests.RpcRequest
	Owner         string           `position:"Body" name:"Owner"`
	EndTime       string           `position:"Body" name:"EndTime"`
	BeginTime     string           `position:"Body" name:"BeginTime"`
	TopicStatuses string           `position:"Body" name:"TopicStatuses"`
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	InstanceId    requests.Integer `position:"Body" name:"InstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	TopicTypes    string           `position:"Body" name:"TopicTypes"`
	NodeId        requests.Integer `position:"Body" name:"NodeId"`
}

// ListTopicsResponse is the response struct for api ListTopics
type ListTopicsResponse struct {
	*responses.BaseResponse
	Success        bool             `json:"Success" xml:"Success"`
	ErrorCode      string           `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string           `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	Data           DataInListTopics `json:"Data" xml:"Data"`
}

// CreateListTopicsRequest creates a request to invoke ListTopics API
func CreateListTopicsRequest() (request *ListTopicsRequest) {
	request = &ListTopicsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListTopics", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTopicsResponse creates a response to parse from ListTopics response
func CreateListTopicsResponse() (response *ListTopicsResponse) {
	response = &ListTopicsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
