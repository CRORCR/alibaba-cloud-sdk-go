package vcs

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

// UnbindPerson invokes the vcs.UnbindPerson API synchronously
// api document: https://help.aliyun.com/api/vcs/unbindperson.html
func (client *Client) UnbindPerson(request *UnbindPersonRequest) (response *UnbindPersonResponse, err error) {
	response = CreateUnbindPersonResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindPersonWithChan invokes the vcs.UnbindPerson API asynchronously
// api document: https://help.aliyun.com/api/vcs/unbindperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindPersonWithChan(request *UnbindPersonRequest) (<-chan *UnbindPersonResponse, <-chan error) {
	responseChan := make(chan *UnbindPersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindPerson(request)
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

// UnbindPersonWithCallback invokes the vcs.UnbindPerson API asynchronously
// api document: https://help.aliyun.com/api/vcs/unbindperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindPersonWithCallback(request *UnbindPersonRequest, callback func(response *UnbindPersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindPersonResponse
		var err error
		defer close(result)
		response, err = client.UnbindPerson(request)
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

// UnbindPersonRequest is the request struct for api UnbindPerson
type UnbindPersonRequest struct {
	*requests.RpcRequest
	IsvSubId  string           `position:"Body" name:"IsvSubId"`
	CorpId    string           `position:"Body" name:"CorpId"`
	ProfileId requests.Integer `position:"Body" name:"ProfileId"`
}

// UnbindPersonResponse is the response struct for api UnbindPerson
type UnbindPersonResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Data      bool   `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindPersonRequest creates a request to invoke UnbindPerson API
func CreateUnbindPersonRequest() (request *UnbindPersonRequest) {
	request = &UnbindPersonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "UnbindPerson", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindPersonResponse creates a response to parse from UnbindPerson response
func CreateUnbindPersonResponse() (response *UnbindPersonResponse) {
	response = &UnbindPersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
