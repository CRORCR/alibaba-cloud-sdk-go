package ocr

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

// RecognizeVINCode invokes the ocr.RecognizeVINCode API synchronously
// api document: https://help.aliyun.com/api/ocr/recognizevincode.html
func (client *Client) RecognizeVINCode(request *RecognizeVINCodeRequest) (response *RecognizeVINCodeResponse, err error) {
	response = CreateRecognizeVINCodeResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeVINCodeWithChan invokes the ocr.RecognizeVINCode API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizevincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeVINCodeWithChan(request *RecognizeVINCodeRequest) (<-chan *RecognizeVINCodeResponse, <-chan error) {
	responseChan := make(chan *RecognizeVINCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeVINCode(request)
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

// RecognizeVINCodeWithCallback invokes the ocr.RecognizeVINCode API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizevincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeVINCodeWithCallback(request *RecognizeVINCodeRequest, callback func(response *RecognizeVINCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeVINCodeResponse
		var err error
		defer close(result)
		response, err = client.RecognizeVINCode(request)
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

// RecognizeVINCodeRequest is the request struct for api RecognizeVINCode
type RecognizeVINCodeRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Query" name:"ImageType"`
	ImageURL  string           `position:"Query" name:"ImageURL"`
}

// RecognizeVINCodeResponse is the response struct for api RecognizeVINCode
type RecognizeVINCodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeVINCodeRequest creates a request to invoke RecognizeVINCode API
func CreateRecognizeVINCodeRequest() (request *RecognizeVINCodeRequest) {
	request = &RecognizeVINCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeVINCode", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeVINCodeResponse creates a response to parse from RecognizeVINCode response
func CreateRecognizeVINCodeResponse() (response *RecognizeVINCodeResponse) {
	response = &RecognizeVINCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
