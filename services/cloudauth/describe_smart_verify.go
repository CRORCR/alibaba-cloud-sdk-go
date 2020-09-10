package cloudauth

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

// DescribeSmartVerify invokes the cloudauth.DescribeSmartVerify API synchronously
// api document: https://help.aliyun.com/api/cloudauth/describesmartverify.html
func (client *Client) DescribeSmartVerify(request *DescribeSmartVerifyRequest) (response *DescribeSmartVerifyResponse, err error) {
	response = CreateDescribeSmartVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSmartVerifyWithChan invokes the cloudauth.DescribeSmartVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describesmartverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSmartVerifyWithChan(request *DescribeSmartVerifyRequest) (<-chan *DescribeSmartVerifyResponse, <-chan error) {
	responseChan := make(chan *DescribeSmartVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSmartVerify(request)
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

// DescribeSmartVerifyWithCallback invokes the cloudauth.DescribeSmartVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describesmartverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSmartVerifyWithCallback(request *DescribeSmartVerifyRequest, callback func(response *DescribeSmartVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSmartVerifyResponse
		var err error
		defer close(result)
		response, err = client.DescribeSmartVerify(request)
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

// DescribeSmartVerifyRequest is the request struct for api DescribeSmartVerify
type DescribeSmartVerifyRequest struct {
	*requests.RpcRequest
	SceneId   requests.Integer `position:"Body" name:"SceneId"`
	CertifyId string           `position:"Body" name:"CertifyId"`
}

// DescribeSmartVerifyResponse is the response struct for api DescribeSmartVerify
type DescribeSmartVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateDescribeSmartVerifyRequest creates a request to invoke DescribeSmartVerify API
func CreateDescribeSmartVerifyRequest() (request *DescribeSmartVerifyRequest) {
	request = &DescribeSmartVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2020-06-18", "DescribeSmartVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSmartVerifyResponse creates a response to parse from DescribeSmartVerify response
func CreateDescribeSmartVerifyResponse() (response *DescribeSmartVerifyResponse) {
	response = &DescribeSmartVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
