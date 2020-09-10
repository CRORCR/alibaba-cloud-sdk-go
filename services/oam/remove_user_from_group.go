package oam

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

// RemoveUserFromGroup invokes the oam.RemoveUserFromGroup API synchronously
// api document: https://help.aliyun.com/api/oam/removeuserfromgroup.html
func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
	response = CreateRemoveUserFromGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveUserFromGroupWithChan invokes the oam.RemoveUserFromGroup API asynchronously
// api document: https://help.aliyun.com/api/oam/removeuserfromgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveUserFromGroupWithChan(request *RemoveUserFromGroupRequest) (<-chan *RemoveUserFromGroupResponse, <-chan error) {
	responseChan := make(chan *RemoveUserFromGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveUserFromGroup(request)
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

// RemoveUserFromGroupWithCallback invokes the oam.RemoveUserFromGroup API asynchronously
// api document: https://help.aliyun.com/api/oam/removeuserfromgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveUserFromGroupWithCallback(request *RemoveUserFromGroupRequest, callback func(response *RemoveUserFromGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveUserFromGroupResponse
		var err error
		defer close(result)
		response, err = client.RemoveUserFromGroup(request)
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

// RemoveUserFromGroupRequest is the request struct for api RemoveUserFromGroup
type RemoveUserFromGroupRequest struct {
	*requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
	UserName  string `position:"Query" name:"UserName"`
}

// RemoveUserFromGroupResponse is the response struct for api RemoveUserFromGroup
type RemoveUserFromGroupResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateRemoveUserFromGroupRequest creates a request to invoke RemoveUserFromGroup API
func CreateRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
	request = &RemoveUserFromGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "RemoveUserFromGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveUserFromGroupResponse creates a response to parse from RemoveUserFromGroup response
func CreateRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
	response = &RemoveUserFromGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
