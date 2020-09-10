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

package dts

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteMigrationJob invokes the dts.DeleteMigrationJob API synchronously
// api document: https://help.aliyun.com/api/dts/deletemigrationjob.html
func (client *Client) DeleteMigrationJob(request *DeleteMigrationJobRequest) (response *DeleteMigrationJobResponse, err error) {
	response = CreateDeleteMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMigrationJobWithChan invokes the dts.DeleteMigrationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/deletemigrationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMigrationJobWithChan(request *DeleteMigrationJobRequest) (<-chan *DeleteMigrationJobResponse, <-chan error) {
	responseChan := make(chan *DeleteMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMigrationJob(request)
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

// DeleteMigrationJobWithCallback invokes the dts.DeleteMigrationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/deletemigrationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMigrationJobWithCallback(request *DeleteMigrationJobRequest, callback func(response *DeleteMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteMigrationJob(request)
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

// DeleteMigrationJobRequest is the request struct for api DeleteMigrationJob
type DeleteMigrationJobRequest struct {
	*requests.RpcRequest
	MigrationJobId string `position:"Query" name:"MigrationJobId"`
	OwnerId        string `position:"Query" name:"OwnerId"`
}

// DeleteMigrationJobResponse is the response struct for api DeleteMigrationJob
type DeleteMigrationJobResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteMigrationJobRequest creates a request to invoke DeleteMigrationJob API
func CreateDeleteMigrationJobRequest() (request *DeleteMigrationJobRequest) {
	request = &DeleteMigrationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "DeleteMigrationJob", "dts", "openAPI")
	return
}

// CreateDeleteMigrationJobResponse creates a response to parse from DeleteMigrationJob response
func CreateDeleteMigrationJobResponse() (response *DeleteMigrationJobResponse) {
	response = &DeleteMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
