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

// GetPortalTempDetail invokes the cloudwf.GetPortalTempDetail API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getportaltempdetail.html
func (client *Client) GetPortalTempDetail(request *GetPortalTempDetailRequest) (response *GetPortalTempDetailResponse, err error) {
	response = CreateGetPortalTempDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetPortalTempDetailWithChan invokes the cloudwf.GetPortalTempDetail API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getportaltempdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPortalTempDetailWithChan(request *GetPortalTempDetailRequest) (<-chan *GetPortalTempDetailResponse, <-chan error) {
	responseChan := make(chan *GetPortalTempDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPortalTempDetail(request)
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

// GetPortalTempDetailWithCallback invokes the cloudwf.GetPortalTempDetail API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getportaltempdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPortalTempDetailWithCallback(request *GetPortalTempDetailRequest, callback func(response *GetPortalTempDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPortalTempDetailResponse
		var err error
		defer close(result)
		response, err = client.GetPortalTempDetail(request)
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

// GetPortalTempDetailRequest is the request struct for api GetPortalTempDetail
type GetPortalTempDetailRequest struct {
	*requests.RpcRequest
	Id       requests.Integer `position:"Query" name:"Id"`
	UniqueId string           `position:"Query" name:"UniqueId"`
}

// GetPortalTempDetailResponse is the response struct for api GetPortalTempDetail
type GetPortalTempDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetPortalTempDetailRequest creates a request to invoke GetPortalTempDetail API
func CreateGetPortalTempDetailRequest() (request *GetPortalTempDetailRequest) {
	request = &GetPortalTempDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetPortalTempDetail", "cloudwf", "openAPI")
	return
}

// CreateGetPortalTempDetailResponse creates a response to parse from GetPortalTempDetail response
func CreateGetPortalTempDetailResponse() (response *GetPortalTempDetailResponse) {
	response = &GetPortalTempDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
