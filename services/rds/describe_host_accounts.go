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

// DescribeHostAccounts invokes the rds.DescribeHostAccounts API synchronously
// api document: https://help.aliyun.com/api/rds/describehostaccounts.html
func (client *Client) DescribeHostAccounts(request *DescribeHostAccountsRequest) (response *DescribeHostAccountsResponse, err error) {
	response = CreateDescribeHostAccountsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHostAccountsWithChan invokes the rds.DescribeHostAccounts API asynchronously
// api document: https://help.aliyun.com/api/rds/describehostaccounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHostAccountsWithChan(request *DescribeHostAccountsRequest) (<-chan *DescribeHostAccountsResponse, <-chan error) {
	responseChan := make(chan *DescribeHostAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHostAccounts(request)
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

// DescribeHostAccountsWithCallback invokes the rds.DescribeHostAccounts API asynchronously
// api document: https://help.aliyun.com/api/rds/describehostaccounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHostAccountsWithCallback(request *DescribeHostAccountsRequest, callback func(response *DescribeHostAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHostAccountsResponse
		var err error
		defer close(result)
		response, err = client.DescribeHostAccounts(request)
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

// DescribeHostAccountsRequest is the request struct for api DescribeHostAccounts
type DescribeHostAccountsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeHostAccountsResponse is the response struct for api DescribeHostAccounts
type DescribeHostAccountsResponse struct {
	*responses.BaseResponse
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Accounts  AccountsInDescribeHostAccounts `json:"Accounts" xml:"Accounts"`
}

// CreateDescribeHostAccountsRequest creates a request to invoke DescribeHostAccounts API
func CreateDescribeHostAccountsRequest() (request *DescribeHostAccountsRequest) {
	request = &DescribeHostAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeHostAccounts", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHostAccountsResponse creates a response to parse from DescribeHostAccounts response
func CreateDescribeHostAccountsResponse() (response *DescribeHostAccountsResponse) {
	response = &DescribeHostAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
