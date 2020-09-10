package vpc

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

// AddCommonBandwidthPackageIp invokes the vpc.AddCommonBandwidthPackageIp API synchronously
// api document: https://help.aliyun.com/api/vpc/addcommonbandwidthpackageip.html
func (client *Client) AddCommonBandwidthPackageIp(request *AddCommonBandwidthPackageIpRequest) (response *AddCommonBandwidthPackageIpResponse, err error) {
	response = CreateAddCommonBandwidthPackageIpResponse()
	err = client.DoAction(request, response)
	return
}

// AddCommonBandwidthPackageIpWithChan invokes the vpc.AddCommonBandwidthPackageIp API asynchronously
// api document: https://help.aliyun.com/api/vpc/addcommonbandwidthpackageip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCommonBandwidthPackageIpWithChan(request *AddCommonBandwidthPackageIpRequest) (<-chan *AddCommonBandwidthPackageIpResponse, <-chan error) {
	responseChan := make(chan *AddCommonBandwidthPackageIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCommonBandwidthPackageIp(request)
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

// AddCommonBandwidthPackageIpWithCallback invokes the vpc.AddCommonBandwidthPackageIp API asynchronously
// api document: https://help.aliyun.com/api/vpc/addcommonbandwidthpackageip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCommonBandwidthPackageIpWithCallback(request *AddCommonBandwidthPackageIpRequest, callback func(response *AddCommonBandwidthPackageIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCommonBandwidthPackageIpResponse
		var err error
		defer close(result)
		response, err = client.AddCommonBandwidthPackageIp(request)
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

// AddCommonBandwidthPackageIpRequest is the request struct for api AddCommonBandwidthPackageIp
type AddCommonBandwidthPackageIpRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	IpType               string           `position:"Query" name:"IpType"`
	IpInstanceId         string           `position:"Query" name:"IpInstanceId"`
}

// AddCommonBandwidthPackageIpResponse is the response struct for api AddCommonBandwidthPackageIp
type AddCommonBandwidthPackageIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddCommonBandwidthPackageIpRequest creates a request to invoke AddCommonBandwidthPackageIp API
func CreateAddCommonBandwidthPackageIpRequest() (request *AddCommonBandwidthPackageIpRequest) {
	request = &AddCommonBandwidthPackageIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AddCommonBandwidthPackageIp", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddCommonBandwidthPackageIpResponse creates a response to parse from AddCommonBandwidthPackageIp response
func CreateAddCommonBandwidthPackageIpResponse() (response *AddCommonBandwidthPackageIpResponse) {
	response = &AddCommonBandwidthPackageIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
