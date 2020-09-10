package elasticsearch

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

// GetRegionConfiguration invokes the elasticsearch.GetRegionConfiguration API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/getregionconfiguration.html
func (client *Client) GetRegionConfiguration(request *GetRegionConfigurationRequest) (response *GetRegionConfigurationResponse, err error) {
	response = CreateGetRegionConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// GetRegionConfigurationWithChan invokes the elasticsearch.GetRegionConfiguration API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/getregionconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRegionConfigurationWithChan(request *GetRegionConfigurationRequest) (<-chan *GetRegionConfigurationResponse, <-chan error) {
	responseChan := make(chan *GetRegionConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRegionConfiguration(request)
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

// GetRegionConfigurationWithCallback invokes the elasticsearch.GetRegionConfiguration API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/getregionconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRegionConfigurationWithCallback(request *GetRegionConfigurationRequest, callback func(response *GetRegionConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRegionConfigurationResponse
		var err error
		defer close(result)
		response, err = client.GetRegionConfiguration(request)
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

// GetRegionConfigurationRequest is the request struct for api GetRegionConfiguration
type GetRegionConfigurationRequest struct {
	*requests.RoaRequest
	ZoneId string `position:"Query" name:"zoneId"`
}

// GetRegionConfigurationResponse is the response struct for api GetRegionConfiguration
type GetRegionConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateGetRegionConfigurationRequest creates a request to invoke GetRegionConfiguration API
func CreateGetRegionConfigurationRequest() (request *GetRegionConfigurationRequest) {
	request = &GetRegionConfigurationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "GetRegionConfiguration", "/openapi/region", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRegionConfigurationResponse creates a response to parse from GetRegionConfiguration response
func CreateGetRegionConfigurationResponse() (response *GetRegionConfigurationResponse) {
	response = &GetRegionConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
