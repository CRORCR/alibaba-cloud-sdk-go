package cloudcallcenter

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

// GetInvokeRole invokes the cloudcallcenter.GetInvokeRole API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinvokerole.html
func (client *Client) GetInvokeRole(request *GetInvokeRoleRequest) (response *GetInvokeRoleResponse, err error) {
	response = CreateGetInvokeRoleResponse()
	err = client.DoAction(request, response)
	return
}

// GetInvokeRoleWithChan invokes the cloudcallcenter.GetInvokeRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinvokerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInvokeRoleWithChan(request *GetInvokeRoleRequest) (<-chan *GetInvokeRoleResponse, <-chan error) {
	responseChan := make(chan *GetInvokeRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInvokeRole(request)
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

// GetInvokeRoleWithCallback invokes the cloudcallcenter.GetInvokeRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinvokerole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInvokeRoleWithCallback(request *GetInvokeRoleRequest, callback func(response *GetInvokeRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInvokeRoleResponse
		var err error
		defer close(result)
		response, err = client.GetInvokeRole(request)
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

// GetInvokeRoleRequest is the request struct for api GetInvokeRole
type GetInvokeRoleRequest struct {
	*requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

// GetInvokeRoleResponse is the response struct for api GetInvokeRole
type GetInvokeRoleResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	HasRole        bool        `json:"HasRole" xml:"HasRole"`
	ServiceRole    ServiceRole `json:"ServiceRole" xml:"ServiceRole"`
}

// CreateGetInvokeRoleRequest creates a request to invoke GetInvokeRole API
func CreateGetInvokeRoleRequest() (request *GetInvokeRoleRequest) {
	request = &GetInvokeRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetInvokeRole", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInvokeRoleResponse creates a response to parse from GetInvokeRole response
func CreateGetInvokeRoleResponse() (response *GetInvokeRoleResponse) {
	response = &GetInvokeRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
