package elasticsearch

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

// GetElastictask invokes the elasticsearch.GetElastictask API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/getelastictask.html
func (client *Client) GetElastictask(request *GetElastictaskRequest) (response *GetElastictaskResponse, err error) {
	response = CreateGetElastictaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetElastictaskWithChan invokes the elasticsearch.GetElastictask API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/getelastictask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetElastictaskWithChan(request *GetElastictaskRequest) (<-chan *GetElastictaskResponse, <-chan error) {
	responseChan := make(chan *GetElastictaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetElastictask(request)
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

// GetElastictaskWithCallback invokes the elasticsearch.GetElastictask API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/getelastictask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetElastictaskWithCallback(request *GetElastictaskRequest, callback func(response *GetElastictaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetElastictaskResponse
		var err error
		defer close(result)
		response, err = client.GetElastictask(request)
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

// GetElastictaskRequest is the request struct for api GetElastictask
type GetElastictaskRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
}

// GetElastictaskResponse is the response struct for api GetElastictask
type GetElastictaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateGetElastictaskRequest creates a request to invoke GetElastictask API
func CreateGetElastictaskRequest() (request *GetElastictaskRequest) {
	request = &GetElastictaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "GetElastictask", "/openapi/instances/[InstanceId]/elastic-task", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetElastictaskResponse creates a response to parse from GetElastictask response
func CreateGetElastictaskResponse() (response *GetElastictaskResponse) {
	response = &GetElastictaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
