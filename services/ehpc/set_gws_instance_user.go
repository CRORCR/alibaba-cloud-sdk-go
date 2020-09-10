package ehpc

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

// SetGWSInstanceUser invokes the ehpc.SetGWSInstanceUser API synchronously
// api document: https://help.aliyun.com/api/ehpc/setgwsinstanceuser.html
func (client *Client) SetGWSInstanceUser(request *SetGWSInstanceUserRequest) (response *SetGWSInstanceUserResponse, err error) {
	response = CreateSetGWSInstanceUserResponse()
	err = client.DoAction(request, response)
	return
}

// SetGWSInstanceUserWithChan invokes the ehpc.SetGWSInstanceUser API asynchronously
// api document: https://help.aliyun.com/api/ehpc/setgwsinstanceuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetGWSInstanceUserWithChan(request *SetGWSInstanceUserRequest) (<-chan *SetGWSInstanceUserResponse, <-chan error) {
	responseChan := make(chan *SetGWSInstanceUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetGWSInstanceUser(request)
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

// SetGWSInstanceUserWithCallback invokes the ehpc.SetGWSInstanceUser API asynchronously
// api document: https://help.aliyun.com/api/ehpc/setgwsinstanceuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetGWSInstanceUserWithCallback(request *SetGWSInstanceUserRequest, callback func(response *SetGWSInstanceUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetGWSInstanceUserResponse
		var err error
		defer close(result)
		response, err = client.SetGWSInstanceUser(request)
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

// SetGWSInstanceUserRequest is the request struct for api SetGWSInstanceUser
type SetGWSInstanceUserRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserUid    string `position:"Query" name:"UserUid"`
	UserName   string `position:"Query" name:"UserName"`
}

// SetGWSInstanceUserResponse is the response struct for api SetGWSInstanceUser
type SetGWSInstanceUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetGWSInstanceUserRequest creates a request to invoke SetGWSInstanceUser API
func CreateSetGWSInstanceUserRequest() (request *SetGWSInstanceUserRequest) {
	request = &SetGWSInstanceUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "SetGWSInstanceUser", "", "")
	request.Method = requests.GET
	return
}

// CreateSetGWSInstanceUserResponse creates a response to parse from SetGWSInstanceUser response
func CreateSetGWSInstanceUserResponse() (response *SetGWSInstanceUserResponse) {
	response = &SetGWSInstanceUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
