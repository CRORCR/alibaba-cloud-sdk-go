package drds

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

// DescribeInstanceSwitchNetwork invokes the drds.DescribeInstanceSwitchNetwork API synchronously
// api document: https://help.aliyun.com/api/drds/describeinstanceswitchnetwork.html
func (client *Client) DescribeInstanceSwitchNetwork(request *DescribeInstanceSwitchNetworkRequest) (response *DescribeInstanceSwitchNetworkResponse, err error) {
	response = CreateDescribeInstanceSwitchNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSwitchNetworkWithChan invokes the drds.DescribeInstanceSwitchNetwork API asynchronously
// api document: https://help.aliyun.com/api/drds/describeinstanceswitchnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSwitchNetworkWithChan(request *DescribeInstanceSwitchNetworkRequest) (<-chan *DescribeInstanceSwitchNetworkResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSwitchNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSwitchNetwork(request)
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

// DescribeInstanceSwitchNetworkWithCallback invokes the drds.DescribeInstanceSwitchNetwork API asynchronously
// api document: https://help.aliyun.com/api/drds/describeinstanceswitchnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSwitchNetworkWithCallback(request *DescribeInstanceSwitchNetworkRequest, callback func(response *DescribeInstanceSwitchNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSwitchNetworkResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSwitchNetwork(request)
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

// DescribeInstanceSwitchNetworkRequest is the request struct for api DescribeInstanceSwitchNetwork
type DescribeInstanceSwitchNetworkRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeInstanceSwitchNetworkResponse is the response struct for api DescribeInstanceSwitchNetwork
type DescribeInstanceSwitchNetworkResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	VpcInfos  VpcInfos `json:"VpcInfos" xml:"VpcInfos"`
}

// CreateDescribeInstanceSwitchNetworkRequest creates a request to invoke DescribeInstanceSwitchNetwork API
func CreateDescribeInstanceSwitchNetworkRequest() (request *DescribeInstanceSwitchNetworkRequest) {
	request = &DescribeInstanceSwitchNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeInstanceSwitchNetwork", "Drds", "openAPI")
	return
}

// CreateDescribeInstanceSwitchNetworkResponse creates a response to parse from DescribeInstanceSwitchNetwork response
func CreateDescribeInstanceSwitchNetworkResponse() (response *DescribeInstanceSwitchNetworkResponse) {
	response = &DescribeInstanceSwitchNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
