package yundun_dbaudit

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

// ModifyInstanceStorage invokes the yundun_dbaudit.ModifyInstanceStorage API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/modifyinstancestorage.html
func (client *Client) ModifyInstanceStorage(request *ModifyInstanceStorageRequest) (response *ModifyInstanceStorageResponse, err error) {
	response = CreateModifyInstanceStorageResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceStorageWithChan invokes the yundun_dbaudit.ModifyInstanceStorage API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/modifyinstancestorage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceStorageWithChan(request *ModifyInstanceStorageRequest) (<-chan *ModifyInstanceStorageResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceStorageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceStorage(request)
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

// ModifyInstanceStorageWithCallback invokes the yundun_dbaudit.ModifyInstanceStorage API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/modifyinstancestorage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceStorageWithCallback(request *ModifyInstanceStorageRequest, callback func(response *ModifyInstanceStorageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceStorageResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceStorage(request)
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

// ModifyInstanceStorageRequest is the request struct for api ModifyInstanceStorage
type ModifyInstanceStorageRequest struct {
	*requests.RpcRequest
	StorageCategory string           `position:"Query" name:"StorageCategory"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	StorageTime     requests.Integer `position:"Query" name:"StorageTime"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	StorageSpace    string           `position:"Query" name:"StorageSpace"`
}

// ModifyInstanceStorageResponse is the response struct for api ModifyInstanceStorage
type ModifyInstanceStorageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceStorageRequest creates a request to invoke ModifyInstanceStorage API
func CreateModifyInstanceStorageRequest() (request *ModifyInstanceStorageRequest) {
	request = &ModifyInstanceStorageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "ModifyInstanceStorage", "dbaudit", "openAPI")
	return
}

// CreateModifyInstanceStorageResponse creates a response to parse from ModifyInstanceStorage response
func CreateModifyInstanceStorageResponse() (response *ModifyInstanceStorageResponse) {
	response = &ModifyInstanceStorageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
