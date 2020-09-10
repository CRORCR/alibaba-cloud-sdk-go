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

// SetDefaultAITemplate invokes the vod.SetDefaultAITemplate API synchronously
func (client *Client) SetDefaultAITemplate(request *SetDefaultAITemplateRequest) (response *SetDefaultAITemplateResponse, err error) {
	response = CreateSetDefaultAITemplateResponse()
	err = client.DoAction(request, response)
	return
}

// SetDefaultAITemplateWithChan invokes the vod.SetDefaultAITemplate API asynchronously
func (client *Client) SetDefaultAITemplateWithChan(request *SetDefaultAITemplateRequest) (<-chan *SetDefaultAITemplateResponse, <-chan error) {
	responseChan := make(chan *SetDefaultAITemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDefaultAITemplate(request)
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

// SetDefaultAITemplateWithCallback invokes the vod.SetDefaultAITemplate API asynchronously
func (client *Client) SetDefaultAITemplateWithCallback(request *SetDefaultAITemplateRequest, callback func(response *SetDefaultAITemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDefaultAITemplateResponse
		var err error
		defer close(result)
		response, err = client.SetDefaultAITemplate(request)
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

// SetDefaultAITemplateRequest is the request struct for api SetDefaultAITemplate
type SetDefaultAITemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId           string           `position:"Query" name:"TemplateId"`
}

// SetDefaultAITemplateResponse is the response struct for api SetDefaultAITemplate
type SetDefaultAITemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateSetDefaultAITemplateRequest creates a request to invoke SetDefaultAITemplate API
func CreateSetDefaultAITemplateRequest() (request *SetDefaultAITemplateRequest) {
	request = &SetDefaultAITemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SetDefaultAITemplate", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDefaultAITemplateResponse creates a response to parse from SetDefaultAITemplate response
func CreateSetDefaultAITemplateResponse() (response *SetDefaultAITemplateResponse) {
	response = &SetDefaultAITemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
