package ivpd

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

// MakeSuperResolutionImage invokes the ivpd.MakeSuperResolutionImage API synchronously
// api document: https://help.aliyun.com/api/ivpd/makesuperresolutionimage.html
func (client *Client) MakeSuperResolutionImage(request *MakeSuperResolutionImageRequest) (response *MakeSuperResolutionImageResponse, err error) {
	response = CreateMakeSuperResolutionImageResponse()
	err = client.DoAction(request, response)
	return
}

// MakeSuperResolutionImageWithChan invokes the ivpd.MakeSuperResolutionImage API asynchronously
// api document: https://help.aliyun.com/api/ivpd/makesuperresolutionimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MakeSuperResolutionImageWithChan(request *MakeSuperResolutionImageRequest) (<-chan *MakeSuperResolutionImageResponse, <-chan error) {
	responseChan := make(chan *MakeSuperResolutionImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MakeSuperResolutionImage(request)
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

// MakeSuperResolutionImageWithCallback invokes the ivpd.MakeSuperResolutionImage API asynchronously
// api document: https://help.aliyun.com/api/ivpd/makesuperresolutionimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MakeSuperResolutionImageWithCallback(request *MakeSuperResolutionImageRequest, callback func(response *MakeSuperResolutionImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MakeSuperResolutionImageResponse
		var err error
		defer close(result)
		response, err = client.MakeSuperResolutionImage(request)
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

// MakeSuperResolutionImageRequest is the request struct for api MakeSuperResolutionImage
type MakeSuperResolutionImageRequest struct {
	*requests.RpcRequest
	Url string `position:"Body" name:"Url"`
}

// MakeSuperResolutionImageResponse is the response struct for api MakeSuperResolutionImage
type MakeSuperResolutionImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateMakeSuperResolutionImageRequest creates a request to invoke MakeSuperResolutionImage API
func CreateMakeSuperResolutionImageRequest() (request *MakeSuperResolutionImageRequest) {
	request = &MakeSuperResolutionImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "MakeSuperResolutionImage", "ivpd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMakeSuperResolutionImageResponse creates a response to parse from MakeSuperResolutionImage response
func CreateMakeSuperResolutionImageResponse() (response *MakeSuperResolutionImageResponse) {
	response = &MakeSuperResolutionImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
