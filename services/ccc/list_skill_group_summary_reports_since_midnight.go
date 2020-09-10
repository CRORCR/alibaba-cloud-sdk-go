package ccc

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

// ListSkillGroupSummaryReportsSinceMidnight invokes the ccc.ListSkillGroupSummaryReportsSinceMidnight API synchronously
// api document: https://help.aliyun.com/api/ccc/listskillgroupsummaryreportssincemidnight.html
func (client *Client) ListSkillGroupSummaryReportsSinceMidnight(request *ListSkillGroupSummaryReportsSinceMidnightRequest) (response *ListSkillGroupSummaryReportsSinceMidnightResponse, err error) {
	response = CreateListSkillGroupSummaryReportsSinceMidnightResponse()
	err = client.DoAction(request, response)
	return
}

// ListSkillGroupSummaryReportsSinceMidnightWithChan invokes the ccc.ListSkillGroupSummaryReportsSinceMidnight API asynchronously
// api document: https://help.aliyun.com/api/ccc/listskillgroupsummaryreportssincemidnight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupSummaryReportsSinceMidnightWithChan(request *ListSkillGroupSummaryReportsSinceMidnightRequest) (<-chan *ListSkillGroupSummaryReportsSinceMidnightResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupSummaryReportsSinceMidnightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroupSummaryReportsSinceMidnight(request)
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

// ListSkillGroupSummaryReportsSinceMidnightWithCallback invokes the ccc.ListSkillGroupSummaryReportsSinceMidnight API asynchronously
// api document: https://help.aliyun.com/api/ccc/listskillgroupsummaryreportssincemidnight.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupSummaryReportsSinceMidnightWithCallback(request *ListSkillGroupSummaryReportsSinceMidnightRequest, callback func(response *ListSkillGroupSummaryReportsSinceMidnightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupSummaryReportsSinceMidnightResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroupSummaryReportsSinceMidnight(request)
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

// ListSkillGroupSummaryReportsSinceMidnightRequest is the request struct for api ListSkillGroupSummaryReportsSinceMidnight
type ListSkillGroupSummaryReportsSinceMidnightRequest struct {
	*requests.RpcRequest
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	SkillGroups string           `position:"Query" name:"SkillGroups"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// ListSkillGroupSummaryReportsSinceMidnightResponse is the response struct for api ListSkillGroupSummaryReportsSinceMidnight
type ListSkillGroupSummaryReportsSinceMidnightResponse struct {
	*responses.BaseResponse
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	Success                      bool                         `json:"Success" xml:"Success"`
	Code                         string                       `json:"Code" xml:"Code"`
	Message                      string                       `json:"Message" xml:"Message"`
	HttpStatusCode               int                          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PagedSkillGroupSummaryReport PagedSkillGroupSummaryReport `json:"PagedSkillGroupSummaryReport" xml:"PagedSkillGroupSummaryReport"`
}

// CreateListSkillGroupSummaryReportsSinceMidnightRequest creates a request to invoke ListSkillGroupSummaryReportsSinceMidnight API
func CreateListSkillGroupSummaryReportsSinceMidnightRequest() (request *ListSkillGroupSummaryReportsSinceMidnightRequest) {
	request = &ListSkillGroupSummaryReportsSinceMidnightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroupSummaryReportsSinceMidnight", "", "")
	return
}

// CreateListSkillGroupSummaryReportsSinceMidnightResponse creates a response to parse from ListSkillGroupSummaryReportsSinceMidnight response
func CreateListSkillGroupSummaryReportsSinceMidnightResponse() (response *ListSkillGroupSummaryReportsSinceMidnightResponse) {
	response = &ListSkillGroupSummaryReportsSinceMidnightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
