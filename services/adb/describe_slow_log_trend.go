package adb

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

// DescribeSlowLogTrend invokes the adb.DescribeSlowLogTrend API synchronously
// api document: https://help.aliyun.com/api/adb/describeslowlogtrend.html
func (client *Client) DescribeSlowLogTrend(request *DescribeSlowLogTrendRequest) (response *DescribeSlowLogTrendResponse, err error) {
	response = CreateDescribeSlowLogTrendResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlowLogTrendWithChan invokes the adb.DescribeSlowLogTrend API asynchronously
// api document: https://help.aliyun.com/api/adb/describeslowlogtrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogTrendWithChan(request *DescribeSlowLogTrendRequest) (<-chan *DescribeSlowLogTrendResponse, <-chan error) {
	responseChan := make(chan *DescribeSlowLogTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlowLogTrend(request)
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

// DescribeSlowLogTrendWithCallback invokes the adb.DescribeSlowLogTrend API asynchronously
// api document: https://help.aliyun.com/api/adb/describeslowlogtrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowLogTrendWithCallback(request *DescribeSlowLogTrendRequest, callback func(response *DescribeSlowLogTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlowLogTrendResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlowLogTrend(request)
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

// DescribeSlowLogTrendRequest is the request struct for api DescribeSlowLogTrend
type DescribeSlowLogTrendRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
}

// DescribeSlowLogTrendResponse is the response struct for api DescribeSlowLogTrend
type DescribeSlowLogTrendResponse struct {
	*responses.BaseResponse
	RequestId   string                      `json:"RequestId" xml:"RequestId"`
	DBClusterId string                      `json:"DBClusterId" xml:"DBClusterId"`
	StartTime   string                      `json:"StartTime" xml:"StartTime"`
	EndTime     string                      `json:"EndTime" xml:"EndTime"`
	Items       ItemsInDescribeSlowLogTrend `json:"Items" xml:"Items"`
}

// CreateDescribeSlowLogTrendRequest creates a request to invoke DescribeSlowLogTrend API
func CreateDescribeSlowLogTrendRequest() (request *DescribeSlowLogTrendRequest) {
	request = &DescribeSlowLogTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeSlowLogTrend", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSlowLogTrendResponse creates a response to parse from DescribeSlowLogTrend response
func CreateDescribeSlowLogTrendResponse() (response *DescribeSlowLogTrendResponse) {
	response = &DescribeSlowLogTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
