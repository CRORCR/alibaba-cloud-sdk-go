package green

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

// RenewWebSiteInstance invokes the green.RenewWebSiteInstance API synchronously
// api document: https://help.aliyun.com/api/green/renewwebsiteinstance.html
func (client *Client) RenewWebSiteInstance(request *RenewWebSiteInstanceRequest) (response *RenewWebSiteInstanceResponse, err error) {
	response = CreateRenewWebSiteInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RenewWebSiteInstanceWithChan invokes the green.RenewWebSiteInstance API asynchronously
// api document: https://help.aliyun.com/api/green/renewwebsiteinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewWebSiteInstanceWithChan(request *RenewWebSiteInstanceRequest) (<-chan *RenewWebSiteInstanceResponse, <-chan error) {
	responseChan := make(chan *RenewWebSiteInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewWebSiteInstance(request)
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

// RenewWebSiteInstanceWithCallback invokes the green.RenewWebSiteInstance API asynchronously
// api document: https://help.aliyun.com/api/green/renewwebsiteinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewWebSiteInstanceWithCallback(request *RenewWebSiteInstanceRequest, callback func(response *RenewWebSiteInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewWebSiteInstanceResponse
		var err error
		defer close(result)
		response, err = client.RenewWebSiteInstance(request)
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

// RenewWebSiteInstanceRequest is the request struct for api RenewWebSiteInstance
type RenewWebSiteInstanceRequest struct {
	*requests.RpcRequest
	ClientToken   string           `position:"Query" name:"ClientToken"`
	OrderNum      requests.Integer `position:"Query" name:"OrderNum"`
	CommodityCode string           `position:"Query" name:"CommodityCode"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Duration      requests.Integer `position:"Query" name:"Duration"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	PricingCycle  string           `position:"Query" name:"PricingCycle"`
	OrderType     string           `position:"Query" name:"OrderType"`
}

// RenewWebSiteInstanceResponse is the response struct for api RenewWebSiteInstance
type RenewWebSiteInstanceResponse struct {
	*responses.BaseResponse
	Code        string                            `json:"Code" xml:"Code"`
	Message     string                            `json:"Message" xml:"Message"`
	OrderId     string                            `json:"OrderId" xml:"OrderId"`
	InstanceId  string                            `json:"InstanceId" xml:"InstanceId"`
	RequestId   string                            `json:"RequestId" xml:"RequestId"`
	InstanceIds InstanceIdsInRenewWebSiteInstance `json:"InstanceIds" xml:"InstanceIds"`
}

// CreateRenewWebSiteInstanceRequest creates a request to invoke RenewWebSiteInstance API
func CreateRenewWebSiteInstanceRequest() (request *RenewWebSiteInstanceRequest) {
	request = &RenewWebSiteInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "RenewWebSiteInstance", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewWebSiteInstanceResponse creates a response to parse from RenewWebSiteInstance response
func CreateRenewWebSiteInstanceResponse() (response *RenewWebSiteInstanceResponse) {
	response = &RenewWebSiteInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
