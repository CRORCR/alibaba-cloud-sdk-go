package cloudcallcenter

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

// ExportVnConversationDetails invokes the cloudcallcenter.ExportVnConversationDetails API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/exportvnconversationdetails.html
func (client *Client) ExportVnConversationDetails(request *ExportVnConversationDetailsRequest) (response *ExportVnConversationDetailsResponse, err error) {
	response = CreateExportVnConversationDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ExportVnConversationDetailsWithChan invokes the cloudcallcenter.ExportVnConversationDetails API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/exportvnconversationdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportVnConversationDetailsWithChan(request *ExportVnConversationDetailsRequest) (<-chan *ExportVnConversationDetailsResponse, <-chan error) {
	responseChan := make(chan *ExportVnConversationDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportVnConversationDetails(request)
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

// ExportVnConversationDetailsWithCallback invokes the cloudcallcenter.ExportVnConversationDetails API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/exportvnconversationdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportVnConversationDetailsWithCallback(request *ExportVnConversationDetailsRequest, callback func(response *ExportVnConversationDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportVnConversationDetailsResponse
		var err error
		defer close(result)
		response, err = client.ExportVnConversationDetails(request)
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

// ExportVnConversationDetailsRequest is the request struct for api ExportVnConversationDetails
type ExportVnConversationDetailsRequest struct {
	*requests.RpcRequest
	BeginTimeLeftRange  requests.Integer `position:"Query" name:"BeginTimeLeftRange"`
	CallingNumber       string           `position:"Query" name:"CallingNumber"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	BeginTimeRightRange requests.Integer `position:"Query" name:"BeginTimeRightRange"`
}

// ExportVnConversationDetailsResponse is the response struct for api ExportVnConversationDetails
type ExportVnConversationDetailsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ExportTaskId string `json:"ExportTaskId" xml:"ExportTaskId"`
}

// CreateExportVnConversationDetailsRequest creates a request to invoke ExportVnConversationDetails API
func CreateExportVnConversationDetailsRequest() (request *ExportVnConversationDetailsRequest) {
	request = &ExportVnConversationDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ExportVnConversationDetails", "", "")
	request.Method = requests.GET
	return
}

// CreateExportVnConversationDetailsResponse creates a response to parse from ExportVnConversationDetails response
func CreateExportVnConversationDetailsResponse() (response *ExportVnConversationDetailsResponse) {
	response = &ExportVnConversationDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
