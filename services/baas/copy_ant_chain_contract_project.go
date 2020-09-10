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

// CopyAntChainContractProject invokes the baas.CopyAntChainContractProject API synchronously
// api document: https://help.aliyun.com/api/baas/copyantchaincontractproject.html
func (client *Client) CopyAntChainContractProject(request *CopyAntChainContractProjectRequest) (response *CopyAntChainContractProjectResponse, err error) {
	response = CreateCopyAntChainContractProjectResponse()
	err = client.DoAction(request, response)
	return
}

// CopyAntChainContractProjectWithChan invokes the baas.CopyAntChainContractProject API asynchronously
// api document: https://help.aliyun.com/api/baas/copyantchaincontractproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyAntChainContractProjectWithChan(request *CopyAntChainContractProjectRequest) (<-chan *CopyAntChainContractProjectResponse, <-chan error) {
	responseChan := make(chan *CopyAntChainContractProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyAntChainContractProject(request)
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

// CopyAntChainContractProjectWithCallback invokes the baas.CopyAntChainContractProject API asynchronously
// api document: https://help.aliyun.com/api/baas/copyantchaincontractproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyAntChainContractProjectWithCallback(request *CopyAntChainContractProjectRequest, callback func(response *CopyAntChainContractProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyAntChainContractProjectResponse
		var err error
		defer close(result)
		response, err = client.CopyAntChainContractProject(request)
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

// CopyAntChainContractProjectRequest is the request struct for api CopyAntChainContractProject
type CopyAntChainContractProjectRequest struct {
	*requests.RpcRequest
	ProjectVersion     string `position:"Body" name:"ProjectVersion"`
	ProjectId          string `position:"Body" name:"ProjectId"`
	ProjectName        string `position:"Body" name:"ProjectName"`
	ProjectDescription string `position:"Body" name:"ProjectDescription"`
}

// CopyAntChainContractProjectResponse is the response struct for api CopyAntChainContractProject
type CopyAntChainContractProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCopyAntChainContractProjectRequest creates a request to invoke CopyAntChainContractProject API
func CreateCopyAntChainContractProjectRequest() (request *CopyAntChainContractProjectRequest) {
	request = &CopyAntChainContractProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "CopyAntChainContractProject", "baas", "openAPI")
	return
}

// CreateCopyAntChainContractProjectResponse creates a response to parse from CopyAntChainContractProject response
func CreateCopyAntChainContractProjectResponse() (response *CopyAntChainContractProjectResponse) {
	response = &CopyAntChainContractProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
