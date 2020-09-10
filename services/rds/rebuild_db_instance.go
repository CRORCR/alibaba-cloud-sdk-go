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

// RebuildDBInstance invokes the rds.RebuildDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/rebuilddbinstance.html
func (client *Client) RebuildDBInstance(request *RebuildDBInstanceRequest) (response *RebuildDBInstanceResponse, err error) {
	response = CreateRebuildDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RebuildDBInstanceWithChan invokes the rds.RebuildDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/rebuilddbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RebuildDBInstanceWithChan(request *RebuildDBInstanceRequest) (<-chan *RebuildDBInstanceResponse, <-chan error) {
	responseChan := make(chan *RebuildDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebuildDBInstance(request)
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

// RebuildDBInstanceWithCallback invokes the rds.RebuildDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/rebuilddbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RebuildDBInstanceWithCallback(request *RebuildDBInstanceRequest, callback func(response *RebuildDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebuildDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.RebuildDBInstance(request)
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

// RebuildDBInstanceRequest is the request struct for api RebuildDBInstance
type RebuildDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	DedicatedHostGroupId string           `position:"Query" name:"DedicatedHostGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RebuildDBInstanceResponse is the response struct for api RebuildDBInstance
type RebuildDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TaskId      int    `json:"TaskId" xml:"TaskId"`
	MigrationId int    `json:"MigrationId" xml:"MigrationId"`
}

// CreateRebuildDBInstanceRequest creates a request to invoke RebuildDBInstance API
func CreateRebuildDBInstanceRequest() (request *RebuildDBInstanceRequest) {
	request = &RebuildDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RebuildDBInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRebuildDBInstanceResponse creates a response to parse from RebuildDBInstance response
func CreateRebuildDBInstanceResponse() (response *RebuildDBInstanceResponse) {
	response = &RebuildDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
