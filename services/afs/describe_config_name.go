package afs

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

// DescribeConfigName invokes the afs.DescribeConfigName API synchronously
// api document: https://help.aliyun.com/api/afs/describeconfigname.html
func (client *Client) DescribeConfigName(request *DescribeConfigNameRequest) (response *DescribeConfigNameResponse, err error) {
	response = CreateDescribeConfigNameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigNameWithChan invokes the afs.DescribeConfigName API asynchronously
// api document: https://help.aliyun.com/api/afs/describeconfigname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigNameWithChan(request *DescribeConfigNameRequest) (<-chan *DescribeConfigNameResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigName(request)
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

// DescribeConfigNameWithCallback invokes the afs.DescribeConfigName API asynchronously
// api document: https://help.aliyun.com/api/afs/describeconfigname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConfigNameWithCallback(request *DescribeConfigNameRequest, callback func(response *DescribeConfigNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigNameResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigName(request)
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

// DescribeConfigNameRequest is the request struct for api DescribeConfigName
type DescribeConfigNameRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeConfigNameResponse is the response struct for api DescribeConfigName
type DescribeConfigNameResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	HasConfig   bool         `json:"HasConfig" xml:"HasConfig"`
	BizCode     string       `json:"BizCode" xml:"BizCode"`
	ConfigNames []ConfigName `json:"ConfigNames" xml:"ConfigNames"`
}

// CreateDescribeConfigNameRequest creates a request to invoke DescribeConfigName API
func CreateDescribeConfigNameRequest() (request *DescribeConfigNameRequest) {
	request = &DescribeConfigNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("afs", "2018-01-12", "DescribeConfigName", "afs", "openAPI")
	return
}

// CreateDescribeConfigNameResponse creates a response to parse from DescribeConfigName response
func CreateDescribeConfigNameResponse() (response *DescribeConfigNameResponse) {
	response = &DescribeConfigNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
