package alikafka

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

// ModifyPartitionNum invokes the alikafka.ModifyPartitionNum API synchronously
// api document: https://help.aliyun.com/api/alikafka/modifypartitionnum.html
func (client *Client) ModifyPartitionNum(request *ModifyPartitionNumRequest) (response *ModifyPartitionNumResponse, err error) {
	response = CreateModifyPartitionNumResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPartitionNumWithChan invokes the alikafka.ModifyPartitionNum API asynchronously
// api document: https://help.aliyun.com/api/alikafka/modifypartitionnum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPartitionNumWithChan(request *ModifyPartitionNumRequest) (<-chan *ModifyPartitionNumResponse, <-chan error) {
	responseChan := make(chan *ModifyPartitionNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPartitionNum(request)
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

// ModifyPartitionNumWithCallback invokes the alikafka.ModifyPartitionNum API asynchronously
// api document: https://help.aliyun.com/api/alikafka/modifypartitionnum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPartitionNumWithCallback(request *ModifyPartitionNumRequest, callback func(response *ModifyPartitionNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPartitionNumResponse
		var err error
		defer close(result)
		response, err = client.ModifyPartitionNum(request)
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

// ModifyPartitionNumRequest is the request struct for api ModifyPartitionNum
type ModifyPartitionNumRequest struct {
	*requests.RpcRequest
	InstanceId      string           `position:"Query" name:"InstanceId"`
	Topic           string           `position:"Query" name:"Topic"`
	AddPartitionNum requests.Integer `position:"Query" name:"AddPartitionNum"`
}

// ModifyPartitionNumResponse is the response struct for api ModifyPartitionNum
type ModifyPartitionNumResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateModifyPartitionNumRequest creates a request to invoke ModifyPartitionNum API
func CreateModifyPartitionNumRequest() (request *ModifyPartitionNumRequest) {
	request = &ModifyPartitionNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "ModifyPartitionNum", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyPartitionNumResponse creates a response to parse from ModifyPartitionNum response
func CreateModifyPartitionNumResponse() (response *ModifyPartitionNumResponse) {
	response = &ModifyPartitionNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
