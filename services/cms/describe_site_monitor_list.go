package cms

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

// DescribeSiteMonitorList invokes the cms.DescribeSiteMonitorList API synchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorlist.html
func (client *Client) DescribeSiteMonitorList(request *DescribeSiteMonitorListRequest) (response *DescribeSiteMonitorListResponse, err error) {
	response = CreateDescribeSiteMonitorListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSiteMonitorListWithChan invokes the cms.DescribeSiteMonitorList API asynchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSiteMonitorListWithChan(request *DescribeSiteMonitorListRequest) (<-chan *DescribeSiteMonitorListResponse, <-chan error) {
	responseChan := make(chan *DescribeSiteMonitorListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSiteMonitorList(request)
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

// DescribeSiteMonitorListWithCallback invokes the cms.DescribeSiteMonitorList API asynchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSiteMonitorListWithCallback(request *DescribeSiteMonitorListRequest, callback func(response *DescribeSiteMonitorListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSiteMonitorListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSiteMonitorList(request)
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

// DescribeSiteMonitorListRequest is the request struct for api DescribeSiteMonitorList
type DescribeSiteMonitorListRequest struct {
	*requests.RpcRequest
	TaskType string           `position:"Query" name:"TaskType"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	Page     requests.Integer `position:"Query" name:"Page"`
	Keyword  string           `position:"Query" name:"Keyword"`
	TaskId   string           `position:"Query" name:"TaskId"`
}

// DescribeSiteMonitorListResponse is the response struct for api DescribeSiteMonitorList
type DescribeSiteMonitorListResponse struct {
	*responses.BaseResponse
	Code         string                                `json:"Code" xml:"Code"`
	Message      string                                `json:"Message" xml:"Message"`
	Success      string                                `json:"Success" xml:"Success"`
	RequestId    string                                `json:"RequestId" xml:"RequestId"`
	PageNumber   int                                   `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                                   `json:"PageSize" xml:"PageSize"`
	TotalCount   int                                   `json:"TotalCount" xml:"TotalCount"`
	SiteMonitors SiteMonitorsInDescribeSiteMonitorList `json:"SiteMonitors" xml:"SiteMonitors"`
}

// CreateDescribeSiteMonitorListRequest creates a request to invoke DescribeSiteMonitorList API
func CreateDescribeSiteMonitorListRequest() (request *DescribeSiteMonitorListRequest) {
	request = &DescribeSiteMonitorListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSiteMonitorList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSiteMonitorListResponse creates a response to parse from DescribeSiteMonitorList response
func CreateDescribeSiteMonitorListResponse() (response *DescribeSiteMonitorListResponse) {
	response = &DescribeSiteMonitorListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
