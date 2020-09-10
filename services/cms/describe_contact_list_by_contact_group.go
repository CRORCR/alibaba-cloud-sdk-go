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

// DescribeContactListByContactGroup invokes the cms.DescribeContactListByContactGroup API synchronously
// api document: https://help.aliyun.com/api/cms/describecontactlistbycontactgroup.html
func (client *Client) DescribeContactListByContactGroup(request *DescribeContactListByContactGroupRequest) (response *DescribeContactListByContactGroupResponse, err error) {
	response = CreateDescribeContactListByContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContactListByContactGroupWithChan invokes the cms.DescribeContactListByContactGroup API asynchronously
// api document: https://help.aliyun.com/api/cms/describecontactlistbycontactgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeContactListByContactGroupWithChan(request *DescribeContactListByContactGroupRequest) (<-chan *DescribeContactListByContactGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeContactListByContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContactListByContactGroup(request)
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

// DescribeContactListByContactGroupWithCallback invokes the cms.DescribeContactListByContactGroup API asynchronously
// api document: https://help.aliyun.com/api/cms/describecontactlistbycontactgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeContactListByContactGroupWithCallback(request *DescribeContactListByContactGroupRequest, callback func(response *DescribeContactListByContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContactListByContactGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeContactListByContactGroup(request)
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

// DescribeContactListByContactGroupRequest is the request struct for api DescribeContactListByContactGroup
type DescribeContactListByContactGroupRequest struct {
	*requests.RpcRequest
	ContactGroupName string `position:"Query" name:"ContactGroupName"`
}

// DescribeContactListByContactGroupResponse is the response struct for api DescribeContactListByContactGroup
type DescribeContactListByContactGroupResponse struct {
	*responses.BaseResponse
	Success   bool                                        `json:"Success" xml:"Success"`
	Code      string                                      `json:"Code" xml:"Code"`
	Message   string                                      `json:"Message" xml:"Message"`
	RequestId string                                      `json:"RequestId" xml:"RequestId"`
	Contacts  ContactsInDescribeContactListByContactGroup `json:"Contacts" xml:"Contacts"`
}

// CreateDescribeContactListByContactGroupRequest creates a request to invoke DescribeContactListByContactGroup API
func CreateDescribeContactListByContactGroupRequest() (request *DescribeContactListByContactGroupRequest) {
	request = &DescribeContactListByContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeContactListByContactGroup", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeContactListByContactGroupResponse creates a response to parse from DescribeContactListByContactGroup response
func CreateDescribeContactListByContactGroupResponse() (response *DescribeContactListByContactGroupResponse) {
	response = &DescribeContactListByContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
