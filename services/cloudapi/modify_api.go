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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyApi invokes the cloudapi.ModifyApi API synchronously
// api document: https://help.aliyun.com/api/cloudapi/modifyapi.html
func (client *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
	response = CreateModifyApiResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyApiWithChan invokes the cloudapi.ModifyApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifyapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyApiWithChan(request *ModifyApiRequest) (<-chan *ModifyApiResponse, <-chan error) {
	responseChan := make(chan *ModifyApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApi(request)
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

// ModifyApiWithCallback invokes the cloudapi.ModifyApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifyapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyApiWithCallback(request *ModifyApiRequest, callback func(response *ModifyApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiResponse
		var err error
		defer close(result)
		response, err = client.ModifyApi(request)
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

// ModifyApiRequest is the request struct for api ModifyApi
type ModifyApiRequest struct {
	*requests.RpcRequest
	WebSocketApiType     string `position:"Query" name:"WebSocketApiType"`
	ErrorCodeSamples     string `position:"Query" name:"ErrorCodeSamples"`
	Description          string `position:"Query" name:"Description"`
	ConstantParameters   string `position:"Query" name:"ConstantParameters"`
	AuthType             string `position:"Query" name:"AuthType"`
	AllowSignatureMethod string `position:"Query" name:"AllowSignatureMethod"`
	ServiceParameters    string `position:"Query" name:"ServiceParameters"`
	FailResultSample     string `position:"Query" name:"FailResultSample"`
	SystemParameters     string `position:"Query" name:"SystemParameters"`
	ServiceParametersMap string `position:"Query" name:"ServiceParametersMap"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	OpenIdConnectConfig  string `position:"Query" name:"OpenIdConnectConfig"`
	RequestParameters    string `position:"Query" name:"RequestParameters"`
	ResultDescriptions   string `position:"Query" name:"ResultDescriptions"`
	Visibility           string `position:"Query" name:"Visibility"`
	GroupId              string `position:"Query" name:"GroupId"`
	ServiceConfig        string `position:"Query" name:"ServiceConfig"`
	ResultType           string `position:"Query" name:"ResultType"`
	ApiName              string `position:"Query" name:"ApiName"`
	ResultSample         string `position:"Query" name:"ResultSample"`
	RequestConfig        string `position:"Query" name:"RequestConfig"`
	ResultBodyModel      string `position:"Query" name:"ResultBodyModel"`
	ApiId                string `position:"Query" name:"ApiId"`
}

// ModifyApiResponse is the response struct for api ModifyApi
type ModifyApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyApiRequest creates a request to invoke ModifyApi API
func CreateModifyApiRequest() (request *ModifyApiRequest) {
	request = &ModifyApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApi", "apigateway", "openAPI")
	return
}

// CreateModifyApiResponse creates a response to parse from ModifyApi response
func CreateModifyApiResponse() (response *ModifyApiResponse) {
	response = &ModifyApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
