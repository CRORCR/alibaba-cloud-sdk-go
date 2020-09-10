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

package dts

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateSynchronizationJob invokes the dts.CreateSynchronizationJob API synchronously
// api document: https://help.aliyun.com/api/dts/createsynchronizationjob.html
func (client *Client) CreateSynchronizationJob(request *CreateSynchronizationJobRequest) (response *CreateSynchronizationJobResponse, err error) {
	response = CreateCreateSynchronizationJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSynchronizationJobWithChan invokes the dts.CreateSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/createsynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSynchronizationJobWithChan(request *CreateSynchronizationJobRequest) (<-chan *CreateSynchronizationJobResponse, <-chan error) {
	responseChan := make(chan *CreateSynchronizationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSynchronizationJob(request)
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

// CreateSynchronizationJobWithCallback invokes the dts.CreateSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/createsynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSynchronizationJobWithCallback(request *CreateSynchronizationJobRequest, callback func(response *CreateSynchronizationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSynchronizationJobResponse
		var err error
		defer close(result)
		response, err = client.CreateSynchronizationJob(request)
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

// CreateSynchronizationJobRequest is the request struct for api CreateSynchronizationJob
type CreateSynchronizationJobRequest struct {
	*requests.RpcRequest
	SourceRegion            string                                      `position:"Query" name:"SourceRegion"`
	DestRegion              string                                      `position:"Query" name:"DestRegion"`
	Topology                string                                      `position:"Query" name:"Topology"`
	SynchronizationJobClass string                                      `position:"Query" name:"SynchronizationJobClass"`
	PayType                 string                                      `position:"Query" name:"PayType"`
	Period                  string                                      `position:"Query" name:"Period"`
	UsedTime                requests.Integer                            `position:"Query" name:"UsedTime"`
	ClientToken             string                                      `position:"Query" name:"ClientToken"`
	NetworkType             string                                      `position:"Query" name:"NetworkType"`
	OwnerId                 string                                      `position:"Query" name:"OwnerId"`
	SourceEndpoint          CreateSynchronizationJobSourceEndpoint      `position:"Query" name:"SourceEndpoint" type:"Struct"`
	DestinationEndpoint     CreateSynchronizationJobDestinationEndpoint `position:"Query" name:"DestinationEndpoint" type:"Struct"`
}

type CreateSynchronizationJobSourceEndpoint struct {
	InstanceType string `name:"InstanceType"`
}

type CreateSynchronizationJobDestinationEndpoint struct {
	InstanceType string `name:"InstanceType"`
}

// CreateSynchronizationJobResponse is the response struct for api CreateSynchronizationJob
type CreateSynchronizationJobResponse struct {
	*responses.BaseResponse
	Success              string `json:"Success" xml:"Success"`
	ErrCode              string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage           string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId            string `json:"RequestId" xml:"RequestId"`
	SynchronizationJobId string `json:"SynchronizationJobId" xml:"SynchronizationJobId"`
}

// CreateCreateSynchronizationJobRequest creates a request to invoke CreateSynchronizationJob API
func CreateCreateSynchronizationJobRequest() (request *CreateSynchronizationJobRequest) {
	request = &CreateSynchronizationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "CreateSynchronizationJob", "dts", "openAPI")
	return
}

// CreateCreateSynchronizationJobResponse creates a response to parse from CreateSynchronizationJob response
func CreateCreateSynchronizationJobResponse() (response *CreateSynchronizationJobResponse) {
	response = &CreateSynchronizationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
