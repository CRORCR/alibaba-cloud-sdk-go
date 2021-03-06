package smc

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

// CreateReplicationJob invokes the smc.CreateReplicationJob API synchronously
// api document: https://help.aliyun.com/api/smc/createreplicationjob.html
func (client *Client) CreateReplicationJob(request *CreateReplicationJobRequest) (response *CreateReplicationJobResponse, err error) {
	response = CreateCreateReplicationJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateReplicationJobWithChan invokes the smc.CreateReplicationJob API asynchronously
// api document: https://help.aliyun.com/api/smc/createreplicationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateReplicationJobWithChan(request *CreateReplicationJobRequest) (<-chan *CreateReplicationJobResponse, <-chan error) {
	responseChan := make(chan *CreateReplicationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateReplicationJob(request)
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

// CreateReplicationJobWithCallback invokes the smc.CreateReplicationJob API asynchronously
// api document: https://help.aliyun.com/api/smc/createreplicationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateReplicationJobWithCallback(request *CreateReplicationJobRequest, callback func(response *CreateReplicationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateReplicationJobResponse
		var err error
		defer close(result)
		response, err = client.CreateReplicationJob(request)
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

// CreateReplicationJobRequest is the request struct for api CreateReplicationJob
type CreateReplicationJobRequest struct {
	*requests.RpcRequest
	TargetType             string                          `position:"Query" name:"TargetType"`
	ClientToken            string                          `position:"Query" name:"ClientToken"`
	Description            string                          `position:"Query" name:"Description"`
	Frequency              requests.Integer                `position:"Query" name:"Frequency"`
	ReplicationParameters  string                          `position:"Query" name:"ReplicationParameters"`
	ImageName              string                          `position:"Query" name:"ImageName"`
	SystemDiskSize         requests.Integer                `position:"Query" name:"SystemDiskSize"`
	InstanceType           string                          `position:"Query" name:"InstanceType"`
	Tag                    *[]CreateReplicationJobTag      `position:"Query" name:"Tag"  type:"Repeated"`
	NetMode                requests.Integer                `position:"Query" name:"NetMode"`
	SourceId               string                          `position:"Query" name:"SourceId"`
	RunOnce                requests.Boolean                `position:"Query" name:"RunOnce"`
	ResourceOwnerAccount   string                          `position:"Query" name:"ResourceOwnerAccount"`
	ValidTime              string                          `position:"Query" name:"ValidTime"`
	OwnerId                requests.Integer                `position:"Query" name:"OwnerId"`
	DataDisk               *[]CreateReplicationJobDataDisk `position:"Query" name:"DataDisk"  type:"Repeated"`
	VSwitchId              string                          `position:"Query" name:"VSwitchId"`
	ScheduledStartTime     string                          `position:"Query" name:"ScheduledStartTime"`
	InstanceId             string                          `position:"Query" name:"InstanceId"`
	VpcId                  string                          `position:"Query" name:"VpcId"`
	Name                   string                          `position:"Query" name:"Name"`
	MaxNumberOfImageToKeep requests.Integer                `position:"Query" name:"MaxNumberOfImageToKeep"`
}

// CreateReplicationJobTag is a repeated param struct in CreateReplicationJobRequest
type CreateReplicationJobTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateReplicationJobDataDisk is a repeated param struct in CreateReplicationJobRequest
type CreateReplicationJobDataDisk struct {
	Size  string `name:"Size"`
	Index string `name:"Index"`
}

// CreateReplicationJobResponse is the response struct for api CreateReplicationJob
type CreateReplicationJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateCreateReplicationJobRequest creates a request to invoke CreateReplicationJob API
func CreateCreateReplicationJobRequest() (request *CreateReplicationJobRequest) {
	request = &CreateReplicationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("smc", "2019-06-01", "CreateReplicationJob", "smc", "openAPI")
	return
}

// CreateCreateReplicationJobResponse creates a response to parse from CreateReplicationJob response
func CreateCreateReplicationJobResponse() (response *CreateReplicationJobResponse) {
	response = &CreateReplicationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
