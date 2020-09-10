package domain

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

// SaveSingleTaskForSaveArtExtension invokes the domain.SaveSingleTaskForSaveArtExtension API synchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforsaveartextension.html
func (client *Client) SaveSingleTaskForSaveArtExtension(request *SaveSingleTaskForSaveArtExtensionRequest) (response *SaveSingleTaskForSaveArtExtensionResponse, err error) {
	response = CreateSaveSingleTaskForSaveArtExtensionResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForSaveArtExtensionWithChan invokes the domain.SaveSingleTaskForSaveArtExtension API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforsaveartextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForSaveArtExtensionWithChan(request *SaveSingleTaskForSaveArtExtensionRequest) (<-chan *SaveSingleTaskForSaveArtExtensionResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForSaveArtExtensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForSaveArtExtension(request)
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

// SaveSingleTaskForSaveArtExtensionWithCallback invokes the domain.SaveSingleTaskForSaveArtExtension API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforsaveartextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForSaveArtExtensionWithCallback(request *SaveSingleTaskForSaveArtExtensionRequest, callback func(response *SaveSingleTaskForSaveArtExtensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForSaveArtExtensionResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForSaveArtExtension(request)
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

// SaveSingleTaskForSaveArtExtensionRequest is the request struct for api SaveSingleTaskForSaveArtExtension
type SaveSingleTaskForSaveArtExtensionRequest struct {
	*requests.RpcRequest
	Subject                 string `position:"Query" name:"Subject"`
	Title                   string `position:"Query" name:"Title"`
	DateOrPeriod            string `position:"Query" name:"DateOrPeriod"`
	Reference               string `position:"Query" name:"Reference"`
	Features                string `position:"Query" name:"Features"`
	InscriptionsAndMarkings string `position:"Query" name:"InscriptionsAndMarkings"`
	ObjectType              string `position:"Query" name:"ObjectType"`
	Lang                    string `position:"Query" name:"Lang"`
	DomainName              string `position:"Query" name:"DomainName"`
	Maker                   string `position:"Query" name:"Maker"`
	MaterialsAndTechniques  string `position:"Query" name:"MaterialsAndTechniques"`
	Dimensions              string `position:"Query" name:"Dimensions"`
}

// SaveSingleTaskForSaveArtExtensionResponse is the response struct for api SaveSingleTaskForSaveArtExtension
type SaveSingleTaskForSaveArtExtensionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForSaveArtExtensionRequest creates a request to invoke SaveSingleTaskForSaveArtExtension API
func CreateSaveSingleTaskForSaveArtExtensionRequest() (request *SaveSingleTaskForSaveArtExtensionRequest) {
	request = &SaveSingleTaskForSaveArtExtensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForSaveArtExtension", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForSaveArtExtensionResponse creates a response to parse from SaveSingleTaskForSaveArtExtension response
func CreateSaveSingleTaskForSaveArtExtensionResponse() (response *SaveSingleTaskForSaveArtExtensionResponse) {
	response = &SaveSingleTaskForSaveArtExtensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
