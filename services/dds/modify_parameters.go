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

// ModifyParameters invokes the dds.ModifyParameters API synchronously
// api document: https://help.aliyun.com/api/dds/modifyparameters.html
func (client *Client) ModifyParameters(request *ModifyParametersRequest) (response *ModifyParametersResponse, err error) {
	response = CreateModifyParametersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyParametersWithChan invokes the dds.ModifyParameters API asynchronously
// api document: https://help.aliyun.com/api/dds/modifyparameters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyParametersWithChan(request *ModifyParametersRequest) (<-chan *ModifyParametersResponse, <-chan error) {
	responseChan := make(chan *ModifyParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyParameters(request)
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

// ModifyParametersWithCallback invokes the dds.ModifyParameters API asynchronously
// api document: https://help.aliyun.com/api/dds/modifyparameters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyParametersWithCallback(request *ModifyParametersRequest, callback func(response *ModifyParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyParametersResponse
		var err error
		defer close(result)
		response, err = client.ModifyParameters(request)
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

// ModifyParametersRequest is the request struct for api ModifyParameters
type ModifyParametersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Parameters           string           `position:"Query" name:"Parameters"`
	CharacterType        string           `position:"Query" name:"CharacterType"`
}

// ModifyParametersResponse is the response struct for api ModifyParameters
type ModifyParametersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyParametersRequest creates a request to invoke ModifyParameters API
func CreateModifyParametersRequest() (request *ModifyParametersRequest) {
	request = &ModifyParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyParameters", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyParametersResponse creates a response to parse from ModifyParameters response
func CreateModifyParametersResponse() (response *ModifyParametersResponse) {
	response = &ModifyParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
