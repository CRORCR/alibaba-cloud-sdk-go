package dataworks_public

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

// RemoveProjectMemberFromRole invokes the dataworks_public.RemoveProjectMemberFromRole API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/removeprojectmemberfromrole.html
func (client *Client) RemoveProjectMemberFromRole(request *RemoveProjectMemberFromRoleRequest) (response *RemoveProjectMemberFromRoleResponse, err error) {
	response = CreateRemoveProjectMemberFromRoleResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveProjectMemberFromRoleWithChan invokes the dataworks_public.RemoveProjectMemberFromRole API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/removeprojectmemberfromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveProjectMemberFromRoleWithChan(request *RemoveProjectMemberFromRoleRequest) (<-chan *RemoveProjectMemberFromRoleResponse, <-chan error) {
	responseChan := make(chan *RemoveProjectMemberFromRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveProjectMemberFromRole(request)
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

// RemoveProjectMemberFromRoleWithCallback invokes the dataworks_public.RemoveProjectMemberFromRole API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/removeprojectmemberfromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveProjectMemberFromRoleWithCallback(request *RemoveProjectMemberFromRoleRequest, callback func(response *RemoveProjectMemberFromRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveProjectMemberFromRoleResponse
		var err error
		defer close(result)
		response, err = client.RemoveProjectMemberFromRole(request)
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

// RemoveProjectMemberFromRoleRequest is the request struct for api RemoveProjectMemberFromRole
type RemoveProjectMemberFromRoleRequest struct {
	*requests.RpcRequest
	RoleCode  string           `position:"Query" name:"RoleCode"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
	UserId    string           `position:"Query" name:"UserId"`
}

// RemoveProjectMemberFromRoleResponse is the response struct for api RemoveProjectMemberFromRole
type RemoveProjectMemberFromRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveProjectMemberFromRoleRequest creates a request to invoke RemoveProjectMemberFromRole API
func CreateRemoveProjectMemberFromRoleRequest() (request *RemoveProjectMemberFromRoleRequest) {
	request = &RemoveProjectMemberFromRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "RemoveProjectMemberFromRole", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveProjectMemberFromRoleResponse creates a response to parse from RemoveProjectMemberFromRole response
func CreateRemoveProjectMemberFromRoleResponse() (response *RemoveProjectMemberFromRoleResponse) {
	response = &RemoveProjectMemberFromRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
