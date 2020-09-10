package unimkt

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

// QueryWithPay invokes the unimkt.QueryWithPay API synchronously
// api document: https://help.aliyun.com/api/unimkt/querywithpay.html
func (client *Client) QueryWithPay(request *QueryWithPayRequest) (response *QueryWithPayResponse, err error) {
	response = CreateQueryWithPayResponse()
	err = client.DoAction(request, response)
	return
}

// QueryWithPayWithChan invokes the unimkt.QueryWithPay API asynchronously
// api document: https://help.aliyun.com/api/unimkt/querywithpay.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryWithPayWithChan(request *QueryWithPayRequest) (<-chan *QueryWithPayResponse, <-chan error) {
	responseChan := make(chan *QueryWithPayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWithPay(request)
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

// QueryWithPayWithCallback invokes the unimkt.QueryWithPay API asynchronously
// api document: https://help.aliyun.com/api/unimkt/querywithpay.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryWithPayWithCallback(request *QueryWithPayRequest, callback func(response *QueryWithPayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWithPayResponse
		var err error
		defer close(result)
		response, err = client.QueryWithPay(request)
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

// QueryWithPayRequest is the request struct for api QueryWithPay
type QueryWithPayRequest struct {
	*requests.RpcRequest
	Extra        string         `position:"Body" name:"Extra"`
	SalePrice    requests.Float `position:"Body" name:"SalePrice"`
	CommodityId  string         `position:"Body" name:"CommodityId"`
	AlipayOpenId string         `position:"Body" name:"AlipayOpenId"`
	DeviceSn     string         `position:"Body" name:"DeviceSn"`
	ChannelId    string         `position:"Body" name:"ChannelId"`
}

// QueryWithPayResponse is the response struct for api QueryWithPay
type QueryWithPayResponse struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	Msg       string `json:"Msg" xml:"Msg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryWithPayRequest creates a request to invoke QueryWithPay API
func CreateQueryWithPayRequest() (request *QueryWithPayRequest) {
	request = &QueryWithPayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-07", "QueryWithPay", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryWithPayResponse creates a response to parse from QueryWithPay response
func CreateQueryWithPayResponse() (response *QueryWithPayResponse) {
	response = &QueryWithPayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
