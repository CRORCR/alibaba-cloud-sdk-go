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

// ModifyDBInstanceSpec invokes the rds.ModifyDBInstanceSpec API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancespec.html
func (client *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
	response = CreateModifyDBInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceSpecWithChan invokes the rds.ModifyDBInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceSpecWithChan(request *ModifyDBInstanceSpecRequest) (<-chan *ModifyDBInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceSpec(request)
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

// ModifyDBInstanceSpecWithCallback invokes the rds.ModifyDBInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceSpecWithCallback(request *ModifyDBInstanceSpecRequest, callback func(response *ModifyDBInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceSpec(request)
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

// ModifyDBInstanceSpecRequest is the request struct for api ModifyDBInstanceSpec
type ModifyDBInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	EffectiveTime         string           `position:"Query" name:"EffectiveTime"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	DBInstanceStorageType string           `position:"Query" name:"DBInstanceStorageType"`
	SourceBiz             string           `position:"Query" name:"SourceBiz"`
	DedicatedHostGroupId  string           `position:"Query" name:"DedicatedHostGroupId"`
	Direction             string           `position:"Query" name:"Direction"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	PayType               string           `position:"Query" name:"PayType"`
}

// ModifyDBInstanceSpecResponse is the response struct for api ModifyDBInstanceSpec
type ModifyDBInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceSpecRequest creates a request to invoke ModifyDBInstanceSpec API
func CreateModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
	request = &ModifyDBInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSpec", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceSpecResponse creates a response to parse from ModifyDBInstanceSpec response
func CreateModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
	response = &ModifyDBInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
