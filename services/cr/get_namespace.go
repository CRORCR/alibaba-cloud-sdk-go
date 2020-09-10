package cr

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

// GetNamespace invokes the cr.GetNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/getnamespace.html
func (client *Client) GetNamespace(request *GetNamespaceRequest) (response *GetNamespaceResponse, err error) {
	response = CreateGetNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetNamespaceWithChan invokes the cr.GetNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNamespaceWithChan(request *GetNamespaceRequest) (<-chan *GetNamespaceResponse, <-chan error) {
	responseChan := make(chan *GetNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNamespace(request)
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

// GetNamespaceWithCallback invokes the cr.GetNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNamespaceWithCallback(request *GetNamespaceRequest, callback func(response *GetNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNamespaceResponse
		var err error
		defer close(result)
		response, err = client.GetNamespace(request)
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

// GetNamespaceRequest is the request struct for api GetNamespace
type GetNamespaceRequest struct {
	*requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

// GetNamespaceResponse is the response struct for api GetNamespace
type GetNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateGetNamespaceRequest creates a request to invoke GetNamespace API
func CreateGetNamespaceRequest() (request *GetNamespaceRequest) {
	request = &GetNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetNamespace", "/namespace/[Namespace]", "acr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetNamespaceResponse creates a response to parse from GetNamespace response
func CreateGetNamespaceResponse() (response *GetNamespaceResponse) {
	response = &GetNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
