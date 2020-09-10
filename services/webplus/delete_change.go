package webplus

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

// DeleteChange invokes the webplus.DeleteChange API synchronously
// api document: https://help.aliyun.com/api/webplus/deletechange.html
func (client *Client) DeleteChange(request *DeleteChangeRequest) (response *DeleteChangeResponse, err error) {
	response = CreateDeleteChangeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteChangeWithChan invokes the webplus.DeleteChange API asynchronously
// api document: https://help.aliyun.com/api/webplus/deletechange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChangeWithChan(request *DeleteChangeRequest) (<-chan *DeleteChangeResponse, <-chan error) {
	responseChan := make(chan *DeleteChangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteChange(request)
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

// DeleteChangeWithCallback invokes the webplus.DeleteChange API asynchronously
// api document: https://help.aliyun.com/api/webplus/deletechange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChangeWithCallback(request *DeleteChangeRequest, callback func(response *DeleteChangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteChangeResponse
		var err error
		defer close(result)
		response, err = client.DeleteChange(request)
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

// DeleteChangeRequest is the request struct for api DeleteChange
type DeleteChangeRequest struct {
	*requests.RoaRequest
	ChangeId string `position:"Query" name:"ChangeId"`
}

// DeleteChangeResponse is the response struct for api DeleteChange
type DeleteChangeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteChangeRequest creates a request to invoke DeleteChange API
func CreateDeleteChangeRequest() (request *DeleteChangeRequest) {
	request = &DeleteChangeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DeleteChange", "/pop/v1/wam/change", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteChangeResponse creates a response to parse from DeleteChange response
func CreateDeleteChangeResponse() (response *DeleteChangeResponse) {
	response = &DeleteChangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
