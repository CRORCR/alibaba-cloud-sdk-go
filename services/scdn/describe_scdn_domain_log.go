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

// DescribeScdnDomainLog invokes the scdn.DescribeScdnDomainLog API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainlog.html
func (client *Client) DescribeScdnDomainLog(request *DescribeScdnDomainLogRequest) (response *DescribeScdnDomainLogResponse, err error) {
	response = CreateDescribeScdnDomainLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainLogWithChan invokes the scdn.DescribeScdnDomainLog API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainLogWithChan(request *DescribeScdnDomainLogRequest) (<-chan *DescribeScdnDomainLogResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainLog(request)
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

// DescribeScdnDomainLogWithCallback invokes the scdn.DescribeScdnDomainLog API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainLogWithCallback(request *DescribeScdnDomainLogRequest, callback func(response *DescribeScdnDomainLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainLog(request)
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

// DescribeScdnDomainLogRequest is the request struct for api DescribeScdnDomainLog
type DescribeScdnDomainLogRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnDomainLogResponse is the response struct for api DescribeScdnDomainLog
type DescribeScdnDomainLogResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	DomainName       string           `json:"DomainName" xml:"DomainName"`
	DomainLogDetails DomainLogDetails `json:"DomainLogDetails" xml:"DomainLogDetails"`
}

// CreateDescribeScdnDomainLogRequest creates a request to invoke DescribeScdnDomainLog API
func CreateDescribeScdnDomainLogRequest() (request *DescribeScdnDomainLogRequest) {
	request = &DescribeScdnDomainLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainLog", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainLogResponse creates a response to parse from DescribeScdnDomainLog response
func CreateDescribeScdnDomainLogResponse() (response *DescribeScdnDomainLogResponse) {
	response = &DescribeScdnDomainLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
