package slb

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

// SetDomainExtensionAttribute invokes the slb.SetDomainExtensionAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/setdomainextensionattribute.html
func (client *Client) SetDomainExtensionAttribute(request *SetDomainExtensionAttributeRequest) (response *SetDomainExtensionAttributeResponse, err error) {
	response = CreateSetDomainExtensionAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// SetDomainExtensionAttributeWithChan invokes the slb.SetDomainExtensionAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/setdomainextensionattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainExtensionAttributeWithChan(request *SetDomainExtensionAttributeRequest) (<-chan *SetDomainExtensionAttributeResponse, <-chan error) {
	responseChan := make(chan *SetDomainExtensionAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomainExtensionAttribute(request)
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

// SetDomainExtensionAttributeWithCallback invokes the slb.SetDomainExtensionAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/setdomainextensionattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainExtensionAttributeWithCallback(request *SetDomainExtensionAttributeRequest, callback func(response *SetDomainExtensionAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainExtensionAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetDomainExtensionAttribute(request)
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

// SetDomainExtensionAttributeRequest is the request struct for api SetDomainExtensionAttribute
type SetDomainExtensionAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string                                          `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer                                `position:"Query" name:"ResourceOwnerId"`
	ServerCertificate    *[]SetDomainExtensionAttributeServerCertificate `position:"Query" name:"ServerCertificate"  type:"Repeated"`
	DomainExtensionId    string                                          `position:"Query" name:"DomainExtensionId"`
	ResourceOwnerAccount string                                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                          `position:"Query" name:"OwnerAccount"`
	CertificateId        *[]string                                       `position:"Query" name:"CertificateId"  type:"Repeated"`
	OwnerId              requests.Integer                                `position:"Query" name:"OwnerId"`
	ServerCertificateId  string                                          `position:"Query" name:"ServerCertificateId"`
	Tags                 string                                          `position:"Query" name:"Tags"`
}

// SetDomainExtensionAttributeServerCertificate is a repeated param struct in SetDomainExtensionAttributeRequest
type SetDomainExtensionAttributeServerCertificate struct {
	BindingType   string `name:"BindingType"`
	CertificateId string `name:"CertificateId"`
	StandardType  string `name:"StandardType"`
}

// SetDomainExtensionAttributeResponse is the response struct for api SetDomainExtensionAttribute
type SetDomainExtensionAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDomainExtensionAttributeRequest creates a request to invoke SetDomainExtensionAttribute API
func CreateSetDomainExtensionAttributeRequest() (request *SetDomainExtensionAttributeRequest) {
	request = &SetDomainExtensionAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetDomainExtensionAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDomainExtensionAttributeResponse creates a response to parse from SetDomainExtensionAttribute response
func CreateSetDomainExtensionAttributeResponse() (response *SetDomainExtensionAttributeResponse) {
	response = &SetDomainExtensionAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
