package edas

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

// UpdateK8sApplicationConfig invokes the edas.UpdateK8sApplicationConfig API synchronously
// api document: https://help.aliyun.com/api/edas/updatek8sapplicationconfig.html
func (client *Client) UpdateK8sApplicationConfig(request *UpdateK8sApplicationConfigRequest) (response *UpdateK8sApplicationConfigResponse, err error) {
	response = CreateUpdateK8sApplicationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateK8sApplicationConfigWithChan invokes the edas.UpdateK8sApplicationConfig API asynchronously
// api document: https://help.aliyun.com/api/edas/updatek8sapplicationconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateK8sApplicationConfigWithChan(request *UpdateK8sApplicationConfigRequest) (<-chan *UpdateK8sApplicationConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateK8sApplicationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateK8sApplicationConfig(request)
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

// UpdateK8sApplicationConfigWithCallback invokes the edas.UpdateK8sApplicationConfig API asynchronously
// api document: https://help.aliyun.com/api/edas/updatek8sapplicationconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateK8sApplicationConfigWithCallback(request *UpdateK8sApplicationConfigRequest, callback func(response *UpdateK8sApplicationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateK8sApplicationConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateK8sApplicationConfig(request)
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

// UpdateK8sApplicationConfigRequest is the request struct for api UpdateK8sApplicationConfig
type UpdateK8sApplicationConfigRequest struct {
	*requests.RoaRequest
	MemoryRequest string           `position:"Query" name:"MemoryRequest"`
	AppId         string           `position:"Query" name:"AppId"`
	CpuRequest    string           `position:"Query" name:"CpuRequest"`
	MemoryLimit   string           `position:"Query" name:"MemoryLimit"`
	ClusterId     string           `position:"Query" name:"ClusterId"`
	CpuLimit      string           `position:"Query" name:"CpuLimit"`
	McpuLimit     string           `position:"Query" name:"McpuLimit"`
	McpuRequest   string           `position:"Query" name:"McpuRequest"`
	Timeout       requests.Integer `position:"Query" name:"Timeout"`
}

// UpdateK8sApplicationConfigResponse is the response struct for api UpdateK8sApplicationConfig
type UpdateK8sApplicationConfigResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateK8sApplicationConfigRequest creates a request to invoke UpdateK8sApplicationConfig API
func CreateUpdateK8sApplicationConfigRequest() (request *UpdateK8sApplicationConfigRequest) {
	request = &UpdateK8sApplicationConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateK8sApplicationConfig", "/pop/v5/k8s/acs/k8s_app_configuration", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateK8sApplicationConfigResponse creates a response to parse from UpdateK8sApplicationConfig response
func CreateUpdateK8sApplicationConfigResponse() (response *UpdateK8sApplicationConfigResponse) {
	response = &UpdateK8sApplicationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
