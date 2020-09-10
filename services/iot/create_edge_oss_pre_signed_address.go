package iot

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

// CreateEdgeOssPreSignedAddress invokes the iot.CreateEdgeOssPreSignedAddress API synchronously
// api document: https://help.aliyun.com/api/iot/createedgeosspresignedaddress.html
func (client *Client) CreateEdgeOssPreSignedAddress(request *CreateEdgeOssPreSignedAddressRequest) (response *CreateEdgeOssPreSignedAddressResponse, err error) {
	response = CreateCreateEdgeOssPreSignedAddressResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEdgeOssPreSignedAddressWithChan invokes the iot.CreateEdgeOssPreSignedAddress API asynchronously
// api document: https://help.aliyun.com/api/iot/createedgeosspresignedaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEdgeOssPreSignedAddressWithChan(request *CreateEdgeOssPreSignedAddressRequest) (<-chan *CreateEdgeOssPreSignedAddressResponse, <-chan error) {
	responseChan := make(chan *CreateEdgeOssPreSignedAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEdgeOssPreSignedAddress(request)
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

// CreateEdgeOssPreSignedAddressWithCallback invokes the iot.CreateEdgeOssPreSignedAddress API asynchronously
// api document: https://help.aliyun.com/api/iot/createedgeosspresignedaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEdgeOssPreSignedAddressWithCallback(request *CreateEdgeOssPreSignedAddressRequest, callback func(response *CreateEdgeOssPreSignedAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEdgeOssPreSignedAddressResponse
		var err error
		defer close(result)
		response, err = client.CreateEdgeOssPreSignedAddress(request)
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

// CreateEdgeOssPreSignedAddressRequest is the request struct for api CreateEdgeOssPreSignedAddress
type CreateEdgeOssPreSignedAddressRequest struct {
	*requests.RpcRequest
	Type            string `position:"Query" name:"Type"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	ResourceVersion string `position:"Query" name:"ResourceVersion"`
	ResourceId      string `position:"Query" name:"ResourceId"`
	FileName        string `position:"Query" name:"FileName"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
}

// CreateEdgeOssPreSignedAddressResponse is the response struct for api CreateEdgeOssPreSignedAddress
type CreateEdgeOssPreSignedAddressResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateEdgeOssPreSignedAddressRequest creates a request to invoke CreateEdgeOssPreSignedAddress API
func CreateCreateEdgeOssPreSignedAddressRequest() (request *CreateEdgeOssPreSignedAddressRequest) {
	request = &CreateEdgeOssPreSignedAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateEdgeOssPreSignedAddress", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEdgeOssPreSignedAddressResponse creates a response to parse from CreateEdgeOssPreSignedAddress response
func CreateCreateEdgeOssPreSignedAddressResponse() (response *CreateEdgeOssPreSignedAddressResponse) {
	response = &CreateEdgeOssPreSignedAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
