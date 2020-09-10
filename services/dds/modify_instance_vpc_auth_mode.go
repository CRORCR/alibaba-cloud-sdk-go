package dds

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

// ModifyInstanceVpcAuthMode invokes the dds.ModifyInstanceVpcAuthMode API synchronously
// api document: https://help.aliyun.com/api/dds/modifyinstancevpcauthmode.html
func (client *Client) ModifyInstanceVpcAuthMode(request *ModifyInstanceVpcAuthModeRequest) (response *ModifyInstanceVpcAuthModeResponse, err error) {
	response = CreateModifyInstanceVpcAuthModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceVpcAuthModeWithChan invokes the dds.ModifyInstanceVpcAuthMode API asynchronously
// api document: https://help.aliyun.com/api/dds/modifyinstancevpcauthmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceVpcAuthModeWithChan(request *ModifyInstanceVpcAuthModeRequest) (<-chan *ModifyInstanceVpcAuthModeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceVpcAuthModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceVpcAuthMode(request)
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

// ModifyInstanceVpcAuthModeWithCallback invokes the dds.ModifyInstanceVpcAuthMode API asynchronously
// api document: https://help.aliyun.com/api/dds/modifyinstancevpcauthmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceVpcAuthModeWithCallback(request *ModifyInstanceVpcAuthModeRequest, callback func(response *ModifyInstanceVpcAuthModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceVpcAuthModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceVpcAuthMode(request)
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

// ModifyInstanceVpcAuthModeRequest is the request struct for api ModifyInstanceVpcAuthMode
type ModifyInstanceVpcAuthModeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	VpcAuthMode          string           `position:"Query" name:"VpcAuthMode"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyInstanceVpcAuthModeResponse is the response struct for api ModifyInstanceVpcAuthMode
type ModifyInstanceVpcAuthModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceVpcAuthModeRequest creates a request to invoke ModifyInstanceVpcAuthMode API
func CreateModifyInstanceVpcAuthModeRequest() (request *ModifyInstanceVpcAuthModeRequest) {
	request = &ModifyInstanceVpcAuthModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyInstanceVpcAuthMode", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceVpcAuthModeResponse creates a response to parse from ModifyInstanceVpcAuthMode response
func CreateModifyInstanceVpcAuthModeResponse() (response *ModifyInstanceVpcAuthModeResponse) {
	response = &ModifyInstanceVpcAuthModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
