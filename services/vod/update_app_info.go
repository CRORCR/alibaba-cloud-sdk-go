package vod

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

// UpdateAppInfo invokes the vod.UpdateAppInfo API synchronously
func (client *Client) UpdateAppInfo(request *UpdateAppInfoRequest) (response *UpdateAppInfoResponse, err error) {
	response = CreateUpdateAppInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAppInfoWithChan invokes the vod.UpdateAppInfo API asynchronously
func (client *Client) UpdateAppInfoWithChan(request *UpdateAppInfoRequest) (<-chan *UpdateAppInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateAppInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAppInfo(request)
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

// UpdateAppInfoWithCallback invokes the vod.UpdateAppInfo API asynchronously
func (client *Client) UpdateAppInfoWithCallback(request *UpdateAppInfoRequest, callback func(response *UpdateAppInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAppInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateAppInfo(request)
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

// UpdateAppInfoRequest is the request struct for api UpdateAppInfo
type UpdateAppInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	ResourceRealOwnerId  requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	AppName              string           `position:"Query" name:"AppName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AppId                string           `position:"Query" name:"AppId"`
	Status               string           `position:"Query" name:"Status"`
}

// UpdateAppInfoResponse is the response struct for api UpdateAppInfo
type UpdateAppInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAppInfoRequest creates a request to invoke UpdateAppInfo API
func CreateUpdateAppInfoRequest() (request *UpdateAppInfoRequest) {
	request = &UpdateAppInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateAppInfo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAppInfoResponse creates a response to parse from UpdateAppInfo response
func CreateUpdateAppInfoResponse() (response *UpdateAppInfoResponse) {
	response = &UpdateAppInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
