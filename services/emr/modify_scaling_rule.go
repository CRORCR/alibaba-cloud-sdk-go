package emr

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

// ModifyScalingRule invokes the emr.ModifyScalingRule API synchronously
// api document: https://help.aliyun.com/api/emr/modifyscalingrule.html
func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (response *ModifyScalingRuleResponse, err error) {
	response = CreateModifyScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingRuleWithChan invokes the emr.ModifyScalingRule API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyscalingrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingRuleWithChan(request *ModifyScalingRuleRequest) (<-chan *ModifyScalingRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingRule(request)
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

// ModifyScalingRuleWithCallback invokes the emr.ModifyScalingRule API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyscalingrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingRuleWithCallback(request *ModifyScalingRuleRequest, callback func(response *ModifyScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingRule(request)
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

// ModifyScalingRuleRequest is the request struct for api ModifyScalingRule
type ModifyScalingRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                      `position:"Query" name:"ResourceOwnerId"`
	RuleName             string                                `position:"Query" name:"RuleName"`
	ScalingRuleId        string                                `position:"Query" name:"ScalingRuleId"`
	RecurrenceEndTime    string                                `position:"Query" name:"RecurrenceEndTime"`
	CloudWatchTrigger    *[]ModifyScalingRuleCloudWatchTrigger `position:"Query" name:"CloudWatchTrigger"  type:"Repeated"`
	Cooldown             requests.Integer                      `position:"Query" name:"Cooldown"`
	LaunchTime           string                                `position:"Query" name:"LaunchTime"`
	AdjustmentValue      requests.Integer                      `position:"Query" name:"AdjustmentValue"`
	AdjustmentType       string                                `position:"Query" name:"AdjustmentType"`
	ClusterId            string                                `position:"Query" name:"ClusterId"`
	LaunchExpirationTime requests.Integer                      `position:"Query" name:"LaunchExpirationTime"`
	RecurrenceValue      string                                `position:"Query" name:"RecurrenceValue"`
	HostGroupId          string                                `position:"Query" name:"HostGroupId"`
	SchedulerTrigger     *[]ModifyScalingRuleSchedulerTrigger  `position:"Query" name:"SchedulerTrigger"  type:"Repeated"`
	RecurrenceType       string                                `position:"Query" name:"RecurrenceType"`
}

// ModifyScalingRuleCloudWatchTrigger is a repeated param struct in ModifyScalingRuleRequest
type ModifyScalingRuleCloudWatchTrigger struct {
	Period             string `name:"Period"`
	EvaluationCount    string `name:"EvaluationCount"`
	Threshold          string `name:"Threshold"`
	MetricName         string `name:"MetricName"`
	ComparisonOperator string `name:"ComparisonOperator"`
	Statistics         string `name:"Statistics"`
}

// ModifyScalingRuleSchedulerTrigger is a repeated param struct in ModifyScalingRuleRequest
type ModifyScalingRuleSchedulerTrigger struct {
	LaunchTime           string `name:"LaunchTime"`
	LaunchExpirationTime string `name:"LaunchExpirationTime"`
	RecurrenceValue      string `name:"RecurrenceValue"`
	RecurrenceEndTime    string `name:"RecurrenceEndTime"`
	RecurrenceType       string `name:"RecurrenceType"`
}

// ModifyScalingRuleResponse is the response struct for api ModifyScalingRule
type ModifyScalingRuleResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ScalingRuleId string `json:"ScalingRuleId" xml:"ScalingRuleId"`
}

// CreateModifyScalingRuleRequest creates a request to invoke ModifyScalingRule API
func CreateModifyScalingRuleRequest() (request *ModifyScalingRuleRequest) {
	request = &ModifyScalingRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyScalingRule", "emr", "openAPI")
	return
}

// CreateModifyScalingRuleResponse creates a response to parse from ModifyScalingRule response
func CreateModifyScalingRuleResponse() (response *ModifyScalingRuleResponse) {
	response = &ModifyScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
