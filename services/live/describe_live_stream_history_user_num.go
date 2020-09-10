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

// DescribeLiveStreamHistoryUserNum invokes the live.DescribeLiveStreamHistoryUserNum API synchronously
// api document: https://help.aliyun.com/api/live/describelivestreamhistoryusernum.html
func (client *Client) DescribeLiveStreamHistoryUserNum(request *DescribeLiveStreamHistoryUserNumRequest) (response *DescribeLiveStreamHistoryUserNumResponse, err error) {
	response = CreateDescribeLiveStreamHistoryUserNumResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamHistoryUserNumWithChan invokes the live.DescribeLiveStreamHistoryUserNum API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamhistoryusernum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamHistoryUserNumWithChan(request *DescribeLiveStreamHistoryUserNumRequest) (<-chan *DescribeLiveStreamHistoryUserNumResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamHistoryUserNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamHistoryUserNum(request)
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

// DescribeLiveStreamHistoryUserNumWithCallback invokes the live.DescribeLiveStreamHistoryUserNum API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamhistoryusernum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamHistoryUserNumWithCallback(request *DescribeLiveStreamHistoryUserNumRequest, callback func(response *DescribeLiveStreamHistoryUserNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamHistoryUserNumResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamHistoryUserNum(request)
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

// DescribeLiveStreamHistoryUserNumRequest is the request struct for api DescribeLiveStreamHistoryUserNum
type DescribeLiveStreamHistoryUserNumRequest struct {
	*requests.RpcRequest
	StartTime     string           `position:"Query" name:"StartTime"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamHistoryUserNumResponse is the response struct for api DescribeLiveStreamHistoryUserNum
type DescribeLiveStreamHistoryUserNumResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	LiveStreamUserNumInfos LiveStreamUserNumInfos `json:"LiveStreamUserNumInfos" xml:"LiveStreamUserNumInfos"`
}

// CreateDescribeLiveStreamHistoryUserNumRequest creates a request to invoke DescribeLiveStreamHistoryUserNum API
func CreateDescribeLiveStreamHistoryUserNumRequest() (request *DescribeLiveStreamHistoryUserNumRequest) {
	request = &DescribeLiveStreamHistoryUserNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamHistoryUserNum", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamHistoryUserNumResponse creates a response to parse from DescribeLiveStreamHistoryUserNum response
func CreateDescribeLiveStreamHistoryUserNumResponse() (response *DescribeLiveStreamHistoryUserNumResponse) {
	response = &DescribeLiveStreamHistoryUserNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
