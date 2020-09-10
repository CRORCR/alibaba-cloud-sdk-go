package acm

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

// DeployConfiguration invokes the acm.DeployConfiguration API synchronously
// api document: https://help.aliyun.com/api/acm/deployconfiguration.html
func (client *Client) DeployConfiguration(request *DeployConfigurationRequest) (response *DeployConfigurationResponse, err error) {
	response = CreateDeployConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// DeployConfigurationWithChan invokes the acm.DeployConfiguration API asynchronously
// api document: https://help.aliyun.com/api/acm/deployconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeployConfigurationWithChan(request *DeployConfigurationRequest) (<-chan *DeployConfigurationResponse, <-chan error) {
	responseChan := make(chan *DeployConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployConfiguration(request)
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

// DeployConfigurationWithCallback invokes the acm.DeployConfiguration API asynchronously
// api document: https://help.aliyun.com/api/acm/deployconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeployConfigurationWithCallback(request *DeployConfigurationRequest, callback func(response *DeployConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployConfigurationResponse
		var err error
		defer close(result)
		response, err = client.DeployConfiguration(request)
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

// DeployConfigurationRequest is the request struct for api DeployConfiguration
type DeployConfigurationRequest struct {
	*requests.RoaRequest
	DataId      string `position:"Body" name:"DataId"`
	AppName     string `position:"Body" name:"AppName"`
	NamespaceId string `position:"Body" name:"NamespaceId"`
	Type        string `position:"Body" name:"Type"`
	Content     string `position:"Body" name:"Content"`
	Group       string `position:"Body" name:"Group"`
	Desc        string `position:"Body" name:"Desc"`
	Tags        string `position:"Body" name:"Tags"`
}

// DeployConfigurationResponse is the response struct for api DeployConfiguration
type DeployConfigurationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeployConfigurationRequest creates a request to invoke DeployConfiguration API
func CreateDeployConfigurationRequest() (request *DeployConfigurationRequest) {
	request = &DeployConfigurationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("acm", "2020-02-06", "DeployConfiguration", "/diamond-ops/pop/configuration", "acms", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateDeployConfigurationResponse creates a response to parse from DeployConfiguration response
func CreateDeployConfigurationResponse() (response *DeployConfigurationResponse) {
	response = &DeployConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
