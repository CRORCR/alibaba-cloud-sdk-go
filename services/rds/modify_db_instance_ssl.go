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

// ModifyDBInstanceSSL invokes the rds.ModifyDBInstanceSSL API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancessl.html
func (client *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (response *ModifyDBInstanceSSLResponse, err error) {
	response = CreateModifyDBInstanceSSLResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceSSLWithChan invokes the rds.ModifyDBInstanceSSL API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancessl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceSSLWithChan(request *ModifyDBInstanceSSLRequest) (<-chan *ModifyDBInstanceSSLResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceSSL(request)
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

// ModifyDBInstanceSSLWithCallback invokes the rds.ModifyDBInstanceSSL API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstancessl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceSSLWithCallback(request *ModifyDBInstanceSSLRequest, callback func(response *ModifyDBInstanceSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceSSLResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceSSL(request)
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

// ModifyDBInstanceSSLRequest is the request struct for api ModifyDBInstanceSSL
type ModifyDBInstanceSSLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionString     string           `position:"Query" name:"ConnectionString"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SSLEnabled           requests.Integer `position:"Query" name:"SSLEnabled"`
}

// ModifyDBInstanceSSLResponse is the response struct for api ModifyDBInstanceSSL
type ModifyDBInstanceSSLResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceSSLRequest creates a request to invoke ModifyDBInstanceSSL API
func CreateModifyDBInstanceSSLRequest() (request *ModifyDBInstanceSSLRequest) {
	request = &ModifyDBInstanceSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSSL", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceSSLResponse creates a response to parse from ModifyDBInstanceSSL response
func CreateModifyDBInstanceSSLResponse() (response *ModifyDBInstanceSSLResponse) {
	response = &ModifyDBInstanceSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
