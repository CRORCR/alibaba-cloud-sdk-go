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

// ModifyReadonlyInstanceDelayReplicationTime invokes the rds.ModifyReadonlyInstanceDelayReplicationTime API synchronously
// api document: https://help.aliyun.com/api/rds/modifyreadonlyinstancedelayreplicationtime.html
func (client *Client) ModifyReadonlyInstanceDelayReplicationTime(request *ModifyReadonlyInstanceDelayReplicationTimeRequest) (response *ModifyReadonlyInstanceDelayReplicationTimeResponse, err error) {
	response = CreateModifyReadonlyInstanceDelayReplicationTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReadonlyInstanceDelayReplicationTimeWithChan invokes the rds.ModifyReadonlyInstanceDelayReplicationTime API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreadonlyinstancedelayreplicationtime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReadonlyInstanceDelayReplicationTimeWithChan(request *ModifyReadonlyInstanceDelayReplicationTimeRequest) (<-chan *ModifyReadonlyInstanceDelayReplicationTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyReadonlyInstanceDelayReplicationTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReadonlyInstanceDelayReplicationTime(request)
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

// ModifyReadonlyInstanceDelayReplicationTimeWithCallback invokes the rds.ModifyReadonlyInstanceDelayReplicationTime API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreadonlyinstancedelayreplicationtime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReadonlyInstanceDelayReplicationTimeWithCallback(request *ModifyReadonlyInstanceDelayReplicationTimeRequest, callback func(response *ModifyReadonlyInstanceDelayReplicationTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReadonlyInstanceDelayReplicationTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyReadonlyInstanceDelayReplicationTime(request)
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

// ModifyReadonlyInstanceDelayReplicationTimeRequest is the request struct for api ModifyReadonlyInstanceDelayReplicationTime
type ModifyReadonlyInstanceDelayReplicationTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ReadSQLReplicationTime string           `position:"Query" name:"ReadSQLReplicationTime"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
}

// ModifyReadonlyInstanceDelayReplicationTimeResponse is the response struct for api ModifyReadonlyInstanceDelayReplicationTime
type ModifyReadonlyInstanceDelayReplicationTimeResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	DBInstanceId           string `json:"DBInstanceId" xml:"DBInstanceId"`
	ReadSQLReplicationTime string `json:"ReadSQLReplicationTime" xml:"ReadSQLReplicationTime"`
	TaskId                 string `json:"TaskId" xml:"TaskId"`
}

// CreateModifyReadonlyInstanceDelayReplicationTimeRequest creates a request to invoke ModifyReadonlyInstanceDelayReplicationTime API
func CreateModifyReadonlyInstanceDelayReplicationTimeRequest() (request *ModifyReadonlyInstanceDelayReplicationTimeRequest) {
	request = &ModifyReadonlyInstanceDelayReplicationTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyReadonlyInstanceDelayReplicationTime", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyReadonlyInstanceDelayReplicationTimeResponse creates a response to parse from ModifyReadonlyInstanceDelayReplicationTime response
func CreateModifyReadonlyInstanceDelayReplicationTimeResponse() (response *ModifyReadonlyInstanceDelayReplicationTimeResponse) {
	response = &ModifyReadonlyInstanceDelayReplicationTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
