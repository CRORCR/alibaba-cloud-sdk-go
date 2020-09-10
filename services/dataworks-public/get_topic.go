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

// GetTopic invokes the dataworks_public.GetTopic API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/gettopic.html
func (client *Client) GetTopic(request *GetTopicRequest) (response *GetTopicResponse, err error) {
	response = CreateGetTopicResponse()
	err = client.DoAction(request, response)
	return
}

// GetTopicWithChan invokes the dataworks_public.GetTopic API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/gettopic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTopicWithChan(request *GetTopicRequest) (<-chan *GetTopicResponse, <-chan error) {
	responseChan := make(chan *GetTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTopic(request)
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

// GetTopicWithCallback invokes the dataworks_public.GetTopic API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/gettopic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTopicWithCallback(request *GetTopicRequest, callback func(response *GetTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTopicResponse
		var err error
		defer close(result)
		response, err = client.GetTopic(request)
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

// GetTopicRequest is the request struct for api GetTopic
type GetTopicRequest struct {
	*requests.RpcRequest
	TopicId requests.Integer `position:"Body" name:"TopicId"`
}

// GetTopicResponse is the response struct for api GetTopic
type GetTopicResponse struct {
	*responses.BaseResponse
	Success        bool           `json:"Success" xml:"Success"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Data           DataInGetTopic `json:"Data" xml:"Data"`
}

// CreateGetTopicRequest creates a request to invoke GetTopic API
func CreateGetTopicRequest() (request *GetTopicRequest) {
	request = &GetTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetTopic", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTopicResponse creates a response to parse from GetTopic response
func CreateGetTopicResponse() (response *GetTopicResponse) {
	response = &GetTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
