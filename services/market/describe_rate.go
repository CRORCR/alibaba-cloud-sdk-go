package market

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

// DescribeRate invokes the market.DescribeRate API synchronously
// api document: https://help.aliyun.com/api/market/describerate.html
func (client *Client) DescribeRate(request *DescribeRateRequest) (response *DescribeRateResponse, err error) {
	response = CreateDescribeRateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRateWithChan invokes the market.DescribeRate API asynchronously
// api document: https://help.aliyun.com/api/market/describerate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRateWithChan(request *DescribeRateRequest) (<-chan *DescribeRateResponse, <-chan error) {
	responseChan := make(chan *DescribeRateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRate(request)
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

// DescribeRateWithCallback invokes the market.DescribeRate API asynchronously
// api document: https://help.aliyun.com/api/market/describerate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRateWithCallback(request *DescribeRateRequest, callback func(response *DescribeRateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRateResponse
		var err error
		defer close(result)
		response, err = client.DescribeRate(request)
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

// DescribeRateRequest is the request struct for api DescribeRate
type DescribeRateRequest struct {
	*requests.RpcRequest
	OrderId string `position:"Query" name:"OrderId"`
}

// DescribeRateResponse is the response struct for api DescribeRate
type DescribeRateResponse struct {
	*responses.BaseResponse
	Id                       int64  `json:"Id" xml:"Id"`
	OrderId                  string `json:"OrderId" xml:"OrderId"`
	InstanceId               string `json:"InstanceId" xml:"InstanceId"`
	AliUid                   int64  `json:"AliUid" xml:"AliUid"`
	ProductId                string `json:"ProductId" xml:"ProductId"`
	Score                    string `json:"Score" xml:"Score"`
	Content                  string `json:"Content" xml:"Content"`
	GmtCreated               int64  `json:"GmtCreated" xml:"GmtCreated"`
	Explaintion              string `json:"Explaintion" xml:"Explaintion"`
	GmtExplaintion           int64  `json:"GmtExplaintion" xml:"GmtExplaintion"`
	AdditionalContent        string `json:"AdditionalContent" xml:"AdditionalContent"`
	GmtAdditional            int64  `json:"GmtAdditional" xml:"GmtAdditional"`
	AdditionalExplaintion    string `json:"AdditionalExplaintion" xml:"AdditionalExplaintion"`
	GmtAdditionalExplaintion int64  `json:"GmtAdditionalExplaintion" xml:"GmtAdditionalExplaintion"`
	Type                     string `json:"Type" xml:"Type"`
	RequestId                string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeRateRequest creates a request to invoke DescribeRate API
func CreateDescribeRateRequest() (request *DescribeRateRequest) {
	request = &DescribeRateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeRate", "yunmarket", "openAPI")
	return
}

// CreateDescribeRateResponse creates a response to parse from DescribeRate response
func CreateDescribeRateResponse() (response *DescribeRateResponse) {
	response = &DescribeRateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
