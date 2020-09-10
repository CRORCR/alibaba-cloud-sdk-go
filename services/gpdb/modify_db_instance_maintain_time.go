package gpdb

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

// ModifyDBInstanceMaintainTime invokes the gpdb.ModifyDBInstanceMaintainTime API synchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstancemaintaintime.html
func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (response *ModifyDBInstanceMaintainTimeResponse, err error) {
	response = CreateModifyDBInstanceMaintainTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceMaintainTimeWithChan invokes the gpdb.ModifyDBInstanceMaintainTime API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstancemaintaintime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceMaintainTimeWithChan(request *ModifyDBInstanceMaintainTimeRequest) (<-chan *ModifyDBInstanceMaintainTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceMaintainTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceMaintainTime(request)
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

// ModifyDBInstanceMaintainTimeWithCallback invokes the gpdb.ModifyDBInstanceMaintainTime API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstancemaintaintime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceMaintainTimeWithCallback(request *ModifyDBInstanceMaintainTimeRequest, callback func(response *ModifyDBInstanceMaintainTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceMaintainTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceMaintainTime(request)
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

// ModifyDBInstanceMaintainTimeRequest is the request struct for api ModifyDBInstanceMaintainTime
type ModifyDBInstanceMaintainTimeRequest struct {
	*requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

// ModifyDBInstanceMaintainTimeResponse is the response struct for api ModifyDBInstanceMaintainTime
type ModifyDBInstanceMaintainTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceMaintainTimeRequest creates a request to invoke ModifyDBInstanceMaintainTime API
func CreateModifyDBInstanceMaintainTimeRequest() (request *ModifyDBInstanceMaintainTimeRequest) {
	request = &ModifyDBInstanceMaintainTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifyDBInstanceMaintainTime", "gpdb", "openAPI")
	return
}

// CreateModifyDBInstanceMaintainTimeResponse creates a response to parse from ModifyDBInstanceMaintainTime response
func CreateModifyDBInstanceMaintainTimeResponse() (response *ModifyDBInstanceMaintainTimeResponse) {
	response = &ModifyDBInstanceMaintainTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
