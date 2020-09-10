package cbn

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

// DescribeCenPrivateZoneRoutes invokes the cbn.DescribeCenPrivateZoneRoutes API synchronously
// api document: https://help.aliyun.com/api/cbn/describecenprivatezoneroutes.html
func (client *Client) DescribeCenPrivateZoneRoutes(request *DescribeCenPrivateZoneRoutesRequest) (response *DescribeCenPrivateZoneRoutesResponse, err error) {
	response = CreateDescribeCenPrivateZoneRoutesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCenPrivateZoneRoutesWithChan invokes the cbn.DescribeCenPrivateZoneRoutes API asynchronously
// api document: https://help.aliyun.com/api/cbn/describecenprivatezoneroutes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCenPrivateZoneRoutesWithChan(request *DescribeCenPrivateZoneRoutesRequest) (<-chan *DescribeCenPrivateZoneRoutesResponse, <-chan error) {
	responseChan := make(chan *DescribeCenPrivateZoneRoutesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCenPrivateZoneRoutes(request)
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

// DescribeCenPrivateZoneRoutesWithCallback invokes the cbn.DescribeCenPrivateZoneRoutes API asynchronously
// api document: https://help.aliyun.com/api/cbn/describecenprivatezoneroutes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCenPrivateZoneRoutesWithCallback(request *DescribeCenPrivateZoneRoutesRequest, callback func(response *DescribeCenPrivateZoneRoutesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCenPrivateZoneRoutesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCenPrivateZoneRoutes(request)
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

// DescribeCenPrivateZoneRoutesRequest is the request struct for api DescribeCenPrivateZoneRoutes
type DescribeCenPrivateZoneRoutesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	AccessRegionId       string           `position:"Query" name:"AccessRegionId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	HostRegionId         string           `position:"Query" name:"HostRegionId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
}

// DescribeCenPrivateZoneRoutesResponse is the response struct for api DescribeCenPrivateZoneRoutes
type DescribeCenPrivateZoneRoutesResponse struct {
	*responses.BaseResponse
	RequestId             string           `json:"RequestId" xml:"RequestId"`
	CenId                 string           `json:"CenId" xml:"CenId"`
	PrivateZoneDnsServers string           `json:"PrivateZoneDnsServers" xml:"PrivateZoneDnsServers"`
	PageNumber            int              `json:"PageNumber" xml:"PageNumber"`
	TotalCount            int              `json:"TotalCount" xml:"TotalCount"`
	PageSize              int              `json:"PageSize" xml:"PageSize"`
	PrivateZoneInfos      PrivateZoneInfos `json:"PrivateZoneInfos" xml:"PrivateZoneInfos"`
}

// CreateDescribeCenPrivateZoneRoutesRequest creates a request to invoke DescribeCenPrivateZoneRoutes API
func CreateDescribeCenPrivateZoneRoutesRequest() (request *DescribeCenPrivateZoneRoutesRequest) {
	request = &DescribeCenPrivateZoneRoutesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenPrivateZoneRoutes", "Cbn", "openAPI")
	return
}

// CreateDescribeCenPrivateZoneRoutesResponse creates a response to parse from DescribeCenPrivateZoneRoutes response
func CreateDescribeCenPrivateZoneRoutesResponse() (response *DescribeCenPrivateZoneRoutesResponse) {
	response = &DescribeCenPrivateZoneRoutesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
