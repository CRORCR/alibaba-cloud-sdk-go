package green

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

// ExportOssResult invokes the green.ExportOssResult API synchronously
// api document: https://help.aliyun.com/api/green/exportossresult.html
func (client *Client) ExportOssResult(request *ExportOssResultRequest) (response *ExportOssResultResponse, err error) {
	response = CreateExportOssResultResponse()
	err = client.DoAction(request, response)
	return
}

// ExportOssResultWithChan invokes the green.ExportOssResult API asynchronously
// api document: https://help.aliyun.com/api/green/exportossresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportOssResultWithChan(request *ExportOssResultRequest) (<-chan *ExportOssResultResponse, <-chan error) {
	responseChan := make(chan *ExportOssResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportOssResult(request)
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

// ExportOssResultWithCallback invokes the green.ExportOssResult API asynchronously
// api document: https://help.aliyun.com/api/green/exportossresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportOssResultWithCallback(request *ExportOssResultRequest, callback func(response *ExportOssResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportOssResultResponse
		var err error
		defer close(result)
		response, err = client.ExportOssResult(request)
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

// ExportOssResultRequest is the request struct for api ExportOssResult
type ExportOssResultRequest struct {
	*requests.RpcRequest
	MinScore     requests.Float   `position:"Query" name:"MinScore"`
	MaxScore     requests.Float   `position:"Query" name:"MaxScore"`
	StartDate    string           `position:"Query" name:"StartDate"`
	Scene        string           `position:"Query" name:"Scene"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Stock        requests.Boolean `position:"Query" name:"Stock"`
	TotalCount   requests.Integer `position:"Query" name:"TotalCount"`
	Suggestion   string           `position:"Query" name:"Suggestion"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	Bucket       string           `position:"Query" name:"Bucket"`
	EndDate      string           `position:"Query" name:"EndDate"`
}

// ExportOssResultResponse is the response struct for api ExportOssResult
type ExportOssResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	FileUrl   string `json:"FileUrl" xml:"FileUrl"`
}

// CreateExportOssResultRequest creates a request to invoke ExportOssResult API
func CreateExportOssResultRequest() (request *ExportOssResultRequest) {
	request = &ExportOssResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "ExportOssResult", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportOssResultResponse creates a response to parse from ExportOssResult response
func CreateExportOssResultResponse() (response *ExportOssResultResponse) {
	response = &ExportOssResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
