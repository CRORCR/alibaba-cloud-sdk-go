package cloudphoto

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

// GetThumbnails invokes the cloudphoto.GetThumbnails API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getthumbnails.html
func (client *Client) GetThumbnails(request *GetThumbnailsRequest) (response *GetThumbnailsResponse, err error) {
	response = CreateGetThumbnailsResponse()
	err = client.DoAction(request, response)
	return
}

// GetThumbnailsWithChan invokes the cloudphoto.GetThumbnails API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getthumbnails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetThumbnailsWithChan(request *GetThumbnailsRequest) (<-chan *GetThumbnailsResponse, <-chan error) {
	responseChan := make(chan *GetThumbnailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetThumbnails(request)
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

// GetThumbnailsWithCallback invokes the cloudphoto.GetThumbnails API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getthumbnails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetThumbnailsWithCallback(request *GetThumbnailsRequest, callback func(response *GetThumbnailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetThumbnailsResponse
		var err error
		defer close(result)
		response, err = client.GetThumbnails(request)
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

// GetThumbnailsRequest is the request struct for api GetThumbnails
type GetThumbnailsRequest struct {
	*requests.RpcRequest
	LibraryId string    `position:"Query" name:"LibraryId"`
	PhotoId   *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
	StoreName string    `position:"Query" name:"StoreName"`
	ZoomType  string    `position:"Query" name:"ZoomType"`
}

// GetThumbnailsResponse is the response struct for api GetThumbnails
type GetThumbnailsResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Action    string                 `json:"Action" xml:"Action"`
	Results   ResultsInGetThumbnails `json:"Results" xml:"Results"`
}

// CreateGetThumbnailsRequest creates a request to invoke GetThumbnails API
func CreateGetThumbnailsRequest() (request *GetThumbnailsRequest) {
	request = &GetThumbnailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetThumbnails", "cloudphoto", "openAPI")
	return
}

// CreateGetThumbnailsResponse creates a response to parse from GetThumbnails response
func CreateGetThumbnailsResponse() (response *GetThumbnailsResponse) {
	response = &GetThumbnailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
