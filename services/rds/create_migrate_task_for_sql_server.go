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

// CreateMigrateTaskForSQLServer invokes the rds.CreateMigrateTaskForSQLServer API synchronously
// api document: https://help.aliyun.com/api/rds/createmigratetaskforsqlserver.html
func (client *Client) CreateMigrateTaskForSQLServer(request *CreateMigrateTaskForSQLServerRequest) (response *CreateMigrateTaskForSQLServerResponse, err error) {
	response = CreateCreateMigrateTaskForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMigrateTaskForSQLServerWithChan invokes the rds.CreateMigrateTaskForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/createmigratetaskforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMigrateTaskForSQLServerWithChan(request *CreateMigrateTaskForSQLServerRequest) (<-chan *CreateMigrateTaskForSQLServerResponse, <-chan error) {
	responseChan := make(chan *CreateMigrateTaskForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMigrateTaskForSQLServer(request)
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

// CreateMigrateTaskForSQLServerWithCallback invokes the rds.CreateMigrateTaskForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/createmigratetaskforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMigrateTaskForSQLServerWithCallback(request *CreateMigrateTaskForSQLServerRequest, callback func(response *CreateMigrateTaskForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMigrateTaskForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.CreateMigrateTaskForSQLServer(request)
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

// CreateMigrateTaskForSQLServerRequest is the request struct for api CreateMigrateTaskForSQLServer
type CreateMigrateTaskForSQLServerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IsOnlineDB           string           `position:"Query" name:"IsOnlineDB"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	TaskType             string           `position:"Query" name:"TaskType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OSSUrls              string           `position:"Query" name:"OSSUrls"`
	DBName               string           `position:"Query" name:"DBName"`
}

// CreateMigrateTaskForSQLServerResponse is the response struct for api CreateMigrateTaskForSQLServer
type CreateMigrateTaskForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceId   string `json:"DBInstanceId" xml:"DBInstanceId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	DBName         string `json:"DBName" xml:"DBName"`
	MigrateIaskId  string `json:"MigrateIaskId" xml:"MigrateIaskId"`
	TaskType       string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateMigrateTaskForSQLServerRequest creates a request to invoke CreateMigrateTaskForSQLServer API
func CreateCreateMigrateTaskForSQLServerRequest() (request *CreateMigrateTaskForSQLServerRequest) {
	request = &CreateMigrateTaskForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateMigrateTaskForSQLServer", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMigrateTaskForSQLServerResponse creates a response to parse from CreateMigrateTaskForSQLServer response
func CreateCreateMigrateTaskForSQLServerResponse() (response *CreateMigrateTaskForSQLServerResponse) {
	response = &CreateMigrateTaskForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
