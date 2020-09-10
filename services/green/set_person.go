package green

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

// SetPerson invokes the green.SetPerson API synchronously
// api document: https://help.aliyun.com/api/green/setperson.html
func (client *Client) SetPerson(request *SetPersonRequest) (response *SetPersonResponse, err error) {
	response = CreateSetPersonResponse()
	err = client.DoAction(request, response)
	return
}

// SetPersonWithChan invokes the green.SetPerson API asynchronously
// api document: https://help.aliyun.com/api/green/setperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetPersonWithChan(request *SetPersonRequest) (<-chan *SetPersonResponse, <-chan error) {
	responseChan := make(chan *SetPersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetPerson(request)
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

// SetPersonWithCallback invokes the green.SetPerson API asynchronously
// api document: https://help.aliyun.com/api/green/setperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetPersonWithCallback(request *SetPersonRequest, callback func(response *SetPersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetPersonResponse
		var err error
		defer close(result)
		response, err = client.SetPerson(request)
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

// SetPersonRequest is the request struct for api SetPerson
type SetPersonRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// SetPersonResponse is the response struct for api SetPerson
type SetPersonResponse struct {
	*responses.BaseResponse
}

// CreateSetPersonRequest creates a request to invoke SetPerson API
func CreateSetPersonRequest() (request *SetPersonRequest) {
	request = &SetPersonRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "SetPerson", "/green/sface/person/update", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetPersonResponse creates a response to parse from SetPerson response
func CreateSetPersonResponse() (response *SetPersonResponse) {
	response = &SetPersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
