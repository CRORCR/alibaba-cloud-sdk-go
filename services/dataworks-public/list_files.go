package dataworks_public

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

// ListFiles invokes the dataworks_public.ListFiles API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/listfiles.html
func (client *Client) ListFiles(request *ListFilesRequest) (response *ListFilesResponse, err error) {
	response = CreateListFilesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFilesWithChan invokes the dataworks_public.ListFiles API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFilesWithChan(request *ListFilesRequest) (<-chan *ListFilesResponse, <-chan error) {
	responseChan := make(chan *ListFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFiles(request)
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

// ListFilesWithCallback invokes the dataworks_public.ListFiles API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFilesWithCallback(request *ListFilesRequest, callback func(response *ListFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFilesResponse
		var err error
		defer close(result)
		response, err = client.ListFiles(request)
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

// ListFilesRequest is the request struct for api ListFiles
type ListFilesRequest struct {
	*requests.RpcRequest
	Owner             string           `position:"Body" name:"Owner"`
	FileTypes         string           `position:"Body" name:"FileTypes"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	PageNumber        requests.Integer `position:"Body" name:"PageNumber"`
	FileFolderPath    string           `position:"Body" name:"FileFolderPath"`
	PageSize          requests.Integer `position:"Body" name:"PageSize"`
	Keyword           string           `position:"Body" name:"Keyword"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	UseType           string           `position:"Body" name:"UseType"`
}

// ListFilesResponse is the response struct for api ListFiles
type ListFilesResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	Success        bool            `json:"Success" xml:"Success"`
	ErrorCode      string          `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string          `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListFiles `json:"Data" xml:"Data"`
}

// CreateListFilesRequest creates a request to invoke ListFiles API
func CreateListFilesRequest() (request *ListFilesRequest) {
	request = &ListFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListFiles", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFilesResponse creates a response to parse from ListFiles response
func CreateListFilesResponse() (response *ListFilesResponse) {
	response = &ListFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
