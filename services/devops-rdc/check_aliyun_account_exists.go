package devops_rdc

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

// CheckAliyunAccountExists invokes the devops_rdc.CheckAliyunAccountExists API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/checkaliyunaccountexists.html
func (client *Client) CheckAliyunAccountExists(request *CheckAliyunAccountExistsRequest) (response *CheckAliyunAccountExistsResponse, err error) {
	response = CreateCheckAliyunAccountExistsResponse()
	err = client.DoAction(request, response)
	return
}

// CheckAliyunAccountExistsWithChan invokes the devops_rdc.CheckAliyunAccountExists API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/checkaliyunaccountexists.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAliyunAccountExistsWithChan(request *CheckAliyunAccountExistsRequest) (<-chan *CheckAliyunAccountExistsResponse, <-chan error) {
	responseChan := make(chan *CheckAliyunAccountExistsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAliyunAccountExists(request)
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

// CheckAliyunAccountExistsWithCallback invokes the devops_rdc.CheckAliyunAccountExists API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/checkaliyunaccountexists.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAliyunAccountExistsWithCallback(request *CheckAliyunAccountExistsRequest, callback func(response *CheckAliyunAccountExistsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAliyunAccountExistsResponse
		var err error
		defer close(result)
		response, err = client.CheckAliyunAccountExists(request)
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

// CheckAliyunAccountExistsRequest is the request struct for api CheckAliyunAccountExists
type CheckAliyunAccountExistsRequest struct {
	*requests.RpcRequest
	UserPk string `position:"Body" name:"UserPk"`
}

// CheckAliyunAccountExistsResponse is the response struct for api CheckAliyunAccountExists
type CheckAliyunAccountExistsResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     bool   `json:"Object" xml:"Object"`
}

// CreateCheckAliyunAccountExistsRequest creates a request to invoke CheckAliyunAccountExists API
func CreateCheckAliyunAccountExistsRequest() (request *CheckAliyunAccountExistsRequest) {
	request = &CheckAliyunAccountExistsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "CheckAliyunAccountExists", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckAliyunAccountExistsResponse creates a response to parse from CheckAliyunAccountExists response
func CreateCheckAliyunAccountExistsResponse() (response *CheckAliyunAccountExistsResponse) {
	response = &CheckAliyunAccountExistsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
