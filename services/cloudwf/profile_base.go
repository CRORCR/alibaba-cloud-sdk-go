package cloudwf

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

// ProfileBase invokes the cloudwf.ProfileBase API synchronously
// api document: https://help.aliyun.com/api/cloudwf/profilebase.html
func (client *Client) ProfileBase(request *ProfileBaseRequest) (response *ProfileBaseResponse, err error) {
	response = CreateProfileBaseResponse()
	err = client.DoAction(request, response)
	return
}

// ProfileBaseWithChan invokes the cloudwf.ProfileBase API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profilebase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileBaseWithChan(request *ProfileBaseRequest) (<-chan *ProfileBaseResponse, <-chan error) {
	responseChan := make(chan *ProfileBaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileBase(request)
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

// ProfileBaseWithCallback invokes the cloudwf.ProfileBase API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profilebase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileBaseWithCallback(request *ProfileBaseRequest, callback func(response *ProfileBaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileBaseResponse
		var err error
		defer close(result)
		response, err = client.ProfileBase(request)
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

// ProfileBaseRequest is the request struct for api ProfileBase
type ProfileBaseRequest struct {
	*requests.RpcRequest
	BeginDate string           `position:"Query" name:"BeginDate"`
	EndDate   string           `position:"Query" name:"EndDate"`
	DataType  requests.Integer `position:"Query" name:"DataType"`
	Gsid      requests.Integer `position:"Query" name:"Gsid"`
}

// ProfileBaseResponse is the response struct for api ProfileBase
type ProfileBaseResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateProfileBaseRequest creates a request to invoke ProfileBase API
func CreateProfileBaseRequest() (request *ProfileBaseRequest) {
	request = &ProfileBaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileBase", "cloudwf", "openAPI")
	return
}

// CreateProfileBaseResponse creates a response to parse from ProfileBase response
func CreateProfileBaseResponse() (response *ProfileBaseResponse) {
	response = &ProfileBaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
