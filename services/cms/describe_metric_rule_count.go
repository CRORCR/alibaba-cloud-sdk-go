package cms

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

// DescribeMetricRuleCount invokes the cms.DescribeMetricRuleCount API synchronously
// api document: https://help.aliyun.com/api/cms/describemetricrulecount.html
func (client *Client) DescribeMetricRuleCount(request *DescribeMetricRuleCountRequest) (response *DescribeMetricRuleCountResponse, err error) {
	response = CreateDescribeMetricRuleCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetricRuleCountWithChan invokes the cms.DescribeMetricRuleCount API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricrulecount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricRuleCountWithChan(request *DescribeMetricRuleCountRequest) (<-chan *DescribeMetricRuleCountResponse, <-chan error) {
	responseChan := make(chan *DescribeMetricRuleCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetricRuleCount(request)
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

// DescribeMetricRuleCountWithCallback invokes the cms.DescribeMetricRuleCount API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricrulecount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricRuleCountWithCallback(request *DescribeMetricRuleCountRequest, callback func(response *DescribeMetricRuleCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetricRuleCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetricRuleCount(request)
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

// DescribeMetricRuleCountRequest is the request struct for api DescribeMetricRuleCount
type DescribeMetricRuleCountRequest struct {
	*requests.RpcRequest
	Namespace  string `position:"Query" name:"Namespace"`
	MetricName string `position:"Query" name:"MetricName"`
}

// DescribeMetricRuleCountResponse is the response struct for api DescribeMetricRuleCount
type DescribeMetricRuleCountResponse struct {
	*responses.BaseResponse
	Success         bool            `json:"Success" xml:"Success"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	MetricRuleCount MetricRuleCount `json:"MetricRuleCount" xml:"MetricRuleCount"`
}

// CreateDescribeMetricRuleCountRequest creates a request to invoke DescribeMetricRuleCount API
func CreateDescribeMetricRuleCountRequest() (request *DescribeMetricRuleCountRequest) {
	request = &DescribeMetricRuleCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMetricRuleCount", "cms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeMetricRuleCountResponse creates a response to parse from DescribeMetricRuleCount response
func CreateDescribeMetricRuleCountResponse() (response *DescribeMetricRuleCountResponse) {
	response = &DescribeMetricRuleCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
