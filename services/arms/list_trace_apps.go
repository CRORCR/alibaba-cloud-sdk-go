package arms

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

// ListTraceApps invokes the arms.ListTraceApps API synchronously
// api document: https://help.aliyun.com/api/arms/listtraceapps.html
func (client *Client) ListTraceApps(request *ListTraceAppsRequest) (response *ListTraceAppsResponse, err error) {
	response = CreateListTraceAppsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTraceAppsWithChan invokes the arms.ListTraceApps API asynchronously
// api document: https://help.aliyun.com/api/arms/listtraceapps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTraceAppsWithChan(request *ListTraceAppsRequest) (<-chan *ListTraceAppsResponse, <-chan error) {
	responseChan := make(chan *ListTraceAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTraceApps(request)
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

// ListTraceAppsWithCallback invokes the arms.ListTraceApps API asynchronously
// api document: https://help.aliyun.com/api/arms/listtraceapps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTraceAppsWithCallback(request *ListTraceAppsRequest, callback func(response *ListTraceAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTraceAppsResponse
		var err error
		defer close(result)
		response, err = client.ListTraceApps(request)
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

// ListTraceAppsRequest is the request struct for api ListTraceApps
type ListTraceAppsRequest struct {
	*requests.RpcRequest
}

// ListTraceAppsResponse is the response struct for api ListTraceApps
type ListTraceAppsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	TraceApps []TraceApp `json:"TraceApps" xml:"TraceApps"`
}

// CreateListTraceAppsRequest creates a request to invoke ListTraceApps API
func CreateListTraceAppsRequest() (request *ListTraceAppsRequest) {
	request = &ListTraceAppsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListTraceApps", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTraceAppsResponse creates a response to parse from ListTraceApps response
func CreateListTraceAppsResponse() (response *ListTraceAppsResponse) {
	response = &ListTraceAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
