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

// ModifyDedicatedInstanceSpec invokes the r_kvstore.ModifyDedicatedInstanceSpec API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydedicatedinstancespec.html
func (client *Client) ModifyDedicatedInstanceSpec(request *ModifyDedicatedInstanceSpecRequest) (response *ModifyDedicatedInstanceSpecResponse, err error) {
	response = CreateModifyDedicatedInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDedicatedInstanceSpecWithChan invokes the r_kvstore.ModifyDedicatedInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydedicatedinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedInstanceSpecWithChan(request *ModifyDedicatedInstanceSpecRequest) (<-chan *ModifyDedicatedInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyDedicatedInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDedicatedInstanceSpec(request)
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

// ModifyDedicatedInstanceSpecWithCallback invokes the r_kvstore.ModifyDedicatedInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydedicatedinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedInstanceSpecWithCallback(request *ModifyDedicatedInstanceSpecRequest, callback func(response *ModifyDedicatedInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDedicatedInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyDedicatedInstanceSpec(request)
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

// ModifyDedicatedInstanceSpecRequest is the request struct for api ModifyDedicatedInstanceSpec
type ModifyDedicatedInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceClass        string           `position:"Query" name:"InstanceClass"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	ForceUpgrade         requests.Boolean `position:"Query" name:"ForceUpgrade"`
}

// ModifyDedicatedInstanceSpecResponse is the response struct for api ModifyDedicatedInstanceSpec
type ModifyDedicatedInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDedicatedInstanceSpecRequest creates a request to invoke ModifyDedicatedInstanceSpec API
func CreateModifyDedicatedInstanceSpecRequest() (request *ModifyDedicatedInstanceSpecRequest) {
	request = &ModifyDedicatedInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyDedicatedInstanceSpec", "redisa", "openAPI")
	return
}

// CreateModifyDedicatedInstanceSpecResponse creates a response to parse from ModifyDedicatedInstanceSpec response
func CreateModifyDedicatedInstanceSpecResponse() (response *ModifyDedicatedInstanceSpecResponse) {
	response = &ModifyDedicatedInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
