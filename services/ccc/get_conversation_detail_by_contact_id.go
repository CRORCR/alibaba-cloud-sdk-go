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

// GetConversationDetailByContactId invokes the ccc.GetConversationDetailByContactId API synchronously
// api document: https://help.aliyun.com/api/ccc/getconversationdetailbycontactid.html
func (client *Client) GetConversationDetailByContactId(request *GetConversationDetailByContactIdRequest) (response *GetConversationDetailByContactIdResponse, err error) {
	response = CreateGetConversationDetailByContactIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetConversationDetailByContactIdWithChan invokes the ccc.GetConversationDetailByContactId API asynchronously
// api document: https://help.aliyun.com/api/ccc/getconversationdetailbycontactid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConversationDetailByContactIdWithChan(request *GetConversationDetailByContactIdRequest) (<-chan *GetConversationDetailByContactIdResponse, <-chan error) {
	responseChan := make(chan *GetConversationDetailByContactIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetConversationDetailByContactId(request)
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

// GetConversationDetailByContactIdWithCallback invokes the ccc.GetConversationDetailByContactId API asynchronously
// api document: https://help.aliyun.com/api/ccc/getconversationdetailbycontactid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConversationDetailByContactIdWithCallback(request *GetConversationDetailByContactIdRequest, callback func(response *GetConversationDetailByContactIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetConversationDetailByContactIdResponse
		var err error
		defer close(result)
		response, err = client.GetConversationDetailByContactId(request)
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

// GetConversationDetailByContactIdRequest is the request struct for api GetConversationDetailByContactId
type GetConversationDetailByContactIdRequest struct {
	*requests.RpcRequest
	ContactId  string           `position:"Query" name:"ContactId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// GetConversationDetailByContactIdResponse is the response struct for api GetConversationDetailByContactId
type GetConversationDetailByContactIdResponse struct {
	*responses.BaseResponse
	RequestId      string                                     `json:"RequestId" xml:"RequestId"`
	Success        bool                                       `json:"Success" xml:"Success"`
	Code           string                                     `json:"Code" xml:"Code"`
	Message        string                                     `json:"Message" xml:"Message"`
	HttpStatusCode int                                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DataList       DataListInGetConversationDetailByContactId `json:"DataList" xml:"DataList"`
}

// CreateGetConversationDetailByContactIdRequest creates a request to invoke GetConversationDetailByContactId API
func CreateGetConversationDetailByContactIdRequest() (request *GetConversationDetailByContactIdRequest) {
	request = &GetConversationDetailByContactIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetConversationDetailByContactId", "", "")
	return
}

// CreateGetConversationDetailByContactIdResponse creates a response to parse from GetConversationDetailByContactId response
func CreateGetConversationDetailByContactIdResponse() (response *GetConversationDetailByContactIdResponse) {
	response = &GetConversationDetailByContactIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
