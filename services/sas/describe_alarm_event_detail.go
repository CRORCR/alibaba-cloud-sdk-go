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

// DescribeAlarmEventDetail invokes the sas.DescribeAlarmEventDetail API synchronously
// api document: https://help.aliyun.com/api/sas/describealarmeventdetail.html
func (client *Client) DescribeAlarmEventDetail(request *DescribeAlarmEventDetailRequest) (response *DescribeAlarmEventDetailResponse, err error) {
	response = CreateDescribeAlarmEventDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlarmEventDetailWithChan invokes the sas.DescribeAlarmEventDetail API asynchronously
// api document: https://help.aliyun.com/api/sas/describealarmeventdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmEventDetailWithChan(request *DescribeAlarmEventDetailRequest) (<-chan *DescribeAlarmEventDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeAlarmEventDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlarmEventDetail(request)
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

// DescribeAlarmEventDetailWithCallback invokes the sas.DescribeAlarmEventDetail API asynchronously
// api document: https://help.aliyun.com/api/sas/describealarmeventdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmEventDetailWithCallback(request *DescribeAlarmEventDetailRequest, callback func(response *DescribeAlarmEventDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlarmEventDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlarmEventDetail(request)
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

// DescribeAlarmEventDetailRequest is the request struct for api DescribeAlarmEventDetail
type DescribeAlarmEventDetailRequest struct {
	*requests.RpcRequest
	AlarmUniqueInfo string `position:"Query" name:"AlarmUniqueInfo"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	From            string `position:"Query" name:"From"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeAlarmEventDetailResponse is the response struct for api DescribeAlarmEventDetail
type DescribeAlarmEventDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAlarmEventDetailRequest creates a request to invoke DescribeAlarmEventDetail API
func CreateDescribeAlarmEventDetailRequest() (request *DescribeAlarmEventDetailRequest) {
	request = &DescribeAlarmEventDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeAlarmEventDetail", "sas", "openAPI")
	return
}

// CreateDescribeAlarmEventDetailResponse creates a response to parse from DescribeAlarmEventDetail response
func CreateDescribeAlarmEventDetailResponse() (response *DescribeAlarmEventDetailResponse) {
	response = &DescribeAlarmEventDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
