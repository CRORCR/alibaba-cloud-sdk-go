package ecs

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

// DescribeLaunchTemplateVersions invokes the ecs.DescribeLaunchTemplateVersions API synchronously
// api document: https://help.aliyun.com/api/ecs/describelaunchtemplateversions.html
func (client *Client) DescribeLaunchTemplateVersions(request *DescribeLaunchTemplateVersionsRequest) (response *DescribeLaunchTemplateVersionsResponse, err error) {
	response = CreateDescribeLaunchTemplateVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLaunchTemplateVersionsWithChan invokes the ecs.DescribeLaunchTemplateVersions API asynchronously
// api document: https://help.aliyun.com/api/ecs/describelaunchtemplateversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLaunchTemplateVersionsWithChan(request *DescribeLaunchTemplateVersionsRequest) (<-chan *DescribeLaunchTemplateVersionsResponse, <-chan error) {
	responseChan := make(chan *DescribeLaunchTemplateVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLaunchTemplateVersions(request)
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

// DescribeLaunchTemplateVersionsWithCallback invokes the ecs.DescribeLaunchTemplateVersions API asynchronously
// api document: https://help.aliyun.com/api/ecs/describelaunchtemplateversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLaunchTemplateVersionsWithCallback(request *DescribeLaunchTemplateVersionsRequest, callback func(response *DescribeLaunchTemplateVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLaunchTemplateVersionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLaunchTemplateVersions(request)
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

// DescribeLaunchTemplateVersionsRequest is the request struct for api DescribeLaunchTemplateVersions
type DescribeLaunchTemplateVersionsRequest struct {
	*requests.RpcRequest
	LaunchTemplateName    string           `position:"Query" name:"LaunchTemplateName"`
	MaxVersion            requests.Integer `position:"Query" name:"MaxVersion"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DefaultVersion        requests.Boolean `position:"Query" name:"DefaultVersion"`
	MinVersion            requests.Integer `position:"Query" name:"MinVersion"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	LaunchTemplateId      string           `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	LaunchTemplateVersion *[]string        `position:"Query" name:"LaunchTemplateVersion"  type:"Repeated"`
	DetailFlag            requests.Boolean `position:"Query" name:"DetailFlag"`
}

// DescribeLaunchTemplateVersionsResponse is the response struct for api DescribeLaunchTemplateVersions
type DescribeLaunchTemplateVersionsResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	TotalCount                int                       `json:"TotalCount" xml:"TotalCount"`
	PageNumber                int                       `json:"PageNumber" xml:"PageNumber"`
	PageSize                  int                       `json:"PageSize" xml:"PageSize"`
	LaunchTemplateVersionSets LaunchTemplateVersionSets `json:"LaunchTemplateVersionSets" xml:"LaunchTemplateVersionSets"`
}

// CreateDescribeLaunchTemplateVersionsRequest creates a request to invoke DescribeLaunchTemplateVersions API
func CreateDescribeLaunchTemplateVersionsRequest() (request *DescribeLaunchTemplateVersionsRequest) {
	request = &DescribeLaunchTemplateVersionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeLaunchTemplateVersions", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLaunchTemplateVersionsResponse creates a response to parse from DescribeLaunchTemplateVersions response
func CreateDescribeLaunchTemplateVersionsResponse() (response *DescribeLaunchTemplateVersionsResponse) {
	response = &DescribeLaunchTemplateVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
