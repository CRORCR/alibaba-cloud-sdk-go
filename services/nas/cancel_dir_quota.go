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

package nas

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// CancelDirQuota invokes the nas.CancelDirQuota API synchronously
// api document: https://help.aliyun.com/api/nas/canceldirquota.html
func (client *Client) CancelDirQuota(request *CancelDirQuotaRequest) (response *CancelDirQuotaResponse, err error) {
	response = CreateCancelDirQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// CancelDirQuotaWithChan invokes the nas.CancelDirQuota API asynchronously
// api document: https://help.aliyun.com/api/nas/canceldirquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelDirQuotaWithChan(request *CancelDirQuotaRequest) (<-chan *CancelDirQuotaResponse, <-chan error) {
	responseChan := make(chan *CancelDirQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelDirQuota(request)
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

// CancelDirQuotaWithCallback invokes the nas.CancelDirQuota API asynchronously
// api document: https://help.aliyun.com/api/nas/canceldirquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelDirQuotaWithCallback(request *CancelDirQuotaRequest, callback func(response *CancelDirQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelDirQuotaResponse
		var err error
		defer close(result)
		response, err = client.CancelDirQuota(request)
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

// CancelDirQuotaRequest is the request struct for api CancelDirQuota
type CancelDirQuotaRequest struct {
	*requests.RpcRequest
	FileSystemId string `position:"Query" name:"FileSystemId"`
	Path         string `position:"Query" name:"Path"`
	UserType     string `position:"Query" name:"UserType"`
	UserId       string `position:"Query" name:"UserId"`
}

// CancelDirQuotaResponse is the response struct for api CancelDirQuota
type CancelDirQuotaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCancelDirQuotaRequest creates a request to invoke CancelDirQuota API
func CreateCancelDirQuotaRequest() (request *CancelDirQuotaRequest) {
	request = &CancelDirQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CancelDirQuota", "nas", "openAPI")
	return
}

// CreateCancelDirQuotaResponse creates a response to parse from CancelDirQuota response
func CreateCancelDirQuotaResponse() (response *CancelDirQuotaResponse) {
	response = &CancelDirQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
