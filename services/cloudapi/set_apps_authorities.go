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

// SetAppsAuthorities invokes the cloudapi.SetAppsAuthorities API synchronously
// api document: https://help.aliyun.com/api/cloudapi/setappsauthorities.html
func (client *Client) SetAppsAuthorities(request *SetAppsAuthoritiesRequest) (response *SetAppsAuthoritiesResponse, err error) {
	response = CreateSetAppsAuthoritiesResponse()
	err = client.DoAction(request, response)
	return
}

// SetAppsAuthoritiesWithChan invokes the cloudapi.SetAppsAuthorities API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/setappsauthorities.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetAppsAuthoritiesWithChan(request *SetAppsAuthoritiesRequest) (<-chan *SetAppsAuthoritiesResponse, <-chan error) {
	responseChan := make(chan *SetAppsAuthoritiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAppsAuthorities(request)
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

// SetAppsAuthoritiesWithCallback invokes the cloudapi.SetAppsAuthorities API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/setappsauthorities.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetAppsAuthoritiesWithCallback(request *SetAppsAuthoritiesRequest, callback func(response *SetAppsAuthoritiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAppsAuthoritiesResponse
		var err error
		defer close(result)
		response, err = client.SetAppsAuthorities(request)
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

// SetAppsAuthoritiesRequest is the request struct for api SetAppsAuthorities
type SetAppsAuthoritiesRequest struct {
	*requests.RpcRequest
	AuthVaildTime string `position:"Query" name:"AuthVaildTime"`
	StageName     string `position:"Query" name:"StageName"`
	GroupId       string `position:"Query" name:"GroupId"`
	Description   string `position:"Query" name:"Description"`
	AuthValidTime string `position:"Query" name:"AuthValidTime"`
	AppIds        string `position:"Query" name:"AppIds"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// SetAppsAuthoritiesResponse is the response struct for api SetAppsAuthorities
type SetAppsAuthoritiesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetAppsAuthoritiesRequest creates a request to invoke SetAppsAuthorities API
func CreateSetAppsAuthoritiesRequest() (request *SetAppsAuthoritiesRequest) {
	request = &SetAppsAuthoritiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SetAppsAuthorities", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetAppsAuthoritiesResponse creates a response to parse from SetAppsAuthorities response
func CreateSetAppsAuthoritiesResponse() (response *SetAppsAuthoritiesResponse) {
	response = &SetAppsAuthoritiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
