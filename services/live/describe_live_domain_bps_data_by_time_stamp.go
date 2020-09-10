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

// DescribeLiveDomainBpsDataByTimeStamp invokes the live.DescribeLiveDomainBpsDataByTimeStamp API synchronously
// api document: https://help.aliyun.com/api/live/describelivedomainbpsdatabytimestamp.html
func (client *Client) DescribeLiveDomainBpsDataByTimeStamp(request *DescribeLiveDomainBpsDataByTimeStampRequest) (response *DescribeLiveDomainBpsDataByTimeStampResponse, err error) {
	response = CreateDescribeLiveDomainBpsDataByTimeStampResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainBpsDataByTimeStampWithChan invokes the live.DescribeLiveDomainBpsDataByTimeStamp API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedomainbpsdatabytimestamp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDomainBpsDataByTimeStampWithChan(request *DescribeLiveDomainBpsDataByTimeStampRequest) (<-chan *DescribeLiveDomainBpsDataByTimeStampResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainBpsDataByTimeStampResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainBpsDataByTimeStamp(request)
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

// DescribeLiveDomainBpsDataByTimeStampWithCallback invokes the live.DescribeLiveDomainBpsDataByTimeStamp API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedomainbpsdatabytimestamp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDomainBpsDataByTimeStampWithCallback(request *DescribeLiveDomainBpsDataByTimeStampRequest, callback func(response *DescribeLiveDomainBpsDataByTimeStampResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainBpsDataByTimeStampResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainBpsDataByTimeStamp(request)
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

// DescribeLiveDomainBpsDataByTimeStampRequest is the request struct for api DescribeLiveDomainBpsDataByTimeStamp
type DescribeLiveDomainBpsDataByTimeStampRequest struct {
	*requests.RpcRequest
	LocationNames string           `position:"Query" name:"LocationNames"`
	IspNames      string           `position:"Query" name:"IspNames"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	TimePoint     string           `position:"Query" name:"TimePoint"`
}

// DescribeLiveDomainBpsDataByTimeStampResponse is the response struct for api DescribeLiveDomainBpsDataByTimeStamp
type DescribeLiveDomainBpsDataByTimeStampResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	DomainName  string      `json:"DomainName" xml:"DomainName"`
	TimeStamp   string      `json:"TimeStamp" xml:"TimeStamp"`
	BpsDataList BpsDataList `json:"BpsDataList" xml:"BpsDataList"`
}

// CreateDescribeLiveDomainBpsDataByTimeStampRequest creates a request to invoke DescribeLiveDomainBpsDataByTimeStamp API
func CreateDescribeLiveDomainBpsDataByTimeStampRequest() (request *DescribeLiveDomainBpsDataByTimeStampRequest) {
	request = &DescribeLiveDomainBpsDataByTimeStampRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainBpsDataByTimeStamp", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainBpsDataByTimeStampResponse creates a response to parse from DescribeLiveDomainBpsDataByTimeStamp response
func CreateDescribeLiveDomainBpsDataByTimeStampResponse() (response *DescribeLiveDomainBpsDataByTimeStampResponse) {
	response = &DescribeLiveDomainBpsDataByTimeStampResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
