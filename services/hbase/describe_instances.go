package hbase

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

// DescribeInstances invokes the hbase.DescribeInstances API synchronously
// api document: https://help.aliyun.com/api/hbase/describeinstances.html
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	response = CreateDescribeInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstancesWithChan invokes the hbase.DescribeInstances API asynchronously
// api document: https://help.aliyun.com/api/hbase/describeinstances.html
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

// DescribeInstancesWithCallback invokes the hbase.DescribeInstances API asynchronously
// api document: https://help.aliyun.com/api/hbase/describeinstances.html
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
	ClusterName string                  `position:"Query" name:"ClusterName"`
	PageNumber  requests.Integer        `position:"Query" name:"PageNumber"`
	PageSize    requests.Integer        `position:"Query" name:"PageSize"`
	Tag         *[]DescribeInstancesTag `position:"Query" name:"Tag"  type:"Repeated"`
	DbType      string                  `position:"Query" name:"DbType"`
}

// DescribeInstancesTag is a repeated param struct in DescribeInstancesRequest
type DescribeInstancesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeInstancesResponse is the response struct for api DescribeInstances
type DescribeInstancesResponse struct {
	*responses.BaseResponse
	RequestId  string                       `json:"RequestId" xml:"RequestId"`
	TotalCount int64                        `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                          `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                          `json:"PageSize" xml:"PageSize"`
	Instances  InstancesInDescribeInstances `json:"Instances" xml:"Instances"`
}

// CreateDescribeInstancesRequest creates a request to invoke DescribeInstances API
func CreateDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeInstances", "hbase", "openAPI")
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
