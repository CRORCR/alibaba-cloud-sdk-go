package alikafka

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

// GetMetaProductList invokes the alikafka.GetMetaProductList API synchronously
// api document: https://help.aliyun.com/api/alikafka/getmetaproductlist.html
func (client *Client) GetMetaProductList(request *GetMetaProductListRequest) (response *GetMetaProductListResponse, err error) {
	response = CreateGetMetaProductListResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaProductListWithChan invokes the alikafka.GetMetaProductList API asynchronously
// api document: https://help.aliyun.com/api/alikafka/getmetaproductlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMetaProductListWithChan(request *GetMetaProductListRequest) (<-chan *GetMetaProductListResponse, <-chan error) {
	responseChan := make(chan *GetMetaProductListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaProductList(request)
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

// GetMetaProductListWithCallback invokes the alikafka.GetMetaProductList API asynchronously
// api document: https://help.aliyun.com/api/alikafka/getmetaproductlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMetaProductListWithCallback(request *GetMetaProductListRequest, callback func(response *GetMetaProductListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaProductListResponse
		var err error
		defer close(result)
		response, err = client.GetMetaProductList(request)
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

// GetMetaProductListRequest is the request struct for api GetMetaProductList
type GetMetaProductListRequest struct {
	*requests.RpcRequest
	ListNormal string `position:"Query" name:"ListNormal"`
}

// GetMetaProductListResponse is the response struct for api GetMetaProductList
type GetMetaProductListResponse struct {
	*responses.BaseResponse
	Success   bool     `json:"Success" xml:"Success"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Code      int      `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	MetaData  MetaData `json:"MetaData" xml:"MetaData"`
}

// CreateGetMetaProductListRequest creates a request to invoke GetMetaProductList API
func CreateGetMetaProductListRequest() (request *GetMetaProductListRequest) {
	request = &GetMetaProductListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetMetaProductList", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMetaProductListResponse creates a response to parse from GetMetaProductList response
func CreateGetMetaProductListResponse() (response *GetMetaProductListResponse) {
	response = &GetMetaProductListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
