package smarthosting

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

// ListManagedPrivateSpaces invokes the smarthosting.ListManagedPrivateSpaces API synchronously
// api document: https://help.aliyun.com/api/smarthosting/listmanagedprivatespaces.html
func (client *Client) ListManagedPrivateSpaces(request *ListManagedPrivateSpacesRequest) (response *ListManagedPrivateSpacesResponse, err error) {
	response = CreateListManagedPrivateSpacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListManagedPrivateSpacesWithChan invokes the smarthosting.ListManagedPrivateSpaces API asynchronously
// api document: https://help.aliyun.com/api/smarthosting/listmanagedprivatespaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListManagedPrivateSpacesWithChan(request *ListManagedPrivateSpacesRequest) (<-chan *ListManagedPrivateSpacesResponse, <-chan error) {
	responseChan := make(chan *ListManagedPrivateSpacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListManagedPrivateSpaces(request)
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

// ListManagedPrivateSpacesWithCallback invokes the smarthosting.ListManagedPrivateSpaces API asynchronously
// api document: https://help.aliyun.com/api/smarthosting/listmanagedprivatespaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListManagedPrivateSpacesWithCallback(request *ListManagedPrivateSpacesRequest, callback func(response *ListManagedPrivateSpacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListManagedPrivateSpacesResponse
		var err error
		defer close(result)
		response, err = client.ListManagedPrivateSpaces(request)
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

// ListManagedPrivateSpacesRequest is the request struct for api ListManagedPrivateSpaces
type ListManagedPrivateSpacesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	ManagedPrivateSpaceName string                         `position:"Query" name:"ManagedPrivateSpaceName"`
	ResourceGroupId         string                         `position:"Query" name:"ResourceGroupId"`
	NextToken               string                         `position:"Query" name:"NextToken"`
	Tag                     *[]ListManagedPrivateSpacesTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount    string                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string                         `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer               `position:"Query" name:"OwnerId"`
	ManagedPrivateSpaceId   *[]string                      `position:"Query" name:"ManagedPrivateSpaceId"  type:"Repeated"`
	ZoneId                  string                         `position:"Query" name:"ZoneId"`
	MaxResults              requests.Integer               `position:"Query" name:"MaxResults"`
}

// ListManagedPrivateSpacesTag is a repeated param struct in ListManagedPrivateSpacesRequest
type ListManagedPrivateSpacesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListManagedPrivateSpacesResponse is the response struct for api ListManagedPrivateSpaces
type ListManagedPrivateSpacesResponse struct {
	*responses.BaseResponse
	NextToken               string                  `json:"NextToken" xml:"NextToken"`
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	TotalCount              int                     `json:"TotalCount" xml:"TotalCount"`
	ManagedPrivateSpaceSets ManagedPrivateSpaceSets `json:"ManagedPrivateSpaceSets" xml:"ManagedPrivateSpaceSets"`
}

// CreateListManagedPrivateSpacesRequest creates a request to invoke ListManagedPrivateSpaces API
func CreateListManagedPrivateSpacesRequest() (request *ListManagedPrivateSpacesRequest) {
	request = &ListManagedPrivateSpacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SmartHosting", "2020-08-01", "ListManagedPrivateSpaces", "smarthosting", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListManagedPrivateSpacesResponse creates a response to parse from ListManagedPrivateSpaces response
func CreateListManagedPrivateSpacesResponse() (response *ListManagedPrivateSpacesResponse) {
	response = &ListManagedPrivateSpacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
