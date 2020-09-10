package scdn

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

// DescribeScdnUserDomains invokes the scdn.DescribeScdnUserDomains API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdnuserdomains.html
func (client *Client) DescribeScdnUserDomains(request *DescribeScdnUserDomainsRequest) (response *DescribeScdnUserDomainsResponse, err error) {
	response = CreateDescribeScdnUserDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnUserDomainsWithChan invokes the scdn.DescribeScdnUserDomains API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnuserdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnUserDomainsWithChan(request *DescribeScdnUserDomainsRequest) (<-chan *DescribeScdnUserDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnUserDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnUserDomains(request)
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

// DescribeScdnUserDomainsWithCallback invokes the scdn.DescribeScdnUserDomains API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnuserdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnUserDomainsWithCallback(request *DescribeScdnUserDomainsRequest, callback func(response *DescribeScdnUserDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnUserDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnUserDomains(request)
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

// DescribeScdnUserDomainsRequest is the request struct for api DescribeScdnUserDomains
type DescribeScdnUserDomainsRequest struct {
	*requests.RpcRequest
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	CheckDomainShow  requests.Boolean `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string           `position:"Query" name:"SecurityToken"`
	ChangeEndTime    string           `position:"Query" name:"ChangeEndTime"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	FuncFilter       string           `position:"Query" name:"FuncFilter"`
	DomainName       string           `position:"Query" name:"DomainName"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	FuncId           string           `position:"Query" name:"FuncId"`
	DomainStatus     string           `position:"Query" name:"DomainStatus"`
	DomainSearchType string           `position:"Query" name:"DomainSearchType"`
	ChangeStartTime  string           `position:"Query" name:"ChangeStartTime"`
}

// DescribeScdnUserDomainsResponse is the response struct for api DescribeScdnUserDomains
type DescribeScdnUserDomainsResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Domains    Domains `json:"Domains" xml:"Domains"`
}

// CreateDescribeScdnUserDomainsRequest creates a request to invoke DescribeScdnUserDomains API
func CreateDescribeScdnUserDomainsRequest() (request *DescribeScdnUserDomainsRequest) {
	request = &DescribeScdnUserDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnUserDomains", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnUserDomainsResponse creates a response to parse from DescribeScdnUserDomains response
func CreateDescribeScdnUserDomainsResponse() (response *DescribeScdnUserDomainsResponse) {
	response = &DescribeScdnUserDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
