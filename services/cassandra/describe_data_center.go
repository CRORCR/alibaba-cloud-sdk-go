package cassandra

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

// DescribeDataCenter invokes the cassandra.DescribeDataCenter API synchronously
// api document: https://help.aliyun.com/api/cassandra/describedatacenter.html
func (client *Client) DescribeDataCenter(request *DescribeDataCenterRequest) (response *DescribeDataCenterResponse, err error) {
	response = CreateDescribeDataCenterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataCenterWithChan invokes the cassandra.DescribeDataCenter API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describedatacenter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataCenterWithChan(request *DescribeDataCenterRequest) (<-chan *DescribeDataCenterResponse, <-chan error) {
	responseChan := make(chan *DescribeDataCenterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataCenter(request)
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

// DescribeDataCenterWithCallback invokes the cassandra.DescribeDataCenter API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describedatacenter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataCenterWithCallback(request *DescribeDataCenterRequest, callback func(response *DescribeDataCenterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataCenterResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataCenter(request)
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

// DescribeDataCenterRequest is the request struct for api DescribeDataCenter
type DescribeDataCenterRequest struct {
	*requests.RpcRequest
	DataCenterId string `position:"Query" name:"DataCenterId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
}

// DescribeDataCenterResponse is the response struct for api DescribeDataCenter
type DescribeDataCenterResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DataCenterId      string `json:"DataCenterId" xml:"DataCenterId"`
	RegionId          string `json:"RegionId" xml:"RegionId"`
	ZoneId            string `json:"ZoneId" xml:"ZoneId"`
	ClusterId         string `json:"ClusterId" xml:"ClusterId"`
	DataCenterName    string `json:"DataCenterName" xml:"DataCenterName"`
	Status            string `json:"Status" xml:"Status"`
	CreatedTime       string `json:"CreatedTime" xml:"CreatedTime"`
	InstanceType      string `json:"InstanceType" xml:"InstanceType"`
	NodeCount         int    `json:"NodeCount" xml:"NodeCount"`
	DiskType          string `json:"DiskType" xml:"DiskType"`
	DiskSize          int    `json:"DiskSize" xml:"DiskSize"`
	VpcId             string `json:"VpcId" xml:"VpcId"`
	VswitchId         string `json:"VswitchId" xml:"VswitchId"`
	PayType           string `json:"PayType" xml:"PayType"`
	CommodityInstance string `json:"CommodityInstance" xml:"CommodityInstance"`
	ExpireTime        string `json:"ExpireTime" xml:"ExpireTime"`
	LockMode          string `json:"LockMode" xml:"LockMode"`
	AutoRenewal       bool   `json:"AutoRenewal" xml:"AutoRenewal"`
	AutoRenewPeriod   int    `json:"AutoRenewPeriod" xml:"AutoRenewPeriod"`
}

// CreateDescribeDataCenterRequest creates a request to invoke DescribeDataCenter API
func CreateDescribeDataCenterRequest() (request *DescribeDataCenterRequest) {
	request = &DescribeDataCenterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeDataCenter", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataCenterResponse creates a response to parse from DescribeDataCenter response
func CreateDescribeDataCenterResponse() (response *DescribeDataCenterResponse) {
	response = &DescribeDataCenterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
