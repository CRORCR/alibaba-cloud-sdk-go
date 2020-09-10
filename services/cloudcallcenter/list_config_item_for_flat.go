package cloudcallcenter

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

// ListConfigItemForFlat invokes the cloudcallcenter.ListConfigItemForFlat API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listconfigitemforflat.html
func (client *Client) ListConfigItemForFlat(request *ListConfigItemForFlatRequest) (response *ListConfigItemForFlatResponse, err error) {
	response = CreateListConfigItemForFlatResponse()
	err = client.DoAction(request, response)
	return
}

// ListConfigItemForFlatWithChan invokes the cloudcallcenter.ListConfigItemForFlat API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listconfigitemforflat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConfigItemForFlatWithChan(request *ListConfigItemForFlatRequest) (<-chan *ListConfigItemForFlatResponse, <-chan error) {
	responseChan := make(chan *ListConfigItemForFlatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConfigItemForFlat(request)
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

// ListConfigItemForFlatWithCallback invokes the cloudcallcenter.ListConfigItemForFlat API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listconfigitemforflat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConfigItemForFlatWithCallback(request *ListConfigItemForFlatRequest, callback func(response *ListConfigItemForFlatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConfigItemForFlatResponse
		var err error
		defer close(result)
		response, err = client.ListConfigItemForFlat(request)
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

// ListConfigItemForFlatRequest is the request struct for api ListConfigItemForFlat
type ListConfigItemForFlatRequest struct {
	*requests.RpcRequest
	InstanceId string    `position:"Query" name:"InstanceId"`
	ConfigItem *[]string `position:"Query" name:"ConfigItem"  type:"Repeated"`
}

// ListConfigItemForFlatResponse is the response struct for api ListConfigItemForFlat
type ListConfigItemForFlatResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ConfigItems    ConfigItems `json:"ConfigItems" xml:"ConfigItems"`
}

// CreateListConfigItemForFlatRequest creates a request to invoke ListConfigItemForFlat API
func CreateListConfigItemForFlatRequest() (request *ListConfigItemForFlatRequest) {
	request = &ListConfigItemForFlatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListConfigItemForFlat", "", "")
	request.Method = requests.POST
	return
}

// CreateListConfigItemForFlatResponse creates a response to parse from ListConfigItemForFlat response
func CreateListConfigItemForFlatResponse() (response *ListConfigItemForFlatResponse) {
	response = &ListConfigItemForFlatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
