package ecs

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

// CreateDemand invokes the ecs.CreateDemand API synchronously
// api document: https://help.aliyun.com/api/ecs/createdemand.html
func (client *Client) CreateDemand(request *CreateDemandRequest) (response *CreateDemandResponse, err error) {
	response = CreateCreateDemandResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDemandWithChan invokes the ecs.CreateDemand API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdemand.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDemandWithChan(request *CreateDemandRequest) (<-chan *CreateDemandResponse, <-chan error) {
	responseChan := make(chan *CreateDemandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDemand(request)
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

// CreateDemandWithCallback invokes the ecs.CreateDemand API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdemand.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDemandWithCallback(request *CreateDemandRequest, callback func(response *CreateDemandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDemandResponse
		var err error
		defer close(result)
		response, err = client.CreateDemand(request)
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

// CreateDemandRequest is the request struct for api CreateDemand
type CreateDemandRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	StartTime            string           `position:"Query" name:"StartTime"`
	DemandDescription    string           `position:"Query" name:"DemandDescription"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	DemandName           string           `position:"Query" name:"DemandName"`
	Amount               requests.Integer `position:"Query" name:"Amount"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// CreateDemandResponse is the response struct for api CreateDemand
type CreateDemandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	DemandId  string `json:"DemandId" xml:"DemandId"`
}

// CreateCreateDemandRequest creates a request to invoke CreateDemand API
func CreateCreateDemandRequest() (request *CreateDemandRequest) {
	request = &CreateDemandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateDemand", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDemandResponse creates a response to parse from CreateDemand response
func CreateCreateDemandResponse() (response *CreateDemandResponse) {
	response = &CreateDemandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
