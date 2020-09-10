package ledgerdb

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

// DisableMember invokes the ledgerdb.DisableMember API synchronously
// api document: https://help.aliyun.com/api/ledgerdb/disablemember.html
func (client *Client) DisableMember(request *DisableMemberRequest) (response *DisableMemberResponse, err error) {
	response = CreateDisableMemberResponse()
	err = client.DoAction(request, response)
	return
}

// DisableMemberWithChan invokes the ledgerdb.DisableMember API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/disablemember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableMemberWithChan(request *DisableMemberRequest) (<-chan *DisableMemberResponse, <-chan error) {
	responseChan := make(chan *DisableMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableMember(request)
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

// DisableMemberWithCallback invokes the ledgerdb.DisableMember API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/disablemember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableMemberWithCallback(request *DisableMemberRequest, callback func(response *DisableMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableMemberResponse
		var err error
		defer close(result)
		response, err = client.DisableMember(request)
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

// DisableMemberRequest is the request struct for api DisableMember
type DisableMemberRequest struct {
	*requests.RpcRequest
	LedgerId string `position:"Body" name:"LedgerId"`
	MemberId string `position:"Body" name:"MemberId"`
}

// DisableMemberResponse is the response struct for api DisableMember
type DisableMemberResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableMemberRequest creates a request to invoke DisableMember API
func CreateDisableMemberRequest() (request *DisableMemberRequest) {
	request = &DisableMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ledgerdb", "2019-11-22", "DisableMember", "ledgerdb", "openAPI")
	return
}

// CreateDisableMemberResponse creates a response to parse from DisableMember response
func CreateDisableMemberResponse() (response *DisableMemberResponse) {
	response = &DisableMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
