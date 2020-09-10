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

// DeleteFile invokes the dataworks_public.DeleteFile API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletefile.html
func (client *Client) DeleteFile(request *DeleteFileRequest) (response *DeleteFileResponse, err error) {
	response = CreateDeleteFileResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFileWithChan invokes the dataworks_public.DeleteFile API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletefile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFileWithChan(request *DeleteFileRequest) (<-chan *DeleteFileResponse, <-chan error) {
	responseChan := make(chan *DeleteFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFile(request)
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

// DeleteFileWithCallback invokes the dataworks_public.DeleteFile API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletefile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFileWithCallback(request *DeleteFileRequest, callback func(response *DeleteFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFileResponse
		var err error
		defer close(result)
		response, err = client.DeleteFile(request)
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

// DeleteFileRequest is the request struct for api DeleteFile
type DeleteFileRequest struct {
	*requests.RpcRequest
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	FileId            requests.Integer `position:"Body" name:"FileId"`
}

// DeleteFileResponse is the response struct for api DeleteFile
type DeleteFileResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteFileRequest creates a request to invoke DeleteFile API
func CreateDeleteFileRequest() (request *DeleteFileRequest) {
	request = &DeleteFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteFile", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFileResponse creates a response to parse from DeleteFile response
func CreateDeleteFileResponse() (response *DeleteFileResponse) {
	response = &DeleteFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
