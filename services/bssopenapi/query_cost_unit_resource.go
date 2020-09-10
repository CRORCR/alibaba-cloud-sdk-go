package bssopenapi

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

// QueryCostUnitResource invokes the bssopenapi.QueryCostUnitResource API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querycostunitresource.html
func (client *Client) QueryCostUnitResource(request *QueryCostUnitResourceRequest) (response *QueryCostUnitResourceResponse, err error) {
	response = CreateQueryCostUnitResourceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCostUnitResourceWithChan invokes the bssopenapi.QueryCostUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querycostunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCostUnitResourceWithChan(request *QueryCostUnitResourceRequest) (<-chan *QueryCostUnitResourceResponse, <-chan error) {
	responseChan := make(chan *QueryCostUnitResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCostUnitResource(request)
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

// QueryCostUnitResourceWithCallback invokes the bssopenapi.QueryCostUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querycostunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCostUnitResourceWithCallback(request *QueryCostUnitResourceRequest, callback func(response *QueryCostUnitResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCostUnitResourceResponse
		var err error
		defer close(result)
		response, err = client.QueryCostUnitResource(request)
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

// QueryCostUnitResourceRequest is the request struct for api QueryCostUnitResource
type QueryCostUnitResourceRequest struct {
	*requests.RpcRequest
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	UnitId   requests.Integer `position:"Query" name:"UnitId"`
	OwnerUid requests.Integer `position:"Query" name:"OwnerUid"`
}

// QueryCostUnitResourceResponse is the response struct for api QueryCostUnitResource
type QueryCostUnitResourceResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Success   bool                        `json:"Success" xml:"Success"`
	Code      string                      `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	Data      DataInQueryCostUnitResource `json:"Data" xml:"Data"`
}

// CreateQueryCostUnitResourceRequest creates a request to invoke QueryCostUnitResource API
func CreateQueryCostUnitResourceRequest() (request *QueryCostUnitResourceRequest) {
	request = &QueryCostUnitResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryCostUnitResource", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryCostUnitResourceResponse creates a response to parse from QueryCostUnitResource response
func CreateQueryCostUnitResourceResponse() (response *QueryCostUnitResourceResponse) {
	response = &QueryCostUnitResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
