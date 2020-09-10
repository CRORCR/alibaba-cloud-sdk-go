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

// ListCallMeasureSummaryReports invokes the cloudcallcenter.ListCallMeasureSummaryReports API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listcallmeasuresummaryreports.html
func (client *Client) ListCallMeasureSummaryReports(request *ListCallMeasureSummaryReportsRequest) (response *ListCallMeasureSummaryReportsResponse, err error) {
	response = CreateListCallMeasureSummaryReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCallMeasureSummaryReportsWithChan invokes the cloudcallcenter.ListCallMeasureSummaryReports API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listcallmeasuresummaryreports.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCallMeasureSummaryReportsWithChan(request *ListCallMeasureSummaryReportsRequest) (<-chan *ListCallMeasureSummaryReportsResponse, <-chan error) {
	responseChan := make(chan *ListCallMeasureSummaryReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCallMeasureSummaryReports(request)
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

// ListCallMeasureSummaryReportsWithCallback invokes the cloudcallcenter.ListCallMeasureSummaryReports API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listcallmeasuresummaryreports.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCallMeasureSummaryReportsWithCallback(request *ListCallMeasureSummaryReportsRequest, callback func(response *ListCallMeasureSummaryReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCallMeasureSummaryReportsResponse
		var err error
		defer close(result)
		response, err = client.ListCallMeasureSummaryReports(request)
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

// ListCallMeasureSummaryReportsRequest is the request struct for api ListCallMeasureSummaryReports
type ListCallMeasureSummaryReportsRequest struct {
	*requests.RpcRequest
	IntervalType string `position:"Query" name:"IntervalType"`
}

// ListCallMeasureSummaryReportsResponse is the response struct for api ListCallMeasureSummaryReports
type ListCallMeasureSummaryReportsResponse struct {
	*responses.BaseResponse
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	Success                      bool                         `json:"Success" xml:"Success"`
	Code                         string                       `json:"Code" xml:"Code"`
	Message                      string                       `json:"Message" xml:"Message"`
	HttpStatusCode               int                          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	CallMeasureSummaryReportList CallMeasureSummaryReportList `json:"CallMeasureSummaryReportList" xml:"CallMeasureSummaryReportList"`
}

// CreateListCallMeasureSummaryReportsRequest creates a request to invoke ListCallMeasureSummaryReports API
func CreateListCallMeasureSummaryReportsRequest() (request *ListCallMeasureSummaryReportsRequest) {
	request = &ListCallMeasureSummaryReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListCallMeasureSummaryReports", "", "")
	request.Method = requests.POST
	return
}

// CreateListCallMeasureSummaryReportsResponse creates a response to parse from ListCallMeasureSummaryReports response
func CreateListCallMeasureSummaryReportsResponse() (response *ListCallMeasureSummaryReportsResponse) {
	response = &ListCallMeasureSummaryReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
