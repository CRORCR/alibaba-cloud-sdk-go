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

package dts

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeConsumerGroup invokes the dts.DescribeConsumerGroup API synchronously
// api document: https://help.aliyun.com/api/dts/describeconsumergroup.html
func (client *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
	response = CreateDescribeConsumerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConsumerGroupWithChan invokes the dts.DescribeConsumerGroup API asynchronously
// api document: https://help.aliyun.com/api/dts/describeconsumergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConsumerGroupWithChan(request *DescribeConsumerGroupRequest) (<-chan *DescribeConsumerGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeConsumerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConsumerGroup(request)
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

// DescribeConsumerGroupWithCallback invokes the dts.DescribeConsumerGroup API asynchronously
// api document: https://help.aliyun.com/api/dts/describeconsumergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConsumerGroupWithCallback(request *DescribeConsumerGroupRequest, callback func(response *DescribeConsumerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConsumerGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeConsumerGroup(request)
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

// DescribeConsumerGroupRequest is the request struct for api DescribeConsumerGroup
type DescribeConsumerGroupRequest struct {
	*requests.RpcRequest
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	PageNum                requests.Integer `position:"Query" name:"PageNum"`
	SubscriptionInstanceId string           `position:"Query" name:"SubscriptionInstanceId"`
	OwnerId                string           `position:"Query" name:"OwnerId"`
}

// DescribeConsumerGroupResponse is the response struct for api DescribeConsumerGroup
type DescribeConsumerGroupResponse struct {
	*responses.BaseResponse
	PageNumber       int                                             `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int                                             `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                                             `json:"PageRecordCount" xml:"PageRecordCount"`
	RequestId        string                                          `json:"RequestId" xml:"RequestId"`
	ConsumerChannels []DescribeConsumerGroupDescribeConsumerChannel0 `json:"ConsumerChannels" xml:"ConsumerChannels"`
}

type DescribeConsumerGroupDescribeConsumerChannel0 struct {
	ConsumerGroupID       string `json:"ConsumerGroupID" xml:"ConsumerGroupID"`
	ConsumerGroupName     string `json:"ConsumerGroupName" xml:"ConsumerGroupName"`
	ConsumptionCheckpoint string `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	UnconsumedData        int64  `json:"UnconsumedData" xml:"UnconsumedData"`
	MessageDelay          int64  `json:"MessageDelay" xml:"MessageDelay"`
	ConsumerGroupUserName string `json:"ConsumerGroupUserName" xml:"ConsumerGroupUserName"`
}

// CreateDescribeConsumerGroupRequest creates a request to invoke DescribeConsumerGroup API
func CreateDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
	request = &DescribeConsumerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "DescribeConsumerGroup", "dts", "openAPI")
	return
}

// CreateDescribeConsumerGroupResponse creates a response to parse from DescribeConsumerGroup response
func CreateDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
	response = &DescribeConsumerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
