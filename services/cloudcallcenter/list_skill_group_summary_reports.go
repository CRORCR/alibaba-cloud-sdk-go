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

// ListSkillGroupSummaryReports invokes the cloudcallcenter.ListSkillGroupSummaryReports API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listskillgroupsummaryreports.html
func (client *Client) ListSkillGroupSummaryReports(request *ListSkillGroupSummaryReportsRequest) (response *ListSkillGroupSummaryReportsResponse, err error) {
	response = CreateListSkillGroupSummaryReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSkillGroupSummaryReportsWithChan invokes the cloudcallcenter.ListSkillGroupSummaryReports API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listskillgroupsummaryreports.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupSummaryReportsWithChan(request *ListSkillGroupSummaryReportsRequest) (<-chan *ListSkillGroupSummaryReportsResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupSummaryReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroupSummaryReports(request)
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

// ListSkillGroupSummaryReportsWithCallback invokes the cloudcallcenter.ListSkillGroupSummaryReports API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listskillgroupsummaryreports.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupSummaryReportsWithCallback(request *ListSkillGroupSummaryReportsRequest, callback func(response *ListSkillGroupSummaryReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupSummaryReportsResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroupSummaryReports(request)
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

// ListSkillGroupSummaryReportsRequest is the request struct for api ListSkillGroupSummaryReports
type ListSkillGroupSummaryReportsRequest struct {
	*requests.RpcRequest
	OrderByOutboundTotalTalkTime string           `position:"Query" name:"OrderByOutboundTotalTalkTime"`
	OrderByInboundTotalTalkTime  string           `position:"Query" name:"OrderByInboundTotalTalkTime"`
	EndTime                      string           `position:"Query" name:"EndTime"`
	StartTime                    string           `position:"Query" name:"StartTime"`
	OrderByInboundTotalCalls     string           `position:"Query" name:"OrderByInboundTotalCalls"`
	PageNumber                   requests.Integer `position:"Query" name:"PageNumber"`
	OrderByTotalTalkTime         string           `position:"Query" name:"OrderByTotalTalkTime"`
	InstanceId                   string           `position:"Query" name:"InstanceId"`
	OrderByOutboundTotalCalls    string           `position:"Query" name:"OrderByOutboundTotalCalls"`
	SkillGroupIds                string           `position:"Query" name:"SkillGroupIds"`
	PageSize                     requests.Integer `position:"Query" name:"PageSize"`
	OrderByOccupancyRate         string           `position:"Query" name:"OrderByOccupancyRate"`
	OrderByTotalCalls            string           `position:"Query" name:"OrderByTotalCalls"`
}

// ListSkillGroupSummaryReportsResponse is the response struct for api ListSkillGroupSummaryReports
type ListSkillGroupSummaryReportsResponse struct {
	*responses.BaseResponse
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	Success        bool                               `json:"Success" xml:"Success"`
	Code           string                             `json:"Code" xml:"Code"`
	Message        string                             `json:"Message" xml:"Message"`
	HttpStatusCode int                                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListSkillGroupSummaryReports `json:"Data" xml:"Data"`
}

// CreateListSkillGroupSummaryReportsRequest creates a request to invoke ListSkillGroupSummaryReports API
func CreateListSkillGroupSummaryReportsRequest() (request *ListSkillGroupSummaryReportsRequest) {
	request = &ListSkillGroupSummaryReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListSkillGroupSummaryReports", "", "")
	request.Method = requests.POST
	return
}

// CreateListSkillGroupSummaryReportsResponse creates a response to parse from ListSkillGroupSummaryReports response
func CreateListSkillGroupSummaryReportsResponse() (response *ListSkillGroupSummaryReportsResponse) {
	response = &ListSkillGroupSummaryReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
