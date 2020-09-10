package dypnsapi

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

// GetMobile invokes the dypnsapi.GetMobile API synchronously
// api document: https://help.aliyun.com/api/dypnsapi/getmobile.html
func (client *Client) GetMobile(request *GetMobileRequest) (response *GetMobileResponse, err error) {
	response = CreateGetMobileResponse()
	err = client.DoAction(request, response)
	return
}

// GetMobileWithChan invokes the dypnsapi.GetMobile API asynchronously
// api document: https://help.aliyun.com/api/dypnsapi/getmobile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMobileWithChan(request *GetMobileRequest) (<-chan *GetMobileResponse, <-chan error) {
	responseChan := make(chan *GetMobileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMobile(request)
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

// GetMobileWithCallback invokes the dypnsapi.GetMobile API asynchronously
// api document: https://help.aliyun.com/api/dypnsapi/getmobile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMobileWithCallback(request *GetMobileRequest, callback func(response *GetMobileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMobileResponse
		var err error
		defer close(result)
		response, err = client.GetMobile(request)
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

// GetMobileRequest is the request struct for api GetMobile
type GetMobileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessToken          string           `position:"Query" name:"AccessToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OutId                string           `position:"Query" name:"OutId"`
}

// GetMobileResponse is the response struct for api GetMobile
type GetMobileResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Code               string             `json:"Code" xml:"Code"`
	Message            string             `json:"Message" xml:"Message"`
	GetMobileResultDTO GetMobileResultDTO `json:"GetMobileResultDTO" xml:"GetMobileResultDTO"`
}

// CreateGetMobileRequest creates a request to invoke GetMobile API
func CreateGetMobileRequest() (request *GetMobileRequest) {
	request = &GetMobileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "GetMobile", "dypns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMobileResponse creates a response to parse from GetMobile response
func CreateGetMobileResponse() (response *GetMobileResponse) {
	response = &GetMobileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
