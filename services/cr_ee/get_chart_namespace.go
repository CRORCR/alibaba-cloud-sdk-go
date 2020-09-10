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

// GetChartNamespace invokes the cr.GetChartNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/getchartnamespace.html
func (client *Client) GetChartNamespace(request *GetChartNamespaceRequest) (response *GetChartNamespaceResponse, err error) {
	response = CreateGetChartNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetChartNamespaceWithChan invokes the cr.GetChartNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getchartnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetChartNamespaceWithChan(request *GetChartNamespaceRequest) (<-chan *GetChartNamespaceResponse, <-chan error) {
	responseChan := make(chan *GetChartNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetChartNamespace(request)
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

// GetChartNamespaceWithCallback invokes the cr.GetChartNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getchartnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetChartNamespaceWithCallback(request *GetChartNamespaceRequest, callback func(response *GetChartNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetChartNamespaceResponse
		var err error
		defer close(result)
		response, err = client.GetChartNamespace(request)
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

// GetChartNamespaceRequest is the request struct for api GetChartNamespace
type GetChartNamespaceRequest struct {
	*requests.RpcRequest
	NamespaceName string `position:"Query" name:"NamespaceName"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// GetChartNamespaceResponse is the response struct for api GetChartNamespace
type GetChartNamespaceResponse struct {
	*responses.BaseResponse
	GetChartNamespaceIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                       string `json:"Code" xml:"Code"`
	RequestId                  string `json:"RequestId" xml:"RequestId"`
	NamespaceName              string `json:"NamespaceName" xml:"NamespaceName"`
	NamespaceStatus            string `json:"NamespaceStatus" xml:"NamespaceStatus"`
	AutoCreateRepo             bool   `json:"AutoCreateRepo" xml:"AutoCreateRepo"`
	DefaultRepoType            string `json:"DefaultRepoType" xml:"DefaultRepoType"`
	InstanceId                 string `json:"InstanceId" xml:"InstanceId"`
	NamespaceId                string `json:"NamespaceId" xml:"NamespaceId"`
}

// CreateGetChartNamespaceRequest creates a request to invoke GetChartNamespace API
func CreateGetChartNamespaceRequest() (request *GetChartNamespaceRequest) {
	request = &GetChartNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "GetChartNamespace", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetChartNamespaceResponse creates a response to parse from GetChartNamespace response
func CreateGetChartNamespaceResponse() (response *GetChartNamespaceResponse) {
	response = &GetChartNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
