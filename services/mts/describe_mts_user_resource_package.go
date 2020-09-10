package mts

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

// DescribeMtsUserResourcePackage invokes the mts.DescribeMtsUserResourcePackage API synchronously
// api document: https://help.aliyun.com/api/mts/describemtsuserresourcepackage.html
func (client *Client) DescribeMtsUserResourcePackage(request *DescribeMtsUserResourcePackageRequest) (response *DescribeMtsUserResourcePackageResponse, err error) {
	response = CreateDescribeMtsUserResourcePackageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMtsUserResourcePackageWithChan invokes the mts.DescribeMtsUserResourcePackage API asynchronously
// api document: https://help.aliyun.com/api/mts/describemtsuserresourcepackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMtsUserResourcePackageWithChan(request *DescribeMtsUserResourcePackageRequest) (<-chan *DescribeMtsUserResourcePackageResponse, <-chan error) {
	responseChan := make(chan *DescribeMtsUserResourcePackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMtsUserResourcePackage(request)
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

// DescribeMtsUserResourcePackageWithCallback invokes the mts.DescribeMtsUserResourcePackage API asynchronously
// api document: https://help.aliyun.com/api/mts/describemtsuserresourcepackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMtsUserResourcePackageWithCallback(request *DescribeMtsUserResourcePackageRequest, callback func(response *DescribeMtsUserResourcePackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMtsUserResourcePackageResponse
		var err error
		defer close(result)
		response, err = client.DescribeMtsUserResourcePackage(request)
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

// DescribeMtsUserResourcePackageRequest is the request struct for api DescribeMtsUserResourcePackage
type DescribeMtsUserResourcePackageRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	ShowLog       string           `position:"Query" name:"ShowLog"`
}

// DescribeMtsUserResourcePackageResponse is the response struct for api DescribeMtsUserResourcePackage
type DescribeMtsUserResourcePackageResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	ResourcePackageInfos ResourcePackageInfos `json:"ResourcePackageInfos" xml:"ResourcePackageInfos"`
}

// CreateDescribeMtsUserResourcePackageRequest creates a request to invoke DescribeMtsUserResourcePackage API
func CreateDescribeMtsUserResourcePackageRequest() (request *DescribeMtsUserResourcePackageRequest) {
	request = &DescribeMtsUserResourcePackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "DescribeMtsUserResourcePackage", "", "")
	return
}

// CreateDescribeMtsUserResourcePackageResponse creates a response to parse from DescribeMtsUserResourcePackage response
func CreateDescribeMtsUserResourcePackageResponse() (response *DescribeMtsUserResourcePackageResponse) {
	response = &DescribeMtsUserResourcePackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
