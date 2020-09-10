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

// DescribeLogBackupFiles invokes the rds.DescribeLogBackupFiles API synchronously
// api document: https://help.aliyun.com/api/rds/describelogbackupfiles.html
func (client *Client) DescribeLogBackupFiles(request *DescribeLogBackupFilesRequest) (response *DescribeLogBackupFilesResponse, err error) {
	response = CreateDescribeLogBackupFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogBackupFilesWithChan invokes the rds.DescribeLogBackupFiles API asynchronously
// api document: https://help.aliyun.com/api/rds/describelogbackupfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLogBackupFilesWithChan(request *DescribeLogBackupFilesRequest) (<-chan *DescribeLogBackupFilesResponse, <-chan error) {
	responseChan := make(chan *DescribeLogBackupFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogBackupFiles(request)
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

// DescribeLogBackupFilesWithCallback invokes the rds.DescribeLogBackupFiles API asynchronously
// api document: https://help.aliyun.com/api/rds/describelogbackupfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLogBackupFilesWithCallback(request *DescribeLogBackupFilesRequest, callback func(response *DescribeLogBackupFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogBackupFilesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogBackupFiles(request)
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

// DescribeLogBackupFilesRequest is the request struct for api DescribeLogBackupFiles
type DescribeLogBackupFilesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLogBackupFilesResponse is the response struct for api DescribeLogBackupFiles
type DescribeLogBackupFilesResponse struct {
	*responses.BaseResponse
	RequestId        string                        `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int                           `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                           `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                           `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalFileSize    int64                         `json:"TotalFileSize" xml:"TotalFileSize"`
	Items            ItemsInDescribeLogBackupFiles `json:"Items" xml:"Items"`
}

// CreateDescribeLogBackupFilesRequest creates a request to invoke DescribeLogBackupFiles API
func CreateDescribeLogBackupFilesRequest() (request *DescribeLogBackupFilesRequest) {
	request = &DescribeLogBackupFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeLogBackupFiles", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLogBackupFilesResponse creates a response to parse from DescribeLogBackupFiles response
func CreateDescribeLogBackupFilesResponse() (response *DescribeLogBackupFilesResponse) {
	response = &DescribeLogBackupFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
