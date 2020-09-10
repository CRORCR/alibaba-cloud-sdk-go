package cloudwf

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

// ProfileTrade invokes the cloudwf.ProfileTrade API synchronously
// api document: https://help.aliyun.com/api/cloudwf/profiletrade.html
func (client *Client) ProfileTrade(request *ProfileTradeRequest) (response *ProfileTradeResponse, err error) {
	response = CreateProfileTradeResponse()
	err = client.DoAction(request, response)
	return
}

// ProfileTradeWithChan invokes the cloudwf.ProfileTrade API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profiletrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileTradeWithChan(request *ProfileTradeRequest) (<-chan *ProfileTradeResponse, <-chan error) {
	responseChan := make(chan *ProfileTradeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileTrade(request)
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

// ProfileTradeWithCallback invokes the cloudwf.ProfileTrade API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profiletrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileTradeWithCallback(request *ProfileTradeRequest, callback func(response *ProfileTradeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileTradeResponse
		var err error
		defer close(result)
		response, err = client.ProfileTrade(request)
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

// ProfileTradeRequest is the request struct for api ProfileTrade
type ProfileTradeRequest struct {
	*requests.RpcRequest
	BeginDate string           `position:"Query" name:"BeginDate"`
	EndDate   string           `position:"Query" name:"EndDate"`
	DataType  requests.Integer `position:"Query" name:"DataType"`
	Gsid      requests.Integer `position:"Query" name:"Gsid"`
}

// ProfileTradeResponse is the response struct for api ProfileTrade
type ProfileTradeResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateProfileTradeRequest creates a request to invoke ProfileTrade API
func CreateProfileTradeRequest() (request *ProfileTradeRequest) {
	request = &ProfileTradeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileTrade", "cloudwf", "openAPI")
	return
}

// CreateProfileTradeResponse creates a response to parse from ProfileTrade response
func CreateProfileTradeResponse() (response *ProfileTradeResponse) {
	response = &ProfileTradeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
