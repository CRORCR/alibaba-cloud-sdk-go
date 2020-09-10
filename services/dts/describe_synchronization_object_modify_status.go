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

package dts

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSynchronizationObjectModifyStatus invokes the dts.DescribeSynchronizationObjectModifyStatus API synchronously
// api document: https://help.aliyun.com/api/dts/describesynchronizationobjectmodifystatus.html
func (client *Client) DescribeSynchronizationObjectModifyStatus(request *DescribeSynchronizationObjectModifyStatusRequest) (response *DescribeSynchronizationObjectModifyStatusResponse, err error) {
	response = CreateDescribeSynchronizationObjectModifyStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSynchronizationObjectModifyStatusWithChan invokes the dts.DescribeSynchronizationObjectModifyStatus API asynchronously
// api document: https://help.aliyun.com/api/dts/describesynchronizationobjectmodifystatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSynchronizationObjectModifyStatusWithChan(request *DescribeSynchronizationObjectModifyStatusRequest) (<-chan *DescribeSynchronizationObjectModifyStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSynchronizationObjectModifyStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSynchronizationObjectModifyStatus(request)
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

// DescribeSynchronizationObjectModifyStatusWithCallback invokes the dts.DescribeSynchronizationObjectModifyStatus API asynchronously
// api document: https://help.aliyun.com/api/dts/describesynchronizationobjectmodifystatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSynchronizationObjectModifyStatusWithCallback(request *DescribeSynchronizationObjectModifyStatusRequest, callback func(response *DescribeSynchronizationObjectModifyStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSynchronizationObjectModifyStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSynchronizationObjectModifyStatus(request)
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

// DescribeSynchronizationObjectModifyStatusRequest is the request struct for api DescribeSynchronizationObjectModifyStatus
type DescribeSynchronizationObjectModifyStatusRequest struct {
	*requests.RpcRequest
	TaskId      string `position:"Query" name:"TaskId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     string `position:"Query" name:"OwnerId"`
}

// DescribeSynchronizationObjectModifyStatusResponse is the response struct for api DescribeSynchronizationObjectModifyStatus
type DescribeSynchronizationObjectModifyStatusResponse struct {
	*responses.BaseResponse
	RequestId                     string                                                                  `json:"RequestId" xml:"RequestId"`
	Status                        string                                                                  `json:"Status" xml:"Status"`
	ErrorMessage                  string                                                                  `json:"ErrorMessage" xml:"ErrorMessage"`
	PrecheckStatus                DescribeSynchronizationObjectModifyStatusPrecheckStatus0                `json:"PrecheckStatus" xml:"PrecheckStatus"`
	StructureInitializationStatus DescribeSynchronizationObjectModifyStatusStructureInitializationStatus0 `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
	DataInitializationStatus      DescribeSynchronizationObjectModifyStatusDataInitializationStatus0      `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	DataSynchronizationStatus     DescribeSynchronizationObjectModifyStatusDataSynchronizationStatus0     `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
}

type DescribeSynchronizationObjectModifyStatusPrecheckStatus0 struct {
	Status  string                                                `json:"Status" xml:"Status"`
	Percent string                                                `json:"Percent" xml:"Percent"`
	Detail  []DescribeSynchronizationObjectModifyStatusCheckItem1 `json:"Detail" xml:"Detail"`
}

type DescribeSynchronizationObjectModifyStatusStructureInitializationStatus0 struct {
	Status       string `json:"Status" xml:"Status"`
	Percent      string `json:"Percent" xml:"Percent"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Progress     string `json:"Progress" xml:"Progress"`
}

type DescribeSynchronizationObjectModifyStatusDataInitializationStatus0 struct {
	Status       string `json:"Status" xml:"Status"`
	Percent      string `json:"Percent" xml:"Percent"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Progress     string `json:"Progress" xml:"Progress"`
}

type DescribeSynchronizationObjectModifyStatusDataSynchronizationStatus0 struct {
	Status       string `json:"Status" xml:"Status"`
	Percent      string `json:"Percent" xml:"Percent"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Delay        string `json:"Delay" xml:"Delay"`
}

type DescribeSynchronizationObjectModifyStatusCheckItem1 struct {
	ItemName     string `json:"ItemName" xml:"ItemName"`
	CheckStatus  string `json:"CheckStatus" xml:"CheckStatus"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RepairMethod string `json:"RepairMethod" xml:"RepairMethod"`
}

// CreateDescribeSynchronizationObjectModifyStatusRequest creates a request to invoke DescribeSynchronizationObjectModifyStatus API
func CreateDescribeSynchronizationObjectModifyStatusRequest() (request *DescribeSynchronizationObjectModifyStatusRequest) {
	request = &DescribeSynchronizationObjectModifyStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "DescribeSynchronizationObjectModifyStatus", "dts", "openAPI")
	return
}

// CreateDescribeSynchronizationObjectModifyStatusResponse creates a response to parse from DescribeSynchronizationObjectModifyStatus response
func CreateDescribeSynchronizationObjectModifyStatusResponse() (response *DescribeSynchronizationObjectModifyStatusResponse) {
	response = &DescribeSynchronizationObjectModifyStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
