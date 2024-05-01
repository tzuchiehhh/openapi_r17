/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ConvergedCharging

import (
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type DefaultApiService service

/*
DefaultApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ChfConvergedChargingChargingDataRequest -

@return PostChargingDataResponse
*/

// PostChargingDataRequest
type PostChargingDataRequest struct {
	ChfConvergedChargingChargingDataRequest *models.ChfConvergedChargingChargingDataRequest
}

func (r *PostChargingDataRequest) SetChfConvergedChargingChargingDataRequest(ChfConvergedChargingChargingDataRequest models.ChfConvergedChargingChargingDataRequest) {
	r.ChfConvergedChargingChargingDataRequest = &ChfConvergedChargingChargingDataRequest
}

type PostChargingDataResponse struct {
	ChfConvergedChargingChargingDataResponse models.ChfConvergedChargingChargingDataResponse
}

type PostChargingDataError struct {
	Location                    string
	Var3gppSbiTargetNfId        string
	PostChargingDataResponse400 models.PostChargingDataResponse400
	PostChargingDataResponse403 models.PostChargingDataResponse403
	PostChargingDataResponse404 models.PostChargingDataResponse404
	ProblemDetails              models.ProblemDetails
	RedirectResponse            models.RedirectResponse
}

func (a *DefaultApiService) PostChargingData(ctx context.Context, request *PostChargingDataRequest) (*PostChargingDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PostChargingDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/chargingdata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.ChfConvergedChargingChargingDataRequest

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.ChfConvergedChargingChargingDataResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 400:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.PostChargingDataResponse400, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 307:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.PostChargingDataResponse403, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.PostChargingDataResponse404, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 405:
		return &localVarReturnValue, nil
	case 408:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 410:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v PostChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}

// PostChargingNotificationRequest
type PostChargingNotificationRequest struct {
	ChargingNotifyRequest *models.ChargingNotifyRequest
}

func NewPostChargingNotificationRequest() *PostChargingNotificationRequest {
	req := &PostChargingNotificationRequest{}

	return req
}

func (r *PostChargingNotificationRequest) SetChargingNotifyRequest(ChargingNotifyRequest models.ChargingNotifyRequest) {
	r.ChargingNotifyRequest = &ChargingNotifyRequest
}

type PostChargingNotificationResponse struct {
	ChargingNotifyResponse models.ChargingNotifyResponse
}

type PostChargingNotificationError struct {
	Location                            string
	Var3gppSbiTargetNfId                string
	PostChargingNotificationResponse400 models.PostChargingNotificationResponse400
	RedirectResponse                    models.RedirectResponse
}

func (a *DefaultApiService) PostChargingNotification(ctx context.Context, uri string, request *PostChargingNotificationRequest) error {
	var (
		localVarHTTPMethod   = strings.ToUpper("POST")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := uri

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ json", "application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.ChargingNotifyRequest

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		return nil
	case 204:
		return nil
	case 307:
		var v PostChargingNotificationError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return apiError
	case 308:
		var v PostChargingNotificationError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return apiError
	case 400:
		var v PostChargingNotificationError
		err = openapi.Deserialize(&v.PostChargingNotificationResponse400, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return err
		}
		apiError.ErrorModel = v
		return apiError
	default:
		return nil
	}
}

// func (a *DefaultApiService) PostChargingNotification(ctx context.Context, uri string, request *PostChargingNotificationRequest) (*PostChargingNotificationResponse, error) {
// 	var (
// 		localVarHTTPMethod   = strings.ToUpper("POST")
// 		localVarPostBody     interface{}
// 		localVarFormFileName string
// 		localVarFileName     string
// 		localVarFileBytes    []byte
// 		localVarReturnValue  PostChargingNotificationResponse
// 	)

// 	// create path and map variables
// 	localVarPath := uri

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	localVarHTTPContentTypes := []string{"application/json"}

// 	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

// 	// to determine the Accept header
// 	localVarHTTPHeaderAccepts := []string{"application/ json", "application/json", "application/problem+json"}

// 	// set Accept header
// 	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
// 	if localVarHTTPHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
// 	}

// 	// body params
// 	localVarPostBody = request.ChargingNotifyRequest

// 	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
// 	if err != nil || localVarHTTPResponse == nil {
// 		return nil, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = localVarHTTPResponse.Body.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	apiError := openapi.GenericOpenAPIError{
// 		RawBody:     localVarBody,
// 		ErrorStatus: localVarHTTPResponse.StatusCode,
// 	}

// 	switch localVarHTTPResponse.StatusCode {
// 	case 200:
// 		err = openapi.Deserialize(&localVarReturnValue.ChargingNotifyResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &localVarReturnValue, nil
// 	case 204:
// 		return &localVarReturnValue, nil
// 	case 307:
// 		var v PostChargingNotificationError
// 		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
// 		if err != nil {
// 			return nil, err
// 		}
// 		v.Location = localVarHTTPResponse.Header.Get("Location")
// 		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
// 		apiError.ErrorModel = v
// 		return nil, apiError
// 	case 308:
// 		var v PostChargingNotificationError
// 		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
// 		if err != nil {
// 			return nil, err
// 		}
// 		v.Location = localVarHTTPResponse.Header.Get("Location")
// 		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
// 		apiError.ErrorModel = v
// 		return nil, apiError
// 	case 400:
// 		var v PostChargingNotificationError
// 		err = openapi.Deserialize(&v.PostChargingNotificationResponse400, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
// 		if err != nil {
// 			return nil, err
// 		}
// 		apiError.ErrorModel = v
// 		return nil, apiError
// 	default:
// 		return &localVarReturnValue, nil
// 	}
// }

/*
DefaultApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ChargingDataRef - a unique identifier for a charging data resource in a PLMN
 * @param ChfConvergedChargingChargingDataRequest -

@return ReleaseChargingDataResponse
*/

// ReleaseChargingDataRequest
type ReleaseChargingDataRequest struct {
	ChargingDataRef                         *string
	ChfConvergedChargingChargingDataRequest *models.ChfConvergedChargingChargingDataRequest
}

func (r *ReleaseChargingDataRequest) SetChargingDataRef(ChargingDataRef string) {
	r.ChargingDataRef = &ChargingDataRef
}
func (r *ReleaseChargingDataRequest) SetChfConvergedChargingChargingDataRequest(ChfConvergedChargingChargingDataRequest models.ChfConvergedChargingChargingDataRequest) {
	r.ChfConvergedChargingChargingDataRequest = &ChfConvergedChargingChargingDataRequest
}

type ReleaseChargingDataResponse struct {
}

type ReleaseChargingDataError struct {
	Location                       string
	Var3gppSbiTargetNfId           string
	ProblemDetails                 models.ProblemDetails
	RedirectResponse               models.RedirectResponse
	ReleaseChargingDataResponse404 models.ReleaseChargingDataResponse404
}

func (a *DefaultApiService) ReleaseChargingData(ctx context.Context, request *ReleaseChargingDataRequest) (*ReleaseChargingDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReleaseChargingDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/chargingdata/{ChargingDataRef}/release"
	localVarPath = strings.Replace(localVarPath, "{"+"ChargingDataRef"+"}", openapi.StringOfValue(*request.ChargingDataRef), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.ChfConvergedChargingChargingDataRequest

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return &localVarReturnValue, nil
	case 307:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ReleaseChargingDataResponse404, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 410:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ReleaseChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}

/*
DefaultApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ChargingDataRef - a unique identifier for a charging data resource in a PLMN
 * @param ChfConvergedChargingChargingDataRequest -

@return UpdateChargingDataResponse
*/

// UpdateChargingDataRequest
type UpdateChargingDataRequest struct {
	ChargingDataRef                         *string
	ChfConvergedChargingChargingDataRequest *models.ChfConvergedChargingChargingDataRequest
}

func (r *UpdateChargingDataRequest) SetChargingDataRef(ChargingDataRef string) {
	r.ChargingDataRef = &ChargingDataRef
}
func (r *UpdateChargingDataRequest) SetChfConvergedChargingChargingDataRequest(ChfConvergedChargingChargingDataRequest models.ChfConvergedChargingChargingDataRequest) {
	r.ChfConvergedChargingChargingDataRequest = &ChfConvergedChargingChargingDataRequest
}

type UpdateChargingDataResponse struct {
	ChfConvergedChargingChargingDataResponse models.ChfConvergedChargingChargingDataResponse
}

type UpdateChargingDataError struct {
	Location                      string
	Var3gppSbiTargetNfId          string
	ProblemDetails                models.ProblemDetails
	RedirectResponse              models.RedirectResponse
	UpdateChargingDataResponse400 models.UpdateChargingDataResponse400
	UpdateChargingDataResponse403 models.UpdateChargingDataResponse403
	UpdateChargingDataResponse404 models.UpdateChargingDataResponse404
}

func (a *DefaultApiService) UpdateChargingData(ctx context.Context, request *UpdateChargingDataRequest) (*UpdateChargingDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateChargingDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/chargingdata/{ChargingDataRef}/update"
	localVarPath = strings.Replace(localVarPath, "{"+"ChargingDataRef"+"}", openapi.StringOfValue(*request.ChargingDataRef), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.ChfConvergedChargingChargingDataRequest

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.ChfConvergedChargingChargingDataResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 307:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.UpdateChargingDataResponse400, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.UpdateChargingDataResponse403, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.UpdateChargingDataResponse404, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 405:
		return &localVarReturnValue, nil
	case 408:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 410:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v UpdateChargingDataError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}