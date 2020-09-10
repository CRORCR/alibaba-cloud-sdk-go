package sas

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

// DescribePropertySoftwareItem invokes the sas.DescribePropertySoftwareItem API synchronously
// api document: https://help.aliyun.com/api/sas/describepropertysoftwareitem.html
func (client *Client) DescribePropertySoftwareItem(request *DescribePropertySoftwareItemRequest) (response *DescribePropertySoftwareItemResponse, err error) {
	response = CreateDescribePropertySoftwareItemResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertySoftwareItemWithChan invokes the sas.DescribePropertySoftwareItem API asynchronously
// api document: https://help.aliyun.com/api/sas/describepropertysoftwareitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePropertySoftwareItemWithChan(request *DescribePropertySoftwareItemRequest) (<-chan *DescribePropertySoftwareItemResponse, <-chan error) {
	responseChan := make(chan *DescribePropertySoftwareItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertySoftwareItem(request)
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

// DescribePropertySoftwareItemWithCallback invokes the sas.DescribePropertySoftwareItem API asynchronously
// api document: https://help.aliyun.com/api/sas/describepropertysoftwareitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePropertySoftwareItemWithCallback(request *DescribePropertySoftwareItemRequest, callback func(response *DescribePropertySoftwareItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertySoftwareItemResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertySoftwareItem(request)
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

// DescribePropertySoftwareItemRequest is the request struct for api DescribePropertySoftwareItem
type DescribePropertySoftwareItemRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ForceFlush  requests.Boolean `position:"Query" name:"ForceFlush"`
}

// DescribePropertySoftwareItemResponse is the response struct for api DescribePropertySoftwareItem
type DescribePropertySoftwareItemResponse struct {
	*responses.BaseResponse
	RequestId     string                 `json:"RequestId" xml:"RequestId"`
	PageInfo      PageInfo               `json:"PageInfo" xml:"PageInfo"`
	PropertyItems []PropertySoftwareItem `json:"PropertyItems" xml:"PropertyItems"`
}

// CreateDescribePropertySoftwareItemRequest creates a request to invoke DescribePropertySoftwareItem API
func CreateDescribePropertySoftwareItemRequest() (request *DescribePropertySoftwareItemRequest) {
	request = &DescribePropertySoftwareItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertySoftwareItem", "sas", "openAPI")
	return
}

// CreateDescribePropertySoftwareItemResponse creates a response to parse from DescribePropertySoftwareItem response
func CreateDescribePropertySoftwareItemResponse() (response *DescribePropertySoftwareItemResponse) {
	response = &DescribePropertySoftwareItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
