package outboundbot

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

// ListTags invokes the outboundbot.ListTags API synchronously
// api document: https://help.aliyun.com/api/outboundbot/listtags.html
func (client *Client) ListTags(request *ListTagsRequest) (response *ListTagsResponse, err error) {
	response = CreateListTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagsWithChan invokes the outboundbot.ListTags API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/listtags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagsWithChan(request *ListTagsRequest) (<-chan *ListTagsResponse, <-chan error) {
	responseChan := make(chan *ListTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTags(request)
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

// ListTagsWithCallback invokes the outboundbot.ListTags API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/listtags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagsWithCallback(request *ListTagsRequest, callback func(response *ListTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagsResponse
		var err error
		defer close(result)
		response, err = client.ListTags(request)
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

// ListTagsRequest is the request struct for api ListTags
type ListTagsRequest struct {
	*requests.RpcRequest
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListTagsResponse is the response struct for api ListTags
type ListTagsResponse struct {
	*responses.BaseResponse
	Code           string     `json:"Code" xml:"Code"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	TagGroups      []TagGroup `json:"TagGroups" xml:"TagGroups"`
	Tags           []Tag      `json:"Tags" xml:"Tags"`
}

// CreateListTagsRequest creates a request to invoke ListTags API
func CreateListTagsRequest() (request *ListTagsRequest) {
	request = &ListTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListTags", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTagsResponse creates a response to parse from ListTags response
func CreateListTagsResponse() (response *ListTagsResponse) {
	response = &ListTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
