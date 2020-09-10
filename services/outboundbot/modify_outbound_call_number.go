package outboundbot

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

// ModifyOutboundCallNumber invokes the outboundbot.ModifyOutboundCallNumber API synchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyoutboundcallnumber.html
func (client *Client) ModifyOutboundCallNumber(request *ModifyOutboundCallNumberRequest) (response *ModifyOutboundCallNumberResponse, err error) {
	response = CreateModifyOutboundCallNumberResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyOutboundCallNumberWithChan invokes the outboundbot.ModifyOutboundCallNumber API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyoutboundcallnumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyOutboundCallNumberWithChan(request *ModifyOutboundCallNumberRequest) (<-chan *ModifyOutboundCallNumberResponse, <-chan error) {
	responseChan := make(chan *ModifyOutboundCallNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyOutboundCallNumber(request)
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

// ModifyOutboundCallNumberWithCallback invokes the outboundbot.ModifyOutboundCallNumber API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyoutboundcallnumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyOutboundCallNumberWithCallback(request *ModifyOutboundCallNumberRequest, callback func(response *ModifyOutboundCallNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyOutboundCallNumberResponse
		var err error
		defer close(result)
		response, err = client.ModifyOutboundCallNumber(request)
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

// ModifyOutboundCallNumberRequest is the request struct for api ModifyOutboundCallNumber
type ModifyOutboundCallNumberRequest struct {
	*requests.RpcRequest
	OutboundCallNumberId string           `position:"Query" name:"OutboundCallNumberId"`
	RateLimitCount       requests.Integer `position:"Query" name:"RateLimitCount"`
	Number               string           `position:"Query" name:"Number"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	RateLimitPeriod      requests.Integer `position:"Query" name:"RateLimitPeriod"`
}

// ModifyOutboundCallNumberResponse is the response struct for api ModifyOutboundCallNumber
type ModifyOutboundCallNumberResponse struct {
	*responses.BaseResponse
	Code               string             `json:"Code" xml:"Code"`
	HttpStatusCode     int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message            string             `json:"Message" xml:"Message"`
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Success            bool               `json:"Success" xml:"Success"`
	OutboundCallNumber OutboundCallNumber `json:"OutboundCallNumber" xml:"OutboundCallNumber"`
}

// CreateModifyOutboundCallNumberRequest creates a request to invoke ModifyOutboundCallNumber API
func CreateModifyOutboundCallNumberRequest() (request *ModifyOutboundCallNumberRequest) {
	request = &ModifyOutboundCallNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyOutboundCallNumber", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyOutboundCallNumberResponse creates a response to parse from ModifyOutboundCallNumber response
func CreateModifyOutboundCallNumberResponse() (response *ModifyOutboundCallNumberResponse) {
	response = &ModifyOutboundCallNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
