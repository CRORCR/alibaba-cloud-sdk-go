package live

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

// DescribeLiveDomainOnlineUserNum invokes the live.DescribeLiveDomainOnlineUserNum API synchronously
// api document: https://help.aliyun.com/api/live/describelivedomainonlineusernum.html
func (client *Client) DescribeLiveDomainOnlineUserNum(request *DescribeLiveDomainOnlineUserNumRequest) (response *DescribeLiveDomainOnlineUserNumResponse, err error) {
	response = CreateDescribeLiveDomainOnlineUserNumResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainOnlineUserNumWithChan invokes the live.DescribeLiveDomainOnlineUserNum API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedomainonlineusernum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDomainOnlineUserNumWithChan(request *DescribeLiveDomainOnlineUserNumRequest) (<-chan *DescribeLiveDomainOnlineUserNumResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainOnlineUserNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainOnlineUserNum(request)
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

// DescribeLiveDomainOnlineUserNumWithCallback invokes the live.DescribeLiveDomainOnlineUserNum API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedomainonlineusernum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDomainOnlineUserNumWithCallback(request *DescribeLiveDomainOnlineUserNumRequest, callback func(response *DescribeLiveDomainOnlineUserNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainOnlineUserNumResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainOnlineUserNum(request)
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

// DescribeLiveDomainOnlineUserNumRequest is the request struct for api DescribeLiveDomainOnlineUserNum
type DescribeLiveDomainOnlineUserNumRequest struct {
	*requests.RpcRequest
	QueryTime  string           `position:"Query" name:"QueryTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainOnlineUserNumResponse is the response struct for api DescribeLiveDomainOnlineUserNum
type DescribeLiveDomainOnlineUserNumResponse struct {
	*responses.BaseResponse
	RequestId      string                                          `json:"RequestId" xml:"RequestId"`
	StreamCount    int                                             `json:"StreamCount" xml:"StreamCount"`
	UserCount      int                                             `json:"UserCount" xml:"UserCount"`
	OnlineUserInfo OnlineUserInfoInDescribeLiveDomainOnlineUserNum `json:"OnlineUserInfo" xml:"OnlineUserInfo"`
}

// CreateDescribeLiveDomainOnlineUserNumRequest creates a request to invoke DescribeLiveDomainOnlineUserNum API
func CreateDescribeLiveDomainOnlineUserNumRequest() (request *DescribeLiveDomainOnlineUserNumRequest) {
	request = &DescribeLiveDomainOnlineUserNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainOnlineUserNum", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainOnlineUserNumResponse creates a response to parse from DescribeLiveDomainOnlineUserNum response
func CreateDescribeLiveDomainOnlineUserNumResponse() (response *DescribeLiveDomainOnlineUserNumResponse) {
	response = &DescribeLiveDomainOnlineUserNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
