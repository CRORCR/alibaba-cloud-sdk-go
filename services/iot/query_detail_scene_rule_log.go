package iot

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

// QueryDetailSceneRuleLog invokes the iot.QueryDetailSceneRuleLog API synchronously
// api document: https://help.aliyun.com/api/iot/querydetailscenerulelog.html
func (client *Client) QueryDetailSceneRuleLog(request *QueryDetailSceneRuleLogRequest) (response *QueryDetailSceneRuleLogResponse, err error) {
	response = CreateQueryDetailSceneRuleLogResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDetailSceneRuleLogWithChan invokes the iot.QueryDetailSceneRuleLog API asynchronously
// api document: https://help.aliyun.com/api/iot/querydetailscenerulelog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDetailSceneRuleLogWithChan(request *QueryDetailSceneRuleLogRequest) (<-chan *QueryDetailSceneRuleLogResponse, <-chan error) {
	responseChan := make(chan *QueryDetailSceneRuleLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDetailSceneRuleLog(request)
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

// QueryDetailSceneRuleLogWithCallback invokes the iot.QueryDetailSceneRuleLog API asynchronously
// api document: https://help.aliyun.com/api/iot/querydetailscenerulelog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDetailSceneRuleLogWithCallback(request *QueryDetailSceneRuleLogRequest, callback func(response *QueryDetailSceneRuleLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDetailSceneRuleLogResponse
		var err error
		defer close(result)
		response, err = client.QueryDetailSceneRuleLog(request)
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

// QueryDetailSceneRuleLogRequest is the request struct for api QueryDetailSceneRuleLog
type QueryDetailSceneRuleLogRequest struct {
	*requests.RpcRequest
	TraceId       string           `position:"Query" name:"TraceId"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	RuleId        string           `position:"Query" name:"RuleId"`
}

// QueryDetailSceneRuleLogResponse is the response struct for api QueryDetailSceneRuleLog
type QueryDetailSceneRuleLogResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string                        `json:"Code" xml:"Code"`
	Data         DataInQueryDetailSceneRuleLog `json:"Data" xml:"Data"`
}

// CreateQueryDetailSceneRuleLogRequest creates a request to invoke QueryDetailSceneRuleLog API
func CreateQueryDetailSceneRuleLogRequest() (request *QueryDetailSceneRuleLogRequest) {
	request = &QueryDetailSceneRuleLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDetailSceneRuleLog", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDetailSceneRuleLogResponse creates a response to parse from QueryDetailSceneRuleLog response
func CreateQueryDetailSceneRuleLogResponse() (response *QueryDetailSceneRuleLogResponse) {
	response = &QueryDetailSceneRuleLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
