package dataworks_public

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

// GetInstanceLog invokes the dataworks_public.GetInstanceLog API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/getinstancelog.html
func (client *Client) GetInstanceLog(request *GetInstanceLogRequest) (response *GetInstanceLogResponse, err error) {
	response = CreateGetInstanceLogResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceLogWithChan invokes the dataworks_public.GetInstanceLog API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getinstancelog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceLogWithChan(request *GetInstanceLogRequest) (<-chan *GetInstanceLogResponse, <-chan error) {
	responseChan := make(chan *GetInstanceLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceLog(request)
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

// GetInstanceLogWithCallback invokes the dataworks_public.GetInstanceLog API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getinstancelog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceLogWithCallback(request *GetInstanceLogRequest, callback func(response *GetInstanceLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceLogResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceLog(request)
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

// GetInstanceLogRequest is the request struct for api GetInstanceLog
type GetInstanceLogRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	InstanceId requests.Integer `position:"Body" name:"InstanceId"`
}

// GetInstanceLogResponse is the response struct for api GetInstanceLog
type GetInstanceLogResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateGetInstanceLogRequest creates a request to invoke GetInstanceLog API
func CreateGetInstanceLogRequest() (request *GetInstanceLogRequest) {
	request = &GetInstanceLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetInstanceLog", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetInstanceLogResponse creates a response to parse from GetInstanceLog response
func CreateGetInstanceLogResponse() (response *GetInstanceLogResponse) {
	response = &GetInstanceLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
