package cs

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

// DescribeClusterDetail invokes the cs.DescribeClusterDetail API synchronously
// api document: https://help.aliyun.com/api/cs/describeclusterdetail.html
func (client *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
	response = CreateDescribeClusterDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterDetailWithChan invokes the cs.DescribeClusterDetail API asynchronously
// api document: https://help.aliyun.com/api/cs/describeclusterdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterDetailWithChan(request *DescribeClusterDetailRequest) (<-chan *DescribeClusterDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterDetail(request)
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

// DescribeClusterDetailWithCallback invokes the cs.DescribeClusterDetail API asynchronously
// api document: https://help.aliyun.com/api/cs/describeclusterdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterDetailWithCallback(request *DescribeClusterDetailRequest, callback func(response *DescribeClusterDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterDetail(request)
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

// DescribeClusterDetailRequest is the request struct for api DescribeClusterDetail
type DescribeClusterDetailRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// DescribeClusterDetailResponse is the response struct for api DescribeClusterDetail
type DescribeClusterDetailResponse struct {
	*responses.BaseResponse
	Name                   string     `json:"name" xml:"name"`
	ClusterId              string     `json:"cluster_id" xml:"cluster_id"`
	RegionId               string     `json:"region_id" xml:"region_id"`
	State                  string     `json:"state" xml:"state"`
	ClusterType            string     `json:"cluster_type" xml:"cluster_type"`
	CurrentVersion         string     `json:"current_version" xml:"current_version"`
	MetaData               string     `json:"meta_data" xml:"meta_data"`
	ResourceGroupId        string     `json:"resource_group_id" xml:"resource_group_id"`
	InstanceType           string     `json:"instance_type" xml:"instance_type"`
	VpcId                  string     `json:"vpc_id" xml:"vpc_id"`
	VswitchId              string     `json:"vswitch_id" xml:"vswitch_id"`
	VswitchCidr            string     `json:"vswitch_cidr" xml:"vswitch_cidr"`
	DataDiskSize           int        `json:"data_disk_size" xml:"data_disk_size"`
	DataDiskCategory       string     `json:"data_disk_category" xml:"data_disk_category"`
	SecurityGroupId        string     `json:"security_group_id" xml:"security_group_id"`
	ZoneId                 string     `json:"zone_id" xml:"zone_id"`
	NetworkMode            string     `json:"network_mode" xml:"network_mode"`
	DockerVersion          string     `json:"docker_version" xml:"docker_version"`
	DeletionProtection     bool       `json:"deletion_protection" xml:"deletion_protection"`
	ExternalLoadbalancerId string     `json:"external_loadbalancer_id" xml:"external_loadbalancer_id"`
	Created                string     `json:"created" xml:"created"`
	Updated                string     `json:"updated" xml:"updated"`
	Size                   string     `json:"size" xml:"size"`
	Tags                   []TagsItem `json:"tags" xml:"tags"`
}

// CreateDescribeClusterDetailRequest creates a request to invoke DescribeClusterDetail API
func CreateDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
	request = &DescribeClusterDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterDetail", "/clusters/[ClusterId]", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeClusterDetailResponse creates a response to parse from DescribeClusterDetail response
func CreateDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
	response = &DescribeClusterDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
