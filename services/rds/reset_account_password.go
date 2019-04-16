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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ResetAccountPassword invokes the rds.ResetAccountPassword API synchronously
// api document: https://help.aliyun.com/api/rds/resetaccountpassword.html
func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
	response = CreateResetAccountPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ResetAccountPasswordWithChan invokes the rds.ResetAccountPassword API asynchronously
// api document: https://help.aliyun.com/api/rds/resetaccountpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAccountPasswordWithChan(request *ResetAccountPasswordRequest) (<-chan *ResetAccountPasswordResponse, <-chan error) {
	responseChan := make(chan *ResetAccountPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetAccountPassword(request)
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

// ResetAccountPasswordWithCallback invokes the rds.ResetAccountPassword API asynchronously
// api document: https://help.aliyun.com/api/rds/resetaccountpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetAccountPasswordWithCallback(request *ResetAccountPasswordRequest, callback func(response *ResetAccountPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetAccountPasswordResponse
		var err error
		defer close(result)
		response, err = client.ResetAccountPassword(request)
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

// ResetAccountPasswordRequest is the request struct for api ResetAccountPassword
type ResetAccountPasswordRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ResetAccountPasswordResponse is the response struct for api ResetAccountPassword
type ResetAccountPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetAccountPasswordRequest creates a request to invoke ResetAccountPassword API
func CreateResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
	request = &ResetAccountPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ResetAccountPassword", "Rds", "openAPI")
	return
}

// CreateResetAccountPasswordResponse creates a response to parse from ResetAccountPassword response
func CreateResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
	response = &ResetAccountPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
