package reid

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

// PullTakeShoesEvent invokes the reid.PullTakeShoesEvent API synchronously
// api document: https://help.aliyun.com/api/reid/pulltakeshoesevent.html
func (client *Client) PullTakeShoesEvent(request *PullTakeShoesEventRequest) (response *PullTakeShoesEventResponse, err error) {
	response = CreatePullTakeShoesEventResponse()
	err = client.DoAction(request, response)
	return
}

// PullTakeShoesEventWithChan invokes the reid.PullTakeShoesEvent API asynchronously
// api document: https://help.aliyun.com/api/reid/pulltakeshoesevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PullTakeShoesEventWithChan(request *PullTakeShoesEventRequest) (<-chan *PullTakeShoesEventResponse, <-chan error) {
	responseChan := make(chan *PullTakeShoesEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PullTakeShoesEvent(request)
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

// PullTakeShoesEventWithCallback invokes the reid.PullTakeShoesEvent API asynchronously
// api document: https://help.aliyun.com/api/reid/pulltakeshoesevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PullTakeShoesEventWithCallback(request *PullTakeShoesEventRequest, callback func(response *PullTakeShoesEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PullTakeShoesEventResponse
		var err error
		defer close(result)
		response, err = client.PullTakeShoesEvent(request)
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

// PullTakeShoesEventRequest is the request struct for api PullTakeShoesEvent
type PullTakeShoesEventRequest struct {
	*requests.RpcRequest
	Date    string           `position:"Body" name:"Date"`
	StoreId requests.Integer `position:"Body" name:"StoreId"`
	SkuId   string           `position:"Body" name:"SkuId"`
}

// PullTakeShoesEventResponse is the response struct for api PullTakeShoesEvent
type PullTakeShoesEventResponse struct {
	*responses.BaseResponse
	StartTime           int64  `json:"StartTime" xml:"StartTime"`
	ErrorCode           string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage        string `json:"ErrorMessage" xml:"ErrorMessage"`
	StoreId             int64  `json:"StoreId" xml:"StoreId"`
	SkuId               string `json:"SkuId" xml:"SkuId"`
	TakeShoesEventCount int    `json:"TakeShoesEventCount" xml:"TakeShoesEventCount"`
	Message             string `json:"Message" xml:"Message"`
	Code                string `json:"Code" xml:"Code"`
	DynamicCode         string `json:"DynamicCode" xml:"DynamicCode"`
	RequestId           string `json:"RequestId" xml:"RequestId"`
	Success             bool   `json:"Success" xml:"Success"`
	DynamicMessage      string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreatePullTakeShoesEventRequest creates a request to invoke PullTakeShoesEvent API
func CreatePullTakeShoesEventRequest() (request *PullTakeShoesEventRequest) {
	request = &PullTakeShoesEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "PullTakeShoesEvent", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePullTakeShoesEventResponse creates a response to parse from PullTakeShoesEvent response
func CreatePullTakeShoesEventResponse() (response *PullTakeShoesEventResponse) {
	response = &PullTakeShoesEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
