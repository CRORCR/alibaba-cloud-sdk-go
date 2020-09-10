package polardb

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

// CreateDatabase invokes the polardb.CreateDatabase API synchronously
// api document: https://help.aliyun.com/api/polardb/createdatabase.html
func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (response *CreateDatabaseResponse, err error) {
	response = CreateCreateDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDatabaseWithChan invokes the polardb.CreateDatabase API asynchronously
// api document: https://help.aliyun.com/api/polardb/createdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDatabaseWithChan(request *CreateDatabaseRequest) (<-chan *CreateDatabaseResponse, <-chan error) {
	responseChan := make(chan *CreateDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDatabase(request)
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

// CreateDatabaseWithCallback invokes the polardb.CreateDatabase API asynchronously
// api document: https://help.aliyun.com/api/polardb/createdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDatabaseWithCallback(request *CreateDatabaseRequest, callback func(response *CreateDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDatabaseResponse
		var err error
		defer close(result)
		response, err = client.CreateDatabase(request)
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

// CreateDatabaseRequest is the request struct for api CreateDatabase
type CreateDatabaseRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountPrivilege     string           `position:"Query" name:"AccountPrivilege"`
	AccountName          string           `position:"Query" name:"AccountName"`
	DBDescription        string           `position:"Query" name:"DBDescription"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Collate              string           `position:"Query" name:"Collate"`
	DBName               string           `position:"Query" name:"DBName"`
	Ctype                string           `position:"Query" name:"Ctype"`
	CharacterSetName     string           `position:"Query" name:"CharacterSetName"`
}

// CreateDatabaseResponse is the response struct for api CreateDatabase
type CreateDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDatabaseRequest creates a request to invoke CreateDatabase API
func CreateCreateDatabaseRequest() (request *CreateDatabaseRequest) {
	request = &CreateDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateDatabase", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDatabaseResponse creates a response to parse from CreateDatabase response
func CreateCreateDatabaseResponse() (response *CreateDatabaseResponse) {
	response = &CreateDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
