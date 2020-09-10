package retailcloud

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

// CloseDeployOrder invokes the retailcloud.CloseDeployOrder API synchronously
// api document: https://help.aliyun.com/api/retailcloud/closedeployorder.html
func (client *Client) CloseDeployOrder(request *CloseDeployOrderRequest) (response *CloseDeployOrderResponse, err error) {
	response = CreateCloseDeployOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CloseDeployOrderWithChan invokes the retailcloud.CloseDeployOrder API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/closedeployorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloseDeployOrderWithChan(request *CloseDeployOrderRequest) (<-chan *CloseDeployOrderResponse, <-chan error) {
	responseChan := make(chan *CloseDeployOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloseDeployOrder(request)
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

// CloseDeployOrderWithCallback invokes the retailcloud.CloseDeployOrder API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/closedeployorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloseDeployOrderWithCallback(request *CloseDeployOrderRequest, callback func(response *CloseDeployOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloseDeployOrderResponse
		var err error
		defer close(result)
		response, err = client.CloseDeployOrder(request)
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

// CloseDeployOrderRequest is the request struct for api CloseDeployOrder
type CloseDeployOrderRequest struct {
	*requests.RpcRequest
	DeployOrderId requests.Integer `position:"Query" name:"DeployOrderId"`
}

// CloseDeployOrderResponse is the response struct for api CloseDeployOrder
type CloseDeployOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCloseDeployOrderRequest creates a request to invoke CloseDeployOrder API
func CreateCloseDeployOrderRequest() (request *CloseDeployOrderRequest) {
	request = &CloseDeployOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "CloseDeployOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateCloseDeployOrderResponse creates a response to parse from CloseDeployOrder response
func CreateCloseDeployOrderResponse() (response *CloseDeployOrderResponse) {
	response = &CloseDeployOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
