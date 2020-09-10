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

// GetFolder invokes the dataworks_public.GetFolder API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/getfolder.html
func (client *Client) GetFolder(request *GetFolderRequest) (response *GetFolderResponse, err error) {
	response = CreateGetFolderResponse()
	err = client.DoAction(request, response)
	return
}

// GetFolderWithChan invokes the dataworks_public.GetFolder API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getfolder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFolderWithChan(request *GetFolderRequest) (<-chan *GetFolderResponse, <-chan error) {
	responseChan := make(chan *GetFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFolder(request)
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

// GetFolderWithCallback invokes the dataworks_public.GetFolder API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getfolder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFolderWithCallback(request *GetFolderRequest, callback func(response *GetFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFolderResponse
		var err error
		defer close(result)
		response, err = client.GetFolder(request)
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

// GetFolderRequest is the request struct for api GetFolder
type GetFolderRequest struct {
	*requests.RpcRequest
	FolderPath        string           `position:"Body" name:"FolderPath"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	FolderId          string           `position:"Body" name:"FolderId"`
}

// GetFolderResponse is the response struct for api GetFolder
type GetFolderResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetFolderRequest creates a request to invoke GetFolder API
func CreateGetFolderRequest() (request *GetFolderRequest) {
	request = &GetFolderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetFolder", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetFolderResponse creates a response to parse from GetFolder response
func CreateGetFolderResponse() (response *GetFolderResponse) {
	response = &GetFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
