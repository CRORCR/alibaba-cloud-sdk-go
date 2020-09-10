package slb

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

// DeleteAccessControlList invokes the slb.DeleteAccessControlList API synchronously
// api document: https://help.aliyun.com/api/slb/deleteaccesscontrollist.html
func (client *Client) DeleteAccessControlList(request *DeleteAccessControlListRequest) (response *DeleteAccessControlListResponse, err error) {
	response = CreateDeleteAccessControlListResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAccessControlListWithChan invokes the slb.DeleteAccessControlList API asynchronously
// api document: https://help.aliyun.com/api/slb/deleteaccesscontrollist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAccessControlListWithChan(request *DeleteAccessControlListRequest) (<-chan *DeleteAccessControlListResponse, <-chan error) {
	responseChan := make(chan *DeleteAccessControlListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAccessControlList(request)
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

// DeleteAccessControlListWithCallback invokes the slb.DeleteAccessControlList API asynchronously
// api document: https://help.aliyun.com/api/slb/deleteaccesscontrollist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAccessControlListWithCallback(request *DeleteAccessControlListRequest, callback func(response *DeleteAccessControlListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAccessControlListResponse
		var err error
		defer close(result)
		response, err = client.DeleteAccessControlList(request)
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

// DeleteAccessControlListRequest is the request struct for api DeleteAccessControlList
type DeleteAccessControlListRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// DeleteAccessControlListResponse is the response struct for api DeleteAccessControlList
type DeleteAccessControlListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteAccessControlListRequest creates a request to invoke DeleteAccessControlList API
func CreateDeleteAccessControlListRequest() (request *DeleteAccessControlListRequest) {
	request = &DeleteAccessControlListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteAccessControlList", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteAccessControlListResponse creates a response to parse from DeleteAccessControlList response
func CreateDeleteAccessControlListResponse() (response *DeleteAccessControlListResponse) {
	response = &DeleteAccessControlListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
