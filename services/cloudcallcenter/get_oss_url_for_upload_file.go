package cloudcallcenter

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

// GetOssUrlForUploadFile invokes the cloudcallcenter.GetOssUrlForUploadFile API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getossurlforuploadfile.html
func (client *Client) GetOssUrlForUploadFile(request *GetOssUrlForUploadFileRequest) (response *GetOssUrlForUploadFileResponse, err error) {
	response = CreateGetOssUrlForUploadFileResponse()
	err = client.DoAction(request, response)
	return
}

// GetOssUrlForUploadFileWithChan invokes the cloudcallcenter.GetOssUrlForUploadFile API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getossurlforuploadfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOssUrlForUploadFileWithChan(request *GetOssUrlForUploadFileRequest) (<-chan *GetOssUrlForUploadFileResponse, <-chan error) {
	responseChan := make(chan *GetOssUrlForUploadFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOssUrlForUploadFile(request)
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

// GetOssUrlForUploadFileWithCallback invokes the cloudcallcenter.GetOssUrlForUploadFile API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getossurlforuploadfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOssUrlForUploadFileWithCallback(request *GetOssUrlForUploadFileRequest, callback func(response *GetOssUrlForUploadFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOssUrlForUploadFileResponse
		var err error
		defer close(result)
		response, err = client.GetOssUrlForUploadFile(request)
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

// GetOssUrlForUploadFileRequest is the request struct for api GetOssUrlForUploadFile
type GetOssUrlForUploadFileRequest struct {
	*requests.RpcRequest
	File string `position:"Query" name:"File"`
}

// GetOssUrlForUploadFileResponse is the response struct for api GetOssUrlForUploadFile
type GetOssUrlForUploadFileResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	CccKey         string `json:"CccKey" xml:"CccKey"`
	PictureUrl     string `json:"PictureUrl" xml:"PictureUrl"`
	Dir            string `json:"Dir" xml:"Dir"`
	Host           string `json:"Host" xml:"Host"`
	Expires        string `json:"Expires" xml:"Expires"`
	OssAccessKeyId string `json:"OssAccessKeyId" xml:"OssAccessKeyId"`
	Policy         string `json:"Policy" xml:"Policy"`
	Signature      string `json:"Signature" xml:"Signature"`
}

// CreateGetOssUrlForUploadFileRequest creates a request to invoke GetOssUrlForUploadFile API
func CreateGetOssUrlForUploadFileRequest() (request *GetOssUrlForUploadFileRequest) {
	request = &GetOssUrlForUploadFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetOssUrlForUploadFile", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOssUrlForUploadFileResponse creates a response to parse from GetOssUrlForUploadFile response
func CreateGetOssUrlForUploadFileResponse() (response *GetOssUrlForUploadFileResponse) {
	response = &GetOssUrlForUploadFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
