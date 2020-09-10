package alidns

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

// DescribeDohSubDomainStatisticsSummary invokes the alidns.DescribeDohSubDomainStatisticsSummary API synchronously
// api document: https://help.aliyun.com/api/alidns/describedohsubdomainstatisticssummary.html
func (client *Client) DescribeDohSubDomainStatisticsSummary(request *DescribeDohSubDomainStatisticsSummaryRequest) (response *DescribeDohSubDomainStatisticsSummaryResponse, err error) {
	response = CreateDescribeDohSubDomainStatisticsSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDohSubDomainStatisticsSummaryWithChan invokes the alidns.DescribeDohSubDomainStatisticsSummary API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedohsubdomainstatisticssummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDohSubDomainStatisticsSummaryWithChan(request *DescribeDohSubDomainStatisticsSummaryRequest) (<-chan *DescribeDohSubDomainStatisticsSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeDohSubDomainStatisticsSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDohSubDomainStatisticsSummary(request)
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

// DescribeDohSubDomainStatisticsSummaryWithCallback invokes the alidns.DescribeDohSubDomainStatisticsSummary API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedohsubdomainstatisticssummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDohSubDomainStatisticsSummaryWithCallback(request *DescribeDohSubDomainStatisticsSummaryRequest, callback func(response *DescribeDohSubDomainStatisticsSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDohSubDomainStatisticsSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeDohSubDomainStatisticsSummary(request)
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

// DescribeDohSubDomainStatisticsSummaryRequest is the request struct for api DescribeDohSubDomainStatisticsSummary
type DescribeDohSubDomainStatisticsSummaryRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OrderBy    string           `position:"Query" name:"OrderBy"`
	StartDate  string           `position:"Query" name:"StartDate"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	EndDate    string           `position:"Query" name:"EndDate"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	SubDomain  string           `position:"Query" name:"SubDomain"`
	Lang       string           `position:"Query" name:"Lang"`
	Direction  string           `position:"Query" name:"Direction"`
}

// DescribeDohSubDomainStatisticsSummaryResponse is the response struct for api DescribeDohSubDomainStatisticsSummary
type DescribeDohSubDomainStatisticsSummaryResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TotalItems int         `json:"TotalItems" xml:"TotalItems"`
	TotalPages int         `json:"TotalPages" xml:"TotalPages"`
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	Statistics []Statistic `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeDohSubDomainStatisticsSummaryRequest creates a request to invoke DescribeDohSubDomainStatisticsSummary API
func CreateDescribeDohSubDomainStatisticsSummaryRequest() (request *DescribeDohSubDomainStatisticsSummaryRequest) {
	request = &DescribeDohSubDomainStatisticsSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDohSubDomainStatisticsSummary", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDohSubDomainStatisticsSummaryResponse creates a response to parse from DescribeDohSubDomainStatisticsSummary response
func CreateDescribeDohSubDomainStatisticsSummaryResponse() (response *DescribeDohSubDomainStatisticsSummaryResponse) {
	response = &DescribeDohSubDomainStatisticsSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
