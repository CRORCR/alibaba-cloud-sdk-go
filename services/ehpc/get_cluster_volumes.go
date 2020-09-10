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

// GetClusterVolumes invokes the ehpc.GetClusterVolumes API synchronously
// api document: https://help.aliyun.com/api/ehpc/getclustervolumes.html
func (client *Client) GetClusterVolumes(request *GetClusterVolumesRequest) (response *GetClusterVolumesResponse, err error) {
	response = CreateGetClusterVolumesResponse()
	err = client.DoAction(request, response)
	return
}

// GetClusterVolumesWithChan invokes the ehpc.GetClusterVolumes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/getclustervolumes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetClusterVolumesWithChan(request *GetClusterVolumesRequest) (<-chan *GetClusterVolumesResponse, <-chan error) {
	responseChan := make(chan *GetClusterVolumesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetClusterVolumes(request)
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

// GetClusterVolumesWithCallback invokes the ehpc.GetClusterVolumes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/getclustervolumes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetClusterVolumesWithCallback(request *GetClusterVolumesRequest, callback func(response *GetClusterVolumesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetClusterVolumesResponse
		var err error
		defer close(result)
		response, err = client.GetClusterVolumes(request)
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

// GetClusterVolumesRequest is the request struct for api GetClusterVolumes
type GetClusterVolumesRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// GetClusterVolumesResponse is the response struct for api GetClusterVolumes
type GetClusterVolumesResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	RegionId  string                     `json:"RegionId" xml:"RegionId"`
	Volumes   VolumesInGetClusterVolumes `json:"Volumes" xml:"Volumes"`
}

// CreateGetClusterVolumesRequest creates a request to invoke GetClusterVolumes API
func CreateGetClusterVolumesRequest() (request *GetClusterVolumesRequest) {
	request = &GetClusterVolumesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetClusterVolumes", "", "")
	request.Method = requests.GET
	return
}

// CreateGetClusterVolumesResponse creates a response to parse from GetClusterVolumes response
func CreateGetClusterVolumesResponse() (response *GetClusterVolumesResponse) {
	response = &GetClusterVolumesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
