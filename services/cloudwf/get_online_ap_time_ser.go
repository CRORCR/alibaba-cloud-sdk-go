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

// GetOnlineApTimeSer invokes the cloudwf.GetOnlineApTimeSer API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getonlineaptimeser.html
func (client *Client) GetOnlineApTimeSer(request *GetOnlineApTimeSerRequest) (response *GetOnlineApTimeSerResponse, err error) {
	response = CreateGetOnlineApTimeSerResponse()
	err = client.DoAction(request, response)
	return
}

// GetOnlineApTimeSerWithChan invokes the cloudwf.GetOnlineApTimeSer API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getonlineaptimeser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOnlineApTimeSerWithChan(request *GetOnlineApTimeSerRequest) (<-chan *GetOnlineApTimeSerResponse, <-chan error) {
	responseChan := make(chan *GetOnlineApTimeSerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOnlineApTimeSer(request)
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

// GetOnlineApTimeSerWithCallback invokes the cloudwf.GetOnlineApTimeSer API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getonlineaptimeser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOnlineApTimeSerWithCallback(request *GetOnlineApTimeSerRequest, callback func(response *GetOnlineApTimeSerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOnlineApTimeSerResponse
		var err error
		defer close(result)
		response, err = client.GetOnlineApTimeSer(request)
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

// GetOnlineApTimeSerRequest is the request struct for api GetOnlineApTimeSer
type GetOnlineApTimeSerRequest struct {
	*requests.RpcRequest
	ZoomStart requests.Integer `position:"Query" name:"ZoomStart"`
	CompanyId requests.Integer `position:"Query" name:"CompanyId"`
	ApgroupId requests.Integer `position:"Query" name:"ApgroupId"`
	Start     requests.Integer `position:"Query" name:"Start"`
	ZoomEnd   requests.Integer `position:"Query" name:"ZoomEnd"`
	End       requests.Integer `position:"Query" name:"End"`
}

// GetOnlineApTimeSerResponse is the response struct for api GetOnlineApTimeSer
type GetOnlineApTimeSerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetOnlineApTimeSerRequest creates a request to invoke GetOnlineApTimeSer API
func CreateGetOnlineApTimeSerRequest() (request *GetOnlineApTimeSerRequest) {
	request = &GetOnlineApTimeSerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetOnlineApTimeSer", "cloudwf", "openAPI")
	return
}

// CreateGetOnlineApTimeSerResponse creates a response to parse from GetOnlineApTimeSer response
func CreateGetOnlineApTimeSerResponse() (response *GetOnlineApTimeSerResponse) {
	response = &GetOnlineApTimeSerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
