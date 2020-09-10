package dm

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

// BatchSendMail invokes the dm.BatchSendMail API synchronously
// api document: https://help.aliyun.com/api/dm/batchsendmail.html
func (client *Client) BatchSendMail(request *BatchSendMailRequest) (response *BatchSendMailResponse, err error) {
	response = CreateBatchSendMailResponse()
	err = client.DoAction(request, response)
	return
}

// BatchSendMailWithChan invokes the dm.BatchSendMail API asynchronously
// api document: https://help.aliyun.com/api/dm/batchsendmail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchSendMailWithChan(request *BatchSendMailRequest) (<-chan *BatchSendMailResponse, <-chan error) {
	responseChan := make(chan *BatchSendMailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchSendMail(request)
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

// BatchSendMailWithCallback invokes the dm.BatchSendMail API asynchronously
// api document: https://help.aliyun.com/api/dm/batchsendmail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchSendMailWithCallback(request *BatchSendMailRequest, callback func(response *BatchSendMailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchSendMailResponse
		var err error
		defer close(result)
		response, err = client.BatchSendMail(request)
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

// BatchSendMailRequest is the request struct for api BatchSendMail
type BatchSendMailRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ReceiversName        string           `position:"Query" name:"ReceiversName"`
	AddressType          requests.Integer `position:"Query" name:"AddressType"`
	TagName              string           `position:"Query" name:"TagName"`
	ReplyAddress         string           `position:"Query" name:"ReplyAddress"`
	ReplyAddressAlias    string           `position:"Query" name:"ReplyAddressAlias"`
	ClickTrace           string           `position:"Query" name:"ClickTrace"`
}

// BatchSendMailResponse is the response struct for api BatchSendMail
type BatchSendMailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	EnvId     string `json:"EnvId" xml:"EnvId"`
}

// CreateBatchSendMailRequest creates a request to invoke BatchSendMail API
func CreateBatchSendMailRequest() (request *BatchSendMailRequest) {
	request = &BatchSendMailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "BatchSendMail", "", "")
	return
}

// CreateBatchSendMailResponse creates a response to parse from BatchSendMail response
func CreateBatchSendMailResponse() (response *BatchSendMailResponse) {
	response = &BatchSendMailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
