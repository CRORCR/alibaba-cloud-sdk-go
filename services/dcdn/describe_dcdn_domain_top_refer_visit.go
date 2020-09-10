package dcdn

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

// DescribeDcdnDomainTopReferVisit invokes the dcdn.DescribeDcdnDomainTopReferVisit API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomaintoprefervisit.html
func (client *Client) DescribeDcdnDomainTopReferVisit(request *DescribeDcdnDomainTopReferVisitRequest) (response *DescribeDcdnDomainTopReferVisitResponse, err error) {
	response = CreateDescribeDcdnDomainTopReferVisitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainTopReferVisitWithChan invokes the dcdn.DescribeDcdnDomainTopReferVisit API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomaintoprefervisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnDomainTopReferVisitWithChan(request *DescribeDcdnDomainTopReferVisitRequest) (<-chan *DescribeDcdnDomainTopReferVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainTopReferVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainTopReferVisit(request)
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

// DescribeDcdnDomainTopReferVisitWithCallback invokes the dcdn.DescribeDcdnDomainTopReferVisit API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomaintoprefervisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnDomainTopReferVisitWithCallback(request *DescribeDcdnDomainTopReferVisitRequest, callback func(response *DescribeDcdnDomainTopReferVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainTopReferVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainTopReferVisit(request)
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

// DescribeDcdnDomainTopReferVisitRequest is the request struct for api DescribeDcdnDomainTopReferVisit
type DescribeDcdnDomainTopReferVisitRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	SortBy        string           `position:"Query" name:"SortBy"`
}

// DescribeDcdnDomainTopReferVisitResponse is the response struct for api DescribeDcdnDomainTopReferVisit
type DescribeDcdnDomainTopReferVisitResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	DomainName   string       `json:"DomainName" xml:"DomainName"`
	StartTime    string       `json:"StartTime" xml:"StartTime"`
	TopReferList TopReferList `json:"TopReferList" xml:"TopReferList"`
}

// CreateDescribeDcdnDomainTopReferVisitRequest creates a request to invoke DescribeDcdnDomainTopReferVisit API
func CreateDescribeDcdnDomainTopReferVisitRequest() (request *DescribeDcdnDomainTopReferVisitRequest) {
	request = &DescribeDcdnDomainTopReferVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainTopReferVisit", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainTopReferVisitResponse creates a response to parse from DescribeDcdnDomainTopReferVisit response
func CreateDescribeDcdnDomainTopReferVisitResponse() (response *DescribeDcdnDomainTopReferVisitResponse) {
	response = &DescribeDcdnDomainTopReferVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
