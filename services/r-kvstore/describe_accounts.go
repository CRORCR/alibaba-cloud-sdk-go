package r_kvstore

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

// DescribeAccounts invokes the r_kvstore.DescribeAccounts API synchronously
func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
	response = CreateDescribeAccountsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccountsWithChan invokes the r_kvstore.DescribeAccounts API asynchronously
func (client *Client) DescribeAccountsWithChan(request *DescribeAccountsRequest) (<-chan *DescribeAccountsResponse, <-chan error) {
	responseChan := make(chan *DescribeAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccounts(request)
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

// DescribeAccountsWithCallback invokes the r_kvstore.DescribeAccounts API asynchronously
func (client *Client) DescribeAccountsWithCallback(request *DescribeAccountsRequest, callback func(response *DescribeAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccountsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccounts(request)
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

// DescribeAccountsRequest is the request struct for api DescribeAccounts
type DescribeAccountsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeAccountsResponse is the response struct for api DescribeAccounts
type DescribeAccountsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Accounts  Accounts `json:"Accounts" xml:"Accounts"`
}

// CreateDescribeAccountsRequest creates a request to invoke DescribeAccounts API
func CreateDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeAccounts", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAccountsResponse creates a response to parse from DescribeAccounts response
func CreateDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
