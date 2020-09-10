package cloudapi

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

// ModifyApiMarketAttributes invokes the cloudapi.ModifyApiMarketAttributes API synchronously
// api document: https://help.aliyun.com/api/cloudapi/modifyapimarketattributes.html
func (client *Client) ModifyApiMarketAttributes(request *ModifyApiMarketAttributesRequest) (response *ModifyApiMarketAttributesResponse, err error) {
	response = CreateModifyApiMarketAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyApiMarketAttributesWithChan invokes the cloudapi.ModifyApiMarketAttributes API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifyapimarketattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyApiMarketAttributesWithChan(request *ModifyApiMarketAttributesRequest) (<-chan *ModifyApiMarketAttributesResponse, <-chan error) {
	responseChan := make(chan *ModifyApiMarketAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApiMarketAttributes(request)
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

// ModifyApiMarketAttributesWithCallback invokes the cloudapi.ModifyApiMarketAttributes API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifyapimarketattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyApiMarketAttributesWithCallback(request *ModifyApiMarketAttributesRequest, callback func(response *ModifyApiMarketAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiMarketAttributesResponse
		var err error
		defer close(result)
		response, err = client.ModifyApiMarketAttributes(request)
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

// ModifyApiMarketAttributesRequest is the request struct for api ModifyApiMarketAttributes
type ModifyApiMarketAttributesRequest struct {
	*requests.RpcRequest
	GroupId            string           `position:"Query" name:"GroupId"`
	SecurityToken      string           `position:"Query" name:"SecurityToken"`
	MarketChargingMode string           `position:"Query" name:"MarketChargingMode"`
	NeedCharging       requests.Boolean `position:"Query" name:"NeedCharging"`
	ApiId              string           `position:"Query" name:"ApiId"`
}

// ModifyApiMarketAttributesResponse is the response struct for api ModifyApiMarketAttributes
type ModifyApiMarketAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyApiMarketAttributesRequest creates a request to invoke ModifyApiMarketAttributes API
func CreateModifyApiMarketAttributesRequest() (request *ModifyApiMarketAttributesRequest) {
	request = &ModifyApiMarketAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApiMarketAttributes", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyApiMarketAttributesResponse creates a response to parse from ModifyApiMarketAttributes response
func CreateModifyApiMarketAttributesResponse() (response *ModifyApiMarketAttributesResponse) {
	response = &ModifyApiMarketAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
