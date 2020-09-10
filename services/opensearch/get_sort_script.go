package opensearch

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

// GetSortScript invokes the opensearch.GetSortScript API synchronously
// api document: https://help.aliyun.com/api/opensearch/getsortscript.html
func (client *Client) GetSortScript(request *GetSortScriptRequest) (response *GetSortScriptResponse, err error) {
	response = CreateGetSortScriptResponse()
	err = client.DoAction(request, response)
	return
}

// GetSortScriptWithChan invokes the opensearch.GetSortScript API asynchronously
// api document: https://help.aliyun.com/api/opensearch/getsortscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSortScriptWithChan(request *GetSortScriptRequest) (<-chan *GetSortScriptResponse, <-chan error) {
	responseChan := make(chan *GetSortScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSortScript(request)
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

// GetSortScriptWithCallback invokes the opensearch.GetSortScript API asynchronously
// api document: https://help.aliyun.com/api/opensearch/getsortscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSortScriptWithCallback(request *GetSortScriptRequest, callback func(response *GetSortScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSortScriptResponse
		var err error
		defer close(result)
		response, err = client.GetSortScript(request)
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

// GetSortScriptRequest is the request struct for api GetSortScript
type GetSortScriptRequest struct {
	*requests.RoaRequest
	AppVersionId     string `position:"Path" name:"appVersionId"`
	ScriptName       string `position:"Path" name:"scriptName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// GetSortScriptResponse is the response struct for api GetSortScript
type GetSortScriptResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"requestId" xml:"requestId"`
	Result    ResultInGetSortScript `json:"result" xml:"result"`
}

// CreateGetSortScriptRequest creates a request to invoke GetSortScript API
func CreateGetSortScriptRequest() (request *GetSortScriptRequest) {
	request = &GetSortScriptRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetSortScript", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appVersionId]/sort-scripts/[scriptName]", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetSortScriptResponse creates a response to parse from GetSortScript response
func CreateGetSortScriptResponse() (response *GetSortScriptResponse) {
	response = &GetSortScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
