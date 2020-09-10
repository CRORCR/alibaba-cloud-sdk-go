package appmallsservice

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

// TaobaoFilmGetSeats invokes the appmallsservice.TaobaoFilmGetSeats API synchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmgetseats.html
func (client *Client) TaobaoFilmGetSeats(request *TaobaoFilmGetSeatsRequest) (response *TaobaoFilmGetSeatsResponse, err error) {
	response = CreateTaobaoFilmGetSeatsResponse()
	err = client.DoAction(request, response)
	return
}

// TaobaoFilmGetSeatsWithChan invokes the appmallsservice.TaobaoFilmGetSeats API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmgetseats.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmGetSeatsWithChan(request *TaobaoFilmGetSeatsRequest) (<-chan *TaobaoFilmGetSeatsResponse, <-chan error) {
	responseChan := make(chan *TaobaoFilmGetSeatsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaobaoFilmGetSeats(request)
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

// TaobaoFilmGetSeatsWithCallback invokes the appmallsservice.TaobaoFilmGetSeats API asynchronously
// api document: https://help.aliyun.com/api/appmallsservice/taobaofilmgetseats.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaobaoFilmGetSeatsWithCallback(request *TaobaoFilmGetSeatsRequest, callback func(response *TaobaoFilmGetSeatsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaobaoFilmGetSeatsResponse
		var err error
		defer close(result)
		response, err = client.TaobaoFilmGetSeats(request)
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

// TaobaoFilmGetSeatsRequest is the request struct for api TaobaoFilmGetSeats
type TaobaoFilmGetSeatsRequest struct {
	*requests.RpcRequest
	ScheduleId requests.Integer `position:"Query" name:"ScheduleId"`
	ParamsJson string           `position:"Query" name:"ParamsJson"`
}

// TaobaoFilmGetSeatsResponse is the response struct for api TaobaoFilmGetSeats
type TaobaoFilmGetSeatsResponse struct {
	*responses.BaseResponse
	ErrorCode string  `json:"ErrorCode" xml:"ErrorCode"`
	Msg       string  `json:"Msg" xml:"Msg"`
	SubCode   string  `json:"SubCode" xml:"SubCode"`
	SubMsg    string  `json:"SubMsg" xml:"SubMsg"`
	LogsId    string  `json:"LogsId" xml:"LogsId"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	SeatMap   SeatMap `json:"SeatMap" xml:"SeatMap"`
}

// CreateTaobaoFilmGetSeatsRequest creates a request to invoke TaobaoFilmGetSeats API
func CreateTaobaoFilmGetSeatsRequest() (request *TaobaoFilmGetSeatsRequest) {
	request = &TaobaoFilmGetSeatsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AppMallsService", "2018-02-24", "TaobaoFilmGetSeats", "", "")
	return
}

// CreateTaobaoFilmGetSeatsResponse creates a response to parse from TaobaoFilmGetSeats response
func CreateTaobaoFilmGetSeatsResponse() (response *TaobaoFilmGetSeatsResponse) {
	response = &TaobaoFilmGetSeatsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
