package mse

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

// ListEurekaInstances invokes the mse.ListEurekaInstances API synchronously
// api document: https://help.aliyun.com/api/mse/listeurekainstances.html
func (client *Client) ListEurekaInstances(request *ListEurekaInstancesRequest) (response *ListEurekaInstancesResponse, err error) {
	response = CreateListEurekaInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListEurekaInstancesWithChan invokes the mse.ListEurekaInstances API asynchronously
// api document: https://help.aliyun.com/api/mse/listeurekainstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEurekaInstancesWithChan(request *ListEurekaInstancesRequest) (<-chan *ListEurekaInstancesResponse, <-chan error) {
	responseChan := make(chan *ListEurekaInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEurekaInstances(request)
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

// ListEurekaInstancesWithCallback invokes the mse.ListEurekaInstances API asynchronously
// api document: https://help.aliyun.com/api/mse/listeurekainstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEurekaInstancesWithCallback(request *ListEurekaInstancesRequest, callback func(response *ListEurekaInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEurekaInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListEurekaInstances(request)
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

// ListEurekaInstancesRequest is the request struct for api ListEurekaInstances
type ListEurekaInstancesRequest struct {
	*requests.RpcRequest
	ClusterId   string           `position:"Query" name:"ClusterId"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	RequestPars string           `position:"Query" name:"RequestPars"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ServiceName string           `position:"Query" name:"ServiceName"`
}

// ListEurekaInstancesResponse is the response struct for api ListEurekaInstances
type ListEurekaInstancesResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Success    bool             `json:"Success" xml:"Success"`
	Message    string           `json:"Message" xml:"Message"`
	ErrorCode  string           `json:"ErrorCode" xml:"ErrorCode"`
	PageNumber int              `json:"PageNumber" xml:"PageNumber"`
	PageSize   int              `json:"PageSize" xml:"PageSize"`
	TotalCount int              `json:"TotalCount" xml:"TotalCount"`
	HttpCode   string           `json:"HttpCode" xml:"HttpCode"`
	Data       []EurekaInstance `json:"Data" xml:"Data"`
}

// CreateListEurekaInstancesRequest creates a request to invoke ListEurekaInstances API
func CreateListEurekaInstancesRequest() (request *ListEurekaInstancesRequest) {
	request = &ListEurekaInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListEurekaInstances", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListEurekaInstancesResponse creates a response to parse from ListEurekaInstances response
func CreateListEurekaInstancesResponse() (response *ListEurekaInstancesResponse) {
	response = &ListEurekaInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
