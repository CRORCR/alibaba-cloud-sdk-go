package cloudcallcenter

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

// GetAcquaintance invokes the cloudcallcenter.GetAcquaintance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getacquaintance.html
func (client *Client) GetAcquaintance(request *GetAcquaintanceRequest) (response *GetAcquaintanceResponse, err error) {
	response = CreateGetAcquaintanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetAcquaintanceWithChan invokes the cloudcallcenter.GetAcquaintance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getacquaintance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAcquaintanceWithChan(request *GetAcquaintanceRequest) (<-chan *GetAcquaintanceResponse, <-chan error) {
	responseChan := make(chan *GetAcquaintanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAcquaintance(request)
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

// GetAcquaintanceWithCallback invokes the cloudcallcenter.GetAcquaintance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getacquaintance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAcquaintanceWithCallback(request *GetAcquaintanceRequest, callback func(response *GetAcquaintanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAcquaintanceResponse
		var err error
		defer close(result)
		response, err = client.GetAcquaintance(request)
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

// GetAcquaintanceRequest is the request struct for api GetAcquaintance
type GetAcquaintanceRequest struct {
	*requests.RpcRequest
	Number             string           `position:"Query" name:"Number"`
	LimitCount         requests.Integer `position:"Query" name:"LimitCount"`
	ContactDisposition requests.Integer `position:"Query" name:"ContactDisposition"`
	ContactType        requests.Integer `position:"Query" name:"ContactType"`
	Days               requests.Integer `position:"Query" name:"Days"`
	Tenant             string           `position:"Query" name:"Tenant"`
}

// GetAcquaintanceResponse is the response struct for api GetAcquaintance
type GetAcquaintanceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	AgentNo        string `json:"AgentNo" xml:"AgentNo"`
	AgentStatus    int    `json:"AgentStatus" xml:"AgentStatus"`
}

// CreateGetAcquaintanceRequest creates a request to invoke GetAcquaintance API
func CreateGetAcquaintanceRequest() (request *GetAcquaintanceRequest) {
	request = &GetAcquaintanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetAcquaintance", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAcquaintanceResponse creates a response to parse from GetAcquaintance response
func CreateGetAcquaintanceResponse() (response *GetAcquaintanceResponse) {
	response = &GetAcquaintanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
