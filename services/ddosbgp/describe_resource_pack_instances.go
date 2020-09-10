package ddosbgp

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

// DescribeResourcePackInstances invokes the ddosbgp.DescribeResourcePackInstances API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeresourcepackinstances.html
func (client *Client) DescribeResourcePackInstances(request *DescribeResourcePackInstancesRequest) (response *DescribeResourcePackInstancesResponse, err error) {
	response = CreateDescribeResourcePackInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourcePackInstancesWithChan invokes the ddosbgp.DescribeResourcePackInstances API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeresourcepackinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourcePackInstancesWithChan(request *DescribeResourcePackInstancesRequest) (<-chan *DescribeResourcePackInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeResourcePackInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourcePackInstances(request)
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

// DescribeResourcePackInstancesWithCallback invokes the ddosbgp.DescribeResourcePackInstances API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeresourcepackinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeResourcePackInstancesWithCallback(request *DescribeResourcePackInstancesRequest, callback func(response *DescribeResourcePackInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourcePackInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourcePackInstances(request)
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

// DescribeResourcePackInstancesRequest is the request struct for api DescribeResourcePackInstances
type DescribeResourcePackInstancesRequest struct {
	*requests.RpcRequest
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeResourcePackInstancesResponse is the response struct for api DescribeResourcePackInstances
type DescribeResourcePackInstancesResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	TotalCount    int            `json:"TotalCount" xml:"TotalCount"`
	ResourcePacks []ResourcePack `json:"ResourcePacks" xml:"ResourcePacks"`
}

// CreateDescribeResourcePackInstancesRequest creates a request to invoke DescribeResourcePackInstances API
func CreateDescribeResourcePackInstancesRequest() (request *DescribeResourcePackInstancesRequest) {
	request = &DescribeResourcePackInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeResourcePackInstances", "ddosbgp", "openAPI")
	return
}

// CreateDescribeResourcePackInstancesResponse creates a response to parse from DescribeResourcePackInstances response
func CreateDescribeResourcePackInstancesResponse() (response *DescribeResourcePackInstancesResponse) {
	response = &DescribeResourcePackInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
