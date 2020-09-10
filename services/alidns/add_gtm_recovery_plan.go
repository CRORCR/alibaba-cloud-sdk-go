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

// AddGtmRecoveryPlan invokes the alidns.AddGtmRecoveryPlan API synchronously
// api document: https://help.aliyun.com/api/alidns/addgtmrecoveryplan.html
func (client *Client) AddGtmRecoveryPlan(request *AddGtmRecoveryPlanRequest) (response *AddGtmRecoveryPlanResponse, err error) {
	response = CreateAddGtmRecoveryPlanResponse()
	err = client.DoAction(request, response)
	return
}

// AddGtmRecoveryPlanWithChan invokes the alidns.AddGtmRecoveryPlan API asynchronously
// api document: https://help.aliyun.com/api/alidns/addgtmrecoveryplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddGtmRecoveryPlanWithChan(request *AddGtmRecoveryPlanRequest) (<-chan *AddGtmRecoveryPlanResponse, <-chan error) {
	responseChan := make(chan *AddGtmRecoveryPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddGtmRecoveryPlan(request)
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

// AddGtmRecoveryPlanWithCallback invokes the alidns.AddGtmRecoveryPlan API asynchronously
// api document: https://help.aliyun.com/api/alidns/addgtmrecoveryplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddGtmRecoveryPlanWithCallback(request *AddGtmRecoveryPlanRequest, callback func(response *AddGtmRecoveryPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddGtmRecoveryPlanResponse
		var err error
		defer close(result)
		response, err = client.AddGtmRecoveryPlan(request)
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

// AddGtmRecoveryPlanRequest is the request struct for api AddGtmRecoveryPlan
type AddGtmRecoveryPlanRequest struct {
	*requests.RpcRequest
	FaultAddrPool string `position:"Query" name:"FaultAddrPool"`
	Remark        string `position:"Query" name:"Remark"`
	UserClientIp  string `position:"Query" name:"UserClientIp"`
	Name          string `position:"Query" name:"Name"`
	Lang          string `position:"Query" name:"Lang"`
}

// AddGtmRecoveryPlanResponse is the response struct for api AddGtmRecoveryPlan
type AddGtmRecoveryPlanResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RecoveryPlanId string `json:"RecoveryPlanId" xml:"RecoveryPlanId"`
}

// CreateAddGtmRecoveryPlanRequest creates a request to invoke AddGtmRecoveryPlan API
func CreateAddGtmRecoveryPlanRequest() (request *AddGtmRecoveryPlanRequest) {
	request = &AddGtmRecoveryPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "AddGtmRecoveryPlan", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddGtmRecoveryPlanResponse creates a response to parse from AddGtmRecoveryPlan response
func CreateAddGtmRecoveryPlanResponse() (response *AddGtmRecoveryPlanResponse) {
	response = &AddGtmRecoveryPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
