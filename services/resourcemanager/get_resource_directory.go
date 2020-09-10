package resourcemanager

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

// GetResourceDirectory invokes the resourcemanager.GetResourceDirectory API synchronously
// api document: https://help.aliyun.com/api/resourcemanager/getresourcedirectory.html
func (client *Client) GetResourceDirectory(request *GetResourceDirectoryRequest) (response *GetResourceDirectoryResponse, err error) {
	response = CreateGetResourceDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceDirectoryWithChan invokes the resourcemanager.GetResourceDirectory API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/getresourcedirectory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourceDirectoryWithChan(request *GetResourceDirectoryRequest) (<-chan *GetResourceDirectoryResponse, <-chan error) {
	responseChan := make(chan *GetResourceDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceDirectory(request)
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

// GetResourceDirectoryWithCallback invokes the resourcemanager.GetResourceDirectory API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/getresourcedirectory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourceDirectoryWithCallback(request *GetResourceDirectoryRequest, callback func(response *GetResourceDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceDirectoryResponse
		var err error
		defer close(result)
		response, err = client.GetResourceDirectory(request)
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

// GetResourceDirectoryRequest is the request struct for api GetResourceDirectory
type GetResourceDirectoryRequest struct {
	*requests.RpcRequest
}

// GetResourceDirectoryResponse is the response struct for api GetResourceDirectory
type GetResourceDirectoryResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ResourceDirectory ResourceDirectory `json:"ResourceDirectory" xml:"ResourceDirectory"`
}

// CreateGetResourceDirectoryRequest creates a request to invoke GetResourceDirectory API
func CreateGetResourceDirectoryRequest() (request *GetResourceDirectoryRequest) {
	request = &GetResourceDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "GetResourceDirectory", "resourcemanager", "openAPI")
	return
}

// CreateGetResourceDirectoryResponse creates a response to parse from GetResourceDirectory response
func CreateGetResourceDirectoryResponse() (response *GetResourceDirectoryResponse) {
	response = &GetResourceDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
