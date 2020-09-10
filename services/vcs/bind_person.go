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

// BindPerson invokes the vcs.BindPerson API synchronously
// api document: https://help.aliyun.com/api/vcs/bindperson.html
func (client *Client) BindPerson(request *BindPersonRequest) (response *BindPersonResponse, err error) {
	response = CreateBindPersonResponse()
	err = client.DoAction(request, response)
	return
}

// BindPersonWithChan invokes the vcs.BindPerson API asynchronously
// api document: https://help.aliyun.com/api/vcs/bindperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindPersonWithChan(request *BindPersonRequest) (<-chan *BindPersonResponse, <-chan error) {
	responseChan := make(chan *BindPersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindPerson(request)
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

// BindPersonWithCallback invokes the vcs.BindPerson API asynchronously
// api document: https://help.aliyun.com/api/vcs/bindperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindPersonWithCallback(request *BindPersonRequest, callback func(response *BindPersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindPersonResponse
		var err error
		defer close(result)
		response, err = client.BindPerson(request)
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

// BindPersonRequest is the request struct for api BindPerson
type BindPersonRequest struct {
	*requests.RpcRequest
	IsvSubId           string           `position:"Body" name:"IsvSubId"`
	CorpId             string           `position:"Body" name:"CorpId"`
	PersonMatchingRate string           `position:"Body" name:"PersonMatchingRate"`
	ProfileId          requests.Integer `position:"Body" name:"ProfileId"`
	PersonId           string           `position:"Body" name:"PersonId"`
}

// BindPersonResponse is the response struct for api BindPerson
type BindPersonResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Data      bool   `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindPersonRequest creates a request to invoke BindPerson API
func CreateBindPersonRequest() (request *BindPersonRequest) {
	request = &BindPersonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "BindPerson", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindPersonResponse creates a response to parse from BindPerson response
func CreateBindPersonResponse() (response *BindPersonResponse) {
	response = &BindPersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
