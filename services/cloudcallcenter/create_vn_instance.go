package cloudcallcenter

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

// CreateVnInstance invokes the cloudcallcenter.CreateVnInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createvninstance.html
func (client *Client) CreateVnInstance(request *CreateVnInstanceRequest) (response *CreateVnInstanceResponse, err error) {
	response = CreateCreateVnInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVnInstanceWithChan invokes the cloudcallcenter.CreateVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createvninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVnInstanceWithChan(request *CreateVnInstanceRequest) (<-chan *CreateVnInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateVnInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVnInstance(request)
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

// CreateVnInstanceWithCallback invokes the cloudcallcenter.CreateVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createvninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVnInstanceWithCallback(request *CreateVnInstanceRequest, callback func(response *CreateVnInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVnInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateVnInstance(request)
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

// CreateVnInstanceRequest is the request struct for api CreateVnInstance
type CreateVnInstanceRequest struct {
	*requests.RpcRequest
	Description       string           `position:"Query" name:"Description"`
	Concurrency       requests.Integer `position:"Query" name:"Concurrency"`
	ChatbotInstanceId string           `position:"Query" name:"ChatbotInstanceId"`
	Name              string           `position:"Query" name:"Name"`
	NluServiceType    string           `position:"Query" name:"NluServiceType"`
}

// CreateVnInstanceResponse is the response struct for api CreateVnInstance
type CreateVnInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
}

// CreateCreateVnInstanceRequest creates a request to invoke CreateVnInstance API
func CreateCreateVnInstanceRequest() (request *CreateVnInstanceRequest) {
	request = &CreateVnInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateVnInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateCreateVnInstanceResponse creates a response to parse from CreateVnInstance response
func CreateCreateVnInstanceResponse() (response *CreateVnInstanceResponse) {
	response = &CreateVnInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
