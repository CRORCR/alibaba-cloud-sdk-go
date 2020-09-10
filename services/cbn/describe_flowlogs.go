package cbn

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

// DescribeFlowlogs invokes the cbn.DescribeFlowlogs API synchronously
// api document: https://help.aliyun.com/api/cbn/describeflowlogs.html
func (client *Client) DescribeFlowlogs(request *DescribeFlowlogsRequest) (response *DescribeFlowlogsResponse, err error) {
	response = CreateDescribeFlowlogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowlogsWithChan invokes the cbn.DescribeFlowlogs API asynchronously
// api document: https://help.aliyun.com/api/cbn/describeflowlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowlogsWithChan(request *DescribeFlowlogsRequest) (<-chan *DescribeFlowlogsResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowlogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowlogs(request)
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

// DescribeFlowlogsWithCallback invokes the cbn.DescribeFlowlogs API asynchronously
// api document: https://help.aliyun.com/api/cbn/describeflowlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowlogsWithCallback(request *DescribeFlowlogsRequest, callback func(response *DescribeFlowlogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowlogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowlogs(request)
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

// DescribeFlowlogsRequest is the request struct for api DescribeFlowlogs
type DescribeFlowlogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	CenId                string           `position:"Query" name:"CenId"`
	Description          string           `position:"Query" name:"Description"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ProjectName          string           `position:"Query" name:"ProjectName"`
	LogStoreName         string           `position:"Query" name:"LogStoreName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
	FlowLogName          string           `position:"Query" name:"FlowLogName"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeFlowlogsResponse is the response struct for api DescribeFlowlogs
type DescribeFlowlogsResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Success    string   `json:"Success" xml:"Success"`
	TotalCount string   `json:"TotalCount" xml:"TotalCount"`
	PageNumber string   `json:"PageNumber" xml:"PageNumber"`
	PageSize   string   `json:"PageSize" xml:"PageSize"`
	FlowLogs   FlowLogs `json:"FlowLogs" xml:"FlowLogs"`
}

// CreateDescribeFlowlogsRequest creates a request to invoke DescribeFlowlogs API
func CreateDescribeFlowlogsRequest() (request *DescribeFlowlogsRequest) {
	request = &DescribeFlowlogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeFlowlogs", "Cbn", "openAPI")
	return
}

// CreateDescribeFlowlogsResponse creates a response to parse from DescribeFlowlogs response
func CreateDescribeFlowlogsResponse() (response *DescribeFlowlogsResponse) {
	response = &DescribeFlowlogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
