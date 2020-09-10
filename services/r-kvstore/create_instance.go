package r_kvstore

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

// CreateInstance invokes the r_kvstore.CreateInstance API synchronously
func (client *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	response = CreateCreateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInstanceWithChan invokes the r_kvstore.CreateInstance API asynchronously
func (client *Client) CreateInstanceWithChan(request *CreateInstanceRequest) (<-chan *CreateInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstance(request)
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

// CreateInstanceWithCallback invokes the r_kvstore.CreateInstance API asynchronously
func (client *Client) CreateInstanceWithCallback(request *CreateInstanceRequest, callback func(response *CreateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateInstance(request)
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

// CreateInstanceRequest is the request struct for api CreateInstance
type CreateInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	Password             string           `position:"Query" name:"Password"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	ShardCount           requests.Integer `position:"Query" name:"ShardCount"`
	AutoRenewPeriod      string           `position:"Query" name:"AutoRenewPeriod"`
	Period               string           `position:"Query" name:"Period"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string           `position:"Query" name:"PrivateIpAddress"`
	InstanceName         string           `position:"Query" name:"InstanceName"`
	AutoRenew            string           `position:"Query" name:"AutoRenew"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	NodeType             string           `position:"Query" name:"NodeType"`
	AutoUseCoupon        string           `position:"Query" name:"AutoUseCoupon"`
	InstanceClass        string           `position:"Query" name:"InstanceClass"`
	Capacity             requests.Integer `position:"Query" name:"Capacity"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	DedicatedHostGroupId string           `position:"Query" name:"DedicatedHostGroupId"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId      string           `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	GlobalInstance       requests.Boolean `position:"Query" name:"GlobalInstance"`
	Token                string           `position:"Query" name:"Token"`
	GlobalInstanceId     string           `position:"Query" name:"GlobalInstanceId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ChargeType           string           `position:"Query" name:"ChargeType"`
	Config               string           `position:"Query" name:"Config"`
}

// CreateInstanceResponse is the response struct for api CreateInstance
type CreateInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	InstanceName     string `json:"InstanceName" xml:"InstanceName"`
	ConnectionDomain string `json:"ConnectionDomain" xml:"ConnectionDomain"`
	Port             int    `json:"Port" xml:"Port"`
	UserName         string `json:"UserName" xml:"UserName"`
	InstanceStatus   string `json:"InstanceStatus" xml:"InstanceStatus"`
	RegionId         string `json:"RegionId" xml:"RegionId"`
	Capacity         int64  `json:"Capacity" xml:"Capacity"`
	QPS              int64  `json:"QPS" xml:"QPS"`
	Bandwidth        int64  `json:"Bandwidth" xml:"Bandwidth"`
	Connections      int64  `json:"Connections" xml:"Connections"`
	ZoneId           string `json:"ZoneId" xml:"ZoneId"`
	Config           string `json:"Config" xml:"Config"`
	ChargeType       string `json:"ChargeType" xml:"ChargeType"`
	EndTime          string `json:"EndTime" xml:"EndTime"`
	NodeType         string `json:"NodeType" xml:"NodeType"`
	NetworkType      string `json:"NetworkType" xml:"NetworkType"`
	VpcId            string `json:"VpcId" xml:"VpcId"`
	VSwitchId        string `json:"VSwitchId" xml:"VSwitchId"`
	PrivateIpAddr    string `json:"PrivateIpAddr" xml:"PrivateIpAddr"`
}

// CreateCreateInstanceRequest creates a request to invoke CreateInstance API
func CreateCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateInstance", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateInstanceResponse creates a response to parse from CreateInstance response
func CreateCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
