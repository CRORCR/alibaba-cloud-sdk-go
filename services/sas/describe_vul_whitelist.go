package sas

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

// DescribeVulWhitelist invokes the sas.DescribeVulWhitelist API synchronously
// api document: https://help.aliyun.com/api/sas/describevulwhitelist.html
func (client *Client) DescribeVulWhitelist(request *DescribeVulWhitelistRequest) (response *DescribeVulWhitelistResponse, err error) {
	response = CreateDescribeVulWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulWhitelistWithChan invokes the sas.DescribeVulWhitelist API asynchronously
// api document: https://help.aliyun.com/api/sas/describevulwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulWhitelistWithChan(request *DescribeVulWhitelistRequest) (<-chan *DescribeVulWhitelistResponse, <-chan error) {
	responseChan := make(chan *DescribeVulWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulWhitelist(request)
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

// DescribeVulWhitelistWithCallback invokes the sas.DescribeVulWhitelist API asynchronously
// api document: https://help.aliyun.com/api/sas/describevulwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulWhitelistWithCallback(request *DescribeVulWhitelistRequest, callback func(response *DescribeVulWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulWhitelistResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulWhitelist(request)
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

// DescribeVulWhitelistRequest is the request struct for api DescribeVulWhitelist
type DescribeVulWhitelistRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeVulWhitelistResponse is the response struct for api DescribeVulWhitelist
type DescribeVulWhitelistResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	PageSize      int            `json:"PageSize" xml:"PageSize"`
	CurrentPage   int            `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount    int            `json:"TotalCount" xml:"TotalCount"`
	VulWhitelists []VulWhitelist `json:"VulWhitelists" xml:"VulWhitelists"`
}

// CreateDescribeVulWhitelistRequest creates a request to invoke DescribeVulWhitelist API
func CreateDescribeVulWhitelistRequest() (request *DescribeVulWhitelistRequest) {
	request = &DescribeVulWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeVulWhitelist", "sas", "openAPI")
	return
}

// CreateDescribeVulWhitelistResponse creates a response to parse from DescribeVulWhitelist response
func CreateDescribeVulWhitelistResponse() (response *DescribeVulWhitelistResponse) {
	response = &DescribeVulWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
