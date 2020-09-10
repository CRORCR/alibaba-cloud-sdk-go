package trademark

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

// QueryTradeMarkApplicationLogs invokes the trademark.QueryTradeMarkApplicationLogs API synchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkapplicationlogs.html
func (client *Client) QueryTradeMarkApplicationLogs(request *QueryTradeMarkApplicationLogsRequest) (response *QueryTradeMarkApplicationLogsResponse, err error) {
	response = CreateQueryTradeMarkApplicationLogsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTradeMarkApplicationLogsWithChan invokes the trademark.QueryTradeMarkApplicationLogs API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkapplicationlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTradeMarkApplicationLogsWithChan(request *QueryTradeMarkApplicationLogsRequest) (<-chan *QueryTradeMarkApplicationLogsResponse, <-chan error) {
	responseChan := make(chan *QueryTradeMarkApplicationLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTradeMarkApplicationLogs(request)
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

// QueryTradeMarkApplicationLogsWithCallback invokes the trademark.QueryTradeMarkApplicationLogs API asynchronously
// api document: https://help.aliyun.com/api/trademark/querytrademarkapplicationlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTradeMarkApplicationLogsWithCallback(request *QueryTradeMarkApplicationLogsRequest, callback func(response *QueryTradeMarkApplicationLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTradeMarkApplicationLogsResponse
		var err error
		defer close(result)
		response, err = client.QueryTradeMarkApplicationLogs(request)
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

// QueryTradeMarkApplicationLogsRequest is the request struct for api QueryTradeMarkApplicationLogs
type QueryTradeMarkApplicationLogsRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"BizId"`
}

// QueryTradeMarkApplicationLogsResponse is the response struct for api QueryTradeMarkApplicationLogs
type QueryTradeMarkApplicationLogsResponse struct {
	*responses.BaseResponse
	RequestId string                              `json:"RequestId" xml:"RequestId"`
	Data      DataInQueryTradeMarkApplicationLogs `json:"Data" xml:"Data"`
}

// CreateQueryTradeMarkApplicationLogsRequest creates a request to invoke QueryTradeMarkApplicationLogs API
func CreateQueryTradeMarkApplicationLogsRequest() (request *QueryTradeMarkApplicationLogsRequest) {
	request = &QueryTradeMarkApplicationLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryTradeMarkApplicationLogs", "trademark", "openAPI")
	return
}

// CreateQueryTradeMarkApplicationLogsResponse creates a response to parse from QueryTradeMarkApplicationLogs response
func CreateQueryTradeMarkApplicationLogsResponse() (response *QueryTradeMarkApplicationLogsResponse) {
	response = &QueryTradeMarkApplicationLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
