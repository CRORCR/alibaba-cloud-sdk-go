package workorder

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

// ListProducts invokes the workorder.ListProducts API synchronously
// api document: https://help.aliyun.com/api/workorder/listproducts.html
func (client *Client) ListProducts(request *ListProductsRequest) (response *ListProductsResponse, err error) {
	response = CreateListProductsResponse()
	err = client.DoAction(request, response)
	return
}

// ListProductsWithChan invokes the workorder.ListProducts API asynchronously
// api document: https://help.aliyun.com/api/workorder/listproducts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListProductsWithChan(request *ListProductsRequest) (<-chan *ListProductsResponse, <-chan error) {
	responseChan := make(chan *ListProductsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProducts(request)
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

// ListProductsWithCallback invokes the workorder.ListProducts API asynchronously
// api document: https://help.aliyun.com/api/workorder/listproducts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListProductsWithCallback(request *ListProductsRequest, callback func(response *ListProductsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProductsResponse
		var err error
		defer close(result)
		response, err = client.ListProducts(request)
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

// ListProductsRequest is the request struct for api ListProducts
type ListProductsRequest struct {
	*requests.RpcRequest
	Language string `position:"Query" name:"Language"`
}

// ListProductsResponse is the response struct for api ListProducts
type ListProductsResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListProductsRequest creates a request to invoke ListProducts API
func CreateListProductsRequest() (request *ListProductsRequest) {
	request = &ListProductsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Workorder", "2020-03-26", "ListProducts", "workorder", "openAPI")
	return
}

// CreateListProductsResponse creates a response to parse from ListProducts response
func CreateListProductsResponse() (response *ListProductsResponse) {
	response = &ListProductsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
