package ccc

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

// DialEx invokes the ccc.DialEx API synchronously
// api document: https://help.aliyun.com/api/ccc/dialex.html
func (client *Client) DialEx(request *DialExRequest) (response *DialExResponse, err error) {
	response = CreateDialExResponse()
	err = client.DoAction(request, response)
	return
}

// DialExWithChan invokes the ccc.DialEx API asynchronously
// api document: https://help.aliyun.com/api/ccc/dialex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DialExWithChan(request *DialExRequest) (<-chan *DialExResponse, <-chan error) {
	responseChan := make(chan *DialExResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DialEx(request)
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

// DialExWithCallback invokes the ccc.DialEx API asynchronously
// api document: https://help.aliyun.com/api/ccc/dialex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DialExWithCallback(request *DialExRequest, callback func(response *DialExResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DialExResponse
		var err error
		defer close(result)
		response, err = client.DialEx(request)
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

// DialExRequest is the request struct for api DialEx
type DialExRequest struct {
	*requests.RpcRequest
	Callee     string           `position:"Query" name:"Callee"`
	RoutPoint  string           `position:"Query" name:"RoutPoint"`
	Caller     string           `position:"Query" name:"Caller"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Provider   string           `position:"Query" name:"Provider"`
	AnswerMode requests.Integer `position:"Query" name:"AnswerMode"`
}

// DialExResponse is the response struct for api DialEx
type DialExResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	StatusCode     string `json:"StatusCode" xml:"StatusCode"`
	StatusDesc     string `json:"StatusDesc" xml:"StatusDesc"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	TimeStamp      string `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateDialExRequest creates a request to invoke DialEx API
func CreateDialExRequest() (request *DialExRequest) {
	request = &DialExRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "DialEx", "", "")
	return
}

// CreateDialExResponse creates a response to parse from DialEx response
func CreateDialExResponse() (response *DialExResponse) {
	response = &DialExResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
