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

// SubmitFile invokes the dataworks_public.SubmitFile API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/submitfile.html
func (client *Client) SubmitFile(request *SubmitFileRequest) (response *SubmitFileResponse, err error) {
	response = CreateSubmitFileResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitFileWithChan invokes the dataworks_public.SubmitFile API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/submitfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitFileWithChan(request *SubmitFileRequest) (<-chan *SubmitFileResponse, <-chan error) {
	responseChan := make(chan *SubmitFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitFile(request)
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

// SubmitFileWithCallback invokes the dataworks_public.SubmitFile API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/submitfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitFileWithCallback(request *SubmitFileRequest, callback func(response *SubmitFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitFileResponse
		var err error
		defer close(result)
		response, err = client.SubmitFile(request)
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

// SubmitFileRequest is the request struct for api SubmitFile
type SubmitFileRequest struct {
	*requests.RpcRequest
	Comment           string           `position:"Body" name:"Comment"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	FileId            requests.Integer `position:"Body" name:"FileId"`
}

// SubmitFileResponse is the response struct for api SubmitFile
type SubmitFileResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data           int64  `json:"Data" xml:"Data"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateSubmitFileRequest creates a request to invoke SubmitFile API
func CreateSubmitFileRequest() (request *SubmitFileRequest) {
	request = &SubmitFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "SubmitFile", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitFileResponse creates a response to parse from SubmitFile response
func CreateSubmitFileResponse() (response *SubmitFileResponse) {
	response = &SubmitFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
