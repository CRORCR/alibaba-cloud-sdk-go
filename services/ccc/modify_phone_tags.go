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

// ModifyPhoneTags invokes the ccc.ModifyPhoneTags API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyphonetags.html
func (client *Client) ModifyPhoneTags(request *ModifyPhoneTagsRequest) (response *ModifyPhoneTagsResponse, err error) {
	response = CreateModifyPhoneTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPhoneTagsWithChan invokes the ccc.ModifyPhoneTags API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyphonetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPhoneTagsWithChan(request *ModifyPhoneTagsRequest) (<-chan *ModifyPhoneTagsResponse, <-chan error) {
	responseChan := make(chan *ModifyPhoneTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPhoneTags(request)
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

// ModifyPhoneTagsWithCallback invokes the ccc.ModifyPhoneTags API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyphonetags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPhoneTagsWithCallback(request *ModifyPhoneTagsRequest, callback func(response *ModifyPhoneTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPhoneTagsResponse
		var err error
		defer close(result)
		response, err = client.ModifyPhoneTags(request)
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

// ModifyPhoneTagsRequest is the request struct for api ModifyPhoneTags
type ModifyPhoneTagsRequest struct {
	*requests.RpcRequest
	InstanceId       string    `position:"Query" name:"InstanceId"`
	SkillGroupIdList *[]string `position:"Query" name:"SkillGroupIdList"  type:"Repeated"`
	ServiceTag       string    `position:"Query" name:"ServiceTag"`
}

// ModifyPhoneTagsResponse is the response struct for api ModifyPhoneTags
type ModifyPhoneTagsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifyPhoneTagsRequest creates a request to invoke ModifyPhoneTags API
func CreateModifyPhoneTagsRequest() (request *ModifyPhoneTagsRequest) {
	request = &ModifyPhoneTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyPhoneTags", "", "")
	return
}

// CreateModifyPhoneTagsResponse creates a response to parse from ModifyPhoneTags response
func CreateModifyPhoneTagsResponse() (response *ModifyPhoneTagsResponse) {
	response = &ModifyPhoneTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
