package alidns

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

// DescribeGtmRecoveryPlan invokes the alidns.DescribeGtmRecoveryPlan API synchronously
// api document: https://help.aliyun.com/api/alidns/describegtmrecoveryplan.html
func (client *Client) DescribeGtmRecoveryPlan(request *DescribeGtmRecoveryPlanRequest) (response *DescribeGtmRecoveryPlanResponse, err error) {
	response = CreateDescribeGtmRecoveryPlanResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmRecoveryPlanWithChan invokes the alidns.DescribeGtmRecoveryPlan API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtmrecoveryplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmRecoveryPlanWithChan(request *DescribeGtmRecoveryPlanRequest) (<-chan *DescribeGtmRecoveryPlanResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmRecoveryPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmRecoveryPlan(request)
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

// DescribeGtmRecoveryPlanWithCallback invokes the alidns.DescribeGtmRecoveryPlan API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtmrecoveryplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmRecoveryPlanWithCallback(request *DescribeGtmRecoveryPlanRequest, callback func(response *DescribeGtmRecoveryPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmRecoveryPlanResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmRecoveryPlan(request)
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

// DescribeGtmRecoveryPlanRequest is the request struct for api DescribeGtmRecoveryPlan
type DescribeGtmRecoveryPlanRequest struct {
	*requests.RpcRequest
	UserClientIp   string           `position:"Query" name:"UserClientIp"`
	RecoveryPlanId requests.Integer `position:"Query" name:"RecoveryPlanId"`
	Lang           string           `position:"Query" name:"Lang"`
}

// DescribeGtmRecoveryPlanResponse is the response struct for api DescribeGtmRecoveryPlan
type DescribeGtmRecoveryPlanResponse struct {
	*responses.BaseResponse
	RequestId             string         `json:"RequestId" xml:"RequestId"`
	RecoveryPlanId        int64          `json:"RecoveryPlanId" xml:"RecoveryPlanId"`
	Name                  string         `json:"Name" xml:"Name"`
	Remark                string         `json:"Remark" xml:"Remark"`
	FaultAddrPoolNum      int            `json:"FaultAddrPoolNum" xml:"FaultAddrPoolNum"`
	Status                string         `json:"Status" xml:"Status"`
	LastExecuteTime       string         `json:"LastExecuteTime" xml:"LastExecuteTime"`
	LastExecuteTimestamp  int64          `json:"LastExecuteTimestamp" xml:"LastExecuteTimestamp"`
	LastRollbackTime      string         `json:"LastRollbackTime" xml:"LastRollbackTime"`
	LastRollbackTimestamp int64          `json:"LastRollbackTimestamp" xml:"LastRollbackTimestamp"`
	CreateTime            string         `json:"CreateTime" xml:"CreateTime"`
	CreateTimestamp       int64          `json:"CreateTimestamp" xml:"CreateTimestamp"`
	UpdateTime            string         `json:"UpdateTime" xml:"UpdateTime"`
	UpdateTimestamp       int64          `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	FaultAddrPools        FaultAddrPools `json:"FaultAddrPools" xml:"FaultAddrPools"`
}

// CreateDescribeGtmRecoveryPlanRequest creates a request to invoke DescribeGtmRecoveryPlan API
func CreateDescribeGtmRecoveryPlanRequest() (request *DescribeGtmRecoveryPlanRequest) {
	request = &DescribeGtmRecoveryPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmRecoveryPlan", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGtmRecoveryPlanResponse creates a response to parse from DescribeGtmRecoveryPlan response
func CreateDescribeGtmRecoveryPlanResponse() (response *DescribeGtmRecoveryPlanResponse) {
	response = &DescribeGtmRecoveryPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
