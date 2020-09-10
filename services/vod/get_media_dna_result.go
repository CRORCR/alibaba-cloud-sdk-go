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

// GetMediaDNAResult invokes the vod.GetMediaDNAResult API synchronously
func (client *Client) GetMediaDNAResult(request *GetMediaDNAResultRequest) (response *GetMediaDNAResultResponse, err error) {
	response = CreateGetMediaDNAResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetMediaDNAResultWithChan invokes the vod.GetMediaDNAResult API asynchronously
func (client *Client) GetMediaDNAResultWithChan(request *GetMediaDNAResultRequest) (<-chan *GetMediaDNAResultResponse, <-chan error) {
	responseChan := make(chan *GetMediaDNAResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMediaDNAResult(request)
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

// GetMediaDNAResultWithCallback invokes the vod.GetMediaDNAResult API asynchronously
func (client *Client) GetMediaDNAResultWithCallback(request *GetMediaDNAResultRequest, callback func(response *GetMediaDNAResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMediaDNAResultResponse
		var err error
		defer close(result)
		response, err = client.GetMediaDNAResult(request)
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

// GetMediaDNAResultRequest is the request struct for api GetMediaDNAResult
type GetMediaDNAResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

// GetMediaDNAResultResponse is the response struct for api GetMediaDNAResult
type GetMediaDNAResultResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	DNAResult DNAResult `json:"DNAResult" xml:"DNAResult"`
}

// CreateGetMediaDNAResultRequest creates a request to invoke GetMediaDNAResult API
func CreateGetMediaDNAResultRequest() (request *GetMediaDNAResultRequest) {
	request = &GetMediaDNAResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetMediaDNAResult", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMediaDNAResultResponse creates a response to parse from GetMediaDNAResult response
func CreateGetMediaDNAResultResponse() (response *GetMediaDNAResultResponse) {
	response = &GetMediaDNAResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
