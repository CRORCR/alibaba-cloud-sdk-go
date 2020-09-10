package ecs

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

// DescribeDisksFullStatus invokes the ecs.DescribeDisksFullStatus API synchronously
// api document: https://help.aliyun.com/api/ecs/describedisksfullstatus.html
func (client *Client) DescribeDisksFullStatus(request *DescribeDisksFullStatusRequest) (response *DescribeDisksFullStatusResponse, err error) {
	response = CreateDescribeDisksFullStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDisksFullStatusWithChan invokes the ecs.DescribeDisksFullStatus API asynchronously
// api document: https://help.aliyun.com/api/ecs/describedisksfullstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDisksFullStatusWithChan(request *DescribeDisksFullStatusRequest) (<-chan *DescribeDisksFullStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeDisksFullStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDisksFullStatus(request)
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

// DescribeDisksFullStatusWithCallback invokes the ecs.DescribeDisksFullStatus API asynchronously
// api document: https://help.aliyun.com/api/ecs/describedisksfullstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDisksFullStatusWithCallback(request *DescribeDisksFullStatusRequest, callback func(response *DescribeDisksFullStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDisksFullStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeDisksFullStatus(request)
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

// DescribeDisksFullStatusRequest is the request struct for api DescribeDisksFullStatus
type DescribeDisksFullStatusRequest struct {
	*requests.RpcRequest
	EventId              *[]string        `position:"Query" name:"EventId"  type:"Repeated"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	EventTimeStart       string           `position:"Query" name:"EventTime.Start"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DiskId               *[]string        `position:"Query" name:"DiskId"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	EventTimeEnd         string           `position:"Query" name:"EventTime.End"`
	HealthStatus         string           `position:"Query" name:"HealthStatus"`
	EventType            string           `position:"Query" name:"EventType"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeDisksFullStatusResponse is the response struct for api DescribeDisksFullStatus
type DescribeDisksFullStatusResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	DiskFullStatusSet DiskFullStatusSet `json:"DiskFullStatusSet" xml:"DiskFullStatusSet"`
}

// CreateDescribeDisksFullStatusRequest creates a request to invoke DescribeDisksFullStatus API
func CreateDescribeDisksFullStatusRequest() (request *DescribeDisksFullStatusRequest) {
	request = &DescribeDisksFullStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDisksFullStatus", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDisksFullStatusResponse creates a response to parse from DescribeDisksFullStatus response
func CreateDescribeDisksFullStatusResponse() (response *DescribeDisksFullStatusResponse) {
	response = &DescribeDisksFullStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
