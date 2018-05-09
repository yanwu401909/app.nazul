package models

import (

)

type ApiRequest struct{
	AuthToken string  `json:"authToken,omitempty"`
	Platform string `json:"platform,omitempty"`
	Agent string `json:"agent,omitempty"`
	Version string `json:"version,omitempty"`
}

type ApiResponse struct{
	ResultCode int32 `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

type EchoResponse struct{
	ApiResponse
	// ResultCode int32 `json:"resultCode,omitempty"`
	// ResultMessage string `json:"resultMessage,omitempty"`
	Data string `json:"data,omitempty"`
}