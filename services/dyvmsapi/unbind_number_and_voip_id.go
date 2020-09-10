package dyvmsapi

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

// UnbindNumberAndVoipId invokes the dyvmsapi.UnbindNumberAndVoipId API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/unbindnumberandvoipid.html
func (client *Client) UnbindNumberAndVoipId(request *UnbindNumberAndVoipIdRequest) (response *UnbindNumberAndVoipIdResponse, err error) {
	response = CreateUnbindNumberAndVoipIdResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindNumberAndVoipIdWithChan invokes the dyvmsapi.UnbindNumberAndVoipId API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/unbindnumberandvoipid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindNumberAndVoipIdWithChan(request *UnbindNumberAndVoipIdRequest) (<-chan *UnbindNumberAndVoipIdResponse, <-chan error) {
	responseChan := make(chan *UnbindNumberAndVoipIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindNumberAndVoipId(request)
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

// UnbindNumberAndVoipIdWithCallback invokes the dyvmsapi.UnbindNumberAndVoipId API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/unbindnumberandvoipid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindNumberAndVoipIdWithCallback(request *UnbindNumberAndVoipIdRequest, callback func(response *UnbindNumberAndVoipIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindNumberAndVoipIdResponse
		var err error
		defer close(result)
		response, err = client.UnbindNumberAndVoipId(request)
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

// UnbindNumberAndVoipIdRequest is the request struct for api UnbindNumberAndVoipId
type UnbindNumberAndVoipIdRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VoipId               string           `position:"Query" name:"VoipId"`
}

// UnbindNumberAndVoipIdResponse is the response struct for api UnbindNumberAndVoipId
type UnbindNumberAndVoipIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Module    string `json:"Module" xml:"Module"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUnbindNumberAndVoipIdRequest creates a request to invoke UnbindNumberAndVoipId API
func CreateUnbindNumberAndVoipIdRequest() (request *UnbindNumberAndVoipIdRequest) {
	request = &UnbindNumberAndVoipIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "UnbindNumberAndVoipId", "dyvms", "openAPI")
	return
}

// CreateUnbindNumberAndVoipIdResponse creates a response to parse from UnbindNumberAndVoipId response
func CreateUnbindNumberAndVoipIdResponse() (response *UnbindNumberAndVoipIdResponse) {
	response = &UnbindNumberAndVoipIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
