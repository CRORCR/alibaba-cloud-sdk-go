package smartag

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

// DescribeSmartAccessGatewayAttribute invokes the smartag.DescribeSmartAccessGatewayAttribute API synchronously
// api document: https://help.aliyun.com/api/smartag/describesmartaccessgatewayattribute.html
func (client *Client) DescribeSmartAccessGatewayAttribute(request *DescribeSmartAccessGatewayAttributeRequest) (response *DescribeSmartAccessGatewayAttributeResponse, err error) {
	response = CreateDescribeSmartAccessGatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSmartAccessGatewayAttributeWithChan invokes the smartag.DescribeSmartAccessGatewayAttribute API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesmartaccessgatewayattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSmartAccessGatewayAttributeWithChan(request *DescribeSmartAccessGatewayAttributeRequest) (<-chan *DescribeSmartAccessGatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeSmartAccessGatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSmartAccessGatewayAttribute(request)
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

// DescribeSmartAccessGatewayAttributeWithCallback invokes the smartag.DescribeSmartAccessGatewayAttribute API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesmartaccessgatewayattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSmartAccessGatewayAttributeWithCallback(request *DescribeSmartAccessGatewayAttributeRequest, callback func(response *DescribeSmartAccessGatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSmartAccessGatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeSmartAccessGatewayAttribute(request)
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

// DescribeSmartAccessGatewayAttributeRequest is the request struct for api DescribeSmartAccessGatewayAttribute
type DescribeSmartAccessGatewayAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// DescribeSmartAccessGatewayAttributeResponse is the response struct for api DescribeSmartAccessGatewayAttribute
type DescribeSmartAccessGatewayAttributeResponse struct {
	*responses.BaseResponse
	RequestId             string                                     `json:"RequestId" xml:"RequestId"`
	SmartAGId             string                                     `json:"SmartAGId" xml:"SmartAGId"`
	Name                  string                                     `json:"Name" xml:"Name"`
	City                  string                                     `json:"City" xml:"City"`
	MaxBandwidth          string                                     `json:"MaxBandwidth" xml:"MaxBandwidth"`
	Status                string                                     `json:"Status" xml:"Status"`
	CidrBlock             string                                     `json:"CidrBlock" xml:"CidrBlock"`
	AssociatedCcnId       string                                     `json:"AssociatedCcnId" xml:"AssociatedCcnId"`
	AssociatedCcnName     string                                     `json:"AssociatedCcnName" xml:"AssociatedCcnName"`
	Description           string                                     `json:"Description" xml:"Description"`
	CreateTime            int64                                      `json:"CreateTime" xml:"CreateTime"`
	EndTime               int64                                      `json:"EndTime" xml:"EndTime"`
	InstanceType          string                                     `json:"InstanceType" xml:"InstanceType"`
	SerialNumber          string                                     `json:"SerialNumber" xml:"SerialNumber"`
	SecurityLockThreshold int                                        `json:"SecurityLockThreshold" xml:"SecurityLockThreshold"`
	DataPlan              int64                                      `json:"DataPlan" xml:"DataPlan"`
	UserCount             int                                        `json:"UserCount" xml:"UserCount"`
	RoutingStrategy       string                                     `json:"RoutingStrategy" xml:"RoutingStrategy"`
	IpsecStatus           string                                     `json:"IpsecStatus" xml:"IpsecStatus"`
	VpnStatus             string                                     `json:"VpnStatus" xml:"VpnStatus"`
	TrafficMasterSn       string                                     `json:"TrafficMasterSn" xml:"TrafficMasterSn"`
	BoxControllerIp       string                                     `json:"BoxControllerIp" xml:"BoxControllerIp"`
	BackupBoxControllerIp string                                     `json:"BackupBoxControllerIp" xml:"BackupBoxControllerIp"`
	UpBandwidthWan        int                                        `json:"UpBandwidthWan" xml:"UpBandwidthWan"`
	UpBandwidth4G         int                                        `json:"UpBandwidth4G" xml:"UpBandwidth4G"`
	EnableOptimization    bool                                       `json:"EnableOptimization" xml:"EnableOptimization"`
	OptimizationType      bool                                       `json:"OptimizationType" xml:"OptimizationType"`
	ResourceGroupId       string                                     `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AccessPointId         string                                     `json:"AccessPointId" xml:"AccessPointId"`
	AclIds                AclIds                                     `json:"AclIds" xml:"AclIds"`
	QosIds                QosIds                                     `json:"QosIds" xml:"QosIds"`
	FlowLogIds            FlowLogIds                                 `json:"FlowLogIds" xml:"FlowLogIds"`
	Links                 LinksInDescribeSmartAccessGatewayAttribute `json:"Links" xml:"Links"`
	Devices               Devices                                    `json:"Devices" xml:"Devices"`
}

// CreateDescribeSmartAccessGatewayAttributeRequest creates a request to invoke DescribeSmartAccessGatewayAttribute API
func CreateDescribeSmartAccessGatewayAttributeRequest() (request *DescribeSmartAccessGatewayAttributeRequest) {
	request = &DescribeSmartAccessGatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSmartAccessGatewayAttribute", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSmartAccessGatewayAttributeResponse creates a response to parse from DescribeSmartAccessGatewayAttribute response
func CreateDescribeSmartAccessGatewayAttributeResponse() (response *DescribeSmartAccessGatewayAttributeResponse) {
	response = &DescribeSmartAccessGatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
