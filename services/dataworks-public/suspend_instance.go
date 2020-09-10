package dataworks_public

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

// SuspendInstance invokes the dataworks_public.SuspendInstance API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/suspendinstance.html
func (client *Client) SuspendInstance(request *SuspendInstanceRequest) (response *SuspendInstanceResponse, err error) {
	response = CreateSuspendInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendInstanceWithChan invokes the dataworks_public.SuspendInstance API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/suspendinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SuspendInstanceWithChan(request *SuspendInstanceRequest) (<-chan *SuspendInstanceResponse, <-chan error) {
	responseChan := make(chan *SuspendInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendInstance(request)
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

// SuspendInstanceWithCallback invokes the dataworks_public.SuspendInstance API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/suspendinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SuspendInstanceWithCallback(request *SuspendInstanceRequest, callback func(response *SuspendInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendInstanceResponse
		var err error
		defer close(result)
		response, err = client.SuspendInstance(request)
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

// SuspendInstanceRequest is the request struct for api SuspendInstance
type SuspendInstanceRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	InstanceId requests.Integer `position:"Body" name:"InstanceId"`
}

// SuspendInstanceResponse is the response struct for api SuspendInstance
type SuspendInstanceResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           bool   `json:"Data" xml:"Data"`
}

// CreateSuspendInstanceRequest creates a request to invoke SuspendInstance API
func CreateSuspendInstanceRequest() (request *SuspendInstanceRequest) {
	request = &SuspendInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "SuspendInstance", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendInstanceResponse creates a response to parse from SuspendInstance response
func CreateSuspendInstanceResponse() (response *SuspendInstanceResponse) {
	response = &SuspendInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
