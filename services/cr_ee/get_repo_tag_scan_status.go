package cr_ee

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

// GetRepoTagScanStatus invokes the cr.GetRepoTagScanStatus API synchronously
// api document: https://help.aliyun.com/api/cr/getrepotagscanstatus.html
func (client *Client) GetRepoTagScanStatus(request *GetRepoTagScanStatusRequest) (response *GetRepoTagScanStatusResponse, err error) {
	response = CreateGetRepoTagScanStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoTagScanStatusWithChan invokes the cr.GetRepoTagScanStatus API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotagscanstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagScanStatusWithChan(request *GetRepoTagScanStatusRequest) (<-chan *GetRepoTagScanStatusResponse, <-chan error) {
	responseChan := make(chan *GetRepoTagScanStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoTagScanStatus(request)
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

// GetRepoTagScanStatusWithCallback invokes the cr.GetRepoTagScanStatus API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotagscanstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagScanStatusWithCallback(request *GetRepoTagScanStatusRequest, callback func(response *GetRepoTagScanStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoTagScanStatusResponse
		var err error
		defer close(result)
		response, err = client.GetRepoTagScanStatus(request)
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

// GetRepoTagScanStatusRequest is the request struct for api GetRepoTagScanStatus
type GetRepoTagScanStatusRequest struct {
	*requests.RpcRequest
	RepoId     string `position:"Query" name:"RepoId"`
	ScanTaskId string `position:"Query" name:"ScanTaskId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Tag        string `position:"Query" name:"Tag"`
}

// GetRepoTagScanStatusResponse is the response struct for api GetRepoTagScanStatus
type GetRepoTagScanStatusResponse struct {
	*responses.BaseResponse
	GetRepoTagScanStatusIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                          string `json:"Code" xml:"Code"`
	RequestId                     string `json:"RequestId" xml:"RequestId"`
	Status                        string `json:"Status" xml:"Status"`
}

// CreateGetRepoTagScanStatusRequest creates a request to invoke GetRepoTagScanStatus API
func CreateGetRepoTagScanStatusRequest() (request *GetRepoTagScanStatusRequest) {
	request = &GetRepoTagScanStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "GetRepoTagScanStatus", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRepoTagScanStatusResponse creates a response to parse from GetRepoTagScanStatus response
func CreateGetRepoTagScanStatusResponse() (response *GetRepoTagScanStatusResponse) {
	response = &GetRepoTagScanStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
