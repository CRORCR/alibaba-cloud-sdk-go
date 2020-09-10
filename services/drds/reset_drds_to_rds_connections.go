package drds

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

// ResetDrdsToRdsConnections invokes the drds.ResetDrdsToRdsConnections API synchronously
// api document: https://help.aliyun.com/api/drds/resetdrdstordsconnections.html
func (client *Client) ResetDrdsToRdsConnections(request *ResetDrdsToRdsConnectionsRequest) (response *ResetDrdsToRdsConnectionsResponse, err error) {
	response = CreateResetDrdsToRdsConnectionsResponse()
	err = client.DoAction(request, response)
	return
}

// ResetDrdsToRdsConnectionsWithChan invokes the drds.ResetDrdsToRdsConnections API asynchronously
// api document: https://help.aliyun.com/api/drds/resetdrdstordsconnections.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetDrdsToRdsConnectionsWithChan(request *ResetDrdsToRdsConnectionsRequest) (<-chan *ResetDrdsToRdsConnectionsResponse, <-chan error) {
	responseChan := make(chan *ResetDrdsToRdsConnectionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetDrdsToRdsConnections(request)
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

// ResetDrdsToRdsConnectionsWithCallback invokes the drds.ResetDrdsToRdsConnections API asynchronously
// api document: https://help.aliyun.com/api/drds/resetdrdstordsconnections.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetDrdsToRdsConnectionsWithCallback(request *ResetDrdsToRdsConnectionsRequest, callback func(response *ResetDrdsToRdsConnectionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetDrdsToRdsConnectionsResponse
		var err error
		defer close(result)
		response, err = client.ResetDrdsToRdsConnections(request)
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

// ResetDrdsToRdsConnectionsRequest is the request struct for api ResetDrdsToRdsConnections
type ResetDrdsToRdsConnectionsRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// ResetDrdsToRdsConnectionsResponse is the response struct for api ResetDrdsToRdsConnections
type ResetDrdsToRdsConnectionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateResetDrdsToRdsConnectionsRequest creates a request to invoke ResetDrdsToRdsConnections API
func CreateResetDrdsToRdsConnectionsRequest() (request *ResetDrdsToRdsConnectionsRequest) {
	request = &ResetDrdsToRdsConnectionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "ResetDrdsToRdsConnections", "Drds", "openAPI")
	return
}

// CreateResetDrdsToRdsConnectionsResponse creates a response to parse from ResetDrdsToRdsConnections response
func CreateResetDrdsToRdsConnectionsResponse() (response *ResetDrdsToRdsConnectionsResponse) {
	response = &ResetDrdsToRdsConnectionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
