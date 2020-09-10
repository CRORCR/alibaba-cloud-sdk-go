package r_kvstore

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

// DescribeMonitorItems invokes the r_kvstore.DescribeMonitorItems API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describemonitoritems.html
func (client *Client) DescribeMonitorItems(request *DescribeMonitorItemsRequest) (response *DescribeMonitorItemsResponse, err error) {
	response = CreateDescribeMonitorItemsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitorItemsWithChan invokes the r_kvstore.DescribeMonitorItems API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describemonitoritems.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitorItemsWithChan(request *DescribeMonitorItemsRequest) (<-chan *DescribeMonitorItemsResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitorItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitorItems(request)
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

// DescribeMonitorItemsWithCallback invokes the r_kvstore.DescribeMonitorItems API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describemonitoritems.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitorItemsWithCallback(request *DescribeMonitorItemsRequest, callback func(response *DescribeMonitorItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitorItemsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitorItems(request)
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

// DescribeMonitorItemsRequest is the request struct for api DescribeMonitorItems
type DescribeMonitorItemsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
}

// DescribeMonitorItemsResponse is the response struct for api DescribeMonitorItems
type DescribeMonitorItemsResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	MonitorItems MonitorItems `json:"MonitorItems" xml:"MonitorItems"`
}

// CreateDescribeMonitorItemsRequest creates a request to invoke DescribeMonitorItems API
func CreateDescribeMonitorItemsRequest() (request *DescribeMonitorItemsRequest) {
	request = &DescribeMonitorItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonitorItems", "redisa", "openAPI")
	return
}

// CreateDescribeMonitorItemsResponse creates a response to parse from DescribeMonitorItems response
func CreateDescribeMonitorItemsResponse() (response *DescribeMonitorItemsResponse) {
	response = &DescribeMonitorItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
