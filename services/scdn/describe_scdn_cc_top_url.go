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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeScdnCcTopUrl invokes the scdn.DescribeScdnCcTopUrl API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdncctopurl.html
func (client *Client) DescribeScdnCcTopUrl(request *DescribeScdnCcTopUrlRequest) (response *DescribeScdnCcTopUrlResponse, err error) {
	response = CreateDescribeScdnCcTopUrlResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnCcTopUrlWithChan invokes the scdn.DescribeScdnCcTopUrl API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdncctopurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnCcTopUrlWithChan(request *DescribeScdnCcTopUrlRequest) (<-chan *DescribeScdnCcTopUrlResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnCcTopUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnCcTopUrl(request)
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

// DescribeScdnCcTopUrlWithCallback invokes the scdn.DescribeScdnCcTopUrl API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdncctopurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnCcTopUrlWithCallback(request *DescribeScdnCcTopUrlRequest, callback func(response *DescribeScdnCcTopUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnCcTopUrlResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnCcTopUrl(request)
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

// DescribeScdnCcTopUrlRequest is the request struct for api DescribeScdnCcTopUrl
type DescribeScdnCcTopUrlRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber string           `position:"Query" name:"PageNumber"`
	PageSize   string           `position:"Query" name:"PageSize"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnCcTopUrlResponse is the response struct for api DescribeScdnCcTopUrl
type DescribeScdnCcTopUrlResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Total             string            `json:"Total" xml:"Total"`
	DomainName        string            `json:"DomainName" xml:"DomainName"`
	AttackUrlDataList AttackUrlDataList `json:"AttackUrlDataList" xml:"AttackUrlDataList"`
}

// CreateDescribeScdnCcTopUrlRequest creates a request to invoke DescribeScdnCcTopUrl API
func CreateDescribeScdnCcTopUrlRequest() (request *DescribeScdnCcTopUrlRequest) {
	request = &DescribeScdnCcTopUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnCcTopUrl", "", "")
	return
}

// CreateDescribeScdnCcTopUrlResponse creates a response to parse from DescribeScdnCcTopUrl response
func CreateDescribeScdnCcTopUrlResponse() (response *DescribeScdnCcTopUrlResponse) {
	response = &DescribeScdnCcTopUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
