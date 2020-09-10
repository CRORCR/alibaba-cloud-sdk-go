package rds

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

// ModifyDBInstanceHAConfig invokes the rds.ModifyDBInstanceHAConfig API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancehaconfig.html
func (client *Client) ModifyDBInstanceHAConfig(request *ModifyDBInstanceHAConfigRequest) (response *ModifyDBInstanceHAConfigResponse, err error) {
	response = CreateModifyDBInstanceHAConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceHAConfigWithChan invokes the rds.ModifyDBInstanceHAConfig API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancehaconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceHAConfigWithChan(request *ModifyDBInstanceHAConfigRequest) (<-chan *ModifyDBInstanceHAConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceHAConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceHAConfig(request)
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

// ModifyDBInstanceHAConfigWithCallback invokes the rds.ModifyDBInstanceHAConfig API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancehaconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceHAConfigWithCallback(request *ModifyDBInstanceHAConfigRequest, callback func(response *ModifyDBInstanceHAConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceHAConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceHAConfig(request)
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

// ModifyDBInstanceHAConfigRequest is the request struct for api ModifyDBInstanceHAConfig
type ModifyDBInstanceHAConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DbInstanceId         string           `position:"Query" name:"DbInstanceId"`
	HAMode               string           `position:"Query" name:"HAMode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SyncMode             string           `position:"Query" name:"SyncMode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBInstanceHAConfigResponse is the response struct for api ModifyDBInstanceHAConfig
type ModifyDBInstanceHAConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceHAConfigRequest creates a request to invoke ModifyDBInstanceHAConfig API
func CreateModifyDBInstanceHAConfigRequest() (request *ModifyDBInstanceHAConfigRequest) {
	request = &ModifyDBInstanceHAConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceHAConfig", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceHAConfigResponse creates a response to parse from ModifyDBInstanceHAConfig response
func CreateModifyDBInstanceHAConfigResponse() (response *ModifyDBInstanceHAConfigResponse) {
	response = &ModifyDBInstanceHAConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
