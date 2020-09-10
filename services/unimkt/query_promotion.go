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

// QueryPromotion invokes the unimkt.QueryPromotion API synchronously
// api document: https://help.aliyun.com/api/unimkt/querypromotion.html
func (client *Client) QueryPromotion(request *QueryPromotionRequest) (response *QueryPromotionResponse, err error) {
	response = CreateQueryPromotionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPromotionWithChan invokes the unimkt.QueryPromotion API asynchronously
// api document: https://help.aliyun.com/api/unimkt/querypromotion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPromotionWithChan(request *QueryPromotionRequest) (<-chan *QueryPromotionResponse, <-chan error) {
	responseChan := make(chan *QueryPromotionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPromotion(request)
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

// QueryPromotionWithCallback invokes the unimkt.QueryPromotion API asynchronously
// api document: https://help.aliyun.com/api/unimkt/querypromotion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPromotionWithCallback(request *QueryPromotionRequest, callback func(response *QueryPromotionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPromotionResponse
		var err error
		defer close(result)
		response, err = client.QueryPromotion(request)
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

// QueryPromotionRequest is the request struct for api QueryPromotion
type QueryPromotionRequest struct {
	*requests.RpcRequest
	Extra        string `position:"Body" name:"Extra"`
	AlipayOpenId string `position:"Body" name:"AlipayOpenId"`
	ChannelId    string `position:"Body" name:"ChannelId"`
}

// QueryPromotionResponse is the response struct for api QueryPromotion
type QueryPromotionResponse struct {
	*responses.BaseResponse
	Status      bool   `json:"Status" xml:"Status"`
	Msg         string `json:"Msg" xml:"Msg"`
	ErrorCode   string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Url         string `json:"Url" xml:"Url"`
	UnionAmount string `json:"UnionAmount" xml:"UnionAmount"`
}

// CreateQueryPromotionRequest creates a request to invoke QueryPromotion API
func CreateQueryPromotionRequest() (request *QueryPromotionRequest) {
	request = &QueryPromotionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-07", "QueryPromotion", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryPromotionResponse creates a response to parse from QueryPromotion response
func CreateQueryPromotionResponse() (response *QueryPromotionResponse) {
	response = &QueryPromotionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
