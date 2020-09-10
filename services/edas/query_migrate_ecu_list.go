package edas

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

// QueryMigrateEcuList invokes the edas.QueryMigrateEcuList API synchronously
// api document: https://help.aliyun.com/api/edas/querymigrateeculist.html
func (client *Client) QueryMigrateEcuList(request *QueryMigrateEcuListRequest) (response *QueryMigrateEcuListResponse, err error) {
	response = CreateQueryMigrateEcuListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMigrateEcuListWithChan invokes the edas.QueryMigrateEcuList API asynchronously
// api document: https://help.aliyun.com/api/edas/querymigrateeculist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMigrateEcuListWithChan(request *QueryMigrateEcuListRequest) (<-chan *QueryMigrateEcuListResponse, <-chan error) {
	responseChan := make(chan *QueryMigrateEcuListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMigrateEcuList(request)
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

// QueryMigrateEcuListWithCallback invokes the edas.QueryMigrateEcuList API asynchronously
// api document: https://help.aliyun.com/api/edas/querymigrateeculist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMigrateEcuListWithCallback(request *QueryMigrateEcuListRequest, callback func(response *QueryMigrateEcuListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMigrateEcuListResponse
		var err error
		defer close(result)
		response, err = client.QueryMigrateEcuList(request)
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

// QueryMigrateEcuListRequest is the request struct for api QueryMigrateEcuList
type QueryMigrateEcuListRequest struct {
	*requests.RoaRequest
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
}

// QueryMigrateEcuListResponse is the response struct for api QueryMigrateEcuList
type QueryMigrateEcuListResponse struct {
	*responses.BaseResponse
	Code          int                                `json:"Code" xml:"Code"`
	Message       string                             `json:"Message" xml:"Message"`
	RequestId     string                             `json:"RequestId" xml:"RequestId"`
	EcuEntityList EcuEntityListInQueryMigrateEcuList `json:"EcuEntityList" xml:"EcuEntityList"`
}

// CreateQueryMigrateEcuListRequest creates a request to invoke QueryMigrateEcuList API
func CreateQueryMigrateEcuListRequest() (request *QueryMigrateEcuListRequest) {
	request = &QueryMigrateEcuListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "QueryMigrateEcuList", "/pop/v5/resource/migrate_ecu_list", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryMigrateEcuListResponse creates a response to parse from QueryMigrateEcuList response
func CreateQueryMigrateEcuListResponse() (response *QueryMigrateEcuListResponse) {
	response = &QueryMigrateEcuListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
