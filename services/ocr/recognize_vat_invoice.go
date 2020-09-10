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

// RecognizeVATInvoice invokes the ocr.RecognizeVATInvoice API synchronously
// api document: https://help.aliyun.com/api/ocr/recognizevatinvoice.html
func (client *Client) RecognizeVATInvoice(request *RecognizeVATInvoiceRequest) (response *RecognizeVATInvoiceResponse, err error) {
	response = CreateRecognizeVATInvoiceResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeVATInvoiceWithChan invokes the ocr.RecognizeVATInvoice API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizevatinvoice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeVATInvoiceWithChan(request *RecognizeVATInvoiceRequest) (<-chan *RecognizeVATInvoiceResponse, <-chan error) {
	responseChan := make(chan *RecognizeVATInvoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeVATInvoice(request)
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

// RecognizeVATInvoiceWithCallback invokes the ocr.RecognizeVATInvoice API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizevatinvoice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeVATInvoiceWithCallback(request *RecognizeVATInvoiceRequest, callback func(response *RecognizeVATInvoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeVATInvoiceResponse
		var err error
		defer close(result)
		response, err = client.RecognizeVATInvoice(request)
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

// RecognizeVATInvoiceRequest is the request struct for api RecognizeVATInvoice
type RecognizeVATInvoiceRequest struct {
	*requests.RpcRequest
	FileType string `position:"Body" name:"FileType"`
	FileURL  string `position:"Body" name:"FileURL"`
}

// RecognizeVATInvoiceResponse is the response struct for api RecognizeVATInvoice
type RecognizeVATInvoiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeVATInvoiceRequest creates a request to invoke RecognizeVATInvoice API
func CreateRecognizeVATInvoiceRequest() (request *RecognizeVATInvoiceRequest) {
	request = &RecognizeVATInvoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeVATInvoice", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeVATInvoiceResponse creates a response to parse from RecognizeVATInvoice response
func CreateRecognizeVATInvoiceResponse() (response *RecognizeVATInvoiceResponse) {
	response = &RecognizeVATInvoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
