package imm

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

// IndexImage invokes the imm.IndexImage API synchronously
// api document: https://help.aliyun.com/api/imm/indeximage.html
func (client *Client) IndexImage(request *IndexImageRequest) (response *IndexImageResponse, err error) {
	response = CreateIndexImageResponse()
	err = client.DoAction(request, response)
	return
}

// IndexImageWithChan invokes the imm.IndexImage API asynchronously
// api document: https://help.aliyun.com/api/imm/indeximage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IndexImageWithChan(request *IndexImageRequest) (<-chan *IndexImageResponse, <-chan error) {
	responseChan := make(chan *IndexImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.IndexImage(request)
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

// IndexImageWithCallback invokes the imm.IndexImage API asynchronously
// api document: https://help.aliyun.com/api/imm/indeximage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) IndexImageWithCallback(request *IndexImageRequest, callback func(response *IndexImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *IndexImageResponse
		var err error
		defer close(result)
		response, err = client.IndexImage(request)
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

// IndexImageRequest is the request struct for api IndexImage
type IndexImageRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	ExternalId      string `position:"Query" name:"ExternalId"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	SourceType      string `position:"Query" name:"SourceType"`
	RealUid         string `position:"Query" name:"RealUid"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	RemarksB        string `position:"Query" name:"RemarksB"`
	RemarksA        string `position:"Query" name:"RemarksA"`
	ImageUri        string `position:"Query" name:"ImageUri"`
	RemarksArrayA   string `position:"Query" name:"RemarksArrayA"`
	RemarksArrayB   string `position:"Query" name:"RemarksArrayB"`
	SourceUri       string `position:"Query" name:"SourceUri"`
	SourcePosition  string `position:"Query" name:"SourcePosition"`
	RemarksD        string `position:"Query" name:"RemarksD"`
	RemarksC        string `position:"Query" name:"RemarksC"`
	SetId           string `position:"Query" name:"SetId"`
}

// IndexImageResponse is the response struct for api IndexImage
type IndexImageResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	SetId         string `json:"SetId" xml:"SetId"`
	ImageUri      string `json:"ImageUri" xml:"ImageUri"`
	RemarksA      string `json:"RemarksA" xml:"RemarksA"`
	RemarksB      string `json:"RemarksB" xml:"RemarksB"`
	CreateTime    string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime    string `json:"ModifyTime" xml:"ModifyTime"`
	RemarksC      string `json:"RemarksC" xml:"RemarksC"`
	RemarksD      string `json:"RemarksD" xml:"RemarksD"`
	ExternalId    string `json:"ExternalId" xml:"ExternalId"`
	RemarksArrayA string `json:"RemarksArrayA" xml:"RemarksArrayA"`
	RemarksArrayB string `json:"RemarksArrayB" xml:"RemarksArrayB"`
}

// CreateIndexImageRequest creates a request to invoke IndexImage API
func CreateIndexImageRequest() (request *IndexImageRequest) {
	request = &IndexImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "IndexImage", "", "")
	request.Method = requests.POST
	return
}

// CreateIndexImageResponse creates a response to parse from IndexImage response
func CreateIndexImageResponse() (response *IndexImageResponse) {
	response = &IndexImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
