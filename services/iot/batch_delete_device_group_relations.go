package iot

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

// BatchDeleteDeviceGroupRelations invokes the iot.BatchDeleteDeviceGroupRelations API synchronously
// api document: https://help.aliyun.com/api/iot/batchdeletedevicegrouprelations.html
func (client *Client) BatchDeleteDeviceGroupRelations(request *BatchDeleteDeviceGroupRelationsRequest) (response *BatchDeleteDeviceGroupRelationsResponse, err error) {
	response = CreateBatchDeleteDeviceGroupRelationsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteDeviceGroupRelationsWithChan invokes the iot.BatchDeleteDeviceGroupRelations API asynchronously
// api document: https://help.aliyun.com/api/iot/batchdeletedevicegrouprelations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteDeviceGroupRelationsWithChan(request *BatchDeleteDeviceGroupRelationsRequest) (<-chan *BatchDeleteDeviceGroupRelationsResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteDeviceGroupRelationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteDeviceGroupRelations(request)
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

// BatchDeleteDeviceGroupRelationsWithCallback invokes the iot.BatchDeleteDeviceGroupRelations API asynchronously
// api document: https://help.aliyun.com/api/iot/batchdeletedevicegrouprelations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteDeviceGroupRelationsWithCallback(request *BatchDeleteDeviceGroupRelationsRequest, callback func(response *BatchDeleteDeviceGroupRelationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteDeviceGroupRelationsResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteDeviceGroupRelations(request)
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

// BatchDeleteDeviceGroupRelationsRequest is the request struct for api BatchDeleteDeviceGroupRelations
type BatchDeleteDeviceGroupRelationsRequest struct {
	*requests.RpcRequest
	IotInstanceId string                                   `position:"Query" name:"IotInstanceId"`
	GroupId       string                                   `position:"Query" name:"GroupId"`
	ApiProduct    string                                   `position:"Body" name:"ApiProduct"`
	ApiRevision   string                                   `position:"Body" name:"ApiRevision"`
	Device        *[]BatchDeleteDeviceGroupRelationsDevice `position:"Query" name:"Device"  type:"Repeated"`
}

// BatchDeleteDeviceGroupRelationsDevice is a repeated param struct in BatchDeleteDeviceGroupRelationsRequest
type BatchDeleteDeviceGroupRelationsDevice struct {
	DeviceName string `name:"DeviceName"`
	ProductKey string `name:"ProductKey"`
}

// BatchDeleteDeviceGroupRelationsResponse is the response struct for api BatchDeleteDeviceGroupRelations
type BatchDeleteDeviceGroupRelationsResponse struct {
	*responses.BaseResponse
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	Success                        bool   `json:"Success" xml:"Success"`
	Code                           string `json:"Code" xml:"Code"`
	ErrorMessage                   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ValidDeviceCount               int    `json:"ValidDeviceCount" xml:"ValidDeviceCount"`
	AlreadyRelatedGroupDeviceCount int    `json:"AlreadyRelatedGroupDeviceCount" xml:"AlreadyRelatedGroupDeviceCount"`
	SuccessDeviceCount             int    `json:"SuccessDeviceCount" xml:"SuccessDeviceCount"`
}

// CreateBatchDeleteDeviceGroupRelationsRequest creates a request to invoke BatchDeleteDeviceGroupRelations API
func CreateBatchDeleteDeviceGroupRelationsRequest() (request *BatchDeleteDeviceGroupRelationsRequest) {
	request = &BatchDeleteDeviceGroupRelationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchDeleteDeviceGroupRelations", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchDeleteDeviceGroupRelationsResponse creates a response to parse from BatchDeleteDeviceGroupRelations response
func CreateBatchDeleteDeviceGroupRelationsResponse() (response *BatchDeleteDeviceGroupRelationsResponse) {
	response = &BatchDeleteDeviceGroupRelationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
