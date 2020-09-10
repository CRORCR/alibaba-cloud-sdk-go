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

// DescribeFlowCategory invokes the emr.DescribeFlowCategory API synchronously
// api document: https://help.aliyun.com/api/emr/describeflowcategory.html
func (client *Client) DescribeFlowCategory(request *DescribeFlowCategoryRequest) (response *DescribeFlowCategoryResponse, err error) {
	response = CreateDescribeFlowCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowCategoryWithChan invokes the emr.DescribeFlowCategory API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflowcategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowCategoryWithChan(request *DescribeFlowCategoryRequest) (<-chan *DescribeFlowCategoryResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowCategory(request)
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

// DescribeFlowCategoryWithCallback invokes the emr.DescribeFlowCategory API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflowcategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowCategoryWithCallback(request *DescribeFlowCategoryRequest, callback func(response *DescribeFlowCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowCategoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowCategory(request)
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

// DescribeFlowCategoryRequest is the request struct for api DescribeFlowCategory
type DescribeFlowCategoryRequest struct {
	*requests.RpcRequest
	Id        string `position:"Query" name:"Id"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DescribeFlowCategoryResponse is the response struct for api DescribeFlowCategory
type DescribeFlowCategoryResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Id           string `json:"Id" xml:"Id"`
	GmtCreate    int64  `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified  int64  `json:"GmtModified" xml:"GmtModified"`
	Name         string `json:"Name" xml:"Name"`
	ParentId     string `json:"ParentId" xml:"ParentId"`
	Type         string `json:"Type" xml:"Type"`
	CategoryType string `json:"CategoryType" xml:"CategoryType"`
	ObjectType   string `json:"ObjectType" xml:"ObjectType"`
	ObjectId     string `json:"ObjectId" xml:"ObjectId"`
	ProjectId    string `json:"ProjectId" xml:"ProjectId"`
}

// CreateDescribeFlowCategoryRequest creates a request to invoke DescribeFlowCategory API
func CreateDescribeFlowCategoryRequest() (request *DescribeFlowCategoryRequest) {
	request = &DescribeFlowCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowCategory", "emr", "openAPI")
	return
}

// CreateDescribeFlowCategoryResponse creates a response to parse from DescribeFlowCategory response
func CreateDescribeFlowCategoryResponse() (response *DescribeFlowCategoryResponse) {
	response = &DescribeFlowCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
