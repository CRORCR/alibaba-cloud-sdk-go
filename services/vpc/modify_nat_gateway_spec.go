package vpc

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

// ModifyNatGatewaySpec invokes the vpc.ModifyNatGatewaySpec API synchronously
// api document: https://help.aliyun.com/api/vpc/modifynatgatewayspec.html
func (client *Client) ModifyNatGatewaySpec(request *ModifyNatGatewaySpecRequest) (response *ModifyNatGatewaySpecResponse, err error) {
	response = CreateModifyNatGatewaySpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNatGatewaySpecWithChan invokes the vpc.ModifyNatGatewaySpec API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifynatgatewayspec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNatGatewaySpecWithChan(request *ModifyNatGatewaySpecRequest) (<-chan *ModifyNatGatewaySpecResponse, <-chan error) {
	responseChan := make(chan *ModifyNatGatewaySpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNatGatewaySpec(request)
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

// ModifyNatGatewaySpecWithCallback invokes the vpc.ModifyNatGatewaySpec API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifynatgatewayspec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNatGatewaySpecWithCallback(request *ModifyNatGatewaySpecRequest, callback func(response *ModifyNatGatewaySpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNatGatewaySpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyNatGatewaySpec(request)
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

// ModifyNatGatewaySpecRequest is the request struct for api ModifyNatGatewaySpec
type ModifyNatGatewaySpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Spec                 string           `position:"Query" name:"Spec"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyNatGatewaySpecResponse is the response struct for api ModifyNatGatewaySpec
type ModifyNatGatewaySpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyNatGatewaySpecRequest creates a request to invoke ModifyNatGatewaySpec API
func CreateModifyNatGatewaySpecRequest() (request *ModifyNatGatewaySpecRequest) {
	request = &ModifyNatGatewaySpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNatGatewaySpec", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyNatGatewaySpecResponse creates a response to parse from ModifyNatGatewaySpec response
func CreateModifyNatGatewaySpecResponse() (response *ModifyNatGatewaySpecResponse) {
	response = &ModifyNatGatewaySpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
