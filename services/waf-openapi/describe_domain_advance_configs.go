package waf_openapi

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

// DescribeDomainAdvanceConfigs invokes the waf_openapi.DescribeDomainAdvanceConfigs API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/describedomainadvanceconfigs.html
func (client *Client) DescribeDomainAdvanceConfigs(request *DescribeDomainAdvanceConfigsRequest) (response *DescribeDomainAdvanceConfigsResponse, err error) {
	response = CreateDescribeDomainAdvanceConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainAdvanceConfigsWithChan invokes the waf_openapi.DescribeDomainAdvanceConfigs API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/describedomainadvanceconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainAdvanceConfigsWithChan(request *DescribeDomainAdvanceConfigsRequest) (<-chan *DescribeDomainAdvanceConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainAdvanceConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainAdvanceConfigs(request)
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

// DescribeDomainAdvanceConfigsWithCallback invokes the waf_openapi.DescribeDomainAdvanceConfigs API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/describedomainadvanceconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainAdvanceConfigsWithCallback(request *DescribeDomainAdvanceConfigsRequest, callback func(response *DescribeDomainAdvanceConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainAdvanceConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainAdvanceConfigs(request)
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

// DescribeDomainAdvanceConfigsRequest is the request struct for api DescribeDomainAdvanceConfigs
type DescribeDomainAdvanceConfigsRequest struct {
	*requests.RpcRequest
	DomainList      string `position:"Query" name:"DomainList"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeDomainAdvanceConfigsResponse is the response struct for api DescribeDomainAdvanceConfigs
type DescribeDomainAdvanceConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	DomainConfigs []DomainConfig `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeDomainAdvanceConfigsRequest creates a request to invoke DescribeDomainAdvanceConfigs API
func CreateDescribeDomainAdvanceConfigsRequest() (request *DescribeDomainAdvanceConfigsRequest) {
	request = &DescribeDomainAdvanceConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeDomainAdvanceConfigs", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainAdvanceConfigsResponse creates a response to parse from DescribeDomainAdvanceConfigs response
func CreateDescribeDomainAdvanceConfigsResponse() (response *DescribeDomainAdvanceConfigsResponse) {
	response = &DescribeDomainAdvanceConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
