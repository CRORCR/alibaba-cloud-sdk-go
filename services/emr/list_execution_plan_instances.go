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

// ListExecutionPlanInstances invokes the emr.ListExecutionPlanInstances API synchronously
// api document: https://help.aliyun.com/api/emr/listexecutionplaninstances.html
func (client *Client) ListExecutionPlanInstances(request *ListExecutionPlanInstancesRequest) (response *ListExecutionPlanInstancesResponse, err error) {
	response = CreateListExecutionPlanInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListExecutionPlanInstancesWithChan invokes the emr.ListExecutionPlanInstances API asynchronously
// api document: https://help.aliyun.com/api/emr/listexecutionplaninstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListExecutionPlanInstancesWithChan(request *ListExecutionPlanInstancesRequest) (<-chan *ListExecutionPlanInstancesResponse, <-chan error) {
	responseChan := make(chan *ListExecutionPlanInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListExecutionPlanInstances(request)
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

// ListExecutionPlanInstancesWithCallback invokes the emr.ListExecutionPlanInstances API asynchronously
// api document: https://help.aliyun.com/api/emr/listexecutionplaninstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListExecutionPlanInstancesWithCallback(request *ListExecutionPlanInstancesRequest, callback func(response *ListExecutionPlanInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListExecutionPlanInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListExecutionPlanInstances(request)
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

// ListExecutionPlanInstancesRequest is the request struct for api ListExecutionPlanInstances
type ListExecutionPlanInstancesRequest struct {
	*requests.RpcRequest
	OnlyLastInstance    requests.Boolean `position:"Query" name:"OnlyLastInstance"`
	ResourceOwnerId     requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanIdList *[]string        `position:"Query" name:"ExecutionPlanIdList"  type:"Repeated"`
	StatusList          *[]string        `position:"Query" name:"StatusList"  type:"Repeated"`
	IsDesc              requests.Boolean `position:"Query" name:"IsDesc"`
	PageNumber          requests.Integer `position:"Query" name:"PageNumber"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
}

// ListExecutionPlanInstancesResponse is the response struct for api ListExecutionPlanInstances
type ListExecutionPlanInstancesResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	PageNumber             int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize               int                    `json:"PageSize" xml:"PageSize"`
	ExecutionPlanInstances ExecutionPlanInstances `json:"ExecutionPlanInstances" xml:"ExecutionPlanInstances"`
}

// CreateListExecutionPlanInstancesRequest creates a request to invoke ListExecutionPlanInstances API
func CreateListExecutionPlanInstancesRequest() (request *ListExecutionPlanInstancesRequest) {
	request = &ListExecutionPlanInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListExecutionPlanInstances", "emr", "openAPI")
	return
}

// CreateListExecutionPlanInstancesResponse creates a response to parse from ListExecutionPlanInstances response
func CreateListExecutionPlanInstancesResponse() (response *ListExecutionPlanInstancesResponse) {
	response = &ListExecutionPlanInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
