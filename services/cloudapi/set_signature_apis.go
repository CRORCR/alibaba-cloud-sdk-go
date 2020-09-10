package cloudapi

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

// SetSignatureApis invokes the cloudapi.SetSignatureApis API synchronously
// api document: https://help.aliyun.com/api/cloudapi/setsignatureapis.html
func (client *Client) SetSignatureApis(request *SetSignatureApisRequest) (response *SetSignatureApisResponse, err error) {
	response = CreateSetSignatureApisResponse()
	err = client.DoAction(request, response)
	return
}

// SetSignatureApisWithChan invokes the cloudapi.SetSignatureApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/setsignatureapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetSignatureApisWithChan(request *SetSignatureApisRequest) (<-chan *SetSignatureApisResponse, <-chan error) {
	responseChan := make(chan *SetSignatureApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetSignatureApis(request)
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

// SetSignatureApisWithCallback invokes the cloudapi.SetSignatureApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/setsignatureapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetSignatureApisWithCallback(request *SetSignatureApisRequest, callback func(response *SetSignatureApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetSignatureApisResponse
		var err error
		defer close(result)
		response, err = client.SetSignatureApis(request)
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

// SetSignatureApisRequest is the request struct for api SetSignatureApis
type SetSignatureApisRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	GroupId       string `position:"Query" name:"GroupId"`
	SignatureId   string `position:"Query" name:"SignatureId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiIds        string `position:"Query" name:"ApiIds"`
}

// SetSignatureApisResponse is the response struct for api SetSignatureApis
type SetSignatureApisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetSignatureApisRequest creates a request to invoke SetSignatureApis API
func CreateSetSignatureApisRequest() (request *SetSignatureApisRequest) {
	request = &SetSignatureApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SetSignatureApis", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetSignatureApisResponse creates a response to parse from SetSignatureApis response
func CreateSetSignatureApisResponse() (response *SetSignatureApisResponse) {
	response = &SetSignatureApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
