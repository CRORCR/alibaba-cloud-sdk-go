package vcs

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

// ListCorps invokes the vcs.ListCorps API synchronously
// api document: https://help.aliyun.com/api/vcs/listcorps.html
func (client *Client) ListCorps(request *ListCorpsRequest) (response *ListCorpsResponse, err error) {
	response = CreateListCorpsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCorpsWithChan invokes the vcs.ListCorps API asynchronously
// api document: https://help.aliyun.com/api/vcs/listcorps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCorpsWithChan(request *ListCorpsRequest) (<-chan *ListCorpsResponse, <-chan error) {
	responseChan := make(chan *ListCorpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCorps(request)
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

// ListCorpsWithCallback invokes the vcs.ListCorps API asynchronously
// api document: https://help.aliyun.com/api/vcs/listcorps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCorpsWithCallback(request *ListCorpsRequest, callback func(response *ListCorpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCorpsResponse
		var err error
		defer close(result)
		response, err = client.ListCorps(request)
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

// ListCorpsRequest is the request struct for api ListCorps
type ListCorpsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
}

// ListCorpsResponse is the response struct for api ListCorps
type ListCorpsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListCorpsRequest creates a request to invoke ListCorps API
func CreateListCorpsRequest() (request *ListCorpsRequest) {
	request = &ListCorpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListCorps", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCorpsResponse creates a response to parse from ListCorps response
func CreateListCorpsResponse() (response *ListCorpsResponse) {
	response = &ListCorpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
