package airec

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

// DescribeMix invokes the airec.DescribeMix API synchronously
// api document: https://help.aliyun.com/api/airec/describemix.html
func (client *Client) DescribeMix(request *DescribeMixRequest) (response *DescribeMixResponse, err error) {
	response = CreateDescribeMixResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMixWithChan invokes the airec.DescribeMix API asynchronously
// api document: https://help.aliyun.com/api/airec/describemix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMixWithChan(request *DescribeMixRequest) (<-chan *DescribeMixResponse, <-chan error) {
	responseChan := make(chan *DescribeMixResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMix(request)
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

// DescribeMixWithCallback invokes the airec.DescribeMix API asynchronously
// api document: https://help.aliyun.com/api/airec/describemix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMixWithCallback(request *DescribeMixRequest, callback func(response *DescribeMixResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMixResponse
		var err error
		defer close(result)
		response, err = client.DescribeMix(request)
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

// DescribeMixRequest is the request struct for api DescribeMix
type DescribeMixRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Name       string `position:"Path" name:"Name"`
}

// DescribeMixResponse is the response struct for api DescribeMix
type DescribeMixResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	Code      string              `json:"Code" xml:"Code"`
	Message   string              `json:"Message" xml:"Message"`
	Result    ResultInDescribeMix `json:"Result" xml:"Result"`
}

// CreateDescribeMixRequest creates a request to invoke DescribeMix API
func CreateDescribeMixRequest() (request *DescribeMixRequest) {
	request = &DescribeMixRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "DescribeMix", "/openapi/instances/[InstanceId]/mixes/[Name]", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeMixResponse creates a response to parse from DescribeMix response
func CreateDescribeMixResponse() (response *DescribeMixResponse) {
	response = &DescribeMixResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
