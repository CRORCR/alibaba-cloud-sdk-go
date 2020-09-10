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

// AddPerson invokes the green.AddPerson API synchronously
// api document: https://help.aliyun.com/api/green/addperson.html
func (client *Client) AddPerson(request *AddPersonRequest) (response *AddPersonResponse, err error) {
	response = CreateAddPersonResponse()
	err = client.DoAction(request, response)
	return
}

// AddPersonWithChan invokes the green.AddPerson API asynchronously
// api document: https://help.aliyun.com/api/green/addperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddPersonWithChan(request *AddPersonRequest) (<-chan *AddPersonResponse, <-chan error) {
	responseChan := make(chan *AddPersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPerson(request)
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

// AddPersonWithCallback invokes the green.AddPerson API asynchronously
// api document: https://help.aliyun.com/api/green/addperson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddPersonWithCallback(request *AddPersonRequest, callback func(response *AddPersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPersonResponse
		var err error
		defer close(result)
		response, err = client.AddPerson(request)
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

// AddPersonRequest is the request struct for api AddPerson
type AddPersonRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// AddPersonResponse is the response struct for api AddPerson
type AddPersonResponse struct {
	*responses.BaseResponse
}

// CreateAddPersonRequest creates a request to invoke AddPerson API
func CreateAddPersonRequest() (request *AddPersonRequest) {
	request = &AddPersonRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "AddPerson", "/green/sface/person/add", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddPersonResponse creates a response to parse from AddPerson response
func CreateAddPersonResponse() (response *AddPersonResponse) {
	response = &AddPersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
