package ccc

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

// GetAgentData invokes the ccc.GetAgentData API synchronously
// api document: https://help.aliyun.com/api/ccc/getagentdata.html
func (client *Client) GetAgentData(request *GetAgentDataRequest) (response *GetAgentDataResponse, err error) {
	response = CreateGetAgentDataResponse()
	err = client.DoAction(request, response)
	return
}

// GetAgentDataWithChan invokes the ccc.GetAgentData API asynchronously
// api document: https://help.aliyun.com/api/ccc/getagentdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAgentDataWithChan(request *GetAgentDataRequest) (<-chan *GetAgentDataResponse, <-chan error) {
	responseChan := make(chan *GetAgentDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAgentData(request)
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

// GetAgentDataWithCallback invokes the ccc.GetAgentData API asynchronously
// api document: https://help.aliyun.com/api/ccc/getagentdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAgentDataWithCallback(request *GetAgentDataRequest, callback func(response *GetAgentDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAgentDataResponse
		var err error
		defer close(result)
		response, err = client.GetAgentData(request)
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

// GetAgentDataRequest is the request struct for api GetAgentData
type GetAgentDataRequest struct {
	*requests.RpcRequest
	StartDay   string           `position:"Query" name:"StartDay"`
	UserId     string           `position:"Query" name:"UserId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	EndDay     string           `position:"Query" name:"EndDay"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// GetAgentDataResponse is the response struct for api GetAgentData
type GetAgentDataResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DataList       DataList `json:"DataList" xml:"DataList"`
}

// CreateGetAgentDataRequest creates a request to invoke GetAgentData API
func CreateGetAgentDataRequest() (request *GetAgentDataRequest) {
	request = &GetAgentDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetAgentData", "", "")
	return
}

// CreateGetAgentDataResponse creates a response to parse from GetAgentData response
func CreateGetAgentDataResponse() (response *GetAgentDataResponse) {
	response = &GetAgentDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
