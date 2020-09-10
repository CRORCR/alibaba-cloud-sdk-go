package ecs

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

// DescribeInstanceRamRole invokes the ecs.DescribeInstanceRamRole API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstanceramrole.html
func (client *Client) DescribeInstanceRamRole(request *DescribeInstanceRamRoleRequest) (response *DescribeInstanceRamRoleResponse, err error) {
	response = CreateDescribeInstanceRamRoleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceRamRoleWithChan invokes the ecs.DescribeInstanceRamRole API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstanceramrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceRamRoleWithChan(request *DescribeInstanceRamRoleRequest) (<-chan *DescribeInstanceRamRoleResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceRamRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceRamRole(request)
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

// DescribeInstanceRamRoleWithCallback invokes the ecs.DescribeInstanceRamRole API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstanceramrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceRamRoleWithCallback(request *DescribeInstanceRamRoleRequest, callback func(response *DescribeInstanceRamRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceRamRoleResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceRamRole(request)
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

// DescribeInstanceRamRoleRequest is the request struct for api DescribeInstanceRamRole
type DescribeInstanceRamRoleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RamRoleName          string           `position:"Query" name:"RamRoleName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
}

// DescribeInstanceRamRoleResponse is the response struct for api DescribeInstanceRamRole
type DescribeInstanceRamRoleResponse struct {
	*responses.BaseResponse
	RequestId           string                                       `json:"RequestId" xml:"RequestId"`
	RegionId            string                                       `json:"RegionId" xml:"RegionId"`
	TotalCount          int                                          `json:"TotalCount" xml:"TotalCount"`
	InstanceRamRoleSets InstanceRamRoleSetsInDescribeInstanceRamRole `json:"InstanceRamRoleSets" xml:"InstanceRamRoleSets"`
}

// CreateDescribeInstanceRamRoleRequest creates a request to invoke DescribeInstanceRamRole API
func CreateDescribeInstanceRamRoleRequest() (request *DescribeInstanceRamRoleRequest) {
	request = &DescribeInstanceRamRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceRamRole", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceRamRoleResponse creates a response to parse from DescribeInstanceRamRole response
func CreateDescribeInstanceRamRoleResponse() (response *DescribeInstanceRamRoleResponse) {
	response = &DescribeInstanceRamRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
