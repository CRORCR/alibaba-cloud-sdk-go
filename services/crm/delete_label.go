package crm

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

// DeleteLabel invokes the crm.DeleteLabel API synchronously
// api document: https://help.aliyun.com/api/crm/deletelabel.html
func (client *Client) DeleteLabel(request *DeleteLabelRequest) (response *DeleteLabelResponse, err error) {
	response = CreateDeleteLabelResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLabelWithChan invokes the crm.DeleteLabel API asynchronously
// api document: https://help.aliyun.com/api/crm/deletelabel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLabelWithChan(request *DeleteLabelRequest) (<-chan *DeleteLabelResponse, <-chan error) {
	responseChan := make(chan *DeleteLabelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLabel(request)
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

// DeleteLabelWithCallback invokes the crm.DeleteLabel API asynchronously
// api document: https://help.aliyun.com/api/crm/deletelabel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLabelWithCallback(request *DeleteLabelRequest, callback func(response *DeleteLabelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLabelResponse
		var err error
		defer close(result)
		response, err = client.DeleteLabel(request)
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

// DeleteLabelRequest is the request struct for api DeleteLabel
type DeleteLabelRequest struct {
	*requests.RpcRequest
	LabelSeries  string `position:"Query" name:"LabelSeries"`
	Organization string `position:"Query" name:"Organization"`
	PK           string `position:"Query" name:"PK"`
	LabelName    string `position:"Query" name:"LabelName"`
	UserName     string `position:"Query" name:"UserName"`
}

// DeleteLabelResponse is the response struct for api DeleteLabel
type DeleteLabelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateDeleteLabelRequest creates a request to invoke DeleteLabel API
func CreateDeleteLabelRequest() (request *DeleteLabelRequest) {
	request = &DeleteLabelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Crm", "2015-04-08", "DeleteLabel", "crm", "openAPI")
	return
}

// CreateDeleteLabelResponse creates a response to parse from DeleteLabel response
func CreateDeleteLabelResponse() (response *DeleteLabelResponse) {
	response = &DeleteLabelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
