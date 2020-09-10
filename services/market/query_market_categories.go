package market

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

// QueryMarketCategories invokes the market.QueryMarketCategories API synchronously
// api document: https://help.aliyun.com/api/market/querymarketcategories.html
func (client *Client) QueryMarketCategories(request *QueryMarketCategoriesRequest) (response *QueryMarketCategoriesResponse, err error) {
	response = CreateQueryMarketCategoriesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMarketCategoriesWithChan invokes the market.QueryMarketCategories API asynchronously
// api document: https://help.aliyun.com/api/market/querymarketcategories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMarketCategoriesWithChan(request *QueryMarketCategoriesRequest) (<-chan *QueryMarketCategoriesResponse, <-chan error) {
	responseChan := make(chan *QueryMarketCategoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMarketCategories(request)
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

// QueryMarketCategoriesWithCallback invokes the market.QueryMarketCategories API asynchronously
// api document: https://help.aliyun.com/api/market/querymarketcategories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMarketCategoriesWithCallback(request *QueryMarketCategoriesRequest, callback func(response *QueryMarketCategoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMarketCategoriesResponse
		var err error
		defer close(result)
		response, err = client.QueryMarketCategories(request)
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

// QueryMarketCategoriesRequest is the request struct for api QueryMarketCategories
type QueryMarketCategoriesRequest struct {
	*requests.RpcRequest
}

// QueryMarketCategoriesResponse is the response struct for api QueryMarketCategories
type QueryMarketCategoriesResponse struct {
	*responses.BaseResponse
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Categories Categories `json:"Categories" xml:"Categories"`
}

// CreateQueryMarketCategoriesRequest creates a request to invoke QueryMarketCategories API
func CreateQueryMarketCategoriesRequest() (request *QueryMarketCategoriesRequest) {
	request = &QueryMarketCategoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "QueryMarketCategories", "yunmarket", "openAPI")
	return
}

// CreateQueryMarketCategoriesResponse creates a response to parse from QueryMarketCategories response
func CreateQueryMarketCategoriesResponse() (response *QueryMarketCategoriesResponse) {
	response = &QueryMarketCategoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
