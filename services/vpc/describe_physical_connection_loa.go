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

// DescribePhysicalConnectionLOA invokes the vpc.DescribePhysicalConnectionLOA API synchronously
// api document: https://help.aliyun.com/api/vpc/describephysicalconnectionloa.html
func (client *Client) DescribePhysicalConnectionLOA(request *DescribePhysicalConnectionLOARequest) (response *DescribePhysicalConnectionLOAResponse, err error) {
	response = CreateDescribePhysicalConnectionLOAResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePhysicalConnectionLOAWithChan invokes the vpc.DescribePhysicalConnectionLOA API asynchronously
// api document: https://help.aliyun.com/api/vpc/describephysicalconnectionloa.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePhysicalConnectionLOAWithChan(request *DescribePhysicalConnectionLOARequest) (<-chan *DescribePhysicalConnectionLOAResponse, <-chan error) {
	responseChan := make(chan *DescribePhysicalConnectionLOAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePhysicalConnectionLOA(request)
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

// DescribePhysicalConnectionLOAWithCallback invokes the vpc.DescribePhysicalConnectionLOA API asynchronously
// api document: https://help.aliyun.com/api/vpc/describephysicalconnectionloa.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePhysicalConnectionLOAWithCallback(request *DescribePhysicalConnectionLOARequest, callback func(response *DescribePhysicalConnectionLOAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePhysicalConnectionLOAResponse
		var err error
		defer close(result)
		response, err = client.DescribePhysicalConnectionLOA(request)
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

// DescribePhysicalConnectionLOARequest is the request struct for api DescribePhysicalConnectionLOA
type DescribePhysicalConnectionLOARequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribePhysicalConnectionLOAResponse is the response struct for api DescribePhysicalConnectionLOA
type DescribePhysicalConnectionLOAResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	PhysicalConnectionLOAType PhysicalConnectionLOAType `json:"PhysicalConnectionLOAType" xml:"PhysicalConnectionLOAType"`
}

// CreateDescribePhysicalConnectionLOARequest creates a request to invoke DescribePhysicalConnectionLOA API
func CreateDescribePhysicalConnectionLOARequest() (request *DescribePhysicalConnectionLOARequest) {
	request = &DescribePhysicalConnectionLOARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribePhysicalConnectionLOA", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePhysicalConnectionLOAResponse creates a response to parse from DescribePhysicalConnectionLOA response
func CreateDescribePhysicalConnectionLOAResponse() (response *DescribePhysicalConnectionLOAResponse) {
	response = &DescribePhysicalConnectionLOAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
