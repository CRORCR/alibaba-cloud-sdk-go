package ens

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

// DescribeServcieSchedule invokes the ens.DescribeServcieSchedule API synchronously
// api document: https://help.aliyun.com/api/ens/describeservcieschedule.html
func (client *Client) DescribeServcieSchedule(request *DescribeServcieScheduleRequest) (response *DescribeServcieScheduleResponse, err error) {
	response = CreateDescribeServcieScheduleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServcieScheduleWithChan invokes the ens.DescribeServcieSchedule API asynchronously
// api document: https://help.aliyun.com/api/ens/describeservcieschedule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServcieScheduleWithChan(request *DescribeServcieScheduleRequest) (<-chan *DescribeServcieScheduleResponse, <-chan error) {
	responseChan := make(chan *DescribeServcieScheduleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServcieSchedule(request)
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

// DescribeServcieScheduleWithCallback invokes the ens.DescribeServcieSchedule API asynchronously
// api document: https://help.aliyun.com/api/ens/describeservcieschedule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServcieScheduleWithCallback(request *DescribeServcieScheduleRequest, callback func(response *DescribeServcieScheduleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServcieScheduleResponse
		var err error
		defer close(result)
		response, err = client.DescribeServcieSchedule(request)
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

// DescribeServcieScheduleRequest is the request struct for api DescribeServcieSchedule
type DescribeServcieScheduleRequest struct {
	*requests.RpcRequest
	Uuid          string `position:"Query" name:"Uuid"`
	PodConfigName string `position:"Query" name:"PodConfigName"`
	AppId         string `position:"Query" name:"AppId"`
}

// DescribeServcieScheduleResponse is the response struct for api DescribeServcieSchedule
type DescribeServcieScheduleResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	InstanceId      string          `json:"InstanceId" xml:"InstanceId"`
	InstanceIp      string          `json:"InstanceIp" xml:"InstanceIp"`
	InstancePort    int             `json:"InstancePort" xml:"InstancePort"`
	Index           int             `json:"Index" xml:"Index"`
	TcpPorts        string          `json:"TcpPorts" xml:"TcpPorts"`
	RequestRepeated bool            `json:"RequestRepeated" xml:"RequestRepeated"`
	PodAbstractInfo PodAbstractInfo `json:"PodAbstractInfo" xml:"PodAbstractInfo"`
}

// CreateDescribeServcieScheduleRequest creates a request to invoke DescribeServcieSchedule API
func CreateDescribeServcieScheduleRequest() (request *DescribeServcieScheduleRequest) {
	request = &DescribeServcieScheduleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeServcieSchedule", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeServcieScheduleResponse creates a response to parse from DescribeServcieSchedule response
func CreateDescribeServcieScheduleResponse() (response *DescribeServcieScheduleResponse) {
	response = &DescribeServcieScheduleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
