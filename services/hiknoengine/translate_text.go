package hiknoengine

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

// TranslateText invokes the hiknoengine.TranslateText API synchronously
// api document: https://help.aliyun.com/api/hiknoengine/translatetext.html
func (client *Client) TranslateText(request *TranslateTextRequest) (response *TranslateTextResponse, err error) {
	response = CreateTranslateTextResponse()
	err = client.DoAction(request, response)
	return
}

// TranslateTextWithChan invokes the hiknoengine.TranslateText API asynchronously
// api document: https://help.aliyun.com/api/hiknoengine/translatetext.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TranslateTextWithChan(request *TranslateTextRequest) (<-chan *TranslateTextResponse, <-chan error) {
	responseChan := make(chan *TranslateTextResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TranslateText(request)
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

// TranslateTextWithCallback invokes the hiknoengine.TranslateText API asynchronously
// api document: https://help.aliyun.com/api/hiknoengine/translatetext.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TranslateTextWithCallback(request *TranslateTextRequest, callback func(response *TranslateTextResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TranslateTextResponse
		var err error
		defer close(result)
		response, err = client.TranslateText(request)
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

// TranslateTextRequest is the request struct for api TranslateText
type TranslateTextRequest struct {
	*requests.RpcRequest
	FromLang string `position:"Body" name:"FromLang"`
	ToLang   string `position:"Body" name:"ToLang"`
	Text     string `position:"Body" name:"Text"`
}

// TranslateTextResponse is the response struct for api TranslateText
type TranslateTextResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTranslateTextRequest creates a request to invoke TranslateText API
func CreateTranslateTextRequest() (request *TranslateTextRequest) {
	request = &TranslateTextRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hiknoengine", "2019-06-25", "TranslateText", "hiknoengine", "openAPI")
	return
}

// CreateTranslateTextResponse creates a response to parse from TranslateText response
func CreateTranslateTextResponse() (response *TranslateTextResponse) {
	response = &TranslateTextResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
