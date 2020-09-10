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

// RefreshObjectCaches invokes the cdn.RefreshObjectCaches API synchronously
// api document: https://help.aliyun.com/api/cdn/refreshobjectcaches.html
func (client *Client) RefreshObjectCaches(request *RefreshObjectCachesRequest) (response *RefreshObjectCachesResponse, err error) {
	response = CreateRefreshObjectCachesResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshObjectCachesWithChan invokes the cdn.RefreshObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/cdn/refreshobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshObjectCachesWithChan(request *RefreshObjectCachesRequest) (<-chan *RefreshObjectCachesResponse, <-chan error) {
	responseChan := make(chan *RefreshObjectCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshObjectCaches(request)
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

// RefreshObjectCachesWithCallback invokes the cdn.RefreshObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/cdn/refreshobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshObjectCachesWithCallback(request *RefreshObjectCachesRequest, callback func(response *RefreshObjectCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshObjectCachesResponse
		var err error
		defer close(result)
		response, err = client.RefreshObjectCaches(request)
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

// RefreshObjectCachesRequest is the request struct for api RefreshObjectCaches
type RefreshObjectCachesRequest struct {
	*requests.RpcRequest
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	ObjectType    string           `position:"Query" name:"ObjectType"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// RefreshObjectCachesResponse is the response struct for api RefreshObjectCaches
type RefreshObjectCachesResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	RefreshTaskId string `json:"RefreshTaskId" xml:"RefreshTaskId"`
}

// CreateRefreshObjectCachesRequest creates a request to invoke RefreshObjectCaches API
func CreateRefreshObjectCachesRequest() (request *RefreshObjectCachesRequest) {
	request = &RefreshObjectCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "RefreshObjectCaches", "", "")
	request.Method = requests.POST
	return
}

// CreateRefreshObjectCachesResponse creates a response to parse from RefreshObjectCaches response
func CreateRefreshObjectCachesResponse() (response *RefreshObjectCachesResponse) {
	response = &RefreshObjectCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
