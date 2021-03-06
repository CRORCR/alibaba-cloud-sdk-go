package rds

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

// ReplaceDedicatedHost invokes the rds.ReplaceDedicatedHost API synchronously
// api document: https://help.aliyun.com/api/rds/replacededicatedhost.html
func (client *Client) ReplaceDedicatedHost(request *ReplaceDedicatedHostRequest) (response *ReplaceDedicatedHostResponse, err error) {
	response = CreateReplaceDedicatedHostResponse()
	err = client.DoAction(request, response)
	return
}

// ReplaceDedicatedHostWithChan invokes the rds.ReplaceDedicatedHost API asynchronously
// api document: https://help.aliyun.com/api/rds/replacededicatedhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReplaceDedicatedHostWithChan(request *ReplaceDedicatedHostRequest) (<-chan *ReplaceDedicatedHostResponse, <-chan error) {
	responseChan := make(chan *ReplaceDedicatedHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplaceDedicatedHost(request)
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

// ReplaceDedicatedHostWithCallback invokes the rds.ReplaceDedicatedHost API asynchronously
// api document: https://help.aliyun.com/api/rds/replacededicatedhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReplaceDedicatedHostWithCallback(request *ReplaceDedicatedHostRequest, callback func(response *ReplaceDedicatedHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplaceDedicatedHostResponse
		var err error
		defer close(result)
		response, err = client.ReplaceDedicatedHost(request)
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

// ReplaceDedicatedHostRequest is the request struct for api ReplaceDedicatedHost
type ReplaceDedicatedHostRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FailoverMode         string           `position:"Query" name:"FailoverMode"`
}

// ReplaceDedicatedHostResponse is the response struct for api ReplaceDedicatedHost
type ReplaceDedicatedHostResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	TaskId          int    `json:"TaskId" xml:"TaskId"`
	DedicatedHostId string `json:"DedicatedHostId" xml:"DedicatedHostId"`
}

// CreateReplaceDedicatedHostRequest creates a request to invoke ReplaceDedicatedHost API
func CreateReplaceDedicatedHostRequest() (request *ReplaceDedicatedHostRequest) {
	request = &ReplaceDedicatedHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ReplaceDedicatedHost", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReplaceDedicatedHostResponse creates a response to parse from ReplaceDedicatedHost response
func CreateReplaceDedicatedHostResponse() (response *ReplaceDedicatedHostResponse) {
	response = &ReplaceDedicatedHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
