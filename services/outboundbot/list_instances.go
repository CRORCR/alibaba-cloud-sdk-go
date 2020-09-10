package outboundbot

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

// ListInstances invokes the outboundbot.ListInstances API synchronously
// api document: https://help.aliyun.com/api/outboundbot/listinstances.html
func (client *Client) ListInstances(request *ListInstancesRequest) (response *ListInstancesResponse, err error) {
	response = CreateListInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstancesWithChan invokes the outboundbot.ListInstances API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/listinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstancesWithChan(request *ListInstancesRequest) (<-chan *ListInstancesResponse, <-chan error) {
	responseChan := make(chan *ListInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstances(request)
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

// ListInstancesWithCallback invokes the outboundbot.ListInstances API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/listinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstancesWithCallback(request *ListInstancesRequest, callback func(response *ListInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListInstances(request)
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

// ListInstancesRequest is the request struct for api ListInstances
type ListInstancesRequest struct {
	*requests.RpcRequest
}

// ListInstancesResponse is the response struct for api ListInstances
type ListInstancesResponse struct {
	*responses.BaseResponse
	Code           string     `json:"Code" xml:"Code"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Instances      []Instance `json:"Instances" xml:"Instances"`
}

// CreateListInstancesRequest creates a request to invoke ListInstances API
func CreateListInstancesRequest() (request *ListInstancesRequest) {
	request = &ListInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListInstances", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstancesResponse creates a response to parse from ListInstances response
func CreateListInstancesResponse() (response *ListInstancesResponse) {
	response = &ListInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
