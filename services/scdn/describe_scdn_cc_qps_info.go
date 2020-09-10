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

// DescribeScdnCcQpsInfo invokes the scdn.DescribeScdnCcQpsInfo API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdnccqpsinfo.html
func (client *Client) DescribeScdnCcQpsInfo(request *DescribeScdnCcQpsInfoRequest) (response *DescribeScdnCcQpsInfoResponse, err error) {
	response = CreateDescribeScdnCcQpsInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnCcQpsInfoWithChan invokes the scdn.DescribeScdnCcQpsInfo API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnccqpsinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnCcQpsInfoWithChan(request *DescribeScdnCcQpsInfoRequest) (<-chan *DescribeScdnCcQpsInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnCcQpsInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnCcQpsInfo(request)
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

// DescribeScdnCcQpsInfoWithCallback invokes the scdn.DescribeScdnCcQpsInfo API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnccqpsinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnCcQpsInfoWithCallback(request *DescribeScdnCcQpsInfoRequest, callback func(response *DescribeScdnCcQpsInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnCcQpsInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnCcQpsInfo(request)
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

// DescribeScdnCcQpsInfoRequest is the request struct for api DescribeScdnCcQpsInfo
type DescribeScdnCcQpsInfoRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnCcQpsInfoResponse is the response struct for api DescribeScdnCcQpsInfo
type DescribeScdnCcQpsInfoResponse struct {
	*responses.BaseResponse
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	Totals     Totals                            `json:"Totals" xml:"Totals"`
	Attacks    Attacks                           `json:"Attacks" xml:"Attacks"`
	TimeScopes TimeScopesInDescribeScdnCcQpsInfo `json:"TimeScopes" xml:"TimeScopes"`
}

// CreateDescribeScdnCcQpsInfoRequest creates a request to invoke DescribeScdnCcQpsInfo API
func CreateDescribeScdnCcQpsInfoRequest() (request *DescribeScdnCcQpsInfoRequest) {
	request = &DescribeScdnCcQpsInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnCcQpsInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeScdnCcQpsInfoResponse creates a response to parse from DescribeScdnCcQpsInfo response
func CreateDescribeScdnCcQpsInfoResponse() (response *DescribeScdnCcQpsInfoResponse) {
	response = &DescribeScdnCcQpsInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
