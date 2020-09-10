package imm

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

// UpdateSet invokes the imm.UpdateSet API synchronously
// api document: https://help.aliyun.com/api/imm/updateset.html
func (client *Client) UpdateSet(request *UpdateSetRequest) (response *UpdateSetResponse, err error) {
	response = CreateUpdateSetResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSetWithChan invokes the imm.UpdateSet API asynchronously
// api document: https://help.aliyun.com/api/imm/updateset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSetWithChan(request *UpdateSetRequest) (<-chan *UpdateSetResponse, <-chan error) {
	responseChan := make(chan *UpdateSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSet(request)
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

// UpdateSetWithCallback invokes the imm.UpdateSet API asynchronously
// api document: https://help.aliyun.com/api/imm/updateset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSetWithCallback(request *UpdateSetRequest, callback func(response *UpdateSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSetResponse
		var err error
		defer close(result)
		response, err = client.UpdateSet(request)
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

// UpdateSetRequest is the request struct for api UpdateSet
type UpdateSetRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetName string `position:"Query" name:"SetName"`
	SetId   string `position:"Query" name:"SetId"`
}

// UpdateSetResponse is the response struct for api UpdateSet
type UpdateSetResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SetId      string `json:"SetId" xml:"SetId"`
	SetName    string `json:"SetName" xml:"SetName"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime string `json:"ModifyTime" xml:"ModifyTime"`
}

// CreateUpdateSetRequest creates a request to invoke UpdateSet API
func CreateUpdateSetRequest() (request *UpdateSetRequest) {
	request = &UpdateSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "UpdateSet", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateSetResponse creates a response to parse from UpdateSet response
func CreateUpdateSetResponse() (response *UpdateSetResponse) {
	response = &UpdateSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
