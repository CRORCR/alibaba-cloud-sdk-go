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

// DisassociateQos invokes the smartag.DisassociateQos API synchronously
// api document: https://help.aliyun.com/api/smartag/disassociateqos.html
func (client *Client) DisassociateQos(request *DisassociateQosRequest) (response *DisassociateQosResponse, err error) {
	response = CreateDisassociateQosResponse()
	err = client.DoAction(request, response)
	return
}

// DisassociateQosWithChan invokes the smartag.DisassociateQos API asynchronously
// api document: https://help.aliyun.com/api/smartag/disassociateqos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisassociateQosWithChan(request *DisassociateQosRequest) (<-chan *DisassociateQosResponse, <-chan error) {
	responseChan := make(chan *DisassociateQosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisassociateQos(request)
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

// DisassociateQosWithCallback invokes the smartag.DisassociateQos API asynchronously
// api document: https://help.aliyun.com/api/smartag/disassociateqos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisassociateQosWithCallback(request *DisassociateQosRequest, callback func(response *DisassociateQosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisassociateQosResponse
		var err error
		defer close(result)
		response, err = client.DisassociateQos(request)
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

// DisassociateQosRequest is the request struct for api DisassociateQos
type DisassociateQosRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	QosId                string           `position:"Query" name:"QosId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// DisassociateQosResponse is the response struct for api DisassociateQos
type DisassociateQosResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisassociateQosRequest creates a request to invoke DisassociateQos API
func CreateDisassociateQosRequest() (request *DisassociateQosRequest) {
	request = &DisassociateQosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DisassociateQos", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisassociateQosResponse creates a response to parse from DisassociateQos response
func CreateDisassociateQosResponse() (response *DisassociateQosResponse) {
	response = &DisassociateQosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
