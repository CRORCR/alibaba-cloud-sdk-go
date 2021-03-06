package green

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

// DescribeWebsiteInstanceKeyUrl invokes the green.DescribeWebsiteInstanceKeyUrl API synchronously
// api document: https://help.aliyun.com/api/green/describewebsiteinstancekeyurl.html
func (client *Client) DescribeWebsiteInstanceKeyUrl(request *DescribeWebsiteInstanceKeyUrlRequest) (response *DescribeWebsiteInstanceKeyUrlResponse, err error) {
	response = CreateDescribeWebsiteInstanceKeyUrlResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebsiteInstanceKeyUrlWithChan invokes the green.DescribeWebsiteInstanceKeyUrl API asynchronously
// api document: https://help.aliyun.com/api/green/describewebsiteinstancekeyurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebsiteInstanceKeyUrlWithChan(request *DescribeWebsiteInstanceKeyUrlRequest) (<-chan *DescribeWebsiteInstanceKeyUrlResponse, <-chan error) {
	responseChan := make(chan *DescribeWebsiteInstanceKeyUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebsiteInstanceKeyUrl(request)
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

// DescribeWebsiteInstanceKeyUrlWithCallback invokes the green.DescribeWebsiteInstanceKeyUrl API asynchronously
// api document: https://help.aliyun.com/api/green/describewebsiteinstancekeyurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebsiteInstanceKeyUrlWithCallback(request *DescribeWebsiteInstanceKeyUrlRequest, callback func(response *DescribeWebsiteInstanceKeyUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebsiteInstanceKeyUrlResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebsiteInstanceKeyUrl(request)
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

// DescribeWebsiteInstanceKeyUrlRequest is the request struct for api DescribeWebsiteInstanceKeyUrl
type DescribeWebsiteInstanceKeyUrlRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeWebsiteInstanceKeyUrlResponse is the response struct for api DescribeWebsiteInstanceKeyUrl
type DescribeWebsiteInstanceKeyUrlResponse struct {
	*responses.BaseResponse
	RequestId                 string   `json:"RequestId" xml:"RequestId"`
	TotalCount                int      `json:"TotalCount" xml:"TotalCount"`
	WebsiteInstanceKeyUrlList []string `json:"WebsiteInstanceKeyUrlList" xml:"WebsiteInstanceKeyUrlList"`
}

// CreateDescribeWebsiteInstanceKeyUrlRequest creates a request to invoke DescribeWebsiteInstanceKeyUrl API
func CreateDescribeWebsiteInstanceKeyUrlRequest() (request *DescribeWebsiteInstanceKeyUrlRequest) {
	request = &DescribeWebsiteInstanceKeyUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeWebsiteInstanceKeyUrl", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWebsiteInstanceKeyUrlResponse creates a response to parse from DescribeWebsiteInstanceKeyUrl response
func CreateDescribeWebsiteInstanceKeyUrlResponse() (response *DescribeWebsiteInstanceKeyUrlResponse) {
	response = &DescribeWebsiteInstanceKeyUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
