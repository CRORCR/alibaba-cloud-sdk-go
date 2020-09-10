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

// RunExecutionPlan invokes the emr.RunExecutionPlan API synchronously
// api document: https://help.aliyun.com/api/emr/runexecutionplan.html
func (client *Client) RunExecutionPlan(request *RunExecutionPlanRequest) (response *RunExecutionPlanResponse, err error) {
	response = CreateRunExecutionPlanResponse()
	err = client.DoAction(request, response)
	return
}

// RunExecutionPlanWithChan invokes the emr.RunExecutionPlan API asynchronously
// api document: https://help.aliyun.com/api/emr/runexecutionplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RunExecutionPlanWithChan(request *RunExecutionPlanRequest) (<-chan *RunExecutionPlanResponse, <-chan error) {
	responseChan := make(chan *RunExecutionPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunExecutionPlan(request)
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

// RunExecutionPlanWithCallback invokes the emr.RunExecutionPlan API asynchronously
// api document: https://help.aliyun.com/api/emr/runexecutionplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RunExecutionPlanWithCallback(request *RunExecutionPlanRequest, callback func(response *RunExecutionPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunExecutionPlanResponse
		var err error
		defer close(result)
		response, err = client.RunExecutionPlan(request)
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

// RunExecutionPlanRequest is the request struct for api RunExecutionPlan
type RunExecutionPlanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Arguments       string           `position:"Query" name:"Arguments"`
	Id              string           `position:"Query" name:"Id"`
}

// RunExecutionPlanResponse is the response struct for api RunExecutionPlan
type RunExecutionPlanResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	ExecutionPlanInstanceId string `json:"ExecutionPlanInstanceId" xml:"ExecutionPlanInstanceId"`
}

// CreateRunExecutionPlanRequest creates a request to invoke RunExecutionPlan API
func CreateRunExecutionPlanRequest() (request *RunExecutionPlanRequest) {
	request = &RunExecutionPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "RunExecutionPlan", "emr", "openAPI")
	return
}

// CreateRunExecutionPlanResponse creates a response to parse from RunExecutionPlan response
func CreateRunExecutionPlanResponse() (response *RunExecutionPlanResponse) {
	response = &RunExecutionPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
