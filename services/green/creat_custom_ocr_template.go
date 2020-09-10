package green

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

// CreatCustomOcrTemplate invokes the green.CreatCustomOcrTemplate API synchronously
// api document: https://help.aliyun.com/api/green/creatcustomocrtemplate.html
func (client *Client) CreatCustomOcrTemplate(request *CreatCustomOcrTemplateRequest) (response *CreatCustomOcrTemplateResponse, err error) {
	response = CreateCreatCustomOcrTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreatCustomOcrTemplateWithChan invokes the green.CreatCustomOcrTemplate API asynchronously
// api document: https://help.aliyun.com/api/green/creatcustomocrtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatCustomOcrTemplateWithChan(request *CreatCustomOcrTemplateRequest) (<-chan *CreatCustomOcrTemplateResponse, <-chan error) {
	responseChan := make(chan *CreatCustomOcrTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatCustomOcrTemplate(request)
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

// CreatCustomOcrTemplateWithCallback invokes the green.CreatCustomOcrTemplate API asynchronously
// api document: https://help.aliyun.com/api/green/creatcustomocrtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatCustomOcrTemplateWithCallback(request *CreatCustomOcrTemplateRequest, callback func(response *CreatCustomOcrTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatCustomOcrTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreatCustomOcrTemplate(request)
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

// CreatCustomOcrTemplateRequest is the request struct for api CreatCustomOcrTemplate
type CreatCustomOcrTemplateRequest struct {
	*requests.RpcRequest
	RecognizeArea string `position:"Query" name:"RecognizeArea"`
	ImgUrl        string `position:"Query" name:"ImgUrl"`
	SourceIp      string `position:"Query" name:"SourceIp"`
	ReferArea     string `position:"Query" name:"ReferArea"`
	Name          string `position:"Query" name:"Name"`
}

// CreatCustomOcrTemplateResponse is the response struct for api CreatCustomOcrTemplate
type CreatCustomOcrTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatCustomOcrTemplateRequest creates a request to invoke CreatCustomOcrTemplate API
func CreateCreatCustomOcrTemplateRequest() (request *CreatCustomOcrTemplateRequest) {
	request = &CreatCustomOcrTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "CreatCustomOcrTemplate", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatCustomOcrTemplateResponse creates a response to parse from CreatCustomOcrTemplate response
func CreateCreatCustomOcrTemplateResponse() (response *CreatCustomOcrTemplateResponse) {
	response = &CreatCustomOcrTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
