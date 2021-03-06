package emr

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

// UpdateUser invokes the emr.UpdateUser API synchronously
// api document: https://help.aliyun.com/api/emr/updateuser.html
func (client *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
	response = CreateUpdateUserResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserWithChan invokes the emr.UpdateUser API asynchronously
// api document: https://help.aliyun.com/api/emr/updateuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateUserWithChan(request *UpdateUserRequest) (<-chan *UpdateUserResponse, <-chan error) {
	responseChan := make(chan *UpdateUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUser(request)
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

// UpdateUserWithCallback invokes the emr.UpdateUser API asynchronously
// api document: https://help.aliyun.com/api/emr/updateuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateUserWithCallback(request *UpdateUserRequest, callback func(response *UpdateUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserResponse
		var err error
		defer close(result)
		response, err = client.UpdateUser(request)
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

// UpdateUserRequest is the request struct for api UpdateUser
type UpdateUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                  `position:"Query" name:"ResourceOwnerId"`
	UserType             string                            `position:"Query" name:"UserType"`
	Description          string                            `position:"Query" name:"Description"`
	UserAccountParamList *[]UpdateUserUserAccountParamList `position:"Query" name:"UserAccountParamList"  type:"Repeated"`
	GroupIdList          *[]string                         `position:"Query" name:"GroupIdList"  type:"Repeated"`
	RoleIdList           *[]string                         `position:"Query" name:"RoleIdList"  type:"Repeated"`
	AliyunUserId         string                            `position:"Query" name:"AliyunUserId"`
	UserName             string                            `position:"Query" name:"UserName"`
	Status               string                            `position:"Query" name:"Status"`
}

// UpdateUserUserAccountParamList is a repeated param struct in UpdateUserRequest
type UpdateUserUserAccountParamList struct {
	AccountType     string `name:"AccountType"`
	AuthType        string `name:"AuthType"`
	AccountPassword string `name:"AccountPassword"`
}

// UpdateUserResponse is the response struct for api UpdateUser
type UpdateUserResponse struct {
	*responses.BaseResponse
	Paging    bool   `json:"Paging" xml:"Paging"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateUpdateUserRequest creates a request to invoke UpdateUser API
func CreateUpdateUserRequest() (request *UpdateUserRequest) {
	request = &UpdateUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "UpdateUser", "emr", "openAPI")
	return
}

// CreateUpdateUserResponse creates a response to parse from UpdateUser response
func CreateUpdateUserResponse() (response *UpdateUserResponse) {
	response = &UpdateUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
