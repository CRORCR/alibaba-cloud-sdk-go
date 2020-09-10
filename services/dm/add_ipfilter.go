package dm

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

// AddIpfilter invokes the dm.AddIpfilter API synchronously
// api document: https://help.aliyun.com/api/dm/addipfilter.html
func (client *Client) AddIpfilter(request *AddIpfilterRequest) (response *AddIpfilterResponse, err error) {
	response = CreateAddIpfilterResponse()
	err = client.DoAction(request, response)
	return
}

// AddIpfilterWithChan invokes the dm.AddIpfilter API asynchronously
// api document: https://help.aliyun.com/api/dm/addipfilter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddIpfilterWithChan(request *AddIpfilterRequest) (<-chan *AddIpfilterResponse, <-chan error) {
	responseChan := make(chan *AddIpfilterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddIpfilter(request)
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

// AddIpfilterWithCallback invokes the dm.AddIpfilter API asynchronously
// api document: https://help.aliyun.com/api/dm/addipfilter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddIpfilterWithCallback(request *AddIpfilterRequest, callback func(response *AddIpfilterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddIpfilterResponse
		var err error
		defer close(result)
		response, err = client.AddIpfilter(request)
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

// AddIpfilterRequest is the request struct for api AddIpfilter
type AddIpfilterRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IpAddress            string           `position:"Query" name:"IpAddress"`
}

// AddIpfilterResponse is the response struct for api AddIpfilter
type AddIpfilterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddIpfilterRequest creates a request to invoke AddIpfilter API
func CreateAddIpfilterRequest() (request *AddIpfilterRequest) {
	request = &AddIpfilterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "AddIpfilter", "", "")
	return
}

// CreateAddIpfilterResponse creates a response to parse from AddIpfilter response
func CreateAddIpfilterResponse() (response *AddIpfilterResponse) {
	response = &AddIpfilterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
