package dm

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

// DescAccountSummary2 invokes the dm.DescAccountSummary2 API synchronously
// api document: https://help.aliyun.com/api/dm/descaccountsummary2.html
func (client *Client) DescAccountSummary2(request *DescAccountSummary2Request) (response *DescAccountSummary2Response, err error) {
	response = CreateDescAccountSummary2Response()
	err = client.DoAction(request, response)
	return
}

// DescAccountSummary2WithChan invokes the dm.DescAccountSummary2 API asynchronously
// api document: https://help.aliyun.com/api/dm/descaccountsummary2.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescAccountSummary2WithChan(request *DescAccountSummary2Request) (<-chan *DescAccountSummary2Response, <-chan error) {
	responseChan := make(chan *DescAccountSummary2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescAccountSummary2(request)
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

// DescAccountSummary2WithCallback invokes the dm.DescAccountSummary2 API asynchronously
// api document: https://help.aliyun.com/api/dm/descaccountsummary2.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescAccountSummary2WithCallback(request *DescAccountSummary2Request, callback func(response *DescAccountSummary2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescAccountSummary2Response
		var err error
		defer close(result)
		response, err = client.DescAccountSummary2(request)
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

// DescAccountSummary2Request is the request struct for api DescAccountSummary2
type DescAccountSummary2Request struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

// DescAccountSummary2Response is the response struct for api DescAccountSummary2
type DescAccountSummary2Response struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	MnsMigrating      int    `json:"MnsMigrating" xml:"MnsMigrating"`
	MnsForceMigrating int    `json:"MnsForceMigrating" xml:"MnsForceMigrating"`
	MnsBag            int    `json:"MnsBag" xml:"MnsBag"`
}

// CreateDescAccountSummary2Request creates a request to invoke DescAccountSummary2 API
func CreateDescAccountSummary2Request() (request *DescAccountSummary2Request) {
	request = &DescAccountSummary2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DescAccountSummary2", "", "")
	return
}

// CreateDescAccountSummary2Response creates a response to parse from DescAccountSummary2 response
func CreateDescAccountSummary2Response() (response *DescAccountSummary2Response) {
	response = &DescAccountSummary2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
