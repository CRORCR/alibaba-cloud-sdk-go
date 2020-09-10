package alidns

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

// UpdateGtmAddressPool invokes the alidns.UpdateGtmAddressPool API synchronously
// api document: https://help.aliyun.com/api/alidns/updategtmaddresspool.html
func (client *Client) UpdateGtmAddressPool(request *UpdateGtmAddressPoolRequest) (response *UpdateGtmAddressPoolResponse, err error) {
	response = CreateUpdateGtmAddressPoolResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGtmAddressPoolWithChan invokes the alidns.UpdateGtmAddressPool API asynchronously
// api document: https://help.aliyun.com/api/alidns/updategtmaddresspool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGtmAddressPoolWithChan(request *UpdateGtmAddressPoolRequest) (<-chan *UpdateGtmAddressPoolResponse, <-chan error) {
	responseChan := make(chan *UpdateGtmAddressPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGtmAddressPool(request)
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

// UpdateGtmAddressPoolWithCallback invokes the alidns.UpdateGtmAddressPool API asynchronously
// api document: https://help.aliyun.com/api/alidns/updategtmaddresspool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGtmAddressPoolWithCallback(request *UpdateGtmAddressPoolRequest, callback func(response *UpdateGtmAddressPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGtmAddressPoolResponse
		var err error
		defer close(result)
		response, err = client.UpdateGtmAddressPool(request)
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

// UpdateGtmAddressPoolRequest is the request struct for api UpdateGtmAddressPool
type UpdateGtmAddressPoolRequest struct {
	*requests.RpcRequest
	Type                string                      `position:"Query" name:"Type"`
	MinAvailableAddrNum requests.Integer            `position:"Query" name:"MinAvailableAddrNum"`
	AddrPoolId          string                      `position:"Query" name:"AddrPoolId"`
	UserClientIp        string                      `position:"Query" name:"UserClientIp"`
	Name                string                      `position:"Query" name:"Name"`
	Lang                string                      `position:"Query" name:"Lang"`
	Addr                *[]UpdateGtmAddressPoolAddr `position:"Query" name:"Addr"  type:"Repeated"`
}

// UpdateGtmAddressPoolAddr is a repeated param struct in UpdateGtmAddressPoolRequest
type UpdateGtmAddressPoolAddr struct {
	Mode      string `name:"Mode"`
	LbaWeight string `name:"LbaWeight"`
	Value     string `name:"Value"`
}

// UpdateGtmAddressPoolResponse is the response struct for api UpdateGtmAddressPool
type UpdateGtmAddressPoolResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateGtmAddressPoolRequest creates a request to invoke UpdateGtmAddressPool API
func CreateUpdateGtmAddressPoolRequest() (request *UpdateGtmAddressPoolRequest) {
	request = &UpdateGtmAddressPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateGtmAddressPool", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGtmAddressPoolResponse creates a response to parse from UpdateGtmAddressPool response
func CreateUpdateGtmAddressPoolResponse() (response *UpdateGtmAddressPoolResponse) {
	response = &UpdateGtmAddressPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
