package live

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

// AddCasterLayout invokes the live.AddCasterLayout API synchronously
// api document: https://help.aliyun.com/api/live/addcasterlayout.html
func (client *Client) AddCasterLayout(request *AddCasterLayoutRequest) (response *AddCasterLayoutResponse, err error) {
	response = CreateAddCasterLayoutResponse()
	err = client.DoAction(request, response)
	return
}

// AddCasterLayoutWithChan invokes the live.AddCasterLayout API asynchronously
// api document: https://help.aliyun.com/api/live/addcasterlayout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCasterLayoutWithChan(request *AddCasterLayoutRequest) (<-chan *AddCasterLayoutResponse, <-chan error) {
	responseChan := make(chan *AddCasterLayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCasterLayout(request)
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

// AddCasterLayoutWithCallback invokes the live.AddCasterLayout API asynchronously
// api document: https://help.aliyun.com/api/live/addcasterlayout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCasterLayoutWithCallback(request *AddCasterLayoutRequest, callback func(response *AddCasterLayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCasterLayoutResponse
		var err error
		defer close(result)
		response, err = client.AddCasterLayout(request)
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

// AddCasterLayoutRequest is the request struct for api AddCasterLayout
type AddCasterLayoutRequest struct {
	*requests.RpcRequest
	BlendList  *[]string                    `position:"Query" name:"BlendList"  type:"Repeated"`
	CasterId   string                       `position:"Query" name:"CasterId"`
	OwnerId    requests.Integer             `position:"Query" name:"OwnerId"`
	AudioLayer *[]AddCasterLayoutAudioLayer `position:"Query" name:"AudioLayer"  type:"Repeated"`
	VideoLayer *[]AddCasterLayoutVideoLayer `position:"Query" name:"VideoLayer"  type:"Repeated"`
	MixList    *[]string                    `position:"Query" name:"MixList"  type:"Repeated"`
}

// AddCasterLayoutAudioLayer is a repeated param struct in AddCasterLayoutRequest
type AddCasterLayoutAudioLayer struct {
	VolumeRate         string `name:"VolumeRate"`
	ValidChannel       string `name:"ValidChannel"`
	FixedDelayDuration string `name:"FixedDelayDuration"`
}

// AddCasterLayoutVideoLayer is a repeated param struct in AddCasterLayoutRequest
type AddCasterLayoutVideoLayer struct {
	FillMode           string    `name:"FillMode"`
	HeightNormalized   string    `name:"HeightNormalized"`
	WidthNormalized    string    `name:"WidthNormalized"`
	PositionRefer      string    `name:"PositionRefer"`
	PositionNormalized *[]string `name:"PositionNormalized" type:"Repeated"`
	FixedDelayDuration string    `name:"FixedDelayDuration"`
}

// AddCasterLayoutResponse is the response struct for api AddCasterLayout
type AddCasterLayoutResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	LayoutId  string `json:"LayoutId" xml:"LayoutId"`
}

// CreateAddCasterLayoutRequest creates a request to invoke AddCasterLayout API
func CreateAddCasterLayoutRequest() (request *AddCasterLayoutRequest) {
	request = &AddCasterLayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddCasterLayout", "", "")
	request.Method = requests.POST
	return
}

// CreateAddCasterLayoutResponse creates a response to parse from AddCasterLayout response
func CreateAddCasterLayoutResponse() (response *AddCasterLayoutResponse) {
	response = &AddCasterLayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
