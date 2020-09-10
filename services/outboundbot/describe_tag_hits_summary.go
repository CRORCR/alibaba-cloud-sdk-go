package outboundbot

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

// DescribeTagHitsSummary invokes the outboundbot.DescribeTagHitsSummary API synchronously
// api document: https://help.aliyun.com/api/outboundbot/describetaghitssummary.html
func (client *Client) DescribeTagHitsSummary(request *DescribeTagHitsSummaryRequest) (response *DescribeTagHitsSummaryResponse, err error) {
	response = CreateDescribeTagHitsSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTagHitsSummaryWithChan invokes the outboundbot.DescribeTagHitsSummary API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/describetaghitssummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTagHitsSummaryWithChan(request *DescribeTagHitsSummaryRequest) (<-chan *DescribeTagHitsSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeTagHitsSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTagHitsSummary(request)
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

// DescribeTagHitsSummaryWithCallback invokes the outboundbot.DescribeTagHitsSummary API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/describetaghitssummary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTagHitsSummaryWithCallback(request *DescribeTagHitsSummaryRequest, callback func(response *DescribeTagHitsSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTagHitsSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeTagHitsSummary(request)
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

// DescribeTagHitsSummaryRequest is the request struct for api DescribeTagHitsSummary
type DescribeTagHitsSummaryRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
}

// DescribeTagHitsSummaryResponse is the response struct for api DescribeTagHitsSummary
type DescribeTagHitsSummaryResponse struct {
	*responses.BaseResponse
	Code           string     `json:"Code" xml:"Code"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	TagGroups      []TagGroup `json:"TagGroups" xml:"TagGroups"`
	TagHitsList    []TagHits  `json:"TagHitsList" xml:"TagHitsList"`
}

// CreateDescribeTagHitsSummaryRequest creates a request to invoke DescribeTagHitsSummary API
func CreateDescribeTagHitsSummaryRequest() (request *DescribeTagHitsSummaryRequest) {
	request = &DescribeTagHitsSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeTagHitsSummary", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTagHitsSummaryResponse creates a response to parse from DescribeTagHitsSummary response
func CreateDescribeTagHitsSummaryResponse() (response *DescribeTagHitsSummaryResponse) {
	response = &DescribeTagHitsSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
