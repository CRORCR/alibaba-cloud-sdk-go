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

// DescribeRiskItemType invokes the sas.DescribeRiskItemType API synchronously
// api document: https://help.aliyun.com/api/sas/describeriskitemtype.html
func (client *Client) DescribeRiskItemType(request *DescribeRiskItemTypeRequest) (response *DescribeRiskItemTypeResponse, err error) {
	response = CreateDescribeRiskItemTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskItemTypeWithChan invokes the sas.DescribeRiskItemType API asynchronously
// api document: https://help.aliyun.com/api/sas/describeriskitemtype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskItemTypeWithChan(request *DescribeRiskItemTypeRequest) (<-chan *DescribeRiskItemTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskItemTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskItemType(request)
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

// DescribeRiskItemTypeWithCallback invokes the sas.DescribeRiskItemType API asynchronously
// api document: https://help.aliyun.com/api/sas/describeriskitemtype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRiskItemTypeWithCallback(request *DescribeRiskItemTypeRequest, callback func(response *DescribeRiskItemTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskItemTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskItemType(request)
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

// DescribeRiskItemTypeRequest is the request struct for api DescribeRiskItemType
type DescribeRiskItemTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
}

// DescribeRiskItemTypeResponse is the response struct for api DescribeRiskItemType
type DescribeRiskItemTypeResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	List      []ItemType `json:"List" xml:"List"`
}

// CreateDescribeRiskItemTypeRequest creates a request to invoke DescribeRiskItemType API
func CreateDescribeRiskItemTypeRequest() (request *DescribeRiskItemTypeRequest) {
	request = &DescribeRiskItemTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeRiskItemType", "sas", "openAPI")
	return
}

// CreateDescribeRiskItemTypeResponse creates a response to parse from DescribeRiskItemType response
func CreateDescribeRiskItemTypeResponse() (response *DescribeRiskItemTypeResponse) {
	response = &DescribeRiskItemTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
