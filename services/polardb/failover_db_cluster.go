package polardb

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

// FailoverDBCluster invokes the polardb.FailoverDBCluster API synchronously
// api document: https://help.aliyun.com/api/polardb/failoverdbcluster.html
func (client *Client) FailoverDBCluster(request *FailoverDBClusterRequest) (response *FailoverDBClusterResponse, err error) {
	response = CreateFailoverDBClusterResponse()
	err = client.DoAction(request, response)
	return
}

// FailoverDBClusterWithChan invokes the polardb.FailoverDBCluster API asynchronously
// api document: https://help.aliyun.com/api/polardb/failoverdbcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FailoverDBClusterWithChan(request *FailoverDBClusterRequest) (<-chan *FailoverDBClusterResponse, <-chan error) {
	responseChan := make(chan *FailoverDBClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FailoverDBCluster(request)
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

// FailoverDBClusterWithCallback invokes the polardb.FailoverDBCluster API asynchronously
// api document: https://help.aliyun.com/api/polardb/failoverdbcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FailoverDBClusterWithCallback(request *FailoverDBClusterRequest, callback func(response *FailoverDBClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FailoverDBClusterResponse
		var err error
		defer close(result)
		response, err = client.FailoverDBCluster(request)
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

// FailoverDBClusterRequest is the request struct for api FailoverDBCluster
type FailoverDBClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TargetDBNodeId       string           `position:"Query" name:"TargetDBNodeId"`
}

// FailoverDBClusterResponse is the response struct for api FailoverDBCluster
type FailoverDBClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateFailoverDBClusterRequest creates a request to invoke FailoverDBCluster API
func CreateFailoverDBClusterRequest() (request *FailoverDBClusterRequest) {
	request = &FailoverDBClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "FailoverDBCluster", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFailoverDBClusterResponse creates a response to parse from FailoverDBCluster response
func CreateFailoverDBClusterResponse() (response *FailoverDBClusterResponse) {
	response = &FailoverDBClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
