package reid

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

// ListPersonByImage invokes the reid.ListPersonByImage API synchronously
// api document: https://help.aliyun.com/api/reid/listpersonbyimage.html
func (client *Client) ListPersonByImage(request *ListPersonByImageRequest) (response *ListPersonByImageResponse, err error) {
	response = CreateListPersonByImageResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersonByImageWithChan invokes the reid.ListPersonByImage API asynchronously
// api document: https://help.aliyun.com/api/reid/listpersonbyimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPersonByImageWithChan(request *ListPersonByImageRequest) (<-chan *ListPersonByImageResponse, <-chan error) {
	responseChan := make(chan *ListPersonByImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersonByImage(request)
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

// ListPersonByImageWithCallback invokes the reid.ListPersonByImage API asynchronously
// api document: https://help.aliyun.com/api/reid/listpersonbyimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPersonByImageWithCallback(request *ListPersonByImageRequest, callback func(response *ListPersonByImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersonByImageResponse
		var err error
		defer close(result)
		response, err = client.ListPersonByImage(request)
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

// ListPersonByImageRequest is the request struct for api ListPersonByImage
type ListPersonByImageRequest struct {
	*requests.RpcRequest
	StoreId  requests.Integer `position:"Body" name:"StoreId"`
	ImageUrl string           `position:"Body" name:"ImageUrl"`
}

// ListPersonByImageResponse is the response struct for api ListPersonByImage
type ListPersonByImageResponse struct {
	*responses.BaseResponse
	ErrorCode               string                  `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage            string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	Success                 bool                    `json:"Success" xml:"Success"`
	PersonSearchResultItems PersonSearchResultItems `json:"PersonSearchResultItems" xml:"PersonSearchResultItems"`
}

// CreateListPersonByImageRequest creates a request to invoke ListPersonByImage API
func CreateListPersonByImageRequest() (request *ListPersonByImageRequest) {
	request = &ListPersonByImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ListPersonByImage", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPersonByImageResponse creates a response to parse from ListPersonByImage response
func CreateListPersonByImageResponse() (response *ListPersonByImageResponse) {
	response = &ListPersonByImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
