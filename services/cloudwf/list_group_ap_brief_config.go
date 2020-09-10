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

// ListGroupApBriefConfig invokes the cloudwf.ListGroupApBriefConfig API synchronously
// api document: https://help.aliyun.com/api/cloudwf/listgroupapbriefconfig.html
func (client *Client) ListGroupApBriefConfig(request *ListGroupApBriefConfigRequest) (response *ListGroupApBriefConfigResponse, err error) {
	response = CreateListGroupApBriefConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListGroupApBriefConfigWithChan invokes the cloudwf.ListGroupApBriefConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listgroupapbriefconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupApBriefConfigWithChan(request *ListGroupApBriefConfigRequest) (<-chan *ListGroupApBriefConfigResponse, <-chan error) {
	responseChan := make(chan *ListGroupApBriefConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGroupApBriefConfig(request)
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

// ListGroupApBriefConfigWithCallback invokes the cloudwf.ListGroupApBriefConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listgroupapbriefconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupApBriefConfigWithCallback(request *ListGroupApBriefConfigRequest, callback func(response *ListGroupApBriefConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGroupApBriefConfigResponse
		var err error
		defer close(result)
		response, err = client.ListGroupApBriefConfig(request)
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

// ListGroupApBriefConfigRequest is the request struct for api ListGroupApBriefConfig
type ListGroupApBriefConfigRequest struct {
	*requests.RpcRequest
	OrderCol   string           `position:"Query" name:"OrderCol"`
	SearchName string           `position:"Query" name:"SearchName"`
	ApgroupId  requests.Integer `position:"Query" name:"ApgroupId"`
	ColCnt     requests.Integer `position:"Query" name:"ColCnt"`
	Length     requests.Integer `position:"Query" name:"Length"`
	PageIndex  requests.Integer `position:"Query" name:"PageIndex"`
	SearchMac  string           `position:"Query" name:"SearchMac"`
	OrderDir   string           `position:"Query" name:"OrderDir"`
}

// ListGroupApBriefConfigResponse is the response struct for api ListGroupApBriefConfig
type ListGroupApBriefConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateListGroupApBriefConfigRequest creates a request to invoke ListGroupApBriefConfig API
func CreateListGroupApBriefConfigRequest() (request *ListGroupApBriefConfigRequest) {
	request = &ListGroupApBriefConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListGroupApBriefConfig", "cloudwf", "openAPI")
	return
}

// CreateListGroupApBriefConfigResponse creates a response to parse from ListGroupApBriefConfig response
func CreateListGroupApBriefConfigResponse() (response *ListGroupApBriefConfigResponse) {
	response = &ListGroupApBriefConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
