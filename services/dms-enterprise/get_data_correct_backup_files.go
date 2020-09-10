package dms_enterprise

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

// GetDataCorrectBackupFiles invokes the dms_enterprise.GetDataCorrectBackupFiles API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatacorrectbackupfiles.html
func (client *Client) GetDataCorrectBackupFiles(request *GetDataCorrectBackupFilesRequest) (response *GetDataCorrectBackupFilesResponse, err error) {
	response = CreateGetDataCorrectBackupFilesResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataCorrectBackupFilesWithChan invokes the dms_enterprise.GetDataCorrectBackupFiles API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatacorrectbackupfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataCorrectBackupFilesWithChan(request *GetDataCorrectBackupFilesRequest) (<-chan *GetDataCorrectBackupFilesResponse, <-chan error) {
	responseChan := make(chan *GetDataCorrectBackupFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataCorrectBackupFiles(request)
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

// GetDataCorrectBackupFilesWithCallback invokes the dms_enterprise.GetDataCorrectBackupFiles API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatacorrectbackupfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataCorrectBackupFilesWithCallback(request *GetDataCorrectBackupFilesRequest, callback func(response *GetDataCorrectBackupFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataCorrectBackupFilesResponse
		var err error
		defer close(result)
		response, err = client.GetDataCorrectBackupFiles(request)
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

// GetDataCorrectBackupFilesRequest is the request struct for api GetDataCorrectBackupFiles
type GetDataCorrectBackupFilesRequest struct {
	*requests.RpcRequest
	ActionDetail map[string]interface{} `position:"Query" name:"ActionDetail"`
	OrderId      requests.Integer       `position:"Query" name:"OrderId"`
	ActionName   string                 `position:"Query" name:"ActionName"`
	Tid          requests.Integer       `position:"Query" name:"Tid"`
}

// GetDataCorrectBackupFilesResponse is the response struct for api GetDataCorrectBackupFiles
type GetDataCorrectBackupFilesResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	ErrorMessage           string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode              string                 `json:"ErrorCode" xml:"ErrorCode"`
	DataCorrectBackupFiles DataCorrectBackupFiles `json:"DataCorrectBackupFiles" xml:"DataCorrectBackupFiles"`
}

// CreateGetDataCorrectBackupFilesRequest creates a request to invoke GetDataCorrectBackupFiles API
func CreateGetDataCorrectBackupFilesRequest() (request *GetDataCorrectBackupFilesRequest) {
	request = &GetDataCorrectBackupFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDataCorrectBackupFiles", "dmsenterprise", "openAPI")
	return
}

// CreateGetDataCorrectBackupFilesResponse creates a response to parse from GetDataCorrectBackupFiles response
func CreateGetDataCorrectBackupFilesResponse() (response *GetDataCorrectBackupFilesResponse) {
	response = &GetDataCorrectBackupFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
