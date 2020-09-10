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

// RemoveUsers invokes the ccc.RemoveUsers API synchronously
// api document: https://help.aliyun.com/api/ccc/removeusers.html
func (client *Client) RemoveUsers(request *RemoveUsersRequest) (response *RemoveUsersResponse, err error) {
	response = CreateRemoveUsersResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveUsersWithChan invokes the ccc.RemoveUsers API asynchronously
// api document: https://help.aliyun.com/api/ccc/removeusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveUsersWithChan(request *RemoveUsersRequest) (<-chan *RemoveUsersResponse, <-chan error) {
	responseChan := make(chan *RemoveUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveUsers(request)
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

// RemoveUsersWithCallback invokes the ccc.RemoveUsers API asynchronously
// api document: https://help.aliyun.com/api/ccc/removeusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveUsersWithCallback(request *RemoveUsersRequest, callback func(response *RemoveUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveUsersResponse
		var err error
		defer close(result)
		response, err = client.RemoveUsers(request)
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

// RemoveUsersRequest is the request struct for api RemoveUsers
type RemoveUsersRequest struct {
	*requests.RpcRequest
	InstanceId string    `position:"Query" name:"InstanceId"`
	UserId     *[]string `position:"Query" name:"UserId"  type:"Repeated"`
}

// RemoveUsersResponse is the response struct for api RemoveUsers
type RemoveUsersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateRemoveUsersRequest creates a request to invoke RemoveUsers API
func CreateRemoveUsersRequest() (request *RemoveUsersRequest) {
	request = &RemoveUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "RemoveUsers", "", "")
	return
}

// CreateRemoveUsersResponse creates a response to parse from RemoveUsers response
func CreateRemoveUsersResponse() (response *RemoveUsersResponse) {
	response = &RemoveUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
