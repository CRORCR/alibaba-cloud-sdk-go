package arms

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

// DescribeTraceLocation invokes the arms.DescribeTraceLocation API synchronously
// api document: https://help.aliyun.com/api/arms/describetracelocation.html
func (client *Client) DescribeTraceLocation(request *DescribeTraceLocationRequest) (response *DescribeTraceLocationResponse, err error) {
	response = CreateDescribeTraceLocationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTraceLocationWithChan invokes the arms.DescribeTraceLocation API asynchronously
// api document: https://help.aliyun.com/api/arms/describetracelocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTraceLocationWithChan(request *DescribeTraceLocationRequest) (<-chan *DescribeTraceLocationResponse, <-chan error) {
	responseChan := make(chan *DescribeTraceLocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTraceLocation(request)
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

// DescribeTraceLocationWithCallback invokes the arms.DescribeTraceLocation API asynchronously
// api document: https://help.aliyun.com/api/arms/describetracelocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTraceLocationWithCallback(request *DescribeTraceLocationRequest, callback func(response *DescribeTraceLocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTraceLocationResponse
		var err error
		defer close(result)
		response, err = client.DescribeTraceLocation(request)
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

// DescribeTraceLocationRequest is the request struct for api DescribeTraceLocation
type DescribeTraceLocationRequest struct {
	*requests.RpcRequest
}

// DescribeTraceLocationResponse is the response struct for api DescribeTraceLocation
type DescribeTraceLocationResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	RegionConfigs []RegionConfig `json:"RegionConfigs" xml:"RegionConfigs"`
}

// CreateDescribeTraceLocationRequest creates a request to invoke DescribeTraceLocation API
func CreateDescribeTraceLocationRequest() (request *DescribeTraceLocationRequest) {
	request = &DescribeTraceLocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DescribeTraceLocation", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTraceLocationResponse creates a response to parse from DescribeTraceLocation response
func CreateDescribeTraceLocationResponse() (response *DescribeTraceLocationResponse) {
	response = &DescribeTraceLocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
