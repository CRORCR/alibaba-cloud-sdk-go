package iot

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

// QueryProductList invokes the iot.QueryProductList API synchronously
// api document: https://help.aliyun.com/api/iot/queryproductlist.html
func (client *Client) QueryProductList(request *QueryProductListRequest) (response *QueryProductListResponse, err error) {
	response = CreateQueryProductListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryProductListWithChan invokes the iot.QueryProductList API asynchronously
// api document: https://help.aliyun.com/api/iot/queryproductlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryProductListWithChan(request *QueryProductListRequest) (<-chan *QueryProductListResponse, <-chan error) {
	responseChan := make(chan *QueryProductListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryProductList(request)
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

// QueryProductListWithCallback invokes the iot.QueryProductList API asynchronously
// api document: https://help.aliyun.com/api/iot/queryproductlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryProductListWithCallback(request *QueryProductListRequest, callback func(response *QueryProductListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryProductListResponse
		var err error
		defer close(result)
		response, err = client.QueryProductList(request)
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

// QueryProductListRequest is the request struct for api QueryProductList
type QueryProductListRequest struct {
	*requests.RpcRequest
	ResourceGroupId     string           `position:"Query" name:"ResourceGroupId"`
	IotInstanceId       string           `position:"Query" name:"IotInstanceId"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	AliyunCommodityCode string           `position:"Query" name:"AliyunCommodityCode"`
	CurrentPage         requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
}

// QueryProductListResponse is the response struct for api QueryProductList
type QueryProductListResponse struct {
	*responses.BaseResponse
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Success      bool                   `json:"Success" xml:"Success"`
	Code         string                 `json:"Code" xml:"Code"`
	ErrorMessage string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryProductList `json:"Data" xml:"Data"`
}

// CreateQueryProductListRequest creates a request to invoke QueryProductList API
func CreateQueryProductListRequest() (request *QueryProductListRequest) {
	request = &QueryProductListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryProductList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryProductListResponse creates a response to parse from QueryProductList response
func CreateQueryProductListResponse() (response *QueryProductListResponse) {
	response = &QueryProductListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
