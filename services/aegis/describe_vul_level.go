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

// DescribeVulLevel invokes the aegis.DescribeVulLevel API synchronously
// api document: https://help.aliyun.com/api/aegis/describevullevel.html
func (client *Client) DescribeVulLevel(request *DescribeVulLevelRequest) (response *DescribeVulLevelResponse, err error) {
	response = CreateDescribeVulLevelResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulLevelWithChan invokes the aegis.DescribeVulLevel API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevullevel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulLevelWithChan(request *DescribeVulLevelRequest) (<-chan *DescribeVulLevelResponse, <-chan error) {
	responseChan := make(chan *DescribeVulLevelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulLevel(request)
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

// DescribeVulLevelWithCallback invokes the aegis.DescribeVulLevel API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevullevel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulLevelWithCallback(request *DescribeVulLevelRequest, callback func(response *DescribeVulLevelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulLevelResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulLevel(request)
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

// DescribeVulLevelRequest is the request struct for api DescribeVulLevel
type DescribeVulLevelRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeVulLevelResponse is the response struct for api DescribeVulLevel
type DescribeVulLevelResponse struct {
	*responses.BaseResponse
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	ConcernLevels []string `json:"ConcernLevels" xml:"ConcernLevels"`
}

// CreateDescribeVulLevelRequest creates a request to invoke DescribeVulLevel API
func CreateDescribeVulLevelRequest() (request *DescribeVulLevelRequest) {
	request = &DescribeVulLevelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeVulLevel", "vipaegis", "openAPI")
	return
}

// CreateDescribeVulLevelResponse creates a response to parse from DescribeVulLevel response
func CreateDescribeVulLevelResponse() (response *DescribeVulLevelResponse) {
	response = &DescribeVulLevelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
