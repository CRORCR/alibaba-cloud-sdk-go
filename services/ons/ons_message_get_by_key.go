package ons

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

// OnsMessageGetByKey invokes the ons.OnsMessageGetByKey API synchronously
// api document: https://help.aliyun.com/api/ons/onsmessagegetbykey.html
func (client *Client) OnsMessageGetByKey(request *OnsMessageGetByKeyRequest) (response *OnsMessageGetByKeyResponse, err error) {
	response = CreateOnsMessageGetByKeyResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMessageGetByKeyWithChan invokes the ons.OnsMessageGetByKey API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmessagegetbykey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMessageGetByKeyWithChan(request *OnsMessageGetByKeyRequest) (<-chan *OnsMessageGetByKeyResponse, <-chan error) {
	responseChan := make(chan *OnsMessageGetByKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMessageGetByKey(request)
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

// OnsMessageGetByKeyWithCallback invokes the ons.OnsMessageGetByKey API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmessagegetbykey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMessageGetByKeyWithCallback(request *OnsMessageGetByKeyRequest, callback func(response *OnsMessageGetByKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMessageGetByKeyResponse
		var err error
		defer close(result)
		response, err = client.OnsMessageGetByKey(request)
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

// OnsMessageGetByKeyRequest is the request struct for api OnsMessageGetByKey
type OnsMessageGetByKeyRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
	Key        string `position:"Query" name:"Key"`
}

// OnsMessageGetByKeyResponse is the response struct for api OnsMessageGetByKey
type OnsMessageGetByKeyResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	HelpUrl   string                   `json:"HelpUrl" xml:"HelpUrl"`
	Data      DataInOnsMessageGetByKey `json:"Data" xml:"Data"`
}

// CreateOnsMessageGetByKeyRequest creates a request to invoke OnsMessageGetByKey API
func CreateOnsMessageGetByKeyRequest() (request *OnsMessageGetByKeyRequest) {
	request = &OnsMessageGetByKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMessageGetByKey", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMessageGetByKeyResponse creates a response to parse from OnsMessageGetByKey response
func CreateOnsMessageGetByKeyResponse() (response *OnsMessageGetByKeyResponse) {
	response = &OnsMessageGetByKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
