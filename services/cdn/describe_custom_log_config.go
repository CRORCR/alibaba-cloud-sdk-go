package cdn

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

// DescribeCustomLogConfig invokes the cdn.DescribeCustomLogConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/describecustomlogconfig.html
func (client *Client) DescribeCustomLogConfig(request *DescribeCustomLogConfigRequest) (response *DescribeCustomLogConfigResponse, err error) {
	response = CreateDescribeCustomLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomLogConfigWithChan invokes the cdn.DescribeCustomLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecustomlogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomLogConfigWithChan(request *DescribeCustomLogConfigRequest) (<-chan *DescribeCustomLogConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomLogConfig(request)
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

// DescribeCustomLogConfigWithCallback invokes the cdn.DescribeCustomLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecustomlogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomLogConfigWithCallback(request *DescribeCustomLogConfigRequest, callback func(response *DescribeCustomLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomLogConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomLogConfig(request)
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

// DescribeCustomLogConfigRequest is the request struct for api DescribeCustomLogConfig
type DescribeCustomLogConfigRequest struct {
	*requests.RpcRequest
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	ConfigId string           `position:"Query" name:"ConfigId"`
}

// DescribeCustomLogConfigResponse is the response struct for api DescribeCustomLogConfig
type DescribeCustomLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Remark    string `json:"Remark" xml:"Remark"`
	Sample    string `json:"Sample" xml:"Sample"`
	Tag       string `json:"Tag" xml:"Tag"`
}

// CreateDescribeCustomLogConfigRequest creates a request to invoke DescribeCustomLogConfig API
func CreateDescribeCustomLogConfigRequest() (request *DescribeCustomLogConfigRequest) {
	request = &DescribeCustomLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCustomLogConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeCustomLogConfigResponse creates a response to parse from DescribeCustomLogConfig response
func CreateDescribeCustomLogConfigResponse() (response *DescribeCustomLogConfigResponse) {
	response = &DescribeCustomLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
