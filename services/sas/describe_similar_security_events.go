package sas

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

// DescribeSimilarSecurityEvents invokes the sas.DescribeSimilarSecurityEvents API synchronously
// api document: https://help.aliyun.com/api/sas/describesimilarsecurityevents.html
func (client *Client) DescribeSimilarSecurityEvents(request *DescribeSimilarSecurityEventsRequest) (response *DescribeSimilarSecurityEventsResponse, err error) {
	response = CreateDescribeSimilarSecurityEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSimilarSecurityEventsWithChan invokes the sas.DescribeSimilarSecurityEvents API asynchronously
// api document: https://help.aliyun.com/api/sas/describesimilarsecurityevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSimilarSecurityEventsWithChan(request *DescribeSimilarSecurityEventsRequest) (<-chan *DescribeSimilarSecurityEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeSimilarSecurityEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSimilarSecurityEvents(request)
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

// DescribeSimilarSecurityEventsWithCallback invokes the sas.DescribeSimilarSecurityEvents API asynchronously
// api document: https://help.aliyun.com/api/sas/describesimilarsecurityevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSimilarSecurityEventsWithCallback(request *DescribeSimilarSecurityEventsRequest, callback func(response *DescribeSimilarSecurityEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSimilarSecurityEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSimilarSecurityEvents(request)
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

// DescribeSimilarSecurityEventsRequest is the request struct for api DescribeSimilarSecurityEvents
type DescribeSimilarSecurityEventsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	TaskId          requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeSimilarSecurityEventsResponse is the response struct for api DescribeSimilarSecurityEvents
type DescribeSimilarSecurityEventsResponse struct {
	*responses.BaseResponse
	RequestId              string                `json:"RequestId" xml:"RequestId"`
	PageInfo               PageInfo              `json:"PageInfo" xml:"PageInfo"`
	SecurityEventsResponse []SimpleSecurityEvent `json:"SecurityEventsResponse" xml:"SecurityEventsResponse"`
}

// CreateDescribeSimilarSecurityEventsRequest creates a request to invoke DescribeSimilarSecurityEvents API
func CreateDescribeSimilarSecurityEventsRequest() (request *DescribeSimilarSecurityEventsRequest) {
	request = &DescribeSimilarSecurityEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeSimilarSecurityEvents", "sas", "openAPI")
	return
}

// CreateDescribeSimilarSecurityEventsResponse creates a response to parse from DescribeSimilarSecurityEvents response
func CreateDescribeSimilarSecurityEventsResponse() (response *DescribeSimilarSecurityEventsResponse) {
	response = &DescribeSimilarSecurityEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
