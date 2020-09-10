package smartag

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

// DescribeSagLanList invokes the smartag.DescribeSagLanList API synchronously
// api document: https://help.aliyun.com/api/smartag/describesaglanlist.html
func (client *Client) DescribeSagLanList(request *DescribeSagLanListRequest) (response *DescribeSagLanListResponse, err error) {
	response = CreateDescribeSagLanListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagLanListWithChan invokes the smartag.DescribeSagLanList API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesaglanlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagLanListWithChan(request *DescribeSagLanListRequest) (<-chan *DescribeSagLanListResponse, <-chan error) {
	responseChan := make(chan *DescribeSagLanListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagLanList(request)
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

// DescribeSagLanListWithCallback invokes the smartag.DescribeSagLanList API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesaglanlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagLanListWithCallback(request *DescribeSagLanListRequest, callback func(response *DescribeSagLanListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagLanListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagLanList(request)
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

// DescribeSagLanListRequest is the request struct for api DescribeSagLanList
type DescribeSagLanListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagLanListResponse is the response struct for api DescribeSagLanList
type DescribeSagLanListResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Lans       []Lan       `json:"Lans" xml:"Lans"`
	TaskStates []TaskState `json:"TaskStates" xml:"TaskStates"`
}

// CreateDescribeSagLanListRequest creates a request to invoke DescribeSagLanList API
func CreateDescribeSagLanListRequest() (request *DescribeSagLanListRequest) {
	request = &DescribeSagLanListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagLanList", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagLanListResponse creates a response to parse from DescribeSagLanList response
func CreateDescribeSagLanListResponse() (response *DescribeSagLanListResponse) {
	response = &DescribeSagLanListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
