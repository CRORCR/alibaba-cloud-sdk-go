package aegis

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

// DescribeScreenHostStatistics invokes the aegis.DescribeScreenHostStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/describescreenhoststatistics.html
func (client *Client) DescribeScreenHostStatistics(request *DescribeScreenHostStatisticsRequest) (response *DescribeScreenHostStatisticsResponse, err error) {
	response = CreateDescribeScreenHostStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScreenHostStatisticsWithChan invokes the aegis.DescribeScreenHostStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreenhoststatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenHostStatisticsWithChan(request *DescribeScreenHostStatisticsRequest) (<-chan *DescribeScreenHostStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeScreenHostStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScreenHostStatistics(request)
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

// DescribeScreenHostStatisticsWithCallback invokes the aegis.DescribeScreenHostStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreenhoststatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenHostStatisticsWithCallback(request *DescribeScreenHostStatisticsRequest, callback func(response *DescribeScreenHostStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScreenHostStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeScreenHostStatistics(request)
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

// DescribeScreenHostStatisticsRequest is the request struct for api DescribeScreenHostStatistics
type DescribeScreenHostStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeScreenHostStatisticsResponse is the response struct for api DescribeScreenHostStatistics
type DescribeScreenHostStatisticsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeScreenHostStatisticsRequest creates a request to invoke DescribeScreenHostStatistics API
func CreateDescribeScreenHostStatisticsRequest() (request *DescribeScreenHostStatisticsRequest) {
	request = &DescribeScreenHostStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeScreenHostStatistics", "vipaegis", "openAPI")
	return
}

// CreateDescribeScreenHostStatisticsResponse creates a response to parse from DescribeScreenHostStatistics response
func CreateDescribeScreenHostStatisticsResponse() (response *DescribeScreenHostStatisticsResponse) {
	response = &DescribeScreenHostStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
