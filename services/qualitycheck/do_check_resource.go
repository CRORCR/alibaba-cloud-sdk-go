package qualitycheck

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

// DoCheckResource invokes the qualitycheck.DoCheckResource API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/docheckresource.html
func (client *Client) DoCheckResource(request *DoCheckResourceRequest) (response *DoCheckResourceResponse, err error) {
	response = CreateDoCheckResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DoCheckResourceWithChan invokes the qualitycheck.DoCheckResource API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/docheckresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DoCheckResourceWithChan(request *DoCheckResourceRequest) (<-chan *DoCheckResourceResponse, <-chan error) {
	responseChan := make(chan *DoCheckResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DoCheckResource(request)
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

// DoCheckResourceWithCallback invokes the qualitycheck.DoCheckResource API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/docheckresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DoCheckResourceWithCallback(request *DoCheckResourceRequest, callback func(response *DoCheckResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DoCheckResourceResponse
		var err error
		defer close(result)
		response, err = client.DoCheckResource(request)
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

// DoCheckResourceRequest is the request struct for api DoCheckResource
type DoCheckResourceRequest struct {
	*requests.RpcRequest
	Country         string           `position:"Query" name:"Country"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Hid             requests.Integer `position:"Query" name:"Hid"`
	Interrupt       requests.Boolean `position:"Query" name:"Interrupt"`
	GmtWakeup       string           `position:"Query" name:"GmtWakeup"`
	TaskExtraData   string           `position:"Query" name:"TaskExtraData"`
	Level           requests.Integer `position:"Query" name:"Level"`
	Message         string           `position:"Query" name:"Message"`
	Success         requests.Boolean `position:"Query" name:"Success"`
	Pk              string           `position:"Query" name:"Pk"`
	Bid             string           `position:"Query" name:"Bid"`
	Prompt          string           `position:"Query" name:"Prompt"`
	TaskIdentifier  string           `position:"Query" name:"TaskIdentifier"`
}

// DoCheckResourceResponse is the response struct for api DoCheckResource
type DoCheckResourceResponse struct {
	*responses.BaseResponse
	Interrupt      string `json:"Interrupt" xml:"Interrupt"`
	Invoker        int64  `json:"Invoker" xml:"Invoker"`
	Pk             string `json:"Pk" xml:"Pk"`
	Bid            string `json:"Bid" xml:"Bid"`
	Hid            int64  `json:"Hid" xml:"Hid"`
	Country        string `json:"Country" xml:"Country"`
	TaskIdentifier string `json:"TaskIdentifier" xml:"TaskIdentifier"`
	GmtWakeup      string `json:"GmtWakeup" xml:"GmtWakeup"`
	Success        bool   `json:"Success" xml:"Success"`
	Message        string `json:"Message" xml:"Message"`
	Level          int64  `json:"Level" xml:"Level"`
	Url            string `json:"Url" xml:"Url"`
	Prompt         string `json:"Prompt" xml:"Prompt"`
}

// CreateDoCheckResourceRequest creates a request to invoke DoCheckResource API
func CreateDoCheckResourceRequest() (request *DoCheckResourceRequest) {
	request = &DoCheckResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DoCheckResource", "", "")
	return
}

// CreateDoCheckResourceResponse creates a response to parse from DoCheckResource response
func CreateDoCheckResourceResponse() (response *DoCheckResourceResponse) {
	response = &DoCheckResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
