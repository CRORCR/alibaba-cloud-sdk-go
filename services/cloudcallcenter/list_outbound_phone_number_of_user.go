package cloudcallcenter

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

// ListOutboundPhoneNumberOfUser invokes the cloudcallcenter.ListOutboundPhoneNumberOfUser API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoutboundphonenumberofuser.html
func (client *Client) ListOutboundPhoneNumberOfUser(request *ListOutboundPhoneNumberOfUserRequest) (response *ListOutboundPhoneNumberOfUserResponse, err error) {
	response = CreateListOutboundPhoneNumberOfUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListOutboundPhoneNumberOfUserWithChan invokes the cloudcallcenter.ListOutboundPhoneNumberOfUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoutboundphonenumberofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOutboundPhoneNumberOfUserWithChan(request *ListOutboundPhoneNumberOfUserRequest) (<-chan *ListOutboundPhoneNumberOfUserResponse, <-chan error) {
	responseChan := make(chan *ListOutboundPhoneNumberOfUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOutboundPhoneNumberOfUser(request)
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

// ListOutboundPhoneNumberOfUserWithCallback invokes the cloudcallcenter.ListOutboundPhoneNumberOfUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoutboundphonenumberofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOutboundPhoneNumberOfUserWithCallback(request *ListOutboundPhoneNumberOfUserRequest, callback func(response *ListOutboundPhoneNumberOfUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOutboundPhoneNumberOfUserResponse
		var err error
		defer close(result)
		response, err = client.ListOutboundPhoneNumberOfUser(request)
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

// ListOutboundPhoneNumberOfUserRequest is the request struct for api ListOutboundPhoneNumberOfUser
type ListOutboundPhoneNumberOfUserRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	UserId     string           `position:"Query" name:"UserId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListOutboundPhoneNumberOfUserResponse is the response struct for api ListOutboundPhoneNumberOfUser
type ListOutboundPhoneNumberOfUserResponse struct {
	*responses.BaseResponse
	RequestId            string                                              `json:"RequestId" xml:"RequestId"`
	Success              bool                                                `json:"Success" xml:"Success"`
	Code                 string                                              `json:"Code" xml:"Code"`
	Message              string                                              `json:"Message" xml:"Message"`
	HttpStatusCode       int                                                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	NumberList           NumberList                                          `json:"NumberList" xml:"NumberList"`
	OutboundPhoneNumbers OutboundPhoneNumbersInListOutboundPhoneNumberOfUser `json:"OutboundPhoneNumbers" xml:"OutboundPhoneNumbers"`
}

// CreateListOutboundPhoneNumberOfUserRequest creates a request to invoke ListOutboundPhoneNumberOfUser API
func CreateListOutboundPhoneNumberOfUserRequest() (request *ListOutboundPhoneNumberOfUserRequest) {
	request = &ListOutboundPhoneNumberOfUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListOutboundPhoneNumberOfUser", "", "")
	request.Method = requests.POST
	return
}

// CreateListOutboundPhoneNumberOfUserResponse creates a response to parse from ListOutboundPhoneNumberOfUser response
func CreateListOutboundPhoneNumberOfUserResponse() (response *ListOutboundPhoneNumberOfUserResponse) {
	response = &ListOutboundPhoneNumberOfUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
