package cs

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

// InstallClusterAddons invokes the cs.InstallClusterAddons API synchronously
// api document: https://help.aliyun.com/api/cs/installclusteraddons.html
func (client *Client) InstallClusterAddons(request *InstallClusterAddonsRequest) (response *InstallClusterAddonsResponse, err error) {
	response = CreateInstallClusterAddonsResponse()
	err = client.DoAction(request, response)
	return
}

// InstallClusterAddonsWithChan invokes the cs.InstallClusterAddons API asynchronously
// api document: https://help.aliyun.com/api/cs/installclusteraddons.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InstallClusterAddonsWithChan(request *InstallClusterAddonsRequest) (<-chan *InstallClusterAddonsResponse, <-chan error) {
	responseChan := make(chan *InstallClusterAddonsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallClusterAddons(request)
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

// InstallClusterAddonsWithCallback invokes the cs.InstallClusterAddons API asynchronously
// api document: https://help.aliyun.com/api/cs/installclusteraddons.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InstallClusterAddonsWithCallback(request *InstallClusterAddonsRequest, callback func(response *InstallClusterAddonsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallClusterAddonsResponse
		var err error
		defer close(result)
		response, err = client.InstallClusterAddons(request)
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

// InstallClusterAddonsRequest is the request struct for api InstallClusterAddons
type InstallClusterAddonsRequest struct {
	*requests.RoaRequest
	Name      string           `position:"Body" name:"name"`
	Disabled  requests.Boolean `position:"Body" name:"disabled"`
	ClusterId string           `position:"Path" name:"ClusterId"`
	Version   string           `position:"Body" name:"version"`
	Config    string           `position:"Body" name:"config"`
	Required  string           `position:"Body" name:"required"`
}

// InstallClusterAddonsResponse is the response struct for api InstallClusterAddons
type InstallClusterAddonsResponse struct {
	*responses.BaseResponse
}

// CreateInstallClusterAddonsRequest creates a request to invoke InstallClusterAddons API
func CreateInstallClusterAddonsRequest() (request *InstallClusterAddonsRequest) {
	request = &InstallClusterAddonsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "InstallClusterAddons", "/clusters/[ClusterId]/components/install", "", "")
	request.Method = requests.POST
	return
}

// CreateInstallClusterAddonsResponse creates a response to parse from InstallClusterAddons response
func CreateInstallClusterAddonsResponse() (response *InstallClusterAddonsResponse) {
	response = &InstallClusterAddonsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
