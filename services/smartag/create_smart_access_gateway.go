package smartag

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

// CreateSmartAccessGateway invokes the smartag.CreateSmartAccessGateway API synchronously
// api document: https://help.aliyun.com/api/smartag/createsmartaccessgateway.html
func (client *Client) CreateSmartAccessGateway(request *CreateSmartAccessGatewayRequest) (response *CreateSmartAccessGatewayResponse, err error) {
	response = CreateCreateSmartAccessGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSmartAccessGatewayWithChan invokes the smartag.CreateSmartAccessGateway API asynchronously
// api document: https://help.aliyun.com/api/smartag/createsmartaccessgateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSmartAccessGatewayWithChan(request *CreateSmartAccessGatewayRequest) (<-chan *CreateSmartAccessGatewayResponse, <-chan error) {
	responseChan := make(chan *CreateSmartAccessGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSmartAccessGateway(request)
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

// CreateSmartAccessGatewayWithCallback invokes the smartag.CreateSmartAccessGateway API asynchronously
// api document: https://help.aliyun.com/api/smartag/createsmartaccessgateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSmartAccessGatewayWithCallback(request *CreateSmartAccessGatewayRequest, callback func(response *CreateSmartAccessGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSmartAccessGatewayResponse
		var err error
		defer close(result)
		response, err = client.CreateSmartAccessGateway(request)
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

// CreateSmartAccessGatewayRequest is the request struct for api CreateSmartAccessGateway
type CreateSmartAccessGatewayRequest struct {
	*requests.RpcRequest
	MaxBandWidth         requests.Integer `position:"Query" name:"MaxBandWidth"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	ReceiverTown         string           `position:"Query" name:"ReceiverTown"`
	ReceiverDistrict     string           `position:"Query" name:"ReceiverDistrict"`
	UserCount            requests.Integer `position:"Query" name:"UserCount"`
	ReceiverAddress      string           `position:"Query" name:"ReceiverAddress"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	BuyerMessage         string           `position:"Query" name:"BuyerMessage"`
	HardWareSpec         string           `position:"Query" name:"HardWareSpec"`
	ReceiverEmail        string           `position:"Query" name:"ReceiverEmail"`
	ReceiverState        string           `position:"Query" name:"ReceiverState"`
	ReceiverCity         string           `position:"Query" name:"ReceiverCity"`
	Period               requests.Integer `position:"Query" name:"Period"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ReceiverMobile       string           `position:"Query" name:"ReceiverMobile"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ReceiverPhone        string           `position:"Query" name:"ReceiverPhone"`
	ReceiverName         string           `position:"Query" name:"ReceiverName"`
	HaType               string           `position:"Query" name:"HaType"`
	Name                 string           `position:"Query" name:"Name"`
	AlreadyHaveSag       requests.Boolean `position:"Query" name:"AlreadyHaveSag"`
	ReceiverCountry      string           `position:"Query" name:"ReceiverCountry"`
	ChargeType           string           `position:"Query" name:"ChargeType"`
	DataPlan             requests.Integer `position:"Query" name:"DataPlan"`
	ReceiverZip          string           `position:"Query" name:"ReceiverZip"`
}

// CreateSmartAccessGatewayResponse is the response struct for api CreateSmartAccessGateway
type CreateSmartAccessGatewayResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	SmartAGId   string `json:"SmartAGId" xml:"SmartAGId"`
	Name        string `json:"Name" xml:"Name"`
	OrderId     string `json:"OrderId" xml:"OrderId"`
	Description string `json:"Description" xml:"Description"`
}

// CreateCreateSmartAccessGatewayRequest creates a request to invoke CreateSmartAccessGateway API
func CreateCreateSmartAccessGatewayRequest() (request *CreateSmartAccessGatewayRequest) {
	request = &CreateSmartAccessGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "CreateSmartAccessGateway", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSmartAccessGatewayResponse creates a response to parse from CreateSmartAccessGateway response
func CreateCreateSmartAccessGatewayResponse() (response *CreateSmartAccessGatewayResponse) {
	response = &CreateSmartAccessGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
