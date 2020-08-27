package ens

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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeInstances invokes the ens.DescribeInstances API synchronously
// api document: https://help.aliyun.com/api/ens/describeinstances.html
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = CreateDescribeInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstancesWithChan invokes the ens.DescribeInstances API asynchronously
// api document: https://help.aliyun.com/api/ens/describeinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstancesWithChan(request *DescribeInstancesRequest) (<-chan *DescribeInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstances(request)
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

// DescribeInstancesWithCallback invokes the ens.DescribeInstances API asynchronously
// api document: https://help.aliyun.com/api/ens/describeinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstancesWithCallback(request *DescribeInstancesRequest, callback func(response *DescribeInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstances(request)
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

// DescribeInstancesRequest is the request struct for api DescribeInstances
type DescribeInstancesRequest struct {
	*requests.RpcRequest
	ImageId              string           `position:"Query" name:"ImageId"`
	SearchKey            string           `position:"Query" name:"SearchKey"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OrderByParams        string           `position:"Query" name:"OrderByParams"`
	EnsRegionId          string           `position:"Query" name:"EnsRegionId"`
	PageSize             string           `position:"Query" name:"PageSize"`
	EnsRegionIds         string           `position:"Query" name:"EnsRegionIds"`
	InstanceResourceType string           `position:"Query" name:"InstanceResourceType"`
	EnsServiceId         string           `position:"Query" name:"EnsServiceId"`
	Version              string           `position:"Query" name:"Version"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	InstanceName         string           `position:"Query" name:"InstanceName"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeInstancesResponse is the response struct for api DescribeInstances
type DescribeInstancesResponse struct {
	*responses.BaseResponse
	RequestId  string                       `json:"RequestId" xml:"RequestId"`
	Code       int                          `json:"Code" xml:"Code"`
	TotalCount int                          `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                          `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                          `json:"PageSize" xml:"PageSize"`
	Instances  InstancesInDescribeInstances `json:"Instances" xml:"Instances"`
}

// CreateDescribeInstancesRequest creates a request to invoke DescribeInstances API
func CreateDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeInstances", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstancesResponse creates a response to parse from DescribeInstances response
func CreateDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
