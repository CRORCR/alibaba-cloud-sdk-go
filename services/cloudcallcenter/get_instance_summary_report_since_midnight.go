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

// GetInstanceSummaryReportSinceMidnight invokes the cloudcallcenter.GetInstanceSummaryReportSinceMidnight API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancesummaryreportsincemidnight.html
func (client *Client) GetInstanceSummaryReportSinceMidnight(request *GetInstanceSummaryReportSinceMidnightRequest) (response *GetInstanceSummaryReportSinceMidnightResponse, err error) {
	response = CreateGetInstanceSummaryReportSinceMidnightResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceSummaryReportSinceMidnightWithChan invokes the cloudcallcenter.GetInstanceSummaryReportSinceMidnight API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancesummaryreportsincemidnight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceSummaryReportSinceMidnightWithChan(request *GetInstanceSummaryReportSinceMidnightRequest) (<-chan *GetInstanceSummaryReportSinceMidnightResponse, <-chan error) {
	responseChan := make(chan *GetInstanceSummaryReportSinceMidnightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceSummaryReportSinceMidnight(request)
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

// GetInstanceSummaryReportSinceMidnightWithCallback invokes the cloudcallcenter.GetInstanceSummaryReportSinceMidnight API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancesummaryreportsincemidnight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceSummaryReportSinceMidnightWithCallback(request *GetInstanceSummaryReportSinceMidnightRequest, callback func(response *GetInstanceSummaryReportSinceMidnightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceSummaryReportSinceMidnightResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceSummaryReportSinceMidnight(request)
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

// GetInstanceSummaryReportSinceMidnightRequest is the request struct for api GetInstanceSummaryReportSinceMidnight
type GetInstanceSummaryReportSinceMidnightRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetInstanceSummaryReportSinceMidnightResponse is the response struct for api GetInstanceSummaryReportSinceMidnight
type GetInstanceSummaryReportSinceMidnightResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Success               bool                  `json:"Success" xml:"Success"`
	Code                  string                `json:"Code" xml:"Code"`
	Message               string                `json:"Message" xml:"Message"`
	HttpStatusCode        int                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	InstanceSummaryReport InstanceSummaryReport `json:"InstanceSummaryReport" xml:"InstanceSummaryReport"`
}

// CreateGetInstanceSummaryReportSinceMidnightRequest creates a request to invoke GetInstanceSummaryReportSinceMidnight API
func CreateGetInstanceSummaryReportSinceMidnightRequest() (request *GetInstanceSummaryReportSinceMidnightRequest) {
	request = &GetInstanceSummaryReportSinceMidnightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetInstanceSummaryReportSinceMidnight", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceSummaryReportSinceMidnightResponse creates a response to parse from GetInstanceSummaryReportSinceMidnight response
func CreateGetInstanceSummaryReportSinceMidnightResponse() (response *GetInstanceSummaryReportSinceMidnightResponse) {
	response = &GetInstanceSummaryReportSinceMidnightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
