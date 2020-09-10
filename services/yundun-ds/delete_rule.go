package yundun_ds

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

// DeleteRule invokes the yundun_ds.DeleteRule API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/deleterule.html
func (client *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
	response = CreateDeleteRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRuleWithChan invokes the yundun_ds.DeleteRule API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/deleterule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRuleWithChan(request *DeleteRuleRequest) (<-chan *DeleteRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRule(request)
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

// DeleteRuleWithCallback invokes the yundun_ds.DeleteRule API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/deleterule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRuleWithCallback(request *DeleteRuleRequest, callback func(response *DeleteRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteRule(request)
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

// DeleteRuleRequest is the request struct for api DeleteRule
type DeleteRuleRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DeleteRuleResponse is the response struct for api DeleteRule
type DeleteRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRuleRequest creates a request to invoke DeleteRule API
func CreateDeleteRuleRequest() (request *DeleteRuleRequest) {
	request = &DeleteRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DeleteRule", "sddp", "openAPI")
	return
}

// CreateDeleteRuleResponse creates a response to parse from DeleteRule response
func CreateDeleteRuleResponse() (response *DeleteRuleResponse) {
	response = &DeleteRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
