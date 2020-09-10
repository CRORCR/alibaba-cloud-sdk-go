package emr

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

// CreateBackup invokes the emr.CreateBackup API synchronously
// api document: https://help.aliyun.com/api/emr/createbackup.html
func (client *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
	response = CreateCreateBackupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackupWithChan invokes the emr.CreateBackup API asynchronously
// api document: https://help.aliyun.com/api/emr/createbackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBackupWithChan(request *CreateBackupRequest) (<-chan *CreateBackupResponse, <-chan error) {
	responseChan := make(chan *CreateBackupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackup(request)
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

// CreateBackupWithCallback invokes the emr.CreateBackup API asynchronously
// api document: https://help.aliyun.com/api/emr/createbackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBackupWithCallback(request *CreateBackupRequest, callback func(response *CreateBackupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackupResponse
		var err error
		defer close(result)
		response, err = client.CreateBackup(request)
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

// CreateBackupRequest is the request struct for api CreateBackup
type CreateBackupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BackupPlanId    string           `position:"Query" name:"BackupPlanId"`
	MetadataType    string           `position:"Query" name:"MetadataType"`
}

// CreateBackupResponse is the response struct for api CreateBackup
type CreateBackupResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	BizId            string `json:"BizId" xml:"BizId"`
	DataSourceId     int64  `json:"DataSourceId" xml:"DataSourceId"`
	TaskType         string `json:"TaskType" xml:"TaskType"`
	TaskStatus       string `json:"TaskStatus" xml:"TaskStatus"`
	StartTime        int64  `json:"StartTime" xml:"StartTime"`
	EndTime          int64  `json:"EndTime" xml:"EndTime"`
	TaskDetail       string `json:"TaskDetail" xml:"TaskDetail"`
	TaskResultDetail string `json:"TaskResultDetail" xml:"TaskResultDetail"`
	TaskProcess      int    `json:"TaskProcess" xml:"TaskProcess"`
	TriggerUser      string `json:"TriggerUser" xml:"TriggerUser"`
	TriggerType      string `json:"TriggerType" xml:"TriggerType"`
	GmtCreate        int64  `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified      int64  `json:"GmtModified" xml:"GmtModified"`
	ClusterBizId     string `json:"ClusterBizId" xml:"ClusterBizId"`
	HostName         string `json:"HostName" xml:"HostName"`
	EcmTaskId        int64  `json:"EcmTaskId" xml:"EcmTaskId"`
}

// CreateCreateBackupRequest creates a request to invoke CreateBackup API
func CreateCreateBackupRequest() (request *CreateBackupRequest) {
	request = &CreateBackupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateBackup", "emr", "openAPI")
	return
}

// CreateCreateBackupResponse creates a response to parse from CreateBackup response
func CreateCreateBackupResponse() (response *CreateBackupResponse) {
	response = &CreateBackupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
