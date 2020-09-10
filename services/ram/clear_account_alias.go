package ram

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

// ClearAccountAlias invokes the ram.ClearAccountAlias API synchronously
// api document: https://help.aliyun.com/api/ram/clearaccountalias.html
func (client *Client) ClearAccountAlias(request *ClearAccountAliasRequest) (response *ClearAccountAliasResponse, err error) {
	response = CreateClearAccountAliasResponse()
	err = client.DoAction(request, response)
	return
}

// ClearAccountAliasWithChan invokes the ram.ClearAccountAlias API asynchronously
// api document: https://help.aliyun.com/api/ram/clearaccountalias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ClearAccountAliasWithChan(request *ClearAccountAliasRequest) (<-chan *ClearAccountAliasResponse, <-chan error) {
	responseChan := make(chan *ClearAccountAliasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClearAccountAlias(request)
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

// ClearAccountAliasWithCallback invokes the ram.ClearAccountAlias API asynchronously
// api document: https://help.aliyun.com/api/ram/clearaccountalias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ClearAccountAliasWithCallback(request *ClearAccountAliasRequest, callback func(response *ClearAccountAliasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClearAccountAliasResponse
		var err error
		defer close(result)
		response, err = client.ClearAccountAlias(request)
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

// ClearAccountAliasRequest is the request struct for api ClearAccountAlias
type ClearAccountAliasRequest struct {
	*requests.RpcRequest
}

// ClearAccountAliasResponse is the response struct for api ClearAccountAlias
type ClearAccountAliasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateClearAccountAliasRequest creates a request to invoke ClearAccountAlias API
func CreateClearAccountAliasRequest() (request *ClearAccountAliasRequest) {
	request = &ClearAccountAliasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ClearAccountAlias", "Ram", "openAPI")
	return
}

// CreateClearAccountAliasResponse creates a response to parse from ClearAccountAlias response
func CreateClearAccountAliasResponse() (response *ClearAccountAliasResponse) {
	response = &ClearAccountAliasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
