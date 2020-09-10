package rtc

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

// ModifyApp invokes the rtc.ModifyApp API synchronously
// api document: https://help.aliyun.com/api/rtc/modifyapp.html
func (client *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
	response = CreateModifyAppResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAppWithChan invokes the rtc.ModifyApp API asynchronously
// api document: https://help.aliyun.com/api/rtc/modifyapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAppWithChan(request *ModifyAppRequest) (<-chan *ModifyAppResponse, <-chan error) {
	responseChan := make(chan *ModifyAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApp(request)
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

// ModifyAppWithCallback invokes the rtc.ModifyApp API asynchronously
// api document: https://help.aliyun.com/api/rtc/modifyapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAppWithCallback(request *ModifyAppRequest, callback func(response *ModifyAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAppResponse
		var err error
		defer close(result)
		response, err = client.ModifyApp(request)
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

// ModifyAppRequest is the request struct for api ModifyApp
type ModifyAppRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppName string           `position:"Query" name:"AppName"`
	AppId   string           `position:"Query" name:"AppId"`
}

// ModifyAppResponse is the response struct for api ModifyApp
type ModifyAppResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAppRequest creates a request to invoke ModifyApp API
func CreateModifyAppRequest() (request *ModifyAppRequest) {
	request = &ModifyAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "ModifyApp", "rtc", "openAPI")
	return
}

// CreateModifyAppResponse creates a response to parse from ModifyApp response
func CreateModifyAppResponse() (response *ModifyAppResponse) {
	response = &ModifyAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
