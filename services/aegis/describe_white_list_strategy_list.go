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

// DescribeWhiteListStrategyList invokes the aegis.DescribeWhiteListStrategyList API synchronously
// api document: https://help.aliyun.com/api/aegis/describewhiteliststrategylist.html
func (client *Client) DescribeWhiteListStrategyList(request *DescribeWhiteListStrategyListRequest) (response *DescribeWhiteListStrategyListResponse, err error) {
	response = CreateDescribeWhiteListStrategyListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhiteListStrategyListWithChan invokes the aegis.DescribeWhiteListStrategyList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhiteliststrategylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListStrategyListWithChan(request *DescribeWhiteListStrategyListRequest) (<-chan *DescribeWhiteListStrategyListResponse, <-chan error) {
	responseChan := make(chan *DescribeWhiteListStrategyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhiteListStrategyList(request)
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

// DescribeWhiteListStrategyListWithCallback invokes the aegis.DescribeWhiteListStrategyList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhiteliststrategylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListStrategyListWithCallback(request *DescribeWhiteListStrategyListRequest, callback func(response *DescribeWhiteListStrategyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhiteListStrategyListResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhiteListStrategyList(request)
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

// DescribeWhiteListStrategyListRequest is the request struct for api DescribeWhiteListStrategyList
type DescribeWhiteListStrategyListRequest struct {
	*requests.RpcRequest
	SourceIp    string `position:"Query" name:"SourceIp"`
	StrategyIds string `position:"Query" name:"StrategyIds"`
	Lang        string `position:"Query" name:"Lang"`
}

// DescribeWhiteListStrategyListResponse is the response struct for api DescribeWhiteListStrategyList
type DescribeWhiteListStrategyListResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Strategies []Strategy `json:"Strategies" xml:"Strategies"`
}

// CreateDescribeWhiteListStrategyListRequest creates a request to invoke DescribeWhiteListStrategyList API
func CreateDescribeWhiteListStrategyListRequest() (request *DescribeWhiteListStrategyListRequest) {
	request = &DescribeWhiteListStrategyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWhiteListStrategyList", "vipaegis", "openAPI")
	return
}

// CreateDescribeWhiteListStrategyListResponse creates a response to parse from DescribeWhiteListStrategyList response
func CreateDescribeWhiteListStrategyListResponse() (response *DescribeWhiteListStrategyListResponse) {
	response = &DescribeWhiteListStrategyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
