// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

const (

	// ErrCodeDetectedLanguageLowConfidenceException for service response error code
	// "DetectedLanguageLowConfidenceException".
	//
	// The confidence that Amazon Comprehend accurately detected the source language
	// is low. If a low confidence level is acceptable for your application, you
	// can use the language in the exception to call Amazon Translate again. For
	// more information, see the DetectDominantLanguage (https://docs.aws.amazon.com/comprehend/latest/dg/API_DetectDominantLanguage.html)
	// operation in the Amazon Comprehend Developer Guide.
	ErrCodeDetectedLanguageLowConfidenceException = "DetectedLanguageLowConfidenceException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal server error occurred. Retry your request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is invalid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// Amazon Translate is unavailable. Retry your request later.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeTextSizeLimitExceededException for service response error code
	// "TextSizeLimitExceededException".
	//
	// The size of the input text exceeds the length constraint for the Text field.
	// Try again with a shorter text.
	ErrCodeTextSizeLimitExceededException = "TextSizeLimitExceededException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// The number of requests exceeds the limit. Resubmit your request later.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnsupportedLanguagePairException for service response error code
	// "UnsupportedLanguagePairException".
	//
	// Amazon Translate cannot translate input text in the source language into
	// this target language. For more information, see how-to-error-msg.
	ErrCodeUnsupportedLanguagePairException = "UnsupportedLanguagePairException"
)
