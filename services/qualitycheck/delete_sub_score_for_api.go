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

// DeleteSubScoreForApi invokes the qualitycheck.DeleteSubScoreForApi API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/deletesubscoreforapi.html
func (client *Client) DeleteSubScoreForApi(request *DeleteSubScoreForApiRequest) (response *DeleteSubScoreForApiResponse, err error) {
	response = CreateDeleteSubScoreForApiResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSubScoreForApiWithChan invokes the qualitycheck.DeleteSubScoreForApi API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/deletesubscoreforapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSubScoreForApiWithChan(request *DeleteSubScoreForApiRequest) (<-chan *DeleteSubScoreForApiResponse, <-chan error) {
	responseChan := make(chan *DeleteSubScoreForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSubScoreForApi(request)
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

// DeleteSubScoreForApiWithCallback invokes the qualitycheck.DeleteSubScoreForApi API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/deletesubscoreforapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSubScoreForApiWithCallback(request *DeleteSubScoreForApiRequest, callback func(response *DeleteSubScoreForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSubScoreForApiResponse
		var err error
		defer close(result)
		response, err = client.DeleteSubScoreForApi(request)
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

// DeleteSubScoreForApiRequest is the request struct for api DeleteSubScoreForApi
type DeleteSubScoreForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// DeleteSubScoreForApiResponse is the response struct for api DeleteSubScoreForApi
type DeleteSubScoreForApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteSubScoreForApiRequest creates a request to invoke DeleteSubScoreForApi API
func CreateDeleteSubScoreForApiRequest() (request *DeleteSubScoreForApiRequest) {
	request = &DeleteSubScoreForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "DeleteSubScoreForApi", "", "")
	return
}

// CreateDeleteSubScoreForApiResponse creates a response to parse from DeleteSubScoreForApi response
func CreateDeleteSubScoreForApiResponse() (response *DeleteSubScoreForApiResponse) {
	response = &DeleteSubScoreForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
