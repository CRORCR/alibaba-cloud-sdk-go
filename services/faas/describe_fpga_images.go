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

// DescribeFpgaImages invokes the faas.DescribeFpgaImages API synchronously
// api document: https://help.aliyun.com/api/faas/describefpgaimages.html
func (client *Client) DescribeFpgaImages(request *DescribeFpgaImagesRequest) (response *DescribeFpgaImagesResponse, err error) {
	response = CreateDescribeFpgaImagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFpgaImagesWithChan invokes the faas.DescribeFpgaImages API asynchronously
// api document: https://help.aliyun.com/api/faas/describefpgaimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFpgaImagesWithChan(request *DescribeFpgaImagesRequest) (<-chan *DescribeFpgaImagesResponse, <-chan error) {
	responseChan := make(chan *DescribeFpgaImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFpgaImages(request)
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

// DescribeFpgaImagesWithCallback invokes the faas.DescribeFpgaImages API asynchronously
// api document: https://help.aliyun.com/api/faas/describefpgaimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFpgaImagesWithCallback(request *DescribeFpgaImagesRequest, callback func(response *DescribeFpgaImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFpgaImagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeFpgaImages(request)
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

// DescribeFpgaImagesRequest is the request struct for api DescribeFpgaImages
type DescribeFpgaImagesRequest struct {
	*requests.RpcRequest
	ECSImageId string `position:"Query" name:"ECSImageId"`
	OwnerAlias string `position:"Query" name:"OwnerAlias"`
}

// DescribeFpgaImagesResponse is the response struct for api DescribeFpgaImages
type DescribeFpgaImagesResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	FpgaImages []FpgaImage `json:"FpgaImages" xml:"FpgaImages"`
}

// CreateDescribeFpgaImagesRequest creates a request to invoke DescribeFpgaImages API
func CreateDescribeFpgaImagesRequest() (request *DescribeFpgaImagesRequest) {
	request = &DescribeFpgaImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "DescribeFpgaImages", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFpgaImagesResponse creates a response to parse from DescribeFpgaImages response
func CreateDescribeFpgaImagesResponse() (response *DescribeFpgaImagesResponse) {
	response = &DescribeFpgaImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
