package das

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

// CreateDiagnosticReport invokes the das.CreateDiagnosticReport API synchronously
// api document: https://help.aliyun.com/api/das/creatediagnosticreport.html
func (client *Client) CreateDiagnosticReport(request *CreateDiagnosticReportRequest) (response *CreateDiagnosticReportResponse, err error) {
	response = CreateCreateDiagnosticReportResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDiagnosticReportWithChan invokes the das.CreateDiagnosticReport API asynchronously
// api document: https://help.aliyun.com/api/das/creatediagnosticreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDiagnosticReportWithChan(request *CreateDiagnosticReportRequest) (<-chan *CreateDiagnosticReportResponse, <-chan error) {
	responseChan := make(chan *CreateDiagnosticReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDiagnosticReport(request)
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

// CreateDiagnosticReportWithCallback invokes the das.CreateDiagnosticReport API asynchronously
// api document: https://help.aliyun.com/api/das/creatediagnosticreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDiagnosticReportWithCallback(request *CreateDiagnosticReportRequest, callback func(response *CreateDiagnosticReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiagnosticReportResponse
		var err error
		defer close(result)
		response, err = client.CreateDiagnosticReport(request)
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

// CreateDiagnosticReportRequest is the request struct for api CreateDiagnosticReport
type CreateDiagnosticReportRequest struct {
	*requests.RpcRequest
	SkipAuth     string `position:"Query" name:"skipAuth"`
	Context      string `position:"Query" name:"__context"`
	Signature    string `position:"Query" name:"signature"`
	EndTime      string `position:"Query" name:"EndTime"`
	StartTime    string `position:"Query" name:"StartTime"`
	UserId       string `position:"Query" name:"UserId"`
	Uid          string `position:"Query" name:"Uid"`
	AccessKey    string `position:"Query" name:"accessKey"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	Timestamp    string `position:"Query" name:"timestamp"`
}

// CreateDiagnosticReportResponse is the response struct for api CreateDiagnosticReport
type CreateDiagnosticReportResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Synchro   string `json:"Synchro" xml:"Synchro"`
}

// CreateCreateDiagnosticReportRequest creates a request to invoke CreateDiagnosticReport API
func CreateCreateDiagnosticReportRequest() (request *CreateDiagnosticReportRequest) {
	request = &CreateDiagnosticReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "CreateDiagnosticReport", "hdm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDiagnosticReportResponse creates a response to parse from CreateDiagnosticReport response
func CreateCreateDiagnosticReportResponse() (response *CreateDiagnosticReportResponse) {
	response = &CreateDiagnosticReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
