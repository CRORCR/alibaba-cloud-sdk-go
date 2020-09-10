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

// CancelClusterUpgrade invokes the cs.CancelClusterUpgrade API synchronously
// api document: https://help.aliyun.com/api/cs/cancelclusterupgrade.html
func (client *Client) CancelClusterUpgrade(request *CancelClusterUpgradeRequest) (response *CancelClusterUpgradeResponse, err error) {
	response = CreateCancelClusterUpgradeResponse()
	err = client.DoAction(request, response)
	return
}

// CancelClusterUpgradeWithChan invokes the cs.CancelClusterUpgrade API asynchronously
// api document: https://help.aliyun.com/api/cs/cancelclusterupgrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelClusterUpgradeWithChan(request *CancelClusterUpgradeRequest) (<-chan *CancelClusterUpgradeResponse, <-chan error) {
	responseChan := make(chan *CancelClusterUpgradeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelClusterUpgrade(request)
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

// CancelClusterUpgradeWithCallback invokes the cs.CancelClusterUpgrade API asynchronously
// api document: https://help.aliyun.com/api/cs/cancelclusterupgrade.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelClusterUpgradeWithCallback(request *CancelClusterUpgradeRequest, callback func(response *CancelClusterUpgradeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelClusterUpgradeResponse
		var err error
		defer close(result)
		response, err = client.CancelClusterUpgrade(request)
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

// CancelClusterUpgradeRequest is the request struct for api CancelClusterUpgrade
type CancelClusterUpgradeRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// CancelClusterUpgradeResponse is the response struct for api CancelClusterUpgrade
type CancelClusterUpgradeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelClusterUpgradeRequest creates a request to invoke CancelClusterUpgrade API
func CreateCancelClusterUpgradeRequest() (request *CancelClusterUpgradeRequest) {
	request = &CancelClusterUpgradeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "CancelClusterUpgrade", "/api/v2/clusters/[ClusterId]/upgrade/cancel", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelClusterUpgradeResponse creates a response to parse from CancelClusterUpgrade response
func CreateCancelClusterUpgradeResponse() (response *CancelClusterUpgradeResponse) {
	response = &CancelClusterUpgradeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
