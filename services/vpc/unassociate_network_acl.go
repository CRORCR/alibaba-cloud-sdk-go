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

// UnassociateNetworkAcl invokes the vpc.UnassociateNetworkAcl API synchronously
// api document: https://help.aliyun.com/api/vpc/unassociatenetworkacl.html
func (client *Client) UnassociateNetworkAcl(request *UnassociateNetworkAclRequest) (response *UnassociateNetworkAclResponse, err error) {
	response = CreateUnassociateNetworkAclResponse()
	err = client.DoAction(request, response)
	return
}

// UnassociateNetworkAclWithChan invokes the vpc.UnassociateNetworkAcl API asynchronously
// api document: https://help.aliyun.com/api/vpc/unassociatenetworkacl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateNetworkAclWithChan(request *UnassociateNetworkAclRequest) (<-chan *UnassociateNetworkAclResponse, <-chan error) {
	responseChan := make(chan *UnassociateNetworkAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociateNetworkAcl(request)
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

// UnassociateNetworkAclWithCallback invokes the vpc.UnassociateNetworkAcl API asynchronously
// api document: https://help.aliyun.com/api/vpc/unassociatenetworkacl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateNetworkAclWithCallback(request *UnassociateNetworkAclRequest, callback func(response *UnassociateNetworkAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociateNetworkAclResponse
		var err error
		defer close(result)
		response, err = client.UnassociateNetworkAcl(request)
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

// UnassociateNetworkAclRequest is the request struct for api UnassociateNetworkAcl
type UnassociateNetworkAclRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                 `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                           `position:"Query" name:"ClientToken"`
	NetworkAclId         string                           `position:"Query" name:"NetworkAclId"`
	Resource             *[]UnassociateNetworkAclResource `position:"Query" name:"Resource"  type:"Repeated"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer                 `position:"Query" name:"OwnerId"`
}

// UnassociateNetworkAclResource is a repeated param struct in UnassociateNetworkAclRequest
type UnassociateNetworkAclResource struct {
	ResourceType string `name:"ResourceType"`
	ResourceId   string `name:"ResourceId"`
}

// UnassociateNetworkAclResponse is the response struct for api UnassociateNetworkAcl
type UnassociateNetworkAclResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnassociateNetworkAclRequest creates a request to invoke UnassociateNetworkAcl API
func CreateUnassociateNetworkAclRequest() (request *UnassociateNetworkAclRequest) {
	request = &UnassociateNetworkAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateNetworkAcl", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnassociateNetworkAclResponse creates a response to parse from UnassociateNetworkAcl response
func CreateUnassociateNetworkAclResponse() (response *UnassociateNetworkAclResponse) {
	response = &UnassociateNetworkAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
