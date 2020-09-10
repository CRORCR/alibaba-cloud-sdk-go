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

// DescribeMetricRuleTargets invokes the cms.DescribeMetricRuleTargets API synchronously
// api document: https://help.aliyun.com/api/cms/describemetricruletargets.html
func (client *Client) DescribeMetricRuleTargets(request *DescribeMetricRuleTargetsRequest) (response *DescribeMetricRuleTargetsResponse, err error) {
	response = CreateDescribeMetricRuleTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetricRuleTargetsWithChan invokes the cms.DescribeMetricRuleTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricruletargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricRuleTargetsWithChan(request *DescribeMetricRuleTargetsRequest) (<-chan *DescribeMetricRuleTargetsResponse, <-chan error) {
	responseChan := make(chan *DescribeMetricRuleTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetricRuleTargets(request)
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

// DescribeMetricRuleTargetsWithCallback invokes the cms.DescribeMetricRuleTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricruletargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricRuleTargetsWithCallback(request *DescribeMetricRuleTargetsRequest, callback func(response *DescribeMetricRuleTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetricRuleTargetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetricRuleTargets(request)
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

// DescribeMetricRuleTargetsRequest is the request struct for api DescribeMetricRuleTargets
type DescribeMetricRuleTargetsRequest struct {
	*requests.RpcRequest
	RuleId string `position:"Query" name:"RuleId"`
}

// DescribeMetricRuleTargetsResponse is the response struct for api DescribeMetricRuleTargets
type DescribeMetricRuleTargetsResponse struct {
	*responses.BaseResponse
	Code      string                             `json:"Code" xml:"Code"`
	Message   string                             `json:"Message" xml:"Message"`
	Success   bool                               `json:"Success" xml:"Success"`
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Targets   TargetsInDescribeMetricRuleTargets `json:"Targets" xml:"Targets"`
}

// CreateDescribeMetricRuleTargetsRequest creates a request to invoke DescribeMetricRuleTargets API
func CreateDescribeMetricRuleTargetsRequest() (request *DescribeMetricRuleTargetsRequest) {
	request = &DescribeMetricRuleTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMetricRuleTargets", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMetricRuleTargetsResponse creates a response to parse from DescribeMetricRuleTargets response
func CreateDescribeMetricRuleTargetsResponse() (response *DescribeMetricRuleTargetsResponse) {
	response = &DescribeMetricRuleTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
