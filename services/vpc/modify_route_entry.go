package vpc

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

// ModifyRouteEntry invokes the vpc.ModifyRouteEntry API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyrouteentry.html
func (client *Client) ModifyRouteEntry(request *ModifyRouteEntryRequest) (response *ModifyRouteEntryResponse, err error) {
	response = CreateModifyRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRouteEntryWithChan invokes the vpc.ModifyRouteEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyrouteentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRouteEntryWithChan(request *ModifyRouteEntryRequest) (<-chan *ModifyRouteEntryResponse, <-chan error) {
	responseChan := make(chan *ModifyRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRouteEntry(request)
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

// ModifyRouteEntryWithCallback invokes the vpc.ModifyRouteEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyrouteentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRouteEntryWithCallback(request *ModifyRouteEntryRequest, callback func(response *ModifyRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.ModifyRouteEntry(request)
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

// ModifyRouteEntryRequest is the request struct for api ModifyRouteEntry
type ModifyRouteEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteEntryName       string           `position:"Query" name:"RouteEntryName"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouteEntryId         string           `position:"Query" name:"RouteEntryId"`
}

// ModifyRouteEntryResponse is the response struct for api ModifyRouteEntry
type ModifyRouteEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRouteEntryRequest creates a request to invoke ModifyRouteEntry API
func CreateModifyRouteEntryRequest() (request *ModifyRouteEntryRequest) {
	request = &ModifyRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyRouteEntryResponse creates a response to parse from ModifyRouteEntry response
func CreateModifyRouteEntryResponse() (response *ModifyRouteEntryResponse) {
	response = &ModifyRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
