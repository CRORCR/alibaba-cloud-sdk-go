package rds

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

// SwitchDBInstanceVpc invokes the rds.SwitchDBInstanceVpc API synchronously
// api document: https://help.aliyun.com/api/rds/switchdbinstancevpc.html
func (client *Client) SwitchDBInstanceVpc(request *SwitchDBInstanceVpcRequest) (response *SwitchDBInstanceVpcResponse, err error) {
	response = CreateSwitchDBInstanceVpcResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchDBInstanceVpcWithChan invokes the rds.SwitchDBInstanceVpc API asynchronously
// api document: https://help.aliyun.com/api/rds/switchdbinstancevpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchDBInstanceVpcWithChan(request *SwitchDBInstanceVpcRequest) (<-chan *SwitchDBInstanceVpcResponse, <-chan error) {
	responseChan := make(chan *SwitchDBInstanceVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchDBInstanceVpc(request)
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

// SwitchDBInstanceVpcWithCallback invokes the rds.SwitchDBInstanceVpc API asynchronously
// api document: https://help.aliyun.com/api/rds/switchdbinstancevpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SwitchDBInstanceVpcWithCallback(request *SwitchDBInstanceVpcRequest, callback func(response *SwitchDBInstanceVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchDBInstanceVpcResponse
		var err error
		defer close(result)
		response, err = client.SwitchDBInstanceVpc(request)
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

// SwitchDBInstanceVpcRequest is the request struct for api SwitchDBInstanceVpc
type SwitchDBInstanceVpcRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string           `position:"Query" name:"PrivateIpAddress"`
	VPCId                string           `position:"Query" name:"VPCId"`
}

// SwitchDBInstanceVpcResponse is the response struct for api SwitchDBInstanceVpc
type SwitchDBInstanceVpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchDBInstanceVpcRequest creates a request to invoke SwitchDBInstanceVpc API
func CreateSwitchDBInstanceVpcRequest() (request *SwitchDBInstanceVpcRequest) {
	request = &SwitchDBInstanceVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "SwitchDBInstanceVpc", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchDBInstanceVpcResponse creates a response to parse from SwitchDBInstanceVpc response
func CreateSwitchDBInstanceVpcResponse() (response *SwitchDBInstanceVpcResponse) {
	response = &SwitchDBInstanceVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
