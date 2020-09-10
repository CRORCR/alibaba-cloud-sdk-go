package faas

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

// CopyFpgaImage invokes the faas.CopyFpgaImage API synchronously
// api document: https://help.aliyun.com/api/faas/copyfpgaimage.html
func (client *Client) CopyFpgaImage(request *CopyFpgaImageRequest) (response *CopyFpgaImageResponse, err error) {
	response = CreateCopyFpgaImageResponse()
	err = client.DoAction(request, response)
	return
}

// CopyFpgaImageWithChan invokes the faas.CopyFpgaImage API asynchronously
// api document: https://help.aliyun.com/api/faas/copyfpgaimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyFpgaImageWithChan(request *CopyFpgaImageRequest) (<-chan *CopyFpgaImageResponse, <-chan error) {
	responseChan := make(chan *CopyFpgaImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyFpgaImage(request)
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

// CopyFpgaImageWithCallback invokes the faas.CopyFpgaImage API asynchronously
// api document: https://help.aliyun.com/api/faas/copyfpgaimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyFpgaImageWithCallback(request *CopyFpgaImageRequest, callback func(response *CopyFpgaImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyFpgaImageResponse
		var err error
		defer close(result)
		response, err = client.CopyFpgaImage(request)
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

// CopyFpgaImageRequest is the request struct for api CopyFpgaImage
type CopyFpgaImageRequest struct {
	*requests.RpcRequest
	TargetRegion      string `position:"Query" name:"TargetRegion"`
	FpgaImageUniqueId string `position:"Query" name:"FpgaImageUniqueId"`
}

// CopyFpgaImageResponse is the response struct for api CopyFpgaImage
type CopyFpgaImageResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	FpgaImageUniqueId string `json:"FpgaImageUniqueId" xml:"FpgaImageUniqueId"`
	Name              string `json:"Name" xml:"Name"`
	CreateTime        string `json:"CreateTime" xml:"CreateTime"`
	Description       string `json:"Description" xml:"Description"`
	ShellUniqueId     string `json:"ShellUniqueId" xml:"ShellUniqueId"`
	State             string `json:"State" xml:"State"`
}

// CreateCopyFpgaImageRequest creates a request to invoke CopyFpgaImage API
func CreateCopyFpgaImageRequest() (request *CopyFpgaImageRequest) {
	request = &CopyFpgaImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "CopyFpgaImage", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCopyFpgaImageResponse creates a response to parse from CopyFpgaImage response
func CreateCopyFpgaImageResponse() (response *CopyFpgaImageResponse) {
	response = &CopyFpgaImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
