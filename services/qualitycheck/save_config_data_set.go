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

// SaveConfigDataSet invokes the qualitycheck.SaveConfigDataSet API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/saveconfigdataset.html
func (client *Client) SaveConfigDataSet(request *SaveConfigDataSetRequest) (response *SaveConfigDataSetResponse, err error) {
	response = CreateSaveConfigDataSetResponse()
	err = client.DoAction(request, response)
	return
}

// SaveConfigDataSetWithChan invokes the qualitycheck.SaveConfigDataSet API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/saveconfigdataset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveConfigDataSetWithChan(request *SaveConfigDataSetRequest) (<-chan *SaveConfigDataSetResponse, <-chan error) {
	responseChan := make(chan *SaveConfigDataSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveConfigDataSet(request)
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

// SaveConfigDataSetWithCallback invokes the qualitycheck.SaveConfigDataSet API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/saveconfigdataset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveConfigDataSetWithCallback(request *SaveConfigDataSetRequest, callback func(response *SaveConfigDataSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveConfigDataSetResponse
		var err error
		defer close(result)
		response, err = client.SaveConfigDataSet(request)
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

// SaveConfigDataSetRequest is the request struct for api SaveConfigDataSet
type SaveConfigDataSetRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// SaveConfigDataSetResponse is the response struct for api SaveConfigDataSet
type SaveConfigDataSetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSaveConfigDataSetRequest creates a request to invoke SaveConfigDataSet API
func CreateSaveConfigDataSetRequest() (request *SaveConfigDataSetRequest) {
	request = &SaveConfigDataSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "SaveConfigDataSet", "", "")
	return
}

// CreateSaveConfigDataSetResponse creates a response to parse from SaveConfigDataSet response
func CreateSaveConfigDataSetResponse() (response *SaveConfigDataSetResponse) {
	response = &SaveConfigDataSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
