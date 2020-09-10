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

// ModifyInstanceMaintainTime invokes the r_kvstore.ModifyInstanceMaintainTime API synchronously
func (client *Client) ModifyInstanceMaintainTime(request *ModifyInstanceMaintainTimeRequest) (response *ModifyInstanceMaintainTimeResponse, err error) {
	response = CreateModifyInstanceMaintainTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceMaintainTimeWithChan invokes the r_kvstore.ModifyInstanceMaintainTime API asynchronously
func (client *Client) ModifyInstanceMaintainTimeWithChan(request *ModifyInstanceMaintainTimeRequest) (<-chan *ModifyInstanceMaintainTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceMaintainTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceMaintainTime(request)
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

// ModifyInstanceMaintainTimeWithCallback invokes the r_kvstore.ModifyInstanceMaintainTime API asynchronously
func (client *Client) ModifyInstanceMaintainTimeWithCallback(request *ModifyInstanceMaintainTimeRequest, callback func(response *ModifyInstanceMaintainTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceMaintainTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceMaintainTime(request)
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

// ModifyInstanceMaintainTimeRequest is the request struct for api ModifyInstanceMaintainTime
type ModifyInstanceMaintainTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	MaintainStartTime    string           `position:"Query" name:"MaintainStartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaintainEndTime      string           `position:"Query" name:"MaintainEndTime"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ModifyInstanceMaintainTimeResponse is the response struct for api ModifyInstanceMaintainTime
type ModifyInstanceMaintainTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceMaintainTimeRequest creates a request to invoke ModifyInstanceMaintainTime API
func CreateModifyInstanceMaintainTimeRequest() (request *ModifyInstanceMaintainTimeRequest) {
	request = &ModifyInstanceMaintainTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceMaintainTime", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceMaintainTimeResponse creates a response to parse from ModifyInstanceMaintainTime response
func CreateModifyInstanceMaintainTimeResponse() (response *ModifyInstanceMaintainTimeResponse) {
	response = &ModifyInstanceMaintainTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
