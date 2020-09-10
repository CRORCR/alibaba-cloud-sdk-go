package cr_ee

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

// CreateChartNamespace invokes the cr.CreateChartNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/createchartnamespace.html
func (client *Client) CreateChartNamespace(request *CreateChartNamespaceRequest) (response *CreateChartNamespaceResponse, err error) {
	response = CreateCreateChartNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateChartNamespaceWithChan invokes the cr.CreateChartNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/createchartnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChartNamespaceWithChan(request *CreateChartNamespaceRequest) (<-chan *CreateChartNamespaceResponse, <-chan error) {
	responseChan := make(chan *CreateChartNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateChartNamespace(request)
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

// CreateChartNamespaceWithCallback invokes the cr.CreateChartNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/createchartnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChartNamespaceWithCallback(request *CreateChartNamespaceRequest, callback func(response *CreateChartNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateChartNamespaceResponse
		var err error
		defer close(result)
		response, err = client.CreateChartNamespace(request)
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

// CreateChartNamespaceRequest is the request struct for api CreateChartNamespace
type CreateChartNamespaceRequest struct {
	*requests.RpcRequest
	NamespaceName   string           `position:"Query" name:"NamespaceName"`
	AutoCreateRepo  requests.Boolean `position:"Query" name:"AutoCreateRepo"`
	DefaultRepoType string           `position:"Query" name:"DefaultRepoType"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
}

// CreateChartNamespaceResponse is the response struct for api CreateChartNamespace
type CreateChartNamespaceResponse struct {
	*responses.BaseResponse
	CreateChartNamespaceIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                          string `json:"Code" xml:"Code"`
	RequestId                     string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateChartNamespaceRequest creates a request to invoke CreateChartNamespace API
func CreateCreateChartNamespaceRequest() (request *CreateChartNamespaceRequest) {
	request = &CreateChartNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "CreateChartNamespace", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateChartNamespaceResponse creates a response to parse from CreateChartNamespace response
func CreateCreateChartNamespaceResponse() (response *CreateChartNamespaceResponse) {
	response = &CreateChartNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
