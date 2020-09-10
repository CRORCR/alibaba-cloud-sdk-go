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

// GetGroupApRadioOnoffProgress invokes the cloudwf.GetGroupApRadioOnoffProgress API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getgroupapradioonoffprogress.html
func (client *Client) GetGroupApRadioOnoffProgress(request *GetGroupApRadioOnoffProgressRequest) (response *GetGroupApRadioOnoffProgressResponse, err error) {
	response = CreateGetGroupApRadioOnoffProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetGroupApRadioOnoffProgressWithChan invokes the cloudwf.GetGroupApRadioOnoffProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getgroupapradioonoffprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetGroupApRadioOnoffProgressWithChan(request *GetGroupApRadioOnoffProgressRequest) (<-chan *GetGroupApRadioOnoffProgressResponse, <-chan error) {
	responseChan := make(chan *GetGroupApRadioOnoffProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGroupApRadioOnoffProgress(request)
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

// GetGroupApRadioOnoffProgressWithCallback invokes the cloudwf.GetGroupApRadioOnoffProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getgroupapradioonoffprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetGroupApRadioOnoffProgressWithCallback(request *GetGroupApRadioOnoffProgressRequest, callback func(response *GetGroupApRadioOnoffProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGroupApRadioOnoffProgressResponse
		var err error
		defer close(result)
		response, err = client.GetGroupApRadioOnoffProgress(request)
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

// GetGroupApRadioOnoffProgressRequest is the request struct for api GetGroupApRadioOnoffProgress
type GetGroupApRadioOnoffProgressRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// GetGroupApRadioOnoffProgressResponse is the response struct for api GetGroupApRadioOnoffProgress
type GetGroupApRadioOnoffProgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetGroupApRadioOnoffProgressRequest creates a request to invoke GetGroupApRadioOnoffProgress API
func CreateGetGroupApRadioOnoffProgressRequest() (request *GetGroupApRadioOnoffProgressRequest) {
	request = &GetGroupApRadioOnoffProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioOnoffProgress", "cloudwf", "openAPI")
	return
}

// CreateGetGroupApRadioOnoffProgressResponse creates a response to parse from GetGroupApRadioOnoffProgress response
func CreateGetGroupApRadioOnoffProgressResponse() (response *GetGroupApRadioOnoffProgressResponse) {
	response = &GetGroupApRadioOnoffProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
