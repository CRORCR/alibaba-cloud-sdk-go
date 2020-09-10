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

// BindApplicationToEdgeInstance invokes the iot.BindApplicationToEdgeInstance API synchronously
// api document: https://help.aliyun.com/api/iot/bindapplicationtoedgeinstance.html
func (client *Client) BindApplicationToEdgeInstance(request *BindApplicationToEdgeInstanceRequest) (response *BindApplicationToEdgeInstanceResponse, err error) {
	response = CreateBindApplicationToEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// BindApplicationToEdgeInstanceWithChan invokes the iot.BindApplicationToEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/bindapplicationtoedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindApplicationToEdgeInstanceWithChan(request *BindApplicationToEdgeInstanceRequest) (<-chan *BindApplicationToEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *BindApplicationToEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindApplicationToEdgeInstance(request)
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

// BindApplicationToEdgeInstanceWithCallback invokes the iot.BindApplicationToEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/bindapplicationtoedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindApplicationToEdgeInstanceWithCallback(request *BindApplicationToEdgeInstanceRequest, callback func(response *BindApplicationToEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindApplicationToEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.BindApplicationToEdgeInstance(request)
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

// BindApplicationToEdgeInstanceRequest is the request struct for api BindApplicationToEdgeInstance
type BindApplicationToEdgeInstanceRequest struct {
	*requests.RpcRequest
	ApplicationVersion string `position:"Query" name:"ApplicationVersion"`
	IotInstanceId      string `position:"Query" name:"IotInstanceId"`
	ApplicationId      string `position:"Query" name:"ApplicationId"`
	InstanceId         string `position:"Query" name:"InstanceId"`
	ApiProduct         string `position:"Body" name:"ApiProduct"`
	ApiRevision        string `position:"Body" name:"ApiRevision"`
}

// BindApplicationToEdgeInstanceResponse is the response struct for api BindApplicationToEdgeInstance
type BindApplicationToEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateBindApplicationToEdgeInstanceRequest creates a request to invoke BindApplicationToEdgeInstance API
func CreateBindApplicationToEdgeInstanceRequest() (request *BindApplicationToEdgeInstanceRequest) {
	request = &BindApplicationToEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BindApplicationToEdgeInstance", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindApplicationToEdgeInstanceResponse creates a response to parse from BindApplicationToEdgeInstance response
func CreateBindApplicationToEdgeInstanceResponse() (response *BindApplicationToEdgeInstanceResponse) {
	response = &BindApplicationToEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
