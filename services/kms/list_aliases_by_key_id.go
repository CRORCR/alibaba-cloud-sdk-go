package kms

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

// ListAliasesByKeyId invokes the kms.ListAliasesByKeyId API synchronously
// api document: https://help.aliyun.com/api/kms/listaliasesbykeyid.html
func (client *Client) ListAliasesByKeyId(request *ListAliasesByKeyIdRequest) (response *ListAliasesByKeyIdResponse, err error) {
	response = CreateListAliasesByKeyIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListAliasesByKeyIdWithChan invokes the kms.ListAliasesByKeyId API asynchronously
// api document: https://help.aliyun.com/api/kms/listaliasesbykeyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAliasesByKeyIdWithChan(request *ListAliasesByKeyIdRequest) (<-chan *ListAliasesByKeyIdResponse, <-chan error) {
	responseChan := make(chan *ListAliasesByKeyIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAliasesByKeyId(request)
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

// ListAliasesByKeyIdWithCallback invokes the kms.ListAliasesByKeyId API asynchronously
// api document: https://help.aliyun.com/api/kms/listaliasesbykeyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAliasesByKeyIdWithCallback(request *ListAliasesByKeyIdRequest, callback func(response *ListAliasesByKeyIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAliasesByKeyIdResponse
		var err error
		defer close(result)
		response, err = client.ListAliasesByKeyId(request)
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

// ListAliasesByKeyIdRequest is the request struct for api ListAliasesByKeyId
type ListAliasesByKeyIdRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	KeyId      string           `position:"Query" name:"KeyId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListAliasesByKeyIdResponse is the response struct for api ListAliasesByKeyId
type ListAliasesByKeyIdResponse struct {
	*responses.BaseResponse
	TotalCount int                         `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                         `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                         `json:"PageSize" xml:"PageSize"`
	RequestId  string                      `json:"RequestId" xml:"RequestId"`
	Aliases    AliasesInListAliasesByKeyId `json:"Aliases" xml:"Aliases"`
}

// CreateListAliasesByKeyIdRequest creates a request to invoke ListAliasesByKeyId API
func CreateListAliasesByKeyIdRequest() (request *ListAliasesByKeyIdRequest) {
	request = &ListAliasesByKeyIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListAliasesByKeyId", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAliasesByKeyIdResponse creates a response to parse from ListAliasesByKeyId response
func CreateListAliasesByKeyIdResponse() (response *ListAliasesByKeyIdResponse) {
	response = &ListAliasesByKeyIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
