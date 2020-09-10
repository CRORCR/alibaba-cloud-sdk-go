package cloudesl

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

// DescribeAlarms invokes the cloudesl.DescribeAlarms API synchronously
// api document: https://help.aliyun.com/api/cloudesl/describealarms.html
func (client *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
	response = CreateDescribeAlarmsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlarmsWithChan invokes the cloudesl.DescribeAlarms API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/describealarms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmsWithChan(request *DescribeAlarmsRequest) (<-chan *DescribeAlarmsResponse, <-chan error) {
	responseChan := make(chan *DescribeAlarmsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlarms(request)
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

// DescribeAlarmsWithCallback invokes the cloudesl.DescribeAlarms API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/describealarms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmsWithCallback(request *DescribeAlarmsRequest, callback func(response *DescribeAlarmsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlarmsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlarms(request)
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

// DescribeAlarmsRequest is the request struct for api DescribeAlarms
type DescribeAlarmsRequest struct {
	*requests.RpcRequest
	StoreId     string           `position:"Body" name:"StoreId"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	AlarmType   string           `position:"Body" name:"AlarmType"`
	AlarmStatus string           `position:"Body" name:"AlarmStatus"`
	ErrorType   string           `position:"Body" name:"ErrorType"`
	AlarmId     string           `position:"Body" name:"AlarmId"`
	DeviceMac   string           `position:"Body" name:"DeviceMac"`
}

// DescribeAlarmsResponse is the response struct for api DescribeAlarms
type DescribeAlarmsResponse struct {
	*responses.BaseResponse
	ErrorMessage   string      `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string      `json:"ErrorCode" xml:"ErrorCode"`
	PageSize       int         `json:"PageSize" xml:"PageSize"`
	Message        string      `json:"Message" xml:"Message"`
	TotalCount     int         `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int         `json:"PageNumber" xml:"PageNumber"`
	DynamicCode    string      `json:"DynamicCode" xml:"DynamicCode"`
	Code           string      `json:"Code" xml:"Code"`
	DynamicMessage string      `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Alarms         []AlarmInfo `json:"Alarms" xml:"Alarms"`
}

// CreateDescribeAlarmsRequest creates a request to invoke DescribeAlarms API
func CreateDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
	request = &DescribeAlarmsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeAlarms", "cloudesl", "openAPI")
	return
}

// CreateDescribeAlarmsResponse creates a response to parse from DescribeAlarms response
func CreateDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
	response = &DescribeAlarmsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
