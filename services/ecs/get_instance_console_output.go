package ecs

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

// GetInstanceConsoleOutput invokes the ecs.GetInstanceConsoleOutput API synchronously
// api document: https://help.aliyun.com/api/ecs/getinstanceconsoleoutput.html
func (client *Client) GetInstanceConsoleOutput(request *GetInstanceConsoleOutputRequest) (response *GetInstanceConsoleOutputResponse, err error) {
	response = CreateGetInstanceConsoleOutputResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceConsoleOutputWithChan invokes the ecs.GetInstanceConsoleOutput API asynchronously
// api document: https://help.aliyun.com/api/ecs/getinstanceconsoleoutput.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceConsoleOutputWithChan(request *GetInstanceConsoleOutputRequest) (<-chan *GetInstanceConsoleOutputResponse, <-chan error) {
	responseChan := make(chan *GetInstanceConsoleOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceConsoleOutput(request)
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

// GetInstanceConsoleOutputWithCallback invokes the ecs.GetInstanceConsoleOutput API asynchronously
// api document: https://help.aliyun.com/api/ecs/getinstanceconsoleoutput.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceConsoleOutputWithCallback(request *GetInstanceConsoleOutputRequest, callback func(response *GetInstanceConsoleOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceConsoleOutputResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceConsoleOutput(request)
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

// GetInstanceConsoleOutputRequest is the request struct for api GetInstanceConsoleOutput
type GetInstanceConsoleOutputRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RemoveSymbols        requests.Boolean `position:"Query" name:"RemoveSymbols"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// GetInstanceConsoleOutputResponse is the response struct for api GetInstanceConsoleOutput
type GetInstanceConsoleOutputResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	ConsoleOutput  string `json:"ConsoleOutput" xml:"ConsoleOutput"`
	LastUpdateTime string `json:"LastUpdateTime" xml:"LastUpdateTime"`
}

// CreateGetInstanceConsoleOutputRequest creates a request to invoke GetInstanceConsoleOutput API
func CreateGetInstanceConsoleOutputRequest() (request *GetInstanceConsoleOutputRequest) {
	request = &GetInstanceConsoleOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "GetInstanceConsoleOutput", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetInstanceConsoleOutputResponse creates a response to parse from GetInstanceConsoleOutput response
func CreateGetInstanceConsoleOutputResponse() (response *GetInstanceConsoleOutputResponse) {
	response = &GetInstanceConsoleOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
