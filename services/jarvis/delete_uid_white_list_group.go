package jarvis

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

// DeleteUidWhiteListGroup invokes the jarvis.DeleteUidWhiteListGroup API synchronously
// api document: https://help.aliyun.com/api/jarvis/deleteuidwhitelistgroup.html
func (client *Client) DeleteUidWhiteListGroup(request *DeleteUidWhiteListGroupRequest) (response *DeleteUidWhiteListGroupResponse, err error) {
	response = CreateDeleteUidWhiteListGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUidWhiteListGroupWithChan invokes the jarvis.DeleteUidWhiteListGroup API asynchronously
// api document: https://help.aliyun.com/api/jarvis/deleteuidwhitelistgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUidWhiteListGroupWithChan(request *DeleteUidWhiteListGroupRequest) (<-chan *DeleteUidWhiteListGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteUidWhiteListGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUidWhiteListGroup(request)
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

// DeleteUidWhiteListGroupWithCallback invokes the jarvis.DeleteUidWhiteListGroup API asynchronously
// api document: https://help.aliyun.com/api/jarvis/deleteuidwhitelistgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUidWhiteListGroupWithCallback(request *DeleteUidWhiteListGroupRequest, callback func(response *DeleteUidWhiteListGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUidWhiteListGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteUidWhiteListGroup(request)
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

// DeleteUidWhiteListGroupRequest is the request struct for api DeleteUidWhiteListGroup
type DeleteUidWhiteListGroupRequest struct {
	*requests.RpcRequest
	GroupIdList string `position:"Query" name:"GroupIdList"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	Lang        string `position:"Query" name:"Lang"`
	SourceCode  string `position:"Query" name:"SourceCode"`
}

// DeleteUidWhiteListGroupResponse is the response struct for api DeleteUidWhiteListGroup
type DeleteUidWhiteListGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateDeleteUidWhiteListGroupRequest creates a request to invoke DeleteUidWhiteListGroup API
func CreateDeleteUidWhiteListGroupRequest() (request *DeleteUidWhiteListGroupRequest) {
	request = &DeleteUidWhiteListGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DeleteUidWhiteListGroup", "jarvis", "openAPI")
	return
}

// CreateDeleteUidWhiteListGroupResponse creates a response to parse from DeleteUidWhiteListGroup response
func CreateDeleteUidWhiteListGroupResponse() (response *DeleteUidWhiteListGroupResponse) {
	response = &DeleteUidWhiteListGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
