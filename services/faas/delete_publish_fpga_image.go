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

// DeletePublishFpgaImage invokes the faas.DeletePublishFpgaImage API synchronously
// api document: https://help.aliyun.com/api/faas/deletepublishfpgaimage.html
func (client *Client) DeletePublishFpgaImage(request *DeletePublishFpgaImageRequest) (response *DeletePublishFpgaImageResponse, err error) {
	response = CreateDeletePublishFpgaImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePublishFpgaImageWithChan invokes the faas.DeletePublishFpgaImage API asynchronously
// api document: https://help.aliyun.com/api/faas/deletepublishfpgaimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePublishFpgaImageWithChan(request *DeletePublishFpgaImageRequest) (<-chan *DeletePublishFpgaImageResponse, <-chan error) {
	responseChan := make(chan *DeletePublishFpgaImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePublishFpgaImage(request)
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

// DeletePublishFpgaImageWithCallback invokes the faas.DeletePublishFpgaImage API asynchronously
// api document: https://help.aliyun.com/api/faas/deletepublishfpgaimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePublishFpgaImageWithCallback(request *DeletePublishFpgaImageRequest, callback func(response *DeletePublishFpgaImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePublishFpgaImageResponse
		var err error
		defer close(result)
		response, err = client.DeletePublishFpgaImage(request)
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

// DeletePublishFpgaImageRequest is the request struct for api DeletePublishFpgaImage
type DeletePublishFpgaImageRequest struct {
	*requests.RpcRequest
	ImageID       string           `position:"Query" name:"ImageID"`
	FpgaImageUUID string           `position:"Query" name:"FpgaImageUUID"`
	CallerUid     requests.Integer `position:"Query" name:"callerUid"`
}

// DeletePublishFpgaImageResponse is the response struct for api DeletePublishFpgaImage
type DeletePublishFpgaImageResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Status        string `json:"Status" xml:"Status"`
	FpgaImageUUID string `json:"FpgaImageUUID" xml:"FpgaImageUUID"`
	Message       string `json:"Message" xml:"Message"`
}

// CreateDeletePublishFpgaImageRequest creates a request to invoke DeletePublishFpgaImage API
func CreateDeletePublishFpgaImageRequest() (request *DeletePublishFpgaImageRequest) {
	request = &DeletePublishFpgaImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2017-08-24", "DeletePublishFpgaImage", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePublishFpgaImageResponse creates a response to parse from DeletePublishFpgaImage response
func CreateDeletePublishFpgaImageResponse() (response *DeletePublishFpgaImageResponse) {
	response = &DeletePublishFpgaImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
