package democenter

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

// GetDemoTrialAuth invokes the democenter.GetDemoTrialAuth API synchronously
// api document: https://help.aliyun.com/api/democenter/getdemotrialauth.html
func (client *Client) GetDemoTrialAuth(request *GetDemoTrialAuthRequest) (response *GetDemoTrialAuthResponse, err error) {
	response = CreateGetDemoTrialAuthResponse()
	err = client.DoAction(request, response)
	return
}

// GetDemoTrialAuthWithChan invokes the democenter.GetDemoTrialAuth API asynchronously
// api document: https://help.aliyun.com/api/democenter/getdemotrialauth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDemoTrialAuthWithChan(request *GetDemoTrialAuthRequest) (<-chan *GetDemoTrialAuthResponse, <-chan error) {
	responseChan := make(chan *GetDemoTrialAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDemoTrialAuth(request)
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

// GetDemoTrialAuthWithCallback invokes the democenter.GetDemoTrialAuth API asynchronously
// api document: https://help.aliyun.com/api/democenter/getdemotrialauth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDemoTrialAuthWithCallback(request *GetDemoTrialAuthRequest, callback func(response *GetDemoTrialAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDemoTrialAuthResponse
		var err error
		defer close(result)
		response, err = client.GetDemoTrialAuth(request)
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

// GetDemoTrialAuthRequest is the request struct for api GetDemoTrialAuth
type GetDemoTrialAuthRequest struct {
	*requests.RpcRequest
	DemoId requests.Integer `position:"Body" name:"DemoId"`
}

// GetDemoTrialAuthResponse is the response struct for api GetDemoTrialAuth
type GetDemoTrialAuthResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Authorization string `json:"Authorization" xml:"Authorization"`
}

// CreateGetDemoTrialAuthRequest creates a request to invoke GetDemoTrialAuth API
func CreateGetDemoTrialAuthRequest() (request *GetDemoTrialAuthRequest) {
	request = &GetDemoTrialAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DemoCenter", "2020-01-21", "GetDemoTrialAuth", "", "")
	return
}

// CreateGetDemoTrialAuthResponse creates a response to parse from GetDemoTrialAuth response
func CreateGetDemoTrialAuthResponse() (response *GetDemoTrialAuthResponse) {
	response = &GetDemoTrialAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
