package baas

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

// CheckFabricOrganizationDomain invokes the baas.CheckFabricOrganizationDomain API synchronously
// api document: https://help.aliyun.com/api/baas/checkfabricorganizationdomain.html
func (client *Client) CheckFabricOrganizationDomain(request *CheckFabricOrganizationDomainRequest) (response *CheckFabricOrganizationDomainResponse, err error) {
	response = CreateCheckFabricOrganizationDomainResponse()
	err = client.DoAction(request, response)
	return
}

// CheckFabricOrganizationDomainWithChan invokes the baas.CheckFabricOrganizationDomain API asynchronously
// api document: https://help.aliyun.com/api/baas/checkfabricorganizationdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckFabricOrganizationDomainWithChan(request *CheckFabricOrganizationDomainRequest) (<-chan *CheckFabricOrganizationDomainResponse, <-chan error) {
	responseChan := make(chan *CheckFabricOrganizationDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckFabricOrganizationDomain(request)
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

// CheckFabricOrganizationDomainWithCallback invokes the baas.CheckFabricOrganizationDomain API asynchronously
// api document: https://help.aliyun.com/api/baas/checkfabricorganizationdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckFabricOrganizationDomainWithCallback(request *CheckFabricOrganizationDomainRequest, callback func(response *CheckFabricOrganizationDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckFabricOrganizationDomainResponse
		var err error
		defer close(result)
		response, err = client.CheckFabricOrganizationDomain(request)
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

// CheckFabricOrganizationDomainRequest is the request struct for api CheckFabricOrganizationDomain
type CheckFabricOrganizationDomainRequest struct {
	*requests.RpcRequest
	DomainCode string `position:"Body" name:"DomainCode"`
	Domain     string `position:"Body" name:"Domain"`
}

// CheckFabricOrganizationDomainResponse is the response struct for api CheckFabricOrganizationDomain
type CheckFabricOrganizationDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCheckFabricOrganizationDomainRequest creates a request to invoke CheckFabricOrganizationDomain API
func CreateCheckFabricOrganizationDomainRequest() (request *CheckFabricOrganizationDomainRequest) {
	request = &CheckFabricOrganizationDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "CheckFabricOrganizationDomain", "baas", "openAPI")
	return
}

// CreateCheckFabricOrganizationDomainResponse creates a response to parse from CheckFabricOrganizationDomain response
func CreateCheckFabricOrganizationDomainResponse() (response *CheckFabricOrganizationDomainResponse) {
	response = &CheckFabricOrganizationDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
