package imm

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

// RefreshOfficeEditToken invokes the imm.RefreshOfficeEditToken API synchronously
// api document: https://help.aliyun.com/api/imm/refreshofficeedittoken.html
func (client *Client) RefreshOfficeEditToken(request *RefreshOfficeEditTokenRequest) (response *RefreshOfficeEditTokenResponse, err error) {
	response = CreateRefreshOfficeEditTokenResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshOfficeEditTokenWithChan invokes the imm.RefreshOfficeEditToken API asynchronously
// api document: https://help.aliyun.com/api/imm/refreshofficeedittoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshOfficeEditTokenWithChan(request *RefreshOfficeEditTokenRequest) (<-chan *RefreshOfficeEditTokenResponse, <-chan error) {
	responseChan := make(chan *RefreshOfficeEditTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshOfficeEditToken(request)
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

// RefreshOfficeEditTokenWithCallback invokes the imm.RefreshOfficeEditToken API asynchronously
// api document: https://help.aliyun.com/api/imm/refreshofficeedittoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshOfficeEditTokenWithCallback(request *RefreshOfficeEditTokenRequest, callback func(response *RefreshOfficeEditTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshOfficeEditTokenResponse
		var err error
		defer close(result)
		response, err = client.RefreshOfficeEditToken(request)
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

// RefreshOfficeEditTokenRequest is the request struct for api RefreshOfficeEditToken
type RefreshOfficeEditTokenRequest struct {
	*requests.RpcRequest
	Project      string `position:"Query" name:"Project"`
	AccessToken  string `position:"Query" name:"AccessToken"`
	RefreshToken string `position:"Query" name:"RefreshToken"`
}

// RefreshOfficeEditTokenResponse is the response struct for api RefreshOfficeEditToken
type RefreshOfficeEditTokenResponse struct {
	*responses.BaseResponse
	AccessToken             string `json:"AccessToken" xml:"AccessToken"`
	AccessTokenExpiredTime  string `json:"AccessTokenExpiredTime" xml:"AccessTokenExpiredTime"`
	RefreshToken            string `json:"RefreshToken" xml:"RefreshToken"`
	RefreshTokenExpiredTime string `json:"RefreshTokenExpiredTime" xml:"RefreshTokenExpiredTime"`
	RequestId               string `json:"RequestId" xml:"RequestId"`
}

// CreateRefreshOfficeEditTokenRequest creates a request to invoke RefreshOfficeEditToken API
func CreateRefreshOfficeEditTokenRequest() (request *RefreshOfficeEditTokenRequest) {
	request = &RefreshOfficeEditTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "RefreshOfficeEditToken", "", "")
	request.Method = requests.POST
	return
}

// CreateRefreshOfficeEditTokenResponse creates a response to parse from RefreshOfficeEditToken response
func CreateRefreshOfficeEditTokenResponse() (response *RefreshOfficeEditTokenResponse) {
	response = &RefreshOfficeEditTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
