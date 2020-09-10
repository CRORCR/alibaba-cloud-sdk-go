package r_kvstore

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

// DescribeCacheAnalysisReportList invokes the r_kvstore.DescribeCacheAnalysisReportList API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecacheanalysisreportlist.html
func (client *Client) DescribeCacheAnalysisReportList(request *DescribeCacheAnalysisReportListRequest) (response *DescribeCacheAnalysisReportListResponse, err error) {
	response = CreateDescribeCacheAnalysisReportListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCacheAnalysisReportListWithChan invokes the r_kvstore.DescribeCacheAnalysisReportList API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecacheanalysisreportlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCacheAnalysisReportListWithChan(request *DescribeCacheAnalysisReportListRequest) (<-chan *DescribeCacheAnalysisReportListResponse, <-chan error) {
	responseChan := make(chan *DescribeCacheAnalysisReportListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCacheAnalysisReportList(request)
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

// DescribeCacheAnalysisReportListWithCallback invokes the r_kvstore.DescribeCacheAnalysisReportList API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecacheanalysisreportlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCacheAnalysisReportListWithCallback(request *DescribeCacheAnalysisReportListRequest, callback func(response *DescribeCacheAnalysisReportListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCacheAnalysisReportListResponse
		var err error
		defer close(result)
		response, err = client.DescribeCacheAnalysisReportList(request)
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

// DescribeCacheAnalysisReportListRequest is the request struct for api DescribeCacheAnalysisReportList
type DescribeCacheAnalysisReportListRequest struct {
	*requests.RpcRequest
	Date                 string           `position:"Query" name:"Date"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumbers          requests.Integer `position:"Query" name:"PageNumbers"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Days                 requests.Integer `position:"Query" name:"Days"`
}

// DescribeCacheAnalysisReportListResponse is the response struct for api DescribeCacheAnalysisReportList
type DescribeCacheAnalysisReportListResponse struct {
	*responses.BaseResponse
	RequestId        string     `json:"RequestId" xml:"RequestId"`
	InstanceId       string     `json:"InstanceId" xml:"InstanceId"`
	TotalRecordCount int        `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumbers      int        `json:"PageNumbers" xml:"PageNumbers"`
	PageRecordCount  int        `json:"PageRecordCount" xml:"PageRecordCount"`
	DailyTasks       DailyTasks `json:"DailyTasks" xml:"DailyTasks"`
}

// CreateDescribeCacheAnalysisReportListRequest creates a request to invoke DescribeCacheAnalysisReportList API
func CreateDescribeCacheAnalysisReportListRequest() (request *DescribeCacheAnalysisReportListRequest) {
	request = &DescribeCacheAnalysisReportListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeCacheAnalysisReportList", "redisa", "openAPI")
	return
}

// CreateDescribeCacheAnalysisReportListResponse creates a response to parse from DescribeCacheAnalysisReportList response
func CreateDescribeCacheAnalysisReportListResponse() (response *DescribeCacheAnalysisReportListResponse) {
	response = &DescribeCacheAnalysisReportListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
