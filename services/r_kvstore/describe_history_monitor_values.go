package r_kvstore

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

// DescribeHistoryMonitorValues invokes the r_kvstore.DescribeHistoryMonitorValues API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describehistorymonitorvalues.html
func (client *Client) DescribeHistoryMonitorValues(request *DescribeHistoryMonitorValuesRequest) (response *DescribeHistoryMonitorValuesResponse, err error) {
	response = CreateDescribeHistoryMonitorValuesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHistoryMonitorValuesWithChan invokes the r_kvstore.DescribeHistoryMonitorValues API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describehistorymonitorvalues.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHistoryMonitorValuesWithChan(request *DescribeHistoryMonitorValuesRequest) (<-chan *DescribeHistoryMonitorValuesResponse, <-chan error) {
	responseChan := make(chan *DescribeHistoryMonitorValuesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHistoryMonitorValues(request)
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

// DescribeHistoryMonitorValuesWithCallback invokes the r_kvstore.DescribeHistoryMonitorValues API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describehistorymonitorvalues.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHistoryMonitorValuesWithCallback(request *DescribeHistoryMonitorValuesRequest, callback func(response *DescribeHistoryMonitorValuesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHistoryMonitorValuesResponse
		var err error
		defer close(result)
		response, err = client.DescribeHistoryMonitorValues(request)
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

// DescribeHistoryMonitorValuesRequest is the request struct for api DescribeHistoryMonitorValues
type DescribeHistoryMonitorValuesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	IntervalForHistory   string           `position:"Query" name:"IntervalForHistory"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	MonitorKeys          string           `position:"Query" name:"MonitorKeys"`
}

// DescribeHistoryMonitorValuesResponse is the response struct for api DescribeHistoryMonitorValues
type DescribeHistoryMonitorValuesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	MonitorHistory string `json:"MonitorHistory" xml:"MonitorHistory"`
}

// CreateDescribeHistoryMonitorValuesRequest creates a request to invoke DescribeHistoryMonitorValues API
func CreateDescribeHistoryMonitorValuesRequest() (request *DescribeHistoryMonitorValuesRequest) {
	request = &DescribeHistoryMonitorValuesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeHistoryMonitorValues", "redisa", "openAPI")
	return
}

// CreateDescribeHistoryMonitorValuesResponse creates a response to parse from DescribeHistoryMonitorValues response
func CreateDescribeHistoryMonitorValuesResponse() (response *DescribeHistoryMonitorValuesResponse) {
	response = &DescribeHistoryMonitorValuesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
