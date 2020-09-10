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
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// SaveApgroupScanConfig invokes the cloudwf.SaveApgroupScanConfig API synchronously
// api document: https://help.aliyun.com/api/cloudwf/saveapgroupscanconfig.html
func (client *Client) SaveApgroupScanConfig(request *SaveApgroupScanConfigRequest) (response *SaveApgroupScanConfigResponse, err error) {
	response = CreateSaveApgroupScanConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SaveApgroupScanConfigWithChan invokes the cloudwf.SaveApgroupScanConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/saveapgroupscanconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveApgroupScanConfigWithChan(request *SaveApgroupScanConfigRequest) (<-chan *SaveApgroupScanConfigResponse, <-chan error) {
	responseChan := make(chan *SaveApgroupScanConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveApgroupScanConfig(request)
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

// SaveApgroupScanConfigWithCallback invokes the cloudwf.SaveApgroupScanConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/saveapgroupscanconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveApgroupScanConfigWithCallback(request *SaveApgroupScanConfigRequest, callback func(response *SaveApgroupScanConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveApgroupScanConfigResponse
		var err error
		defer close(result)
		response, err = client.SaveApgroupScanConfig(request)
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

// SaveApgroupScanConfigRequest is the request struct for api SaveApgroupScanConfig
type SaveApgroupScanConfigRequest struct {
	*requests.RpcRequest
	JsonData  string           `position:"Query" name:"JsonData"`
	ApgroupId requests.Integer `position:"Query" name:"ApgroupId"`
}

// SaveApgroupScanConfigResponse is the response struct for api SaveApgroupScanConfig
type SaveApgroupScanConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateSaveApgroupScanConfigRequest creates a request to invoke SaveApgroupScanConfig API
func CreateSaveApgroupScanConfigRequest() (request *SaveApgroupScanConfigRequest) {
	request = &SaveApgroupScanConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApgroupScanConfig", "cloudwf", "openAPI")
	return
}

// CreateSaveApgroupScanConfigResponse creates a response to parse from SaveApgroupScanConfig response
func CreateSaveApgroupScanConfigResponse() (response *SaveApgroupScanConfigResponse) {
	response = &SaveApgroupScanConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
