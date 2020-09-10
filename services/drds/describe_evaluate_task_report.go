package drds

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

// DescribeEvaluateTaskReport invokes the drds.DescribeEvaluateTaskReport API synchronously
// api document: https://help.aliyun.com/api/drds/describeevaluatetaskreport.html
func (client *Client) DescribeEvaluateTaskReport(request *DescribeEvaluateTaskReportRequest) (response *DescribeEvaluateTaskReportResponse, err error) {
	response = CreateDescribeEvaluateTaskReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEvaluateTaskReportWithChan invokes the drds.DescribeEvaluateTaskReport API asynchronously
// api document: https://help.aliyun.com/api/drds/describeevaluatetaskreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEvaluateTaskReportWithChan(request *DescribeEvaluateTaskReportRequest) (<-chan *DescribeEvaluateTaskReportResponse, <-chan error) {
	responseChan := make(chan *DescribeEvaluateTaskReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEvaluateTaskReport(request)
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

// DescribeEvaluateTaskReportWithCallback invokes the drds.DescribeEvaluateTaskReport API asynchronously
// api document: https://help.aliyun.com/api/drds/describeevaluatetaskreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEvaluateTaskReportWithCallback(request *DescribeEvaluateTaskReportRequest, callback func(response *DescribeEvaluateTaskReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEvaluateTaskReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeEvaluateTaskReport(request)
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

// DescribeEvaluateTaskReportRequest is the request struct for api DescribeEvaluateTaskReport
type DescribeEvaluateTaskReportRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeEvaluateTaskReportResponse is the response struct for api DescribeEvaluateTaskReport
type DescribeEvaluateTaskReportResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	EvaluateResult EvaluateResult `json:"EvaluateResult" xml:"EvaluateResult"`
}

// CreateDescribeEvaluateTaskReportRequest creates a request to invoke DescribeEvaluateTaskReport API
func CreateDescribeEvaluateTaskReportRequest() (request *DescribeEvaluateTaskReportRequest) {
	request = &DescribeEvaluateTaskReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeEvaluateTaskReport", "Drds", "openAPI")
	return
}

// CreateDescribeEvaluateTaskReportResponse creates a response to parse from DescribeEvaluateTaskReport response
func CreateDescribeEvaluateTaskReportResponse() (response *DescribeEvaluateTaskReportResponse) {
	response = &DescribeEvaluateTaskReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
