package cloudwf

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

// ResetAp invokes the cloudwf.ResetAp API synchronously
// api document: https://help.aliyun.com/api/cloudwf/resetap.html
func (client *Client) ResetAp(request *ResetApRequest) (response *ResetApResponse, err error) {
	response = CreateResetApResponse()
	err = client.DoAction(request, response)
	return
}

// ResetApWithChan invokes the cloudwf.ResetAp API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/resetap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetApWithChan(request *ResetApRequest) (<-chan *ResetApResponse, <-chan error) {
	responseChan := make(chan *ResetApResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetAp(request)
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

// ResetApWithCallback invokes the cloudwf.ResetAp API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/resetap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetApWithCallback(request *ResetApRequest, callback func(response *ResetApResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetApResponse
		var err error
		defer close(result)
		response, err = client.ResetAp(request)
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

// ResetApRequest is the request struct for api ResetAp
type ResetApRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// ResetApResponse is the response struct for api ResetAp
type ResetApResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateResetApRequest creates a request to invoke ResetAp API
func CreateResetApRequest() (request *ResetApRequest) {
	request = &ResetApRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ResetAp", "cloudwf", "openAPI")
	return
}

// CreateResetApResponse creates a response to parse from ResetAp response
func CreateResetApResponse() (response *ResetApResponse) {
	response = &ResetApResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
