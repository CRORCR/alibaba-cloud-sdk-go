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

// DescribeRouteServicesInCen invokes the cbn.DescribeRouteServicesInCen API synchronously
// api document: https://help.aliyun.com/api/cbn/describerouteservicesincen.html
func (client *Client) DescribeRouteServicesInCen(request *DescribeRouteServicesInCenRequest) (response *DescribeRouteServicesInCenResponse, err error) {
	response = CreateDescribeRouteServicesInCenResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRouteServicesInCenWithChan invokes the cbn.DescribeRouteServicesInCen API asynchronously
// api document: https://help.aliyun.com/api/cbn/describerouteservicesincen.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouteServicesInCenWithChan(request *DescribeRouteServicesInCenRequest) (<-chan *DescribeRouteServicesInCenResponse, <-chan error) {
	responseChan := make(chan *DescribeRouteServicesInCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouteServicesInCen(request)
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

// DescribeRouteServicesInCenWithCallback invokes the cbn.DescribeRouteServicesInCen API asynchronously
// api document: https://help.aliyun.com/api/cbn/describerouteservicesincen.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouteServicesInCenWithCallback(request *DescribeRouteServicesInCenRequest, callback func(response *DescribeRouteServicesInCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouteServicesInCenResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouteServicesInCen(request)
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

// DescribeRouteServicesInCenRequest is the request struct for api DescribeRouteServicesInCen
type DescribeRouteServicesInCenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	AccessRegionId       string           `position:"Query" name:"AccessRegionId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Host                 string           `position:"Query" name:"Host"`
	HostRegionId         string           `position:"Query" name:"HostRegionId"`
	HostVpcId            string           `position:"Query" name:"HostVpcId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRouteServicesInCenResponse is the response struct for api DescribeRouteServicesInCen
type DescribeRouteServicesInCenResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	TotalCount          int                 `json:"TotalCount" xml:"TotalCount"`
	PageNumber          int                 `json:"PageNumber" xml:"PageNumber"`
	PageSize            int                 `json:"PageSize" xml:"PageSize"`
	RouteServiceEntries RouteServiceEntries `json:"RouteServiceEntries" xml:"RouteServiceEntries"`
}

// CreateDescribeRouteServicesInCenRequest creates a request to invoke DescribeRouteServicesInCen API
func CreateDescribeRouteServicesInCenRequest() (request *DescribeRouteServicesInCenRequest) {
	request = &DescribeRouteServicesInCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeRouteServicesInCen", "Cbn", "openAPI")
	return
}

// CreateDescribeRouteServicesInCenResponse creates a response to parse from DescribeRouteServicesInCen response
func CreateDescribeRouteServicesInCenResponse() (response *DescribeRouteServicesInCenResponse) {
	response = &DescribeRouteServicesInCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
