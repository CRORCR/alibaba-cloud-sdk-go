package retailcloud

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

// CreateDb invokes the retailcloud.CreateDb API synchronously
// api document: https://help.aliyun.com/api/retailcloud/createdb.html
func (client *Client) CreateDb(request *CreateDbRequest) (response *CreateDbResponse, err error) {
	response = CreateCreateDbResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDbWithChan invokes the retailcloud.CreateDb API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/createdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDbWithChan(request *CreateDbRequest) (<-chan *CreateDbResponse, <-chan error) {
	responseChan := make(chan *CreateDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDb(request)
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

// CreateDbWithCallback invokes the retailcloud.CreateDb API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/createdb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDbWithCallback(request *CreateDbRequest, callback func(response *CreateDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDbResponse
		var err error
		defer close(result)
		response, err = client.CreateDb(request)
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

// CreateDbRequest is the request struct for api CreateDb
type CreateDbRequest struct {
	*requests.RpcRequest
	DbName           string `position:"Body" name:"DbName"`
	DbInstanceId     string `position:"Body" name:"DbInstanceId"`
	DbDescription    string `position:"Body" name:"DbDescription"`
	CharacterSetName string `position:"Body" name:"CharacterSetName"`
}

// CreateDbResponse is the response struct for api CreateDb
type CreateDbResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
}

// CreateCreateDbRequest creates a request to invoke CreateDb API
func CreateCreateDbRequest() (request *CreateDbRequest) {
	request = &CreateDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "CreateDb", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDbResponse creates a response to parse from CreateDb response
func CreateCreateDbResponse() (response *CreateDbResponse) {
	response = &CreateDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
