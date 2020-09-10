package alidns

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

// DescribeGtmAccessStrategies invokes the alidns.DescribeGtmAccessStrategies API synchronously
// api document: https://help.aliyun.com/api/alidns/describegtmaccessstrategies.html
func (client *Client) DescribeGtmAccessStrategies(request *DescribeGtmAccessStrategiesRequest) (response *DescribeGtmAccessStrategiesResponse, err error) {
	response = CreateDescribeGtmAccessStrategiesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmAccessStrategiesWithChan invokes the alidns.DescribeGtmAccessStrategies API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtmaccessstrategies.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmAccessStrategiesWithChan(request *DescribeGtmAccessStrategiesRequest) (<-chan *DescribeGtmAccessStrategiesResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmAccessStrategiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmAccessStrategies(request)
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

// DescribeGtmAccessStrategiesWithCallback invokes the alidns.DescribeGtmAccessStrategies API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtmaccessstrategies.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmAccessStrategiesWithCallback(request *DescribeGtmAccessStrategiesRequest, callback func(response *DescribeGtmAccessStrategiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmAccessStrategiesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmAccessStrategies(request)
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

// DescribeGtmAccessStrategiesRequest is the request struct for api DescribeGtmAccessStrategies
type DescribeGtmAccessStrategiesRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DescribeGtmAccessStrategiesResponse is the response struct for api DescribeGtmAccessStrategies
type DescribeGtmAccessStrategiesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalItems int        `json:"TotalItems" xml:"TotalItems"`
	TotalPages int        `json:"TotalPages" xml:"TotalPages"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	Strategies Strategies `json:"Strategies" xml:"Strategies"`
}

// CreateDescribeGtmAccessStrategiesRequest creates a request to invoke DescribeGtmAccessStrategies API
func CreateDescribeGtmAccessStrategiesRequest() (request *DescribeGtmAccessStrategiesRequest) {
	request = &DescribeGtmAccessStrategiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmAccessStrategies", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGtmAccessStrategiesResponse creates a response to parse from DescribeGtmAccessStrategies response
func CreateDescribeGtmAccessStrategiesResponse() (response *DescribeGtmAccessStrategiesResponse) {
	response = &DescribeGtmAccessStrategiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
