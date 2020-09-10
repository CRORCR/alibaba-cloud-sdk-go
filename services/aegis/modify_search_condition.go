package aegis

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

// ModifySearchCondition invokes the aegis.ModifySearchCondition API synchronously
// api document: https://help.aliyun.com/api/aegis/modifysearchcondition.html
func (client *Client) ModifySearchCondition(request *ModifySearchConditionRequest) (response *ModifySearchConditionResponse, err error) {
	response = CreateModifySearchConditionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySearchConditionWithChan invokes the aegis.ModifySearchCondition API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifysearchcondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySearchConditionWithChan(request *ModifySearchConditionRequest) (<-chan *ModifySearchConditionResponse, <-chan error) {
	responseChan := make(chan *ModifySearchConditionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySearchCondition(request)
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

// ModifySearchConditionWithCallback invokes the aegis.ModifySearchCondition API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifysearchcondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySearchConditionWithCallback(request *ModifySearchConditionRequest, callback func(response *ModifySearchConditionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySearchConditionResponse
		var err error
		defer close(result)
		response, err = client.ModifySearchCondition(request)
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

// ModifySearchConditionRequest is the request struct for api ModifySearchCondition
type ModifySearchConditionRequest struct {
	*requests.RpcRequest
	SourceIp         string `position:"Query" name:"SourceIp"`
	Name             string `position:"Query" name:"Name"`
	FilterConditions string `position:"Query" name:"FilterConditions"`
}

// ModifySearchConditionResponse is the response struct for api ModifySearchCondition
type ModifySearchConditionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySearchConditionRequest creates a request to invoke ModifySearchCondition API
func CreateModifySearchConditionRequest() (request *ModifySearchConditionRequest) {
	request = &ModifySearchConditionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifySearchCondition", "vipaegis", "openAPI")
	return
}

// CreateModifySearchConditionResponse creates a response to parse from ModifySearchCondition response
func CreateModifySearchConditionResponse() (response *ModifySearchConditionResponse) {
	response = &ModifySearchConditionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
