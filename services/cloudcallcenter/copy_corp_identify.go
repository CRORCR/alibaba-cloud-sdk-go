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

// CopyCorpIdentify invokes the cloudcallcenter.CopyCorpIdentify API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/copycorpidentify.html
func (client *Client) CopyCorpIdentify(request *CopyCorpIdentifyRequest) (response *CopyCorpIdentifyResponse, err error) {
	response = CreateCopyCorpIdentifyResponse()
	err = client.DoAction(request, response)
	return
}

// CopyCorpIdentifyWithChan invokes the cloudcallcenter.CopyCorpIdentify API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/copycorpidentify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyCorpIdentifyWithChan(request *CopyCorpIdentifyRequest) (<-chan *CopyCorpIdentifyResponse, <-chan error) {
	responseChan := make(chan *CopyCorpIdentifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyCorpIdentify(request)
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

// CopyCorpIdentifyWithCallback invokes the cloudcallcenter.CopyCorpIdentify API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/copycorpidentify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyCorpIdentifyWithCallback(request *CopyCorpIdentifyRequest, callback func(response *CopyCorpIdentifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyCorpIdentifyResponse
		var err error
		defer close(result)
		response, err = client.CopyCorpIdentify(request)
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

// CopyCorpIdentifyRequest is the request struct for api CopyCorpIdentify
type CopyCorpIdentifyRequest struct {
	*requests.RpcRequest
	OrderId            string `position:"Query" name:"OrderId"`
	RegionNameProvince string `position:"Query" name:"RegionNameProvince"`
	RealNameInsId      string `position:"Query" name:"RealNameInsId"`
	OffsiteCertPic     string `position:"Query" name:"OffsiteCertPic"`
	RegionNameCity     string `position:"Query" name:"RegionNameCity"`
}

// CopyCorpIdentifyResponse is the response struct for api CopyCorpIdentify
type CopyCorpIdentifyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateCopyCorpIdentifyRequest creates a request to invoke CopyCorpIdentify API
func CreateCopyCorpIdentifyRequest() (request *CopyCorpIdentifyRequest) {
	request = &CopyCorpIdentifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CopyCorpIdentify", "", "")
	request.Method = requests.POST
	return
}

// CreateCopyCorpIdentifyResponse creates a response to parse from CopyCorpIdentify response
func CreateCopyCorpIdentifyResponse() (response *CopyCorpIdentifyResponse) {
	response = &CopyCorpIdentifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
