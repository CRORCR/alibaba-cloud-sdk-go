package dcdn

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

// DescribeDcdnConfigOfVersion invokes the dcdn.DescribeDcdnConfigOfVersion API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnconfigofversion.html
func (client *Client) DescribeDcdnConfigOfVersion(request *DescribeDcdnConfigOfVersionRequest) (response *DescribeDcdnConfigOfVersionResponse, err error) {
	response = CreateDescribeDcdnConfigOfVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnConfigOfVersionWithChan invokes the dcdn.DescribeDcdnConfigOfVersion API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnconfigofversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnConfigOfVersionWithChan(request *DescribeDcdnConfigOfVersionRequest) (<-chan *DescribeDcdnConfigOfVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnConfigOfVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnConfigOfVersion(request)
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

// DescribeDcdnConfigOfVersionWithCallback invokes the dcdn.DescribeDcdnConfigOfVersion API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnconfigofversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnConfigOfVersionWithCallback(request *DescribeDcdnConfigOfVersionRequest, callback func(response *DescribeDcdnConfigOfVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnConfigOfVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnConfigOfVersion(request)
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

// DescribeDcdnConfigOfVersionRequest is the request struct for api DescribeDcdnConfigOfVersion
type DescribeDcdnConfigOfVersionRequest struct {
	*requests.RpcRequest
	VersionId     string           `position:"Query" name:"VersionId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	FunctionName  string           `position:"Query" name:"FunctionName"`
	GroupId       requests.Integer `position:"Query" name:"GroupId"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	FunctionId    requests.Integer `position:"Query" name:"FunctionId"`
}

// DescribeDcdnConfigOfVersionResponse is the response struct for api DescribeDcdnConfigOfVersion
type DescribeDcdnConfigOfVersionResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	VersionConfigs VersionConfigs `json:"VersionConfigs" xml:"VersionConfigs"`
}

// CreateDescribeDcdnConfigOfVersionRequest creates a request to invoke DescribeDcdnConfigOfVersion API
func CreateDescribeDcdnConfigOfVersionRequest() (request *DescribeDcdnConfigOfVersionRequest) {
	request = &DescribeDcdnConfigOfVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnConfigOfVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnConfigOfVersionResponse creates a response to parse from DescribeDcdnConfigOfVersion response
func CreateDescribeDcdnConfigOfVersionResponse() (response *DescribeDcdnConfigOfVersionResponse) {
	response = &DescribeDcdnConfigOfVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
