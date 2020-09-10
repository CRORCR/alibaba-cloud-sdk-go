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

// GetPublicAccessUrls invokes the cloudphoto.GetPublicAccessUrls API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getpublicaccessurls.html
func (client *Client) GetPublicAccessUrls(request *GetPublicAccessUrlsRequest) (response *GetPublicAccessUrlsResponse, err error) {
	response = CreateGetPublicAccessUrlsResponse()
	err = client.DoAction(request, response)
	return
}

// GetPublicAccessUrlsWithChan invokes the cloudphoto.GetPublicAccessUrls API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getpublicaccessurls.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPublicAccessUrlsWithChan(request *GetPublicAccessUrlsRequest) (<-chan *GetPublicAccessUrlsResponse, <-chan error) {
	responseChan := make(chan *GetPublicAccessUrlsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPublicAccessUrls(request)
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

// GetPublicAccessUrlsWithCallback invokes the cloudphoto.GetPublicAccessUrls API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getpublicaccessurls.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPublicAccessUrlsWithCallback(request *GetPublicAccessUrlsRequest, callback func(response *GetPublicAccessUrlsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPublicAccessUrlsResponse
		var err error
		defer close(result)
		response, err = client.GetPublicAccessUrls(request)
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

// GetPublicAccessUrlsRequest is the request struct for api GetPublicAccessUrls
type GetPublicAccessUrlsRequest struct {
	*requests.RpcRequest
	DomainType string    `position:"Query" name:"DomainType"`
	LibraryId  string    `position:"Query" name:"LibraryId"`
	PhotoId    *[]string `position:"Query" name:"PhotoId"  type:"Repeated"`
	StoreName  string    `position:"Query" name:"StoreName"`
	ZoomType   string    `position:"Query" name:"ZoomType"`
}

// GetPublicAccessUrlsResponse is the response struct for api GetPublicAccessUrls
type GetPublicAccessUrlsResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Action    string   `json:"Action" xml:"Action"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateGetPublicAccessUrlsRequest creates a request to invoke GetPublicAccessUrls API
func CreateGetPublicAccessUrlsRequest() (request *GetPublicAccessUrlsRequest) {
	request = &GetPublicAccessUrlsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPublicAccessUrls", "cloudphoto", "openAPI")
	return
}

// CreateGetPublicAccessUrlsResponse creates a response to parse from GetPublicAccessUrls response
func CreateGetPublicAccessUrlsResponse() (response *GetPublicAccessUrlsResponse) {
	response = &GetPublicAccessUrlsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
