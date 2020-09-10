package alikafka

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

// UpgradePrePayOrder invokes the alikafka.UpgradePrePayOrder API synchronously
// api document: https://help.aliyun.com/api/alikafka/upgradeprepayorder.html
func (client *Client) UpgradePrePayOrder(request *UpgradePrePayOrderRequest) (response *UpgradePrePayOrderResponse, err error) {
	response = CreateUpgradePrePayOrderResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradePrePayOrderWithChan invokes the alikafka.UpgradePrePayOrder API asynchronously
// api document: https://help.aliyun.com/api/alikafka/upgradeprepayorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradePrePayOrderWithChan(request *UpgradePrePayOrderRequest) (<-chan *UpgradePrePayOrderResponse, <-chan error) {
	responseChan := make(chan *UpgradePrePayOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradePrePayOrder(request)
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

// UpgradePrePayOrderWithCallback invokes the alikafka.UpgradePrePayOrder API asynchronously
// api document: https://help.aliyun.com/api/alikafka/upgradeprepayorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradePrePayOrderWithCallback(request *UpgradePrePayOrderRequest, callback func(response *UpgradePrePayOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradePrePayOrderResponse
		var err error
		defer close(result)
		response, err = client.UpgradePrePayOrder(request)
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

// UpgradePrePayOrderRequest is the request struct for api UpgradePrePayOrder
type UpgradePrePayOrderRequest struct {
	*requests.RpcRequest
	DiskSize   requests.Integer `position:"Query" name:"DiskSize"`
	IoMax      requests.Integer `position:"Query" name:"IoMax"`
	TopicQuota requests.Integer `position:"Query" name:"TopicQuota"`
	EipMax     requests.Integer `position:"Query" name:"EipMax"`
	SpecType   string           `position:"Query" name:"SpecType"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

// UpgradePrePayOrderResponse is the response struct for api UpgradePrePayOrder
type UpgradePrePayOrderResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpgradePrePayOrderRequest creates a request to invoke UpgradePrePayOrder API
func CreateUpgradePrePayOrderRequest() (request *UpgradePrePayOrderRequest) {
	request = &UpgradePrePayOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "UpgradePrePayOrder", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradePrePayOrderResponse creates a response to parse from UpgradePrePayOrder response
func CreateUpgradePrePayOrderResponse() (response *UpgradePrePayOrderResponse) {
	response = &UpgradePrePayOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
