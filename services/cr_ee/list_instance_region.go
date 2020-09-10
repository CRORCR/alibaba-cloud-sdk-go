package cr_ee

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

// ListInstanceRegion invokes the cr.ListInstanceRegion API synchronously
// api document: https://help.aliyun.com/api/cr/listinstanceregion.html
func (client *Client) ListInstanceRegion(request *ListInstanceRegionRequest) (response *ListInstanceRegionResponse, err error) {
	response = CreateListInstanceRegionResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstanceRegionWithChan invokes the cr.ListInstanceRegion API asynchronously
// api document: https://help.aliyun.com/api/cr/listinstanceregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstanceRegionWithChan(request *ListInstanceRegionRequest) (<-chan *ListInstanceRegionResponse, <-chan error) {
	responseChan := make(chan *ListInstanceRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstanceRegion(request)
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

// ListInstanceRegionWithCallback invokes the cr.ListInstanceRegion API asynchronously
// api document: https://help.aliyun.com/api/cr/listinstanceregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstanceRegionWithCallback(request *ListInstanceRegionRequest, callback func(response *ListInstanceRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstanceRegionResponse
		var err error
		defer close(result)
		response, err = client.ListInstanceRegion(request)
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

// ListInstanceRegionRequest is the request struct for api ListInstanceRegion
type ListInstanceRegionRequest struct {
	*requests.RpcRequest
	Lang string `position:"Query" name:"Lang"`
}

// ListInstanceRegionResponse is the response struct for api ListInstanceRegion
type ListInstanceRegionResponse struct {
	*responses.BaseResponse
	ListInstanceRegionIsSuccess bool          `json:"IsSuccess" xml:"IsSuccess"`
	Code                        string        `json:"Code" xml:"Code"`
	RequestId                   string        `json:"RequestId" xml:"RequestId"`
	Regions                     []RegionsItem `json:"Regions" xml:"Regions"`
}

// CreateListInstanceRegionRequest creates a request to invoke ListInstanceRegion API
func CreateListInstanceRegionRequest() (request *ListInstanceRegionRequest) {
	request = &ListInstanceRegionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListInstanceRegion", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstanceRegionResponse creates a response to parse from ListInstanceRegion response
func CreateListInstanceRegionResponse() (response *ListInstanceRegionResponse) {
	response = &ListInstanceRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
