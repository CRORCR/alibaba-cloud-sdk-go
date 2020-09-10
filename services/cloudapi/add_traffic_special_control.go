package cloudapi

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

// AddTrafficSpecialControl invokes the cloudapi.AddTrafficSpecialControl API synchronously
// api document: https://help.aliyun.com/api/cloudapi/addtrafficspecialcontrol.html
func (client *Client) AddTrafficSpecialControl(request *AddTrafficSpecialControlRequest) (response *AddTrafficSpecialControlResponse, err error) {
	response = CreateAddTrafficSpecialControlResponse()
	err = client.DoAction(request, response)
	return
}

// AddTrafficSpecialControlWithChan invokes the cloudapi.AddTrafficSpecialControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/addtrafficspecialcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTrafficSpecialControlWithChan(request *AddTrafficSpecialControlRequest) (<-chan *AddTrafficSpecialControlResponse, <-chan error) {
	responseChan := make(chan *AddTrafficSpecialControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTrafficSpecialControl(request)
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

// AddTrafficSpecialControlWithCallback invokes the cloudapi.AddTrafficSpecialControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/addtrafficspecialcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTrafficSpecialControlWithCallback(request *AddTrafficSpecialControlRequest, callback func(response *AddTrafficSpecialControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTrafficSpecialControlResponse
		var err error
		defer close(result)
		response, err = client.AddTrafficSpecialControl(request)
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

// AddTrafficSpecialControlRequest is the request struct for api AddTrafficSpecialControl
type AddTrafficSpecialControlRequest struct {
	*requests.RpcRequest
	TrafficControlId string           `position:"Query" name:"TrafficControlId"`
	SpecialKey       string           `position:"Query" name:"SpecialKey"`
	TrafficValue     requests.Integer `position:"Query" name:"TrafficValue"`
	SecurityToken    string           `position:"Query" name:"SecurityToken"`
	SpecialType      string           `position:"Query" name:"SpecialType"`
}

// AddTrafficSpecialControlResponse is the response struct for api AddTrafficSpecialControl
type AddTrafficSpecialControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddTrafficSpecialControlRequest creates a request to invoke AddTrafficSpecialControl API
func CreateAddTrafficSpecialControlRequest() (request *AddTrafficSpecialControlRequest) {
	request = &AddTrafficSpecialControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "AddTrafficSpecialControl", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddTrafficSpecialControlResponse creates a response to parse from AddTrafficSpecialControl response
func CreateAddTrafficSpecialControlResponse() (response *AddTrafficSpecialControlResponse) {
	response = &AddTrafficSpecialControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
