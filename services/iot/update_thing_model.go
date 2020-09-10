package iot

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

// UpdateThingModel invokes the iot.UpdateThingModel API synchronously
// api document: https://help.aliyun.com/api/iot/updatethingmodel.html
func (client *Client) UpdateThingModel(request *UpdateThingModelRequest) (response *UpdateThingModelResponse, err error) {
	response = CreateUpdateThingModelResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateThingModelWithChan invokes the iot.UpdateThingModel API asynchronously
// api document: https://help.aliyun.com/api/iot/updatethingmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateThingModelWithChan(request *UpdateThingModelRequest) (<-chan *UpdateThingModelResponse, <-chan error) {
	responseChan := make(chan *UpdateThingModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateThingModel(request)
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

// UpdateThingModelWithCallback invokes the iot.UpdateThingModel API asynchronously
// api document: https://help.aliyun.com/api/iot/updatethingmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateThingModelWithCallback(request *UpdateThingModelRequest, callback func(response *UpdateThingModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateThingModelResponse
		var err error
		defer close(result)
		response, err = client.UpdateThingModel(request)
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

// UpdateThingModelRequest is the request struct for api UpdateThingModel
type UpdateThingModelRequest struct {
	*requests.RpcRequest
	IotInstanceId  string `position:"Query" name:"IotInstanceId"`
	Identifier     string `position:"Query" name:"Identifier"`
	ProductKey     string `position:"Query" name:"ProductKey"`
	ApiProduct     string `position:"Body" name:"ApiProduct"`
	ThingModelJson string `position:"Query" name:"ThingModelJson"`
	ApiRevision    string `position:"Body" name:"ApiRevision"`
}

// UpdateThingModelResponse is the response struct for api UpdateThingModel
type UpdateThingModelResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateThingModelRequest creates a request to invoke UpdateThingModel API
func CreateUpdateThingModelRequest() (request *UpdateThingModelRequest) {
	request = &UpdateThingModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateThingModel", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateThingModelResponse creates a response to parse from UpdateThingModel response
func CreateUpdateThingModelResponse() (response *UpdateThingModelResponse) {
	response = &UpdateThingModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
