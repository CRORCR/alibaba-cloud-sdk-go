package ddoscoo

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

// DescribeWebAccessLogEmptyCount invokes the ddoscoo.DescribeWebAccessLogEmptyCount API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describewebaccesslogemptycount.html
func (client *Client) DescribeWebAccessLogEmptyCount(request *DescribeWebAccessLogEmptyCountRequest) (response *DescribeWebAccessLogEmptyCountResponse, err error) {
	response = CreateDescribeWebAccessLogEmptyCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebAccessLogEmptyCountWithChan invokes the ddoscoo.DescribeWebAccessLogEmptyCount API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describewebaccesslogemptycount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebAccessLogEmptyCountWithChan(request *DescribeWebAccessLogEmptyCountRequest) (<-chan *DescribeWebAccessLogEmptyCountResponse, <-chan error) {
	responseChan := make(chan *DescribeWebAccessLogEmptyCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebAccessLogEmptyCount(request)
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

// DescribeWebAccessLogEmptyCountWithCallback invokes the ddoscoo.DescribeWebAccessLogEmptyCount API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describewebaccesslogemptycount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebAccessLogEmptyCountWithCallback(request *DescribeWebAccessLogEmptyCountRequest, callback func(response *DescribeWebAccessLogEmptyCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebAccessLogEmptyCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebAccessLogEmptyCount(request)
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

// DescribeWebAccessLogEmptyCountRequest is the request struct for api DescribeWebAccessLogEmptyCount
type DescribeWebAccessLogEmptyCountRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeWebAccessLogEmptyCountResponse is the response struct for api DescribeWebAccessLogEmptyCount
type DescribeWebAccessLogEmptyCountResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	AvailableCount int    `json:"AvailableCount" xml:"AvailableCount"`
}

// CreateDescribeWebAccessLogEmptyCountRequest creates a request to invoke DescribeWebAccessLogEmptyCount API
func CreateDescribeWebAccessLogEmptyCountRequest() (request *DescribeWebAccessLogEmptyCountRequest) {
	request = &DescribeWebAccessLogEmptyCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebAccessLogEmptyCount", "ddoscoo", "openAPI")
	return
}

// CreateDescribeWebAccessLogEmptyCountResponse creates a response to parse from DescribeWebAccessLogEmptyCount response
func CreateDescribeWebAccessLogEmptyCountResponse() (response *DescribeWebAccessLogEmptyCountResponse) {
	response = &DescribeWebAccessLogEmptyCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
