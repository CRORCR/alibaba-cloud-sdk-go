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

// GetInstanceByShop invokes the cloudwf.GetInstanceByShop API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getinstancebyshop.html
func (client *Client) GetInstanceByShop(request *GetInstanceByShopRequest) (response *GetInstanceByShopResponse, err error) {
	response = CreateGetInstanceByShopResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceByShopWithChan invokes the cloudwf.GetInstanceByShop API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getinstancebyshop.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceByShopWithChan(request *GetInstanceByShopRequest) (<-chan *GetInstanceByShopResponse, <-chan error) {
	responseChan := make(chan *GetInstanceByShopResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceByShop(request)
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

// GetInstanceByShopWithCallback invokes the cloudwf.GetInstanceByShop API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getinstancebyshop.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceByShopWithCallback(request *GetInstanceByShopRequest, callback func(response *GetInstanceByShopResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceByShopResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceByShop(request)
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

// GetInstanceByShopRequest is the request struct for api GetInstanceByShop
type GetInstanceByShopRequest struct {
	*requests.RpcRequest
	ShopId requests.Integer `position:"Query" name:"ShopId"`
}

// GetInstanceByShopResponse is the response struct for api GetInstanceByShop
type GetInstanceByShopResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetInstanceByShopRequest creates a request to invoke GetInstanceByShop API
func CreateGetInstanceByShopRequest() (request *GetInstanceByShopRequest) {
	request = &GetInstanceByShopRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetInstanceByShop", "cloudwf", "openAPI")
	return
}

// CreateGetInstanceByShopResponse creates a response to parse from GetInstanceByShop response
func CreateGetInstanceByShopResponse() (response *GetInstanceByShopResponse) {
	response = &GetInstanceByShopResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
