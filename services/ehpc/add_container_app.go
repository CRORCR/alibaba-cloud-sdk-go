package ehpc

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

// AddContainerApp invokes the ehpc.AddContainerApp API synchronously
// api document: https://help.aliyun.com/api/ehpc/addcontainerapp.html
func (client *Client) AddContainerApp(request *AddContainerAppRequest) (response *AddContainerAppResponse, err error) {
	response = CreateAddContainerAppResponse()
	err = client.DoAction(request, response)
	return
}

// AddContainerAppWithChan invokes the ehpc.AddContainerApp API asynchronously
// api document: https://help.aliyun.com/api/ehpc/addcontainerapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddContainerAppWithChan(request *AddContainerAppRequest) (<-chan *AddContainerAppResponse, <-chan error) {
	responseChan := make(chan *AddContainerAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddContainerApp(request)
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

// AddContainerAppWithCallback invokes the ehpc.AddContainerApp API asynchronously
// api document: https://help.aliyun.com/api/ehpc/addcontainerapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddContainerAppWithCallback(request *AddContainerAppRequest, callback func(response *AddContainerAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddContainerAppResponse
		var err error
		defer close(result)
		response, err = client.AddContainerApp(request)
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

// AddContainerAppRequest is the request struct for api AddContainerApp
type AddContainerAppRequest struct {
	*requests.RpcRequest
	Description   string `position:"Query" name:"Description"`
	Repository    string `position:"Query" name:"Repository"`
	ContainerType string `position:"Query" name:"ContainerType"`
	Name          string `position:"Query" name:"Name"`
	ImageTag      string `position:"Query" name:"ImageTag"`
}

// AddContainerAppResponse is the response struct for api AddContainerApp
type AddContainerAppResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ContainerId ContainerId `json:"ContainerId" xml:"ContainerId"`
}

// CreateAddContainerAppRequest creates a request to invoke AddContainerApp API
func CreateAddContainerAppRequest() (request *AddContainerAppRequest) {
	request = &AddContainerAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "AddContainerApp", "", "")
	request.Method = requests.GET
	return
}

// CreateAddContainerAppResponse creates a response to parse from AddContainerApp response
func CreateAddContainerAppResponse() (response *AddContainerAppResponse) {
	response = &AddContainerAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
