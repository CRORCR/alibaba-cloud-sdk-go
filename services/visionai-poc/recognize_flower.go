package visionai_poc

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

// RecognizeFlower invokes the visionai_poc.RecognizeFlower API synchronously
// api document: https://help.aliyun.com/api/visionai-poc/recognizeflower.html
func (client *Client) RecognizeFlower(request *RecognizeFlowerRequest) (response *RecognizeFlowerResponse, err error) {
	response = CreateRecognizeFlowerResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeFlowerWithChan invokes the visionai_poc.RecognizeFlower API asynchronously
// api document: https://help.aliyun.com/api/visionai-poc/recognizeflower.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeFlowerWithChan(request *RecognizeFlowerRequest) (<-chan *RecognizeFlowerResponse, <-chan error) {
	responseChan := make(chan *RecognizeFlowerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeFlower(request)
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

// RecognizeFlowerWithCallback invokes the visionai_poc.RecognizeFlower API asynchronously
// api document: https://help.aliyun.com/api/visionai-poc/recognizeflower.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeFlowerWithCallback(request *RecognizeFlowerRequest, callback func(response *RecognizeFlowerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeFlowerResponse
		var err error
		defer close(result)
		response, err = client.RecognizeFlower(request)
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

// RecognizeFlowerRequest is the request struct for api RecognizeFlower
type RecognizeFlowerRequest struct {
	*requests.RpcRequest
	Method   string `position:"Body" name:"Method"`
	ImageUrl string `position:"Body" name:"ImageUrl"`
	Url      string `position:"Body" name:"Url"`
}

// RecognizeFlowerResponse is the response struct for api RecognizeFlower
type RecognizeFlowerResponse struct {
	*responses.BaseResponse
	Code     int                       `json:"Code" xml:"Code"`
	Message  string                    `json:"Message" xml:"Message"`
	Success  bool                      `json:"Success" xml:"Success"`
	Response ResponseInRecognizeFlower `json:"Response" xml:"Response"`
}

// CreateRecognizeFlowerRequest creates a request to invoke RecognizeFlower API
func CreateRecognizeFlowerRequest() (request *RecognizeFlowerRequest) {
	request = &RecognizeFlowerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("visionai-poc", "2020-04-08", "RecognizeFlower", "", "")
	return
}

// CreateRecognizeFlowerResponse creates a response to parse from RecognizeFlower response
func CreateRecognizeFlowerResponse() (response *RecognizeFlowerResponse) {
	response = &RecognizeFlowerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
