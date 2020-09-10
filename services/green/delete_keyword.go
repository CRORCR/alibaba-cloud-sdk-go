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

// DeleteKeyword invokes the green.DeleteKeyword API synchronously
// api document: https://help.aliyun.com/api/green/deletekeyword.html
func (client *Client) DeleteKeyword(request *DeleteKeywordRequest) (response *DeleteKeywordResponse, err error) {
	response = CreateDeleteKeywordResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteKeywordWithChan invokes the green.DeleteKeyword API asynchronously
// api document: https://help.aliyun.com/api/green/deletekeyword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteKeywordWithChan(request *DeleteKeywordRequest) (<-chan *DeleteKeywordResponse, <-chan error) {
	responseChan := make(chan *DeleteKeywordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteKeyword(request)
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

// DeleteKeywordWithCallback invokes the green.DeleteKeyword API asynchronously
// api document: https://help.aliyun.com/api/green/deletekeyword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteKeywordWithCallback(request *DeleteKeywordRequest, callback func(response *DeleteKeywordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteKeywordResponse
		var err error
		defer close(result)
		response, err = client.DeleteKeyword(request)
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

// DeleteKeywordRequest is the request struct for api DeleteKeyword
type DeleteKeywordRequest struct {
	*requests.RpcRequest
	Keywords     string `position:"Query" name:"Keywords"`
	KeywordLibId string `position:"Query" name:"KeywordLibId"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	Ids          string `position:"Query" name:"Ids"`
	Lang         string `position:"Query" name:"Lang"`
}

// DeleteKeywordResponse is the response struct for api DeleteKeyword
type DeleteKeywordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteKeywordRequest creates a request to invoke DeleteKeyword API
func CreateDeleteKeywordRequest() (request *DeleteKeywordRequest) {
	request = &DeleteKeywordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DeleteKeyword", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteKeywordResponse creates a response to parse from DeleteKeyword response
func CreateDeleteKeywordResponse() (response *DeleteKeywordResponse) {
	response = &DeleteKeywordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
