package emr

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

// ModifyResourcePool invokes the emr.ModifyResourcePool API synchronously
// api document: https://help.aliyun.com/api/emr/modifyresourcepool.html
func (client *Client) ModifyResourcePool(request *ModifyResourcePoolRequest) (response *ModifyResourcePoolResponse, err error) {
	response = CreateModifyResourcePoolResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyResourcePoolWithChan invokes the emr.ModifyResourcePool API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyresourcepool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyResourcePoolWithChan(request *ModifyResourcePoolRequest) (<-chan *ModifyResourcePoolResponse, <-chan error) {
	responseChan := make(chan *ModifyResourcePoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyResourcePool(request)
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

// ModifyResourcePoolWithCallback invokes the emr.ModifyResourcePool API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyresourcepool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyResourcePoolWithCallback(request *ModifyResourcePoolRequest, callback func(response *ModifyResourcePoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyResourcePoolResponse
		var err error
		defer close(result)
		response, err = client.ModifyResourcePool(request)
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

// ModifyResourcePoolRequest is the request struct for api ModifyResourcePool
type ModifyResourcePoolRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer            `position:"Query" name:"ResourceOwnerId"`
	Active          requests.Boolean            `position:"Query" name:"Active"`
	ClusterId       string                      `position:"Query" name:"ClusterId"`
	Yarnsiteconfig  string                      `position:"Query" name:"Yarnsiteconfig"`
	Name            string                      `position:"Query" name:"Name"`
	Id              string                      `position:"Query" name:"Id"`
	Config          *[]ModifyResourcePoolConfig `position:"Query" name:"Config"  type:"Repeated"`
}

// ModifyResourcePoolConfig is a repeated param struct in ModifyResourcePoolRequest
type ModifyResourcePoolConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	Note        string `name:"Note"`
	ConfigValue string `name:"ConfigValue"`
	Id          string `name:"Id"`
	Category    string `name:"Category"`
}

// ModifyResourcePoolResponse is the response struct for api ModifyResourcePool
type ModifyResourcePoolResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyResourcePoolRequest creates a request to invoke ModifyResourcePool API
func CreateModifyResourcePoolRequest() (request *ModifyResourcePoolRequest) {
	request = &ModifyResourcePoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyResourcePool", "emr", "openAPI")
	return
}

// CreateModifyResourcePoolResponse creates a response to parse from ModifyResourcePool response
func CreateModifyResourcePoolResponse() (response *ModifyResourcePoolResponse) {
	response = &ModifyResourcePoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
