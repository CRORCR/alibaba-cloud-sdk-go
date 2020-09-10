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

// DescribeStrategyExecDetail invokes the sas.DescribeStrategyExecDetail API synchronously
// api document: https://help.aliyun.com/api/sas/describestrategyexecdetail.html
func (client *Client) DescribeStrategyExecDetail(request *DescribeStrategyExecDetailRequest) (response *DescribeStrategyExecDetailResponse, err error) {
	response = CreateDescribeStrategyExecDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStrategyExecDetailWithChan invokes the sas.DescribeStrategyExecDetail API asynchronously
// api document: https://help.aliyun.com/api/sas/describestrategyexecdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStrategyExecDetailWithChan(request *DescribeStrategyExecDetailRequest) (<-chan *DescribeStrategyExecDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeStrategyExecDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStrategyExecDetail(request)
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

// DescribeStrategyExecDetailWithCallback invokes the sas.DescribeStrategyExecDetail API asynchronously
// api document: https://help.aliyun.com/api/sas/describestrategyexecdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStrategyExecDetailWithCallback(request *DescribeStrategyExecDetailRequest, callback func(response *DescribeStrategyExecDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStrategyExecDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeStrategyExecDetail(request)
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

// DescribeStrategyExecDetailRequest is the request struct for api DescribeStrategyExecDetail
type DescribeStrategyExecDetailRequest struct {
	*requests.RpcRequest
	SourceIp   string           `position:"Query" name:"SourceIp"`
	StrategyId requests.Integer `position:"Query" name:"StrategyId"`
}

// DescribeStrategyExecDetailResponse is the response struct for api DescribeStrategyExecDetail
type DescribeStrategyExecDetailResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	StartTime      string      `json:"StartTime" xml:"StartTime"`
	EndTime        string      `json:"EndTime" xml:"EndTime"`
	Source         string      `json:"Source" xml:"Source"`
	Percent        string      `json:"Percent" xml:"Percent"`
	SuccessCount   int         `json:"SuccessCount" xml:"SuccessCount"`
	FailCount      int         `json:"FailCount" xml:"FailCount"`
	InProcessCount int         `json:"InProcessCount" xml:"InProcessCount"`
	FailedEcsList  []FailedEcs `json:"FailedEcsList" xml:"FailedEcsList"`
}

// CreateDescribeStrategyExecDetailRequest creates a request to invoke DescribeStrategyExecDetail API
func CreateDescribeStrategyExecDetailRequest() (request *DescribeStrategyExecDetailRequest) {
	request = &DescribeStrategyExecDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeStrategyExecDetail", "sas", "openAPI")
	return
}

// CreateDescribeStrategyExecDetailResponse creates a response to parse from DescribeStrategyExecDetail response
func CreateDescribeStrategyExecDetailResponse() (response *DescribeStrategyExecDetailResponse) {
	response = &DescribeStrategyExecDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
