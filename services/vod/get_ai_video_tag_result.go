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

// GetAIVideoTagResult invokes the vod.GetAIVideoTagResult API synchronously
func (client *Client) GetAIVideoTagResult(request *GetAIVideoTagResultRequest) (response *GetAIVideoTagResultResponse, err error) {
	response = CreateGetAIVideoTagResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetAIVideoTagResultWithChan invokes the vod.GetAIVideoTagResult API asynchronously
func (client *Client) GetAIVideoTagResultWithChan(request *GetAIVideoTagResultRequest) (<-chan *GetAIVideoTagResultResponse, <-chan error) {
	responseChan := make(chan *GetAIVideoTagResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAIVideoTagResult(request)
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

// GetAIVideoTagResultWithCallback invokes the vod.GetAIVideoTagResult API asynchronously
func (client *Client) GetAIVideoTagResultWithCallback(request *GetAIVideoTagResultRequest, callback func(response *GetAIVideoTagResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAIVideoTagResultResponse
		var err error
		defer close(result)
		response, err = client.GetAIVideoTagResult(request)
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

// GetAIVideoTagResultRequest is the request struct for api GetAIVideoTagResult
type GetAIVideoTagResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

// GetAIVideoTagResultResponse is the response struct for api GetAIVideoTagResult
type GetAIVideoTagResultResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	VideoTagResult VideoTagResult `json:"VideoTagResult" xml:"VideoTagResult"`
}

// CreateGetAIVideoTagResultRequest creates a request to invoke GetAIVideoTagResult API
func CreateGetAIVideoTagResultRequest() (request *GetAIVideoTagResultRequest) {
	request = &GetAIVideoTagResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAIVideoTagResult", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAIVideoTagResultResponse creates a response to parse from GetAIVideoTagResult response
func CreateGetAIVideoTagResultResponse() (response *GetAIVideoTagResultResponse) {
	response = &GetAIVideoTagResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
