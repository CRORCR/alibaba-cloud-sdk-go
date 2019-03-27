package cloudwf

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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CheckRootPermission invokes the cloudwf.CheckRootPermission API synchronously
// api document: https://help.aliyun.com/api/cloudwf/checkrootpermission.html
func (client *Client) CheckRootPermission(request *CheckRootPermissionRequest) (response *CheckRootPermissionResponse, err error) {
	response = CreateCheckRootPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// CheckRootPermissionWithChan invokes the cloudwf.CheckRootPermission API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/checkrootpermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRootPermissionWithChan(request *CheckRootPermissionRequest) (<-chan *CheckRootPermissionResponse, <-chan error) {
	responseChan := make(chan *CheckRootPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckRootPermission(request)
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

// CheckRootPermissionWithCallback invokes the cloudwf.CheckRootPermission API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/checkrootpermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRootPermissionWithCallback(request *CheckRootPermissionRequest, callback func(response *CheckRootPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckRootPermissionResponse
		var err error
		defer close(result)
		response, err = client.CheckRootPermission(request)
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

// CheckRootPermissionRequest is the request struct for api CheckRootPermission
type CheckRootPermissionRequest struct {
	*requests.RpcRequest
}

// CheckRootPermissionResponse is the response struct for api CheckRootPermission
type CheckRootPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateCheckRootPermissionRequest creates a request to invoke CheckRootPermission API
func CreateCheckRootPermissionRequest() (request *CheckRootPermissionRequest) {
	request = &CheckRootPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "CheckRootPermission", "cloudwf", "openAPI")
	return
}

// CreateCheckRootPermissionResponse creates a response to parse from CheckRootPermission response
func CreateCheckRootPermissionResponse() (response *CheckRootPermissionResponse) {
	response = &CheckRootPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}