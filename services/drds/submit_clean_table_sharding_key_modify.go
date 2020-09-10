package drds

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

// SubmitCleanTableShardingKeyModify invokes the drds.SubmitCleanTableShardingKeyModify API synchronously
// api document: https://help.aliyun.com/api/drds/submitcleantableshardingkeymodify.html
func (client *Client) SubmitCleanTableShardingKeyModify(request *SubmitCleanTableShardingKeyModifyRequest) (response *SubmitCleanTableShardingKeyModifyResponse, err error) {
	response = CreateSubmitCleanTableShardingKeyModifyResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitCleanTableShardingKeyModifyWithChan invokes the drds.SubmitCleanTableShardingKeyModify API asynchronously
// api document: https://help.aliyun.com/api/drds/submitcleantableshardingkeymodify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitCleanTableShardingKeyModifyWithChan(request *SubmitCleanTableShardingKeyModifyRequest) (<-chan *SubmitCleanTableShardingKeyModifyResponse, <-chan error) {
	responseChan := make(chan *SubmitCleanTableShardingKeyModifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitCleanTableShardingKeyModify(request)
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

// SubmitCleanTableShardingKeyModifyWithCallback invokes the drds.SubmitCleanTableShardingKeyModify API asynchronously
// api document: https://help.aliyun.com/api/drds/submitcleantableshardingkeymodify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitCleanTableShardingKeyModifyWithCallback(request *SubmitCleanTableShardingKeyModifyRequest, callback func(response *SubmitCleanTableShardingKeyModifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitCleanTableShardingKeyModifyResponse
		var err error
		defer close(result)
		response, err = client.SubmitCleanTableShardingKeyModify(request)
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

// SubmitCleanTableShardingKeyModifyRequest is the request struct for api SubmitCleanTableShardingKeyModify
type SubmitCleanTableShardingKeyModifyRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	TaskId         string `position:"Query" name:"TaskId"`
}

// SubmitCleanTableShardingKeyModifyResponse is the response struct for api SubmitCleanTableShardingKeyModify
type SubmitCleanTableShardingKeyModifyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateSubmitCleanTableShardingKeyModifyRequest creates a request to invoke SubmitCleanTableShardingKeyModify API
func CreateSubmitCleanTableShardingKeyModifyRequest() (request *SubmitCleanTableShardingKeyModifyRequest) {
	request = &SubmitCleanTableShardingKeyModifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SubmitCleanTableShardingKeyModify", "Drds", "openAPI")
	return
}

// CreateSubmitCleanTableShardingKeyModifyResponse creates a response to parse from SubmitCleanTableShardingKeyModify response
func CreateSubmitCleanTableShardingKeyModifyResponse() (response *SubmitCleanTableShardingKeyModifyResponse) {
	response = &SubmitCleanTableShardingKeyModifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
