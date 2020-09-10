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

// ListConfig invokes the ccc.ListConfig API synchronously
// api document: https://help.aliyun.com/api/ccc/listconfig.html
func (client *Client) ListConfig(request *ListConfigRequest) (response *ListConfigResponse, err error) {
	response = CreateListConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListConfigWithChan invokes the ccc.ListConfig API asynchronously
// api document: https://help.aliyun.com/api/ccc/listconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConfigWithChan(request *ListConfigRequest) (<-chan *ListConfigResponse, <-chan error) {
	responseChan := make(chan *ListConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConfig(request)
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

// ListConfigWithCallback invokes the ccc.ListConfig API asynchronously
// api document: https://help.aliyun.com/api/ccc/listconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConfigWithCallback(request *ListConfigRequest, callback func(response *ListConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConfigResponse
		var err error
		defer close(result)
		response, err = client.ListConfig(request)
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

// ListConfigRequest is the request struct for api ListConfig
type ListConfigRequest struct {
	*requests.RpcRequest
	InstanceId string    `position:"Query" name:"InstanceId"`
	ConfigItem *[]string `position:"Query" name:"ConfigItem"  type:"Repeated"`
}

// ListConfigResponse is the response struct for api ListConfig
type ListConfigResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ConfigItems    ConfigItems `json:"ConfigItems" xml:"ConfigItems"`
}

// CreateListConfigRequest creates a request to invoke ListConfig API
func CreateListConfigRequest() (request *ListConfigRequest) {
	request = &ListConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListConfig", "", "")
	return
}

// CreateListConfigResponse creates a response to parse from ListConfig response
func CreateListConfigResponse() (response *ListConfigResponse) {
	response = &ListConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
