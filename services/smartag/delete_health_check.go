package smartag

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

// DeleteHealthCheck invokes the smartag.DeleteHealthCheck API synchronously
// api document: https://help.aliyun.com/api/smartag/deletehealthcheck.html
func (client *Client) DeleteHealthCheck(request *DeleteHealthCheckRequest) (response *DeleteHealthCheckResponse, err error) {
	response = CreateDeleteHealthCheckResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHealthCheckWithChan invokes the smartag.DeleteHealthCheck API asynchronously
// api document: https://help.aliyun.com/api/smartag/deletehealthcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHealthCheckWithChan(request *DeleteHealthCheckRequest) (<-chan *DeleteHealthCheckResponse, <-chan error) {
	responseChan := make(chan *DeleteHealthCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHealthCheck(request)
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

// DeleteHealthCheckWithCallback invokes the smartag.DeleteHealthCheck API asynchronously
// api document: https://help.aliyun.com/api/smartag/deletehealthcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHealthCheckWithCallback(request *DeleteHealthCheckRequest, callback func(response *DeleteHealthCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHealthCheckResponse
		var err error
		defer close(result)
		response, err = client.DeleteHealthCheck(request)
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

// DeleteHealthCheckRequest is the request struct for api DeleteHealthCheck
type DeleteHealthCheckRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HcInstanceId         string           `position:"Query" name:"HcInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteHealthCheckResponse is the response struct for api DeleteHealthCheck
type DeleteHealthCheckResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHealthCheckRequest creates a request to invoke DeleteHealthCheck API
func CreateDeleteHealthCheckRequest() (request *DeleteHealthCheckRequest) {
	request = &DeleteHealthCheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteHealthCheck", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteHealthCheckResponse creates a response to parse from DeleteHealthCheck response
func CreateDeleteHealthCheckResponse() (response *DeleteHealthCheckResponse) {
	response = &DeleteHealthCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
