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

// DescribeDomainAttackEvents invokes the ddoscoo.DescribeDomainAttackEvents API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomainattackevents.html
func (client *Client) DescribeDomainAttackEvents(request *DescribeDomainAttackEventsRequest) (response *DescribeDomainAttackEventsResponse, err error) {
	response = CreateDescribeDomainAttackEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainAttackEventsWithChan invokes the ddoscoo.DescribeDomainAttackEvents API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomainattackevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainAttackEventsWithChan(request *DescribeDomainAttackEventsRequest) (<-chan *DescribeDomainAttackEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainAttackEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainAttackEvents(request)
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

// DescribeDomainAttackEventsWithCallback invokes the ddoscoo.DescribeDomainAttackEvents API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomainattackevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainAttackEventsWithCallback(request *DescribeDomainAttackEventsRequest, callback func(response *DescribeDomainAttackEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainAttackEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainAttackEvents(request)
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

// DescribeDomainAttackEventsRequest is the request struct for api DescribeDomainAttackEvents
type DescribeDomainAttackEventsRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	Domain          string           `position:"Query" name:"Domain"`
}

// DescribeDomainAttackEventsResponse is the response struct for api DescribeDomainAttackEvents
type DescribeDomainAttackEventsResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	TotalCount         int64  `json:"TotalCount" xml:"TotalCount"`
	DomainAttackEvents []Data `json:"DomainAttackEvents" xml:"DomainAttackEvents"`
}

// CreateDescribeDomainAttackEventsRequest creates a request to invoke DescribeDomainAttackEvents API
func CreateDescribeDomainAttackEventsRequest() (request *DescribeDomainAttackEventsRequest) {
	request = &DescribeDomainAttackEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeDomainAttackEvents", "ddoscoo", "openAPI")
	return
}

// CreateDescribeDomainAttackEventsResponse creates a response to parse from DescribeDomainAttackEvents response
func CreateDescribeDomainAttackEventsResponse() (response *DescribeDomainAttackEventsResponse) {
	response = &DescribeDomainAttackEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
