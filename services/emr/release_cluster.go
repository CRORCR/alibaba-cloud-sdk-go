package emr

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

// ReleaseCluster invokes the emr.ReleaseCluster API synchronously
// api document: https://help.aliyun.com/api/emr/releasecluster.html
func (client *Client) ReleaseCluster(request *ReleaseClusterRequest) (response *ReleaseClusterResponse, err error) {
	response = CreateReleaseClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseClusterWithChan invokes the emr.ReleaseCluster API asynchronously
// api document: https://help.aliyun.com/api/emr/releasecluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseClusterWithChan(request *ReleaseClusterRequest) (<-chan *ReleaseClusterResponse, <-chan error) {
	responseChan := make(chan *ReleaseClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseCluster(request)
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

// ReleaseClusterWithCallback invokes the emr.ReleaseCluster API asynchronously
// api document: https://help.aliyun.com/api/emr/releasecluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseClusterWithCallback(request *ReleaseClusterRequest, callback func(response *ReleaseClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseClusterResponse
		var err error
		defer close(result)
		response, err = client.ReleaseCluster(request)
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

// ReleaseClusterRequest is the request struct for api ReleaseCluster
type ReleaseClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ForceRelease    requests.Boolean `position:"Query" name:"ForceRelease"`
	Id              string           `position:"Query" name:"Id"`
}

// ReleaseClusterResponse is the response struct for api ReleaseCluster
type ReleaseClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseClusterRequest creates a request to invoke ReleaseCluster API
func CreateReleaseClusterRequest() (request *ReleaseClusterRequest) {
	request = &ReleaseClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ReleaseCluster", "emr", "openAPI")
	return
}

// CreateReleaseClusterResponse creates a response to parse from ReleaseCluster response
func CreateReleaseClusterResponse() (response *ReleaseClusterResponse) {
	response = &ReleaseClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
