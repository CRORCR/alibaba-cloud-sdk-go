package ehpc

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

// GetCloudMetricLogs invokes the ehpc.GetCloudMetricLogs API synchronously
// api document: https://help.aliyun.com/api/ehpc/getcloudmetriclogs.html
func (client *Client) GetCloudMetricLogs(request *GetCloudMetricLogsRequest) (response *GetCloudMetricLogsResponse, err error) {
	response = CreateGetCloudMetricLogsResponse()
	err = client.DoAction(request, response)
	return
}

// GetCloudMetricLogsWithChan invokes the ehpc.GetCloudMetricLogs API asynchronously
// api document: https://help.aliyun.com/api/ehpc/getcloudmetriclogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCloudMetricLogsWithChan(request *GetCloudMetricLogsRequest) (<-chan *GetCloudMetricLogsResponse, <-chan error) {
	responseChan := make(chan *GetCloudMetricLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCloudMetricLogs(request)
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

// GetCloudMetricLogsWithCallback invokes the ehpc.GetCloudMetricLogs API asynchronously
// api document: https://help.aliyun.com/api/ehpc/getcloudmetriclogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCloudMetricLogsWithCallback(request *GetCloudMetricLogsRequest, callback func(response *GetCloudMetricLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCloudMetricLogsResponse
		var err error
		defer close(result)
		response, err = client.GetCloudMetricLogs(request)
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

// GetCloudMetricLogsRequest is the request struct for api GetCloudMetricLogs
type GetCloudMetricLogsRequest struct {
	*requests.RpcRequest
	MetricScope         string           `position:"Query" name:"MetricScope"`
	ClusterId           string           `position:"Query" name:"ClusterId"`
	AggregationInterval requests.Integer `position:"Query" name:"AggregationInterval"`
	Reverse             requests.Boolean `position:"Query" name:"Reverse"`
	AggregationType     string           `position:"Query" name:"AggregationType"`
	Filter              string           `position:"Query" name:"Filter"`
	MetricCategories    string           `position:"Query" name:"MetricCategories"`
	From                requests.Integer `position:"Query" name:"From"`
	To                  requests.Integer `position:"Query" name:"To"`
}

// GetCloudMetricLogsResponse is the response struct for api GetCloudMetricLogs
type GetCloudMetricLogsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	MetricLogs MetricLogs `json:"MetricLogs" xml:"MetricLogs"`
}

// CreateGetCloudMetricLogsRequest creates a request to invoke GetCloudMetricLogs API
func CreateGetCloudMetricLogsRequest() (request *GetCloudMetricLogsRequest) {
	request = &GetCloudMetricLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetCloudMetricLogs", "", "")
	request.Method = requests.GET
	return
}

// CreateGetCloudMetricLogsResponse creates a response to parse from GetCloudMetricLogs response
func CreateGetCloudMetricLogsResponse() (response *GetCloudMetricLogsResponse) {
	response = &GetCloudMetricLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
