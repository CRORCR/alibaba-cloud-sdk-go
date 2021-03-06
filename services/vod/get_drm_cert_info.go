package vod

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

// GetDRMCertInfo invokes the vod.GetDRMCertInfo API synchronously
func (client *Client) GetDRMCertInfo(request *GetDRMCertInfoRequest) (response *GetDRMCertInfoResponse, err error) {
	response = CreateGetDRMCertInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetDRMCertInfoWithChan invokes the vod.GetDRMCertInfo API asynchronously
func (client *Client) GetDRMCertInfoWithChan(request *GetDRMCertInfoRequest) (<-chan *GetDRMCertInfoResponse, <-chan error) {
	responseChan := make(chan *GetDRMCertInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDRMCertInfo(request)
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

// GetDRMCertInfoWithCallback invokes the vod.GetDRMCertInfo API asynchronously
func (client *Client) GetDRMCertInfoWithCallback(request *GetDRMCertInfoRequest, callback func(response *GetDRMCertInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDRMCertInfoResponse
		var err error
		defer close(result)
		response, err = client.GetDRMCertInfo(request)
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

// GetDRMCertInfoRequest is the request struct for api GetDRMCertInfo
type GetDRMCertInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string           `position:"Query" name:"VideoId"`
	CertId               string           `position:"Query" name:"CertId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GetDRMCertInfoResponse is the response struct for api GetDRMCertInfo
type GetDRMCertInfoResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DRMCertInfo string `json:"DRMCertInfo" xml:"DRMCertInfo"`
}

// CreateGetDRMCertInfoRequest creates a request to invoke GetDRMCertInfo API
func CreateGetDRMCertInfoRequest() (request *GetDRMCertInfoRequest) {
	request = &GetDRMCertInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetDRMCertInfo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDRMCertInfoResponse creates a response to parse from GetDRMCertInfo response
func CreateGetDRMCertInfoResponse() (response *GetDRMCertInfoResponse) {
	response = &GetDRMCertInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
