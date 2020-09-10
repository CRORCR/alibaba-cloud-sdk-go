package yundun_bastionhost

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

// ConfigInstancePort invokes the yundun_bastionhost.ConfigInstancePort API synchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/configinstanceport.html
func (client *Client) ConfigInstancePort(request *ConfigInstancePortRequest) (response *ConfigInstancePortResponse, err error) {
	response = CreateConfigInstancePortResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigInstancePortWithChan invokes the yundun_bastionhost.ConfigInstancePort API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/configinstanceport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigInstancePortWithChan(request *ConfigInstancePortRequest) (<-chan *ConfigInstancePortResponse, <-chan error) {
	responseChan := make(chan *ConfigInstancePortResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigInstancePort(request)
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

// ConfigInstancePortWithCallback invokes the yundun_bastionhost.ConfigInstancePort API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/configinstanceport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigInstancePortWithCallback(request *ConfigInstancePortRequest, callback func(response *ConfigInstancePortResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigInstancePortResponse
		var err error
		defer close(result)
		response, err = client.ConfigInstancePort(request)
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

// ConfigInstancePortRequest is the request struct for api ConfigInstancePort
type ConfigInstancePortRequest struct {
	*requests.RpcRequest
	StandardPort requests.Integer `position:"Query" name:"StandardPort"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	CustomPort   requests.Integer `position:"Query" name:"CustomPort"`
	Lang         string           `position:"Query" name:"Lang"`
}

// ConfigInstancePortResponse is the response struct for api ConfigInstancePort
type ConfigInstancePortResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigInstancePortRequest creates a request to invoke ConfigInstancePort API
func CreateConfigInstancePortRequest() (request *ConfigInstancePortRequest) {
	request = &ConfigInstancePortRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-bastionhost", "2018-10-10", "ConfigInstancePort", "bastionhost", "openAPI")
	return
}

// CreateConfigInstancePortResponse creates a response to parse from ConfigInstancePort response
func CreateConfigInstancePortResponse() (response *ConfigInstancePortResponse) {
	response = &ConfigInstancePortResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
