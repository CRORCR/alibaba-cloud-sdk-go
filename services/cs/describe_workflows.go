package cs

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

// DescribeWorkflows invokes the cs.DescribeWorkflows API synchronously
// api document: https://help.aliyun.com/api/cs/describeworkflows.html
func (client *Client) DescribeWorkflows(request *DescribeWorkflowsRequest) (response *DescribeWorkflowsResponse, err error) {
	response = CreateDescribeWorkflowsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWorkflowsWithChan invokes the cs.DescribeWorkflows API asynchronously
// api document: https://help.aliyun.com/api/cs/describeworkflows.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWorkflowsWithChan(request *DescribeWorkflowsRequest) (<-chan *DescribeWorkflowsResponse, <-chan error) {
	responseChan := make(chan *DescribeWorkflowsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWorkflows(request)
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

// DescribeWorkflowsWithCallback invokes the cs.DescribeWorkflows API asynchronously
// api document: https://help.aliyun.com/api/cs/describeworkflows.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWorkflowsWithCallback(request *DescribeWorkflowsRequest, callback func(response *DescribeWorkflowsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWorkflowsResponse
		var err error
		defer close(result)
		response, err = client.DescribeWorkflows(request)
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

// DescribeWorkflowsRequest is the request struct for api DescribeWorkflows
type DescribeWorkflowsRequest struct {
	*requests.RoaRequest
}

// DescribeWorkflowsResponse is the response struct for api DescribeWorkflows
type DescribeWorkflowsResponse struct {
	*responses.BaseResponse
}

// CreateDescribeWorkflowsRequest creates a request to invoke DescribeWorkflows API
func CreateDescribeWorkflowsRequest() (request *DescribeWorkflowsRequest) {
	request = &DescribeWorkflowsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeWorkflows", "/gs/workflows", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeWorkflowsResponse creates a response to parse from DescribeWorkflows response
func CreateDescribeWorkflowsResponse() (response *DescribeWorkflowsResponse) {
	response = &DescribeWorkflowsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
