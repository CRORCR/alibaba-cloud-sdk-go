package yundun_ds

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

// DescribeTransferEventCounts invokes the yundun_ds.DescribeTransferEventCounts API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describetransfereventcounts.html
func (client *Client) DescribeTransferEventCounts(request *DescribeTransferEventCountsRequest) (response *DescribeTransferEventCountsResponse, err error) {
	response = CreateDescribeTransferEventCountsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTransferEventCountsWithChan invokes the yundun_ds.DescribeTransferEventCounts API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describetransfereventcounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTransferEventCountsWithChan(request *DescribeTransferEventCountsRequest) (<-chan *DescribeTransferEventCountsResponse, <-chan error) {
	responseChan := make(chan *DescribeTransferEventCountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTransferEventCounts(request)
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

// DescribeTransferEventCountsWithCallback invokes the yundun_ds.DescribeTransferEventCounts API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describetransfereventcounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTransferEventCountsWithCallback(request *DescribeTransferEventCountsRequest, callback func(response *DescribeTransferEventCountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTransferEventCountsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTransferEventCounts(request)
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

// DescribeTransferEventCountsRequest is the request struct for api DescribeTransferEventCounts
type DescribeTransferEventCountsRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DescribeTransferEventCountsResponse is the response struct for api DescribeTransferEventCounts
type DescribeTransferEventCountsResponse struct {
	*responses.BaseResponse
	RequestId              string     `json:"RequestId" xml:"RequestId"`
	TransferEventCountList []Transfer `json:"TransferEventCountList" xml:"TransferEventCountList"`
}

// CreateDescribeTransferEventCountsRequest creates a request to invoke DescribeTransferEventCounts API
func CreateDescribeTransferEventCountsRequest() (request *DescribeTransferEventCountsRequest) {
	request = &DescribeTransferEventCountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeTransferEventCounts", "sddp", "openAPI")
	return
}

// CreateDescribeTransferEventCountsResponse creates a response to parse from DescribeTransferEventCounts response
func CreateDescribeTransferEventCountsResponse() (response *DescribeTransferEventCountsResponse) {
	response = &DescribeTransferEventCountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
