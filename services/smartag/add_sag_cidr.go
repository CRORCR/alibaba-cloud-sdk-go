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

// AddSagCidr invokes the smartag.AddSagCidr API synchronously
// api document: https://help.aliyun.com/api/smartag/addsagcidr.html
func (client *Client) AddSagCidr(request *AddSagCidrRequest) (response *AddSagCidrResponse, err error) {
	response = CreateAddSagCidrResponse()
	err = client.DoAction(request, response)
	return
}

// AddSagCidrWithChan invokes the smartag.AddSagCidr API asynchronously
// api document: https://help.aliyun.com/api/smartag/addsagcidr.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSagCidrWithChan(request *AddSagCidrRequest) (<-chan *AddSagCidrResponse, <-chan error) {
	responseChan := make(chan *AddSagCidrResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSagCidr(request)
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

// AddSagCidrWithCallback invokes the smartag.AddSagCidr API asynchronously
// api document: https://help.aliyun.com/api/smartag/addsagcidr.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSagCidrWithCallback(request *AddSagCidrRequest, callback func(response *AddSagCidrResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSagCidrResponse
		var err error
		defer close(result)
		response, err = client.AddSagCidr(request)
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

// AddSagCidrRequest is the request struct for api AddSagCidr
type AddSagCidrRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Cidr                 string           `position:"Query" name:"Cidr"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	EnableBackup         requests.Boolean `position:"Query" name:"EnableBackup"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// AddSagCidrResponse is the response struct for api AddSagCidr
type AddSagCidrResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddSagCidrRequest creates a request to invoke AddSagCidr API
func CreateAddSagCidrRequest() (request *AddSagCidrRequest) {
	request = &AddSagCidrRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "AddSagCidr", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddSagCidrResponse creates a response to parse from AddSagCidr response
func CreateAddSagCidrResponse() (response *AddSagCidrResponse) {
	response = &AddSagCidrResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
