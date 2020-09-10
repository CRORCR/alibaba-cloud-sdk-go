package ddoscoo

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

// DescribeSceneDefenseObjects invokes the ddoscoo.DescribeSceneDefenseObjects API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describescenedefenseobjects.html
func (client *Client) DescribeSceneDefenseObjects(request *DescribeSceneDefenseObjectsRequest) (response *DescribeSceneDefenseObjectsResponse, err error) {
	response = CreateDescribeSceneDefenseObjectsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSceneDefenseObjectsWithChan invokes the ddoscoo.DescribeSceneDefenseObjects API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describescenedefenseobjects.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSceneDefenseObjectsWithChan(request *DescribeSceneDefenseObjectsRequest) (<-chan *DescribeSceneDefenseObjectsResponse, <-chan error) {
	responseChan := make(chan *DescribeSceneDefenseObjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSceneDefenseObjects(request)
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

// DescribeSceneDefenseObjectsWithCallback invokes the ddoscoo.DescribeSceneDefenseObjects API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describescenedefenseobjects.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSceneDefenseObjectsWithCallback(request *DescribeSceneDefenseObjectsRequest, callback func(response *DescribeSceneDefenseObjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSceneDefenseObjectsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSceneDefenseObjects(request)
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

// DescribeSceneDefenseObjectsRequest is the request struct for api DescribeSceneDefenseObjects
type DescribeSceneDefenseObjectsRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	PolicyId        string `position:"Query" name:"PolicyId"`
}

// DescribeSceneDefenseObjectsResponse is the response struct for api DescribeSceneDefenseObjects
type DescribeSceneDefenseObjectsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	Objects   []Object `json:"Objects" xml:"Objects"`
}

// CreateDescribeSceneDefenseObjectsRequest creates a request to invoke DescribeSceneDefenseObjects API
func CreateDescribeSceneDefenseObjectsRequest() (request *DescribeSceneDefenseObjectsRequest) {
	request = &DescribeSceneDefenseObjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeSceneDefenseObjects", "ddoscoo", "openAPI")
	return
}

// CreateDescribeSceneDefenseObjectsResponse creates a response to parse from DescribeSceneDefenseObjects response
func CreateDescribeSceneDefenseObjectsResponse() (response *DescribeSceneDefenseObjectsResponse) {
	response = &DescribeSceneDefenseObjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
