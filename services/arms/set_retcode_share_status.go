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

// SetRetcodeShareStatus invokes the arms.SetRetcodeShareStatus API synchronously
// api document: https://help.aliyun.com/api/arms/setretcodesharestatus.html
func (client *Client) SetRetcodeShareStatus(request *SetRetcodeShareStatusRequest) (response *SetRetcodeShareStatusResponse, err error) {
	response = CreateSetRetcodeShareStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetRetcodeShareStatusWithChan invokes the arms.SetRetcodeShareStatus API asynchronously
// api document: https://help.aliyun.com/api/arms/setretcodesharestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetRetcodeShareStatusWithChan(request *SetRetcodeShareStatusRequest) (<-chan *SetRetcodeShareStatusResponse, <-chan error) {
	responseChan := make(chan *SetRetcodeShareStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetRetcodeShareStatus(request)
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

// SetRetcodeShareStatusWithCallback invokes the arms.SetRetcodeShareStatus API asynchronously
// api document: https://help.aliyun.com/api/arms/setretcodesharestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetRetcodeShareStatusWithCallback(request *SetRetcodeShareStatusRequest, callback func(response *SetRetcodeShareStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetRetcodeShareStatusResponse
		var err error
		defer close(result)
		response, err = client.SetRetcodeShareStatus(request)
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

// SetRetcodeShareStatusRequest is the request struct for api SetRetcodeShareStatus
type SetRetcodeShareStatusRequest struct {
	*requests.RpcRequest
	Pid    string           `position:"Query" name:"Pid"`
	Status requests.Boolean `position:"Query" name:"Status"`
}

// SetRetcodeShareStatusResponse is the response struct for api SetRetcodeShareStatus
type SetRetcodeShareStatusResponse struct {
	*responses.BaseResponse
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	SetRetcodeShareStatusIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateSetRetcodeShareStatusRequest creates a request to invoke SetRetcodeShareStatus API
func CreateSetRetcodeShareStatusRequest() (request *SetRetcodeShareStatusRequest) {
	request = &SetRetcodeShareStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SetRetcodeShareStatus", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetRetcodeShareStatusResponse creates a response to parse from SetRetcodeShareStatus response
func CreateSetRetcodeShareStatusResponse() (response *SetRetcodeShareStatusResponse) {
	response = &SetRetcodeShareStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
