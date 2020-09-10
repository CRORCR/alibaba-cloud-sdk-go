package r_kvstore

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

// ModifySecurityGroupConfiguration invokes the r_kvstore.ModifySecurityGroupConfiguration API synchronously
func (client *Client) ModifySecurityGroupConfiguration(request *ModifySecurityGroupConfigurationRequest) (response *ModifySecurityGroupConfigurationResponse, err error) {
	response = CreateModifySecurityGroupConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySecurityGroupConfigurationWithChan invokes the r_kvstore.ModifySecurityGroupConfiguration API asynchronously
func (client *Client) ModifySecurityGroupConfigurationWithChan(request *ModifySecurityGroupConfigurationRequest) (<-chan *ModifySecurityGroupConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityGroupConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityGroupConfiguration(request)
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

// ModifySecurityGroupConfigurationWithCallback invokes the r_kvstore.ModifySecurityGroupConfiguration API asynchronously
func (client *Client) ModifySecurityGroupConfigurationWithCallback(request *ModifySecurityGroupConfigurationRequest, callback func(response *ModifySecurityGroupConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityGroupConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityGroupConfiguration(request)
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

// ModifySecurityGroupConfigurationRequest is the request struct for api ModifySecurityGroupConfiguration
type ModifySecurityGroupConfigurationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupId      string           `position:"Query" name:"SecurityGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifySecurityGroupConfigurationResponse is the response struct for api ModifySecurityGroupConfiguration
type ModifySecurityGroupConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySecurityGroupConfigurationRequest creates a request to invoke ModifySecurityGroupConfiguration API
func CreateModifySecurityGroupConfigurationRequest() (request *ModifySecurityGroupConfigurationRequest) {
	request = &ModifySecurityGroupConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifySecurityGroupConfiguration", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySecurityGroupConfigurationResponse creates a response to parse from ModifySecurityGroupConfiguration response
func CreateModifySecurityGroupConfigurationResponse() (response *ModifySecurityGroupConfigurationResponse) {
	response = &ModifySecurityGroupConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
