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

// CreateSkillGroup invokes the ccc.CreateSkillGroup API synchronously
// api document: https://help.aliyun.com/api/ccc/createskillgroup.html
func (client *Client) CreateSkillGroup(request *CreateSkillGroupRequest) (response *CreateSkillGroupResponse, err error) {
	response = CreateCreateSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSkillGroupWithChan invokes the ccc.CreateSkillGroup API asynchronously
// api document: https://help.aliyun.com/api/ccc/createskillgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSkillGroupWithChan(request *CreateSkillGroupRequest) (<-chan *CreateSkillGroupResponse, <-chan error) {
	responseChan := make(chan *CreateSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSkillGroup(request)
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

// CreateSkillGroupWithCallback invokes the ccc.CreateSkillGroup API asynchronously
// api document: https://help.aliyun.com/api/ccc/createskillgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSkillGroupWithCallback(request *CreateSkillGroupRequest, callback func(response *CreateSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateSkillGroup(request)
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

// CreateSkillGroupRequest is the request struct for api CreateSkillGroup
type CreateSkillGroupRequest struct {
	*requests.RpcRequest
	AllowPrivateOutboundNumber requests.Boolean `position:"Query" name:"AllowPrivateOutboundNumber"`
	Description                string           `position:"Query" name:"Description"`
	RoutingStrategy            string           `position:"Query" name:"RoutingStrategy"`
	UserId                     *[]string        `position:"Query" name:"UserId"  type:"Repeated"`
	SkillLevel                 *[]string        `position:"Query" name:"SkillLevel"  type:"Repeated"`
	InstanceId                 string           `position:"Query" name:"InstanceId"`
	OutboundPhoneNumberId      *[]string        `position:"Query" name:"OutboundPhoneNumberId"  type:"Repeated"`
	Name                       string           `position:"Query" name:"Name"`
}

// CreateSkillGroupResponse is the response struct for api CreateSkillGroup
type CreateSkillGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SkillGroupId   string `json:"SkillGroupId" xml:"SkillGroupId"`
}

// CreateCreateSkillGroupRequest creates a request to invoke CreateSkillGroup API
func CreateCreateSkillGroupRequest() (request *CreateSkillGroupRequest) {
	request = &CreateSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateSkillGroup", "", "")
	return
}

// CreateCreateSkillGroupResponse creates a response to parse from CreateSkillGroup response
func CreateCreateSkillGroupResponse() (response *CreateSkillGroupResponse) {
	response = &CreateSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
