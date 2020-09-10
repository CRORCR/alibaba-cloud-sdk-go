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

// DescribeFabricChaincodeUploadPolicy invokes the baas.DescribeFabricChaincodeUploadPolicy API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricchaincodeuploadpolicy.html
func (client *Client) DescribeFabricChaincodeUploadPolicy(request *DescribeFabricChaincodeUploadPolicyRequest) (response *DescribeFabricChaincodeUploadPolicyResponse, err error) {
	response = CreateDescribeFabricChaincodeUploadPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricChaincodeUploadPolicyWithChan invokes the baas.DescribeFabricChaincodeUploadPolicy API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricchaincodeuploadpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricChaincodeUploadPolicyWithChan(request *DescribeFabricChaincodeUploadPolicyRequest) (<-chan *DescribeFabricChaincodeUploadPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricChaincodeUploadPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricChaincodeUploadPolicy(request)
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

// DescribeFabricChaincodeUploadPolicyWithCallback invokes the baas.DescribeFabricChaincodeUploadPolicy API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricchaincodeuploadpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricChaincodeUploadPolicyWithCallback(request *DescribeFabricChaincodeUploadPolicyRequest, callback func(response *DescribeFabricChaincodeUploadPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricChaincodeUploadPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricChaincodeUploadPolicy(request)
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

// DescribeFabricChaincodeUploadPolicyRequest is the request struct for api DescribeFabricChaincodeUploadPolicy
type DescribeFabricChaincodeUploadPolicyRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Body" name:"OrganizationId"`
}

// DescribeFabricChaincodeUploadPolicyResponse is the response struct for api DescribeFabricChaincodeUploadPolicy
type DescribeFabricChaincodeUploadPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeFabricChaincodeUploadPolicyRequest creates a request to invoke DescribeFabricChaincodeUploadPolicy API
func CreateDescribeFabricChaincodeUploadPolicyRequest() (request *DescribeFabricChaincodeUploadPolicyRequest) {
	request = &DescribeFabricChaincodeUploadPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricChaincodeUploadPolicy", "baas", "openAPI")
	return
}

// CreateDescribeFabricChaincodeUploadPolicyResponse creates a response to parse from DescribeFabricChaincodeUploadPolicy response
func CreateDescribeFabricChaincodeUploadPolicyResponse() (response *DescribeFabricChaincodeUploadPolicyResponse) {
	response = &DescribeFabricChaincodeUploadPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
