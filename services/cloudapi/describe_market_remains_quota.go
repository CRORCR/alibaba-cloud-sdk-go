package cloudapi

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

// DescribeMarketRemainsQuota invokes the cloudapi.DescribeMarketRemainsQuota API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describemarketremainsquota.html
func (client *Client) DescribeMarketRemainsQuota(request *DescribeMarketRemainsQuotaRequest) (response *DescribeMarketRemainsQuotaResponse, err error) {
	response = CreateDescribeMarketRemainsQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMarketRemainsQuotaWithChan invokes the cloudapi.DescribeMarketRemainsQuota API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describemarketremainsquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMarketRemainsQuotaWithChan(request *DescribeMarketRemainsQuotaRequest) (<-chan *DescribeMarketRemainsQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeMarketRemainsQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMarketRemainsQuota(request)
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

// DescribeMarketRemainsQuotaWithCallback invokes the cloudapi.DescribeMarketRemainsQuota API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describemarketremainsquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMarketRemainsQuotaWithCallback(request *DescribeMarketRemainsQuotaRequest, callback func(response *DescribeMarketRemainsQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMarketRemainsQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeMarketRemainsQuota(request)
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

// DescribeMarketRemainsQuotaRequest is the request struct for api DescribeMarketRemainsQuota
type DescribeMarketRemainsQuotaRequest struct {
	*requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeMarketRemainsQuotaResponse is the response struct for api DescribeMarketRemainsQuota
type DescribeMarketRemainsQuotaResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	RemainsQuota int64  `json:"RemainsQuota" xml:"RemainsQuota"`
}

// CreateDescribeMarketRemainsQuotaRequest creates a request to invoke DescribeMarketRemainsQuota API
func CreateDescribeMarketRemainsQuotaRequest() (request *DescribeMarketRemainsQuotaRequest) {
	request = &DescribeMarketRemainsQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeMarketRemainsQuota", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMarketRemainsQuotaResponse creates a response to parse from DescribeMarketRemainsQuota response
func CreateDescribeMarketRemainsQuotaResponse() (response *DescribeMarketRemainsQuotaResponse) {
	response = &DescribeMarketRemainsQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
