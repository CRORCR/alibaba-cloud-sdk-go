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

// RepairGroupAp invokes the cloudwf.RepairGroupAp API synchronously
// api document: https://help.aliyun.com/api/cloudwf/repairgroupap.html
func (client *Client) RepairGroupAp(request *RepairGroupApRequest) (response *RepairGroupApResponse, err error) {
	response = CreateRepairGroupApResponse()
	err = client.DoAction(request, response)
	return
}

// RepairGroupApWithChan invokes the cloudwf.RepairGroupAp API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/repairgroupap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RepairGroupApWithChan(request *RepairGroupApRequest) (<-chan *RepairGroupApResponse, <-chan error) {
	responseChan := make(chan *RepairGroupApResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RepairGroupAp(request)
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

// RepairGroupApWithCallback invokes the cloudwf.RepairGroupAp API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/repairgroupap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RepairGroupApWithCallback(request *RepairGroupApRequest, callback func(response *RepairGroupApResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RepairGroupApResponse
		var err error
		defer close(result)
		response, err = client.RepairGroupAp(request)
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

// RepairGroupApRequest is the request struct for api RepairGroupAp
type RepairGroupApRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// RepairGroupApResponse is the response struct for api RepairGroupAp
type RepairGroupApResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateRepairGroupApRequest creates a request to invoke RepairGroupAp API
func CreateRepairGroupApRequest() (request *RepairGroupApRequest) {
	request = &RepairGroupApRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "RepairGroupAp", "cloudwf", "openAPI")
	return
}

// CreateRepairGroupApResponse creates a response to parse from RepairGroupAp response
func CreateRepairGroupApResponse() (response *RepairGroupApResponse) {
	response = &RepairGroupApResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
