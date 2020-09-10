package polardb

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

// ModifyLogBackupPolicy invokes the polardb.ModifyLogBackupPolicy API synchronously
// api document: https://help.aliyun.com/api/polardb/modifylogbackuppolicy.html
func (client *Client) ModifyLogBackupPolicy(request *ModifyLogBackupPolicyRequest) (response *ModifyLogBackupPolicyResponse, err error) {
	response = CreateModifyLogBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLogBackupPolicyWithChan invokes the polardb.ModifyLogBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/polardb/modifylogbackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogBackupPolicyWithChan(request *ModifyLogBackupPolicyRequest) (<-chan *ModifyLogBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyLogBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLogBackupPolicy(request)
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

// ModifyLogBackupPolicyWithCallback invokes the polardb.ModifyLogBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/polardb/modifylogbackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogBackupPolicyWithCallback(request *ModifyLogBackupPolicyRequest, callback func(response *ModifyLogBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLogBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyLogBackupPolicy(request)
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

// ModifyLogBackupPolicyRequest is the request struct for api ModifyLogBackupPolicy
type ModifyLogBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId              string           `position:"Query" name:"DBClusterId"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	LogBackupRetentionPeriod string           `position:"Query" name:"LogBackupRetentionPeriod"`
}

// ModifyLogBackupPolicyResponse is the response struct for api ModifyLogBackupPolicy
type ModifyLogBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLogBackupPolicyRequest creates a request to invoke ModifyLogBackupPolicy API
func CreateModifyLogBackupPolicyRequest() (request *ModifyLogBackupPolicyRequest) {
	request = &ModifyLogBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyLogBackupPolicy", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLogBackupPolicyResponse creates a response to parse from ModifyLogBackupPolicy response
func CreateModifyLogBackupPolicyResponse() (response *ModifyLogBackupPolicyResponse) {
	response = &ModifyLogBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
