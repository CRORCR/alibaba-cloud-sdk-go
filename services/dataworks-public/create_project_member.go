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

// CreateProjectMember invokes the dataworks_public.CreateProjectMember API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/createprojectmember.html
func (client *Client) CreateProjectMember(request *CreateProjectMemberRequest) (response *CreateProjectMemberResponse, err error) {
	response = CreateCreateProjectMemberResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProjectMemberWithChan invokes the dataworks_public.CreateProjectMember API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/createprojectmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProjectMemberWithChan(request *CreateProjectMemberRequest) (<-chan *CreateProjectMemberResponse, <-chan error) {
	responseChan := make(chan *CreateProjectMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProjectMember(request)
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

// CreateProjectMemberWithCallback invokes the dataworks_public.CreateProjectMember API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/createprojectmember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProjectMemberWithCallback(request *CreateProjectMemberRequest, callback func(response *CreateProjectMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProjectMemberResponse
		var err error
		defer close(result)
		response, err = client.CreateProjectMember(request)
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

// CreateProjectMemberRequest is the request struct for api CreateProjectMember
type CreateProjectMemberRequest struct {
	*requests.RpcRequest
	RoleCode    string           `position:"Query" name:"RoleCode"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
	UserId      string           `position:"Query" name:"UserId"`
}

// CreateProjectMemberResponse is the response struct for api CreateProjectMember
type CreateProjectMemberResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateProjectMemberRequest creates a request to invoke CreateProjectMember API
func CreateCreateProjectMemberRequest() (request *CreateProjectMemberRequest) {
	request = &CreateProjectMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateProjectMember", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProjectMemberResponse creates a response to parse from CreateProjectMember response
func CreateCreateProjectMemberResponse() (response *CreateProjectMemberResponse) {
	response = &CreateProjectMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
