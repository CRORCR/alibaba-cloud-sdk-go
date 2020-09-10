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

// GetContactIdentifyByOutBoundTaskId invokes the cloudcallcenter.GetContactIdentifyByOutBoundTaskId API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactidentifybyoutboundtaskid.html
func (client *Client) GetContactIdentifyByOutBoundTaskId(request *GetContactIdentifyByOutBoundTaskIdRequest) (response *GetContactIdentifyByOutBoundTaskIdResponse, err error) {
	response = CreateGetContactIdentifyByOutBoundTaskIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetContactIdentifyByOutBoundTaskIdWithChan invokes the cloudcallcenter.GetContactIdentifyByOutBoundTaskId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactidentifybyoutboundtaskid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContactIdentifyByOutBoundTaskIdWithChan(request *GetContactIdentifyByOutBoundTaskIdRequest) (<-chan *GetContactIdentifyByOutBoundTaskIdResponse, <-chan error) {
	responseChan := make(chan *GetContactIdentifyByOutBoundTaskIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetContactIdentifyByOutBoundTaskId(request)
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

// GetContactIdentifyByOutBoundTaskIdWithCallback invokes the cloudcallcenter.GetContactIdentifyByOutBoundTaskId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactidentifybyoutboundtaskid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContactIdentifyByOutBoundTaskIdWithCallback(request *GetContactIdentifyByOutBoundTaskIdRequest, callback func(response *GetContactIdentifyByOutBoundTaskIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetContactIdentifyByOutBoundTaskIdResponse
		var err error
		defer close(result)
		response, err = client.GetContactIdentifyByOutBoundTaskId(request)
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

// GetContactIdentifyByOutBoundTaskIdRequest is the request struct for api GetContactIdentifyByOutBoundTaskId
type GetContactIdentifyByOutBoundTaskIdRequest struct {
	*requests.RpcRequest
	InstanceId     string `position:"Query" name:"InstanceId"`
	OutboundTaskId string `position:"Query" name:"OutboundTaskId"`
}

// GetContactIdentifyByOutBoundTaskIdResponse is the response struct for api GetContactIdentifyByOutBoundTaskId
type GetContactIdentifyByOutBoundTaskIdResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         bool            `json:"Success" xml:"Success"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	HttpStatusCode  int             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ContactIdentity ContactIdentity `json:"ContactIdentity" xml:"ContactIdentity"`
}

// CreateGetContactIdentifyByOutBoundTaskIdRequest creates a request to invoke GetContactIdentifyByOutBoundTaskId API
func CreateGetContactIdentifyByOutBoundTaskIdRequest() (request *GetContactIdentifyByOutBoundTaskIdRequest) {
	request = &GetContactIdentifyByOutBoundTaskIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetContactIdentifyByOutBoundTaskId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetContactIdentifyByOutBoundTaskIdResponse creates a response to parse from GetContactIdentifyByOutBoundTaskId response
func CreateGetContactIdentifyByOutBoundTaskIdResponse() (response *GetContactIdentifyByOutBoundTaskIdResponse) {
	response = &GetContactIdentifyByOutBoundTaskIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
