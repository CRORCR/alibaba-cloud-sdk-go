package vod

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

// ListAITemplate invokes the vod.ListAITemplate API synchronously
func (client *Client) ListAITemplate(request *ListAITemplateRequest) (response *ListAITemplateResponse, err error) {
	response = CreateListAITemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ListAITemplateWithChan invokes the vod.ListAITemplate API asynchronously
func (client *Client) ListAITemplateWithChan(request *ListAITemplateRequest) (<-chan *ListAITemplateResponse, <-chan error) {
	responseChan := make(chan *ListAITemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAITemplate(request)
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

// ListAITemplateWithCallback invokes the vod.ListAITemplate API asynchronously
func (client *Client) ListAITemplateWithCallback(request *ListAITemplateRequest, callback func(response *ListAITemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAITemplateResponse
		var err error
		defer close(result)
		response, err = client.ListAITemplate(request)
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

// ListAITemplateRequest is the request struct for api ListAITemplate
type ListAITemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateType         string           `position:"Query" name:"TemplateType"`
}

// ListAITemplateResponse is the response struct for api ListAITemplate
type ListAITemplateResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	TemplateInfoList []TemplateInfoListItem `json:"TemplateInfoList" xml:"TemplateInfoList"`
}

// CreateListAITemplateRequest creates a request to invoke ListAITemplate API
func CreateListAITemplateRequest() (request *ListAITemplateRequest) {
	request = &ListAITemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListAITemplate", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAITemplateResponse creates a response to parse from ListAITemplate response
func CreateListAITemplateResponse() (response *ListAITemplateResponse) {
	response = &ListAITemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
