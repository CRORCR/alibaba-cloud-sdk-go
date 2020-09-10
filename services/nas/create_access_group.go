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

package nas

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateAccessGroup invokes the nas.CreateAccessGroup API synchronously
// api document: https://help.aliyun.com/api/nas/createaccessgroup.html
func (client *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (response *CreateAccessGroupResponse, err error) {
	response = CreateCreateAccessGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAccessGroupWithChan invokes the nas.CreateAccessGroup API asynchronously
// api document: https://help.aliyun.com/api/nas/createaccessgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccessGroupWithChan(request *CreateAccessGroupRequest) (<-chan *CreateAccessGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAccessGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAccessGroup(request)
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

// CreateAccessGroupWithCallback invokes the nas.CreateAccessGroup API asynchronously
// api document: https://help.aliyun.com/api/nas/createaccessgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccessGroupWithCallback(request *CreateAccessGroupRequest, callback func(response *CreateAccessGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAccessGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAccessGroup(request)
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

// CreateAccessGroupRequest is the request struct for api CreateAccessGroup
type CreateAccessGroupRequest struct {
	*requests.RpcRequest
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessGroupType string `position:"Query" name:"AccessGroupType"`
	Description     string `position:"Query" name:"Description"`
	FileSystemType  string `position:"Query" name:"FileSystemType"`
}

// CreateAccessGroupResponse is the response struct for api CreateAccessGroup
type CreateAccessGroupResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	AccessGroupName string `json:"AccessGroupName" xml:"AccessGroupName"`
}

// CreateCreateAccessGroupRequest creates a request to invoke CreateAccessGroup API
func CreateCreateAccessGroupRequest() (request *CreateAccessGroupRequest) {
	request = &CreateAccessGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateAccessGroup", "nas", "openAPI")
	return
}

// CreateCreateAccessGroupResponse creates a response to parse from CreateAccessGroup response
func CreateCreateAccessGroupResponse() (response *CreateAccessGroupResponse) {
	response = &CreateAccessGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
