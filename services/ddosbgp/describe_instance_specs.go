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

// DescribeInstanceSpecs invokes the ddosbgp.DescribeInstanceSpecs API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeinstancespecs.html
func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
	response = CreateDescribeInstanceSpecsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSpecsWithChan invokes the ddosbgp.DescribeInstanceSpecs API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeinstancespecs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSpecsWithChan(request *DescribeInstanceSpecsRequest) (<-chan *DescribeInstanceSpecsResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSpecsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSpecs(request)
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

// DescribeInstanceSpecsWithCallback invokes the ddosbgp.DescribeInstanceSpecs API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeinstancespecs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSpecsWithCallback(request *DescribeInstanceSpecsRequest, callback func(response *DescribeInstanceSpecsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSpecsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSpecs(request)
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

// DescribeInstanceSpecsRequest is the request struct for api DescribeInstanceSpecs
type DescribeInstanceSpecsRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	InstanceIdList  string `position:"Query" name:"InstanceIdList"`
	DdosRegionId    string `position:"Query" name:"DdosRegionId"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeInstanceSpecsResponse is the response struct for api DescribeInstanceSpecs
type DescribeInstanceSpecsResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	InstanceSpecs []InstanceSpec `json:"InstanceSpecs" xml:"InstanceSpecs"`
}

// CreateDescribeInstanceSpecsRequest creates a request to invoke DescribeInstanceSpecs API
func CreateDescribeInstanceSpecsRequest() (request *DescribeInstanceSpecsRequest) {
	request = &DescribeInstanceSpecsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeInstanceSpecs", "ddosbgp", "openAPI")
	return
}

// CreateDescribeInstanceSpecsResponse creates a response to parse from DescribeInstanceSpecs response
func CreateDescribeInstanceSpecsResponse() (response *DescribeInstanceSpecsResponse) {
	response = &DescribeInstanceSpecsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
