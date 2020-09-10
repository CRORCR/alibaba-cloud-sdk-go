package cloudwf

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

// GroupIntime invokes the cloudwf.GroupIntime API synchronously
// api document: https://help.aliyun.com/api/cloudwf/groupintime.html
func (client *Client) GroupIntime(request *GroupIntimeRequest) (response *GroupIntimeResponse, err error) {
	response = CreateGroupIntimeResponse()
	err = client.DoAction(request, response)
	return
}

// GroupIntimeWithChan invokes the cloudwf.GroupIntime API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/groupintime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GroupIntimeWithChan(request *GroupIntimeRequest) (<-chan *GroupIntimeResponse, <-chan error) {
	responseChan := make(chan *GroupIntimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GroupIntime(request)
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

// GroupIntimeWithCallback invokes the cloudwf.GroupIntime API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/groupintime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GroupIntimeWithCallback(request *GroupIntimeRequest, callback func(response *GroupIntimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GroupIntimeResponse
		var err error
		defer close(result)
		response, err = client.GroupIntime(request)
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

// GroupIntimeRequest is the request struct for api GroupIntime
type GroupIntimeRequest struct {
	*requests.RpcRequest
	Gsid requests.Integer `position:"Query" name:"Gsid"`
}

// GroupIntimeResponse is the response struct for api GroupIntime
type GroupIntimeResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGroupIntimeRequest creates a request to invoke GroupIntime API
func CreateGroupIntimeRequest() (request *GroupIntimeRequest) {
	request = &GroupIntimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GroupIntime", "cloudwf", "openAPI")
	return
}

// CreateGroupIntimeResponse creates a response to parse from GroupIntime response
func CreateGroupIntimeResponse() (response *GroupIntimeResponse) {
	response = &GroupIntimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
