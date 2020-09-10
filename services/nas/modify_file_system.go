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

package nas

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyFileSystem invokes the nas.ModifyFileSystem API synchronously
// api document: https://help.aliyun.com/api/nas/modifyfilesystem.html
func (client *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (response *ModifyFileSystemResponse, err error) {
	response = CreateModifyFileSystemResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFileSystemWithChan invokes the nas.ModifyFileSystem API asynchronously
// api document: https://help.aliyun.com/api/nas/modifyfilesystem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFileSystemWithChan(request *ModifyFileSystemRequest) (<-chan *ModifyFileSystemResponse, <-chan error) {
	responseChan := make(chan *ModifyFileSystemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFileSystem(request)
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

// ModifyFileSystemWithCallback invokes the nas.ModifyFileSystem API asynchronously
// api document: https://help.aliyun.com/api/nas/modifyfilesystem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFileSystemWithCallback(request *ModifyFileSystemRequest, callback func(response *ModifyFileSystemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFileSystemResponse
		var err error
		defer close(result)
		response, err = client.ModifyFileSystem(request)
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

// ModifyFileSystemRequest is the request struct for api ModifyFileSystem
type ModifyFileSystemRequest struct {
	*requests.RpcRequest
	FileSystemId string `position:"Query" name:"FileSystemId"`
	Description  string `position:"Query" name:"Description"`
}

// ModifyFileSystemResponse is the response struct for api ModifyFileSystem
type ModifyFileSystemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyFileSystemRequest creates a request to invoke ModifyFileSystem API
func CreateModifyFileSystemRequest() (request *ModifyFileSystemRequest) {
	request = &ModifyFileSystemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ModifyFileSystem", "nas", "openAPI")
	return
}

// CreateModifyFileSystemResponse creates a response to parse from ModifyFileSystem response
func CreateModifyFileSystemResponse() (response *ModifyFileSystemResponse) {
	response = &ModifyFileSystemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
