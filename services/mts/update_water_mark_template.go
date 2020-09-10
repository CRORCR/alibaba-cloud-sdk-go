package mts

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

// UpdateWaterMarkTemplate invokes the mts.UpdateWaterMarkTemplate API synchronously
// api document: https://help.aliyun.com/api/mts/updatewatermarktemplate.html
func (client *Client) UpdateWaterMarkTemplate(request *UpdateWaterMarkTemplateRequest) (response *UpdateWaterMarkTemplateResponse, err error) {
	response = CreateUpdateWaterMarkTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWaterMarkTemplateWithChan invokes the mts.UpdateWaterMarkTemplate API asynchronously
// api document: https://help.aliyun.com/api/mts/updatewatermarktemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWaterMarkTemplateWithChan(request *UpdateWaterMarkTemplateRequest) (<-chan *UpdateWaterMarkTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateWaterMarkTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWaterMarkTemplate(request)
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

// UpdateWaterMarkTemplateWithCallback invokes the mts.UpdateWaterMarkTemplate API asynchronously
// api document: https://help.aliyun.com/api/mts/updatewatermarktemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWaterMarkTemplateWithCallback(request *UpdateWaterMarkTemplateRequest, callback func(response *UpdateWaterMarkTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWaterMarkTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateWaterMarkTemplate(request)
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

// UpdateWaterMarkTemplateRequest is the request struct for api UpdateWaterMarkTemplate
type UpdateWaterMarkTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	WaterMarkTemplateId  string           `position:"Query" name:"WaterMarkTemplateId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	Config               string           `position:"Query" name:"Config"`
}

// UpdateWaterMarkTemplateResponse is the response struct for api UpdateWaterMarkTemplate
type UpdateWaterMarkTemplateResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	WaterMarkTemplate WaterMarkTemplate `json:"WaterMarkTemplate" xml:"WaterMarkTemplate"`
}

// CreateUpdateWaterMarkTemplateRequest creates a request to invoke UpdateWaterMarkTemplate API
func CreateUpdateWaterMarkTemplateRequest() (request *UpdateWaterMarkTemplateRequest) {
	request = &UpdateWaterMarkTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateWaterMarkTemplate", "", "")
	return
}

// CreateUpdateWaterMarkTemplateResponse creates a response to parse from UpdateWaterMarkTemplate response
func CreateUpdateWaterMarkTemplateResponse() (response *UpdateWaterMarkTemplateResponse) {
	response = &UpdateWaterMarkTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
