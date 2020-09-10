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

// GetRepoListByNamespace invokes the cr.GetRepoListByNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/getrepolistbynamespace.html
func (client *Client) GetRepoListByNamespace(request *GetRepoListByNamespaceRequest) (response *GetRepoListByNamespaceResponse, err error) {
	response = CreateGetRepoListByNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoListByNamespaceWithChan invokes the cr.GetRepoListByNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepolistbynamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoListByNamespaceWithChan(request *GetRepoListByNamespaceRequest) (<-chan *GetRepoListByNamespaceResponse, <-chan error) {
	responseChan := make(chan *GetRepoListByNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoListByNamespace(request)
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

// GetRepoListByNamespaceWithCallback invokes the cr.GetRepoListByNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepolistbynamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoListByNamespaceWithCallback(request *GetRepoListByNamespaceRequest, callback func(response *GetRepoListByNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoListByNamespaceResponse
		var err error
		defer close(result)
		response, err = client.GetRepoListByNamespace(request)
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

// GetRepoListByNamespaceRequest is the request struct for api GetRepoListByNamespace
type GetRepoListByNamespaceRequest struct {
	*requests.RoaRequest
	RepoNamespace string           `position:"Path" name:"RepoNamespace"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Page          requests.Integer `position:"Query" name:"Page"`
	Status        string           `position:"Query" name:"Status"`
}

// GetRepoListByNamespaceResponse is the response struct for api GetRepoListByNamespace
type GetRepoListByNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoListByNamespaceRequest creates a request to invoke GetRepoListByNamespace API
func CreateGetRepoListByNamespaceRequest() (request *GetRepoListByNamespaceRequest) {
	request = &GetRepoListByNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoListByNamespace", "/repos/[RepoNamespace]", "acr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRepoListByNamespaceResponse creates a response to parse from GetRepoListByNamespace response
func CreateGetRepoListByNamespaceResponse() (response *GetRepoListByNamespaceResponse) {
	response = &GetRepoListByNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
