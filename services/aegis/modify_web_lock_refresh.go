package aegis

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

// ModifyWebLockRefresh invokes the aegis.ModifyWebLockRefresh API synchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockrefresh.html
func (client *Client) ModifyWebLockRefresh(request *ModifyWebLockRefreshRequest) (response *ModifyWebLockRefreshResponse, err error) {
	response = CreateModifyWebLockRefreshResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebLockRefreshWithChan invokes the aegis.ModifyWebLockRefresh API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockrefresh.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebLockRefreshWithChan(request *ModifyWebLockRefreshRequest) (<-chan *ModifyWebLockRefreshResponse, <-chan error) {
	responseChan := make(chan *ModifyWebLockRefreshResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebLockRefresh(request)
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

// ModifyWebLockRefreshWithCallback invokes the aegis.ModifyWebLockRefresh API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyweblockrefresh.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWebLockRefreshWithCallback(request *ModifyWebLockRefreshRequest, callback func(response *ModifyWebLockRefreshResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebLockRefreshResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebLockRefresh(request)
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

// ModifyWebLockRefreshRequest is the request struct for api ModifyWebLockRefresh
type ModifyWebLockRefreshRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
	Uuid     string `position:"Query" name:"Uuid"`
}

// ModifyWebLockRefreshResponse is the response struct for api ModifyWebLockRefresh
type ModifyWebLockRefreshResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebLockRefreshRequest creates a request to invoke ModifyWebLockRefresh API
func CreateModifyWebLockRefreshRequest() (request *ModifyWebLockRefreshRequest) {
	request = &ModifyWebLockRefreshRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyWebLockRefresh", "vipaegis", "openAPI")
	return
}

// CreateModifyWebLockRefreshResponse creates a response to parse from ModifyWebLockRefresh response
func CreateModifyWebLockRefreshResponse() (response *ModifyWebLockRefreshResponse) {
	response = &ModifyWebLockRefreshResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
