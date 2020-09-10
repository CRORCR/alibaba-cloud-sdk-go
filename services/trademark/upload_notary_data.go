package trademark

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

// UploadNotaryData invokes the trademark.UploadNotaryData API synchronously
// api document: https://help.aliyun.com/api/trademark/uploadnotarydata.html
func (client *Client) UploadNotaryData(request *UploadNotaryDataRequest) (response *UploadNotaryDataResponse, err error) {
	response = CreateUploadNotaryDataResponse()
	err = client.DoAction(request, response)
	return
}

// UploadNotaryDataWithChan invokes the trademark.UploadNotaryData API asynchronously
// api document: https://help.aliyun.com/api/trademark/uploadnotarydata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadNotaryDataWithChan(request *UploadNotaryDataRequest) (<-chan *UploadNotaryDataResponse, <-chan error) {
	responseChan := make(chan *UploadNotaryDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadNotaryData(request)
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

// UploadNotaryDataWithCallback invokes the trademark.UploadNotaryData API asynchronously
// api document: https://help.aliyun.com/api/trademark/uploadnotarydata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadNotaryDataWithCallback(request *UploadNotaryDataRequest, callback func(response *UploadNotaryDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadNotaryDataResponse
		var err error
		defer close(result)
		response, err = client.UploadNotaryData(request)
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

// UploadNotaryDataRequest is the request struct for api UploadNotaryData
type UploadNotaryDataRequest struct {
	*requests.RpcRequest
	UploadContext string           `position:"Query" name:"UploadContext"`
	BizOrderNo    string           `position:"Query" name:"BizOrderNo"`
	NotaryType    requests.Integer `position:"Query" name:"NotaryType"`
}

// UploadNotaryDataResponse is the response struct for api UploadNotaryData
type UploadNotaryDataResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	UserAuthUrl string `json:"UserAuthUrl" xml:"UserAuthUrl"`
}

// CreateUploadNotaryDataRequest creates a request to invoke UploadNotaryData API
func CreateUploadNotaryDataRequest() (request *UploadNotaryDataRequest) {
	request = &UploadNotaryDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "UploadNotaryData", "trademark", "openAPI")
	return
}

// CreateUploadNotaryDataResponse creates a response to parse from UploadNotaryData response
func CreateUploadNotaryDataResponse() (response *UploadNotaryDataResponse) {
	response = &UploadNotaryDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
