package ehpc

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

// InstallSoftware invokes the ehpc.InstallSoftware API synchronously
// api document: https://help.aliyun.com/api/ehpc/installsoftware.html
func (client *Client) InstallSoftware(request *InstallSoftwareRequest) (response *InstallSoftwareResponse, err error) {
	response = CreateInstallSoftwareResponse()
	err = client.DoAction(request, response)
	return
}

// InstallSoftwareWithChan invokes the ehpc.InstallSoftware API asynchronously
// api document: https://help.aliyun.com/api/ehpc/installsoftware.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InstallSoftwareWithChan(request *InstallSoftwareRequest) (<-chan *InstallSoftwareResponse, <-chan error) {
	responseChan := make(chan *InstallSoftwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallSoftware(request)
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

// InstallSoftwareWithCallback invokes the ehpc.InstallSoftware API asynchronously
// api document: https://help.aliyun.com/api/ehpc/installsoftware.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InstallSoftwareWithCallback(request *InstallSoftwareRequest, callback func(response *InstallSoftwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallSoftwareResponse
		var err error
		defer close(result)
		response, err = client.InstallSoftware(request)
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

// InstallSoftwareRequest is the request struct for api InstallSoftware
type InstallSoftwareRequest struct {
	*requests.RpcRequest
	ClusterId   string `position:"Query" name:"ClusterId"`
	Application string `position:"Query" name:"Application"`
}

// InstallSoftwareResponse is the response struct for api InstallSoftware
type InstallSoftwareResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInstallSoftwareRequest creates a request to invoke InstallSoftware API
func CreateInstallSoftwareRequest() (request *InstallSoftwareRequest) {
	request = &InstallSoftwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "InstallSoftware", "", "")
	request.Method = requests.GET
	return
}

// CreateInstallSoftwareResponse creates a response to parse from InstallSoftware response
func CreateInstallSoftwareResponse() (response *InstallSoftwareResponse) {
	response = &InstallSoftwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
