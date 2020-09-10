package rds

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

// DescribeAvailableDedicatedHostClasses invokes the rds.DescribeAvailableDedicatedHostClasses API synchronously
// api document: https://help.aliyun.com/api/rds/describeavailablededicatedhostclasses.html
func (client *Client) DescribeAvailableDedicatedHostClasses(request *DescribeAvailableDedicatedHostClassesRequest) (response *DescribeAvailableDedicatedHostClassesResponse, err error) {
	response = CreateDescribeAvailableDedicatedHostClassesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableDedicatedHostClassesWithChan invokes the rds.DescribeAvailableDedicatedHostClasses API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailablededicatedhostclasses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableDedicatedHostClassesWithChan(request *DescribeAvailableDedicatedHostClassesRequest) (<-chan *DescribeAvailableDedicatedHostClassesResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableDedicatedHostClassesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableDedicatedHostClasses(request)
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

// DescribeAvailableDedicatedHostClassesWithCallback invokes the rds.DescribeAvailableDedicatedHostClasses API asynchronously
// api document: https://help.aliyun.com/api/rds/describeavailablededicatedhostclasses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableDedicatedHostClassesWithCallback(request *DescribeAvailableDedicatedHostClassesRequest, callback func(response *DescribeAvailableDedicatedHostClassesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableDedicatedHostClassesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableDedicatedHostClasses(request)
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

// DescribeAvailableDedicatedHostClassesRequest is the request struct for api DescribeAvailableDedicatedHostClasses
type DescribeAvailableDedicatedHostClassesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	StorageType          string           `position:"Query" name:"StorageType"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeAvailableDedicatedHostClassesResponse is the response struct for api DescribeAvailableDedicatedHostClasses
type DescribeAvailableDedicatedHostClassesResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	HostClasses HostClasses `json:"HostClasses" xml:"HostClasses"`
}

// CreateDescribeAvailableDedicatedHostClassesRequest creates a request to invoke DescribeAvailableDedicatedHostClasses API
func CreateDescribeAvailableDedicatedHostClassesRequest() (request *DescribeAvailableDedicatedHostClassesRequest) {
	request = &DescribeAvailableDedicatedHostClassesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAvailableDedicatedHostClasses", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAvailableDedicatedHostClassesResponse creates a response to parse from DescribeAvailableDedicatedHostClasses response
func CreateDescribeAvailableDedicatedHostClassesResponse() (response *DescribeAvailableDedicatedHostClassesResponse) {
	response = &DescribeAvailableDedicatedHostClassesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
