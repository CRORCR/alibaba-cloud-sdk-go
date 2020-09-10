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

// ModifySecurityCheckScheduleConfig invokes the sas.ModifySecurityCheckScheduleConfig API synchronously
// api document: https://help.aliyun.com/api/sas/modifysecuritycheckscheduleconfig.html
func (client *Client) ModifySecurityCheckScheduleConfig(request *ModifySecurityCheckScheduleConfigRequest) (response *ModifySecurityCheckScheduleConfigResponse, err error) {
	response = CreateModifySecurityCheckScheduleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySecurityCheckScheduleConfigWithChan invokes the sas.ModifySecurityCheckScheduleConfig API asynchronously
// api document: https://help.aliyun.com/api/sas/modifysecuritycheckscheduleconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityCheckScheduleConfigWithChan(request *ModifySecurityCheckScheduleConfigRequest) (<-chan *ModifySecurityCheckScheduleConfigResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityCheckScheduleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityCheckScheduleConfig(request)
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

// ModifySecurityCheckScheduleConfigWithCallback invokes the sas.ModifySecurityCheckScheduleConfig API asynchronously
// api document: https://help.aliyun.com/api/sas/modifysecuritycheckscheduleconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityCheckScheduleConfigWithCallback(request *ModifySecurityCheckScheduleConfigRequest, callback func(response *ModifySecurityCheckScheduleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityCheckScheduleConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityCheckScheduleConfig(request)
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

// ModifySecurityCheckScheduleConfigRequest is the request struct for api ModifySecurityCheckScheduleConfig
type ModifySecurityCheckScheduleConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	DaysOfWeek      string           `position:"Query" name:"DaysOfWeek"`
	Lang            string           `position:"Query" name:"Lang"`
}

// ModifySecurityCheckScheduleConfigResponse is the response struct for api ModifySecurityCheckScheduleConfig
type ModifySecurityCheckScheduleConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySecurityCheckScheduleConfigRequest creates a request to invoke ModifySecurityCheckScheduleConfig API
func CreateModifySecurityCheckScheduleConfigRequest() (request *ModifySecurityCheckScheduleConfigRequest) {
	request = &ModifySecurityCheckScheduleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifySecurityCheckScheduleConfig", "sas", "openAPI")
	return
}

// CreateModifySecurityCheckScheduleConfigResponse creates a response to parse from ModifySecurityCheckScheduleConfig response
func CreateModifySecurityCheckScheduleConfigResponse() (response *ModifySecurityCheckScheduleConfigResponse) {
	response = &ModifySecurityCheckScheduleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
