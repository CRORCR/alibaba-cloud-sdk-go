package cs

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

// CreateTemplate invokes the cs.CreateTemplate API synchronously
// api document: https://help.aliyun.com/api/cs/createtemplate.html
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	response = CreateCreateTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTemplateWithChan invokes the cs.CreateTemplate API asynchronously
// api document: https://help.aliyun.com/api/cs/createtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTemplateWithChan(request *CreateTemplateRequest) (<-chan *CreateTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTemplate(request)
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

// CreateTemplateWithCallback invokes the cs.CreateTemplate API asynchronously
// api document: https://help.aliyun.com/api/cs/createtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTemplateWithCallback(request *CreateTemplateRequest, callback func(response *CreateTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateTemplate(request)
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

// CreateTemplateRequest is the request struct for api CreateTemplate
type CreateTemplateRequest struct {
	*requests.RoaRequest
	Template     string `position:"Body" name:"template"`
	Name         string `position:"Body" name:"name"`
	TemplateType string `position:"Body" name:"template_type"`
	Tags         string `position:"Body" name:"tags"`
}

// CreateTemplateResponse is the response struct for api CreateTemplate
type CreateTemplateResponse struct {
	*responses.BaseResponse
	TemplateId string `json:"template_id" xml:"template_id"`
}

// CreateCreateTemplateRequest creates a request to invoke CreateTemplate API
func CreateCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "CreateTemplate", "/templates", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTemplateResponse creates a response to parse from CreateTemplate response
func CreateCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
