package live

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

// EditHtmlResource invokes the live.EditHtmlResource API synchronously
// api document: https://help.aliyun.com/api/live/edithtmlresource.html
func (client *Client) EditHtmlResource(request *EditHtmlResourceRequest) (response *EditHtmlResourceResponse, err error) {
	response = CreateEditHtmlResourceResponse()
	err = client.DoAction(request, response)
	return
}

// EditHtmlResourceWithChan invokes the live.EditHtmlResource API asynchronously
// api document: https://help.aliyun.com/api/live/edithtmlresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EditHtmlResourceWithChan(request *EditHtmlResourceRequest) (<-chan *EditHtmlResourceResponse, <-chan error) {
	responseChan := make(chan *EditHtmlResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditHtmlResource(request)
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

// EditHtmlResourceWithCallback invokes the live.EditHtmlResource API asynchronously
// api document: https://help.aliyun.com/api/live/edithtmlresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EditHtmlResourceWithCallback(request *EditHtmlResourceRequest, callback func(response *EditHtmlResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditHtmlResourceResponse
		var err error
		defer close(result)
		response, err = client.EditHtmlResource(request)
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

// EditHtmlResourceRequest is the request struct for api EditHtmlResource
type EditHtmlResourceRequest struct {
	*requests.RpcRequest
	HtmlUrl        string           `position:"Query" name:"HtmlUrl"`
	CasterId       string           `position:"Query" name:"CasterId"`
	HtmlContent    string           `position:"Query" name:"htmlContent"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	HtmlResourceId string           `position:"Query" name:"HtmlResourceId"`
	Config         string           `position:"Query" name:"Config"`
}

// EditHtmlResourceResponse is the response struct for api EditHtmlResource
type EditHtmlResourceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HtmlResourceId string `json:"HtmlResourceId" xml:"HtmlResourceId"`
}

// CreateEditHtmlResourceRequest creates a request to invoke EditHtmlResource API
func CreateEditHtmlResourceRequest() (request *EditHtmlResourceRequest) {
	request = &EditHtmlResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "EditHtmlResource", "", "")
	request.Method = requests.POST
	return
}

// CreateEditHtmlResourceResponse creates a response to parse from EditHtmlResource response
func CreateEditHtmlResourceResponse() (response *EditHtmlResourceResponse) {
	response = &EditHtmlResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
