package trademark

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

// ApplyNotaryPost invokes the trademark.ApplyNotaryPost API synchronously
// api document: https://help.aliyun.com/api/trademark/applynotarypost.html
func (client *Client) ApplyNotaryPost(request *ApplyNotaryPostRequest) (response *ApplyNotaryPostResponse, err error) {
	response = CreateApplyNotaryPostResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyNotaryPostWithChan invokes the trademark.ApplyNotaryPost API asynchronously
// api document: https://help.aliyun.com/api/trademark/applynotarypost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyNotaryPostWithChan(request *ApplyNotaryPostRequest) (<-chan *ApplyNotaryPostResponse, <-chan error) {
	responseChan := make(chan *ApplyNotaryPostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyNotaryPost(request)
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

// ApplyNotaryPostWithCallback invokes the trademark.ApplyNotaryPost API asynchronously
// api document: https://help.aliyun.com/api/trademark/applynotarypost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyNotaryPostWithCallback(request *ApplyNotaryPostRequest, callback func(response *ApplyNotaryPostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyNotaryPostResponse
		var err error
		defer close(result)
		response, err = client.ApplyNotaryPost(request)
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

// ApplyNotaryPostRequest is the request struct for api ApplyNotaryPost
type ApplyNotaryPostRequest struct {
	*requests.RpcRequest
	ReceiverName    string           `position:"Query" name:"ReceiverName"`
	ReceiverPhone   string           `position:"Query" name:"ReceiverPhone"`
	NotaryOrderId   requests.Integer `position:"Query" name:"NotaryOrderId"`
	ReceiverAddress string           `position:"Query" name:"ReceiverAddress"`
}

// ApplyNotaryPostResponse is the response struct for api ApplyNotaryPost
type ApplyNotaryPostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateApplyNotaryPostRequest creates a request to invoke ApplyNotaryPost API
func CreateApplyNotaryPostRequest() (request *ApplyNotaryPostRequest) {
	request = &ApplyNotaryPostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "ApplyNotaryPost", "trademark", "openAPI")
	return
}

// CreateApplyNotaryPostResponse creates a response to parse from ApplyNotaryPost response
func CreateApplyNotaryPostResponse() (response *ApplyNotaryPostResponse) {
	response = &ApplyNotaryPostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
