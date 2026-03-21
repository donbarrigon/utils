package herror

const (
	// 4xx Client Errors
	StatusTextBadRequest                   = "Bad Request"                     // 400
	StatusTextUnauthorized                 = "Unauthorized"                    // 401
	StatusTextPaymentRequired              = "Payment Required"                // 402
	StatusTextForbidden                    = "Forbidden"                       // 403
	StatusTextNotFound                     = "Not Found"                       // 404
	StatusTextMethodNotAllowed             = "Method Not Allowed"              // 405
	StatusTextNotAcceptable                = "Not Acceptable"                  // 406
	StatusTextProxyAuthRequired            = "Proxy Authentication Required"   // 407
	StatusTextRequestTimeout               = "Request Timeout"                 // 408
	StatusTextConflict                     = "Conflict"                        // 409
	StatusTextGone                         = "Gone"                            // 410
	StatusTextLengthRequired               = "Length Required"                 // 411
	StatusTextPreconditionFailed           = "Precondition Failed"             // 412
	StatusTextRequestEntityTooLarge        = "Request Entity Too Large"        // 413
	StatusTextRequestURITooLong            = "Request URI Too Long"            // 414
	StatusTextUnsupportedMediaType         = "Unsupported Media Type"          // 415
	StatusTextRequestedRangeNotSatisfiable = "Requested Range Not Satisfiable" // 416
	StatusTextExpectationFailed            = "Expectation Failed"              // 417
	StatusTextTeapot                       = "I'm a teapot"                    // 418
	StatusTextMisdirectedRequest           = "Misdirected Request"             // 421
	StatusTextUnprocessableEntity          = "Unprocessable Entity"            // 422
	StatusTextLocked                       = "Locked"                          // 423
	StatusTextFailedDependency             = "Failed Dependency"               // 424
	StatusTextTooEarly                     = "Too Early"                       // 425
	StatusTextUpgradeRequired              = "Upgrade Required"                // 426
	StatusTextPreconditionRequired         = "Precondition Required"           // 428
	StatusTextTooManyRequests              = "Too Many Requests"               // 429
	StatusTextRequestHeaderFieldsTooLarge  = "Request Header Fields Too Large" // 431
	StatusTextUnavailableForLegalReasons   = "Unavailable For Legal Reasons"   // 451

	// 5xx Server Errors
	StatusTextInternalServerError           = "Internal Server Error"           // 500
	StatusTextNotImplemented                = "Not Implemented"                 // 501
	StatusTextBadGateway                    = "Bad Gateway"                     // 502
	StatusTextServiceUnavailable            = "Service Unavailable"             // 503
	StatusTextGatewayTimeout                = "Gateway Timeout"                 // 504
	StatusTextHTTPVersionNotSupported       = "HTTP Version Not Supported"      // 505
	StatusTextVariantAlsoNegotiates         = "Variant Also Negotiates"         // 506
	StatusTextInsufficientStorage           = "Insufficient Storage"            // 507
	StatusTextLoopDetected                  = "Loop Detected"                   // 508
	StatusTextNotExtended                   = "Not Extended"                    // 510
	StatusTextNetworkAuthenticationRequired = "Network Authentication Required" // 511

	// 4xx Client Errors
	StatusMessageBadRequest                   = "The request is malformed or missing required parameters. Please verify the syntax and try again."
	StatusMessageUnauthorized                 = "Authentication is required. Please provide valid credentials to access this resource."
	StatusMessagePaymentRequired              = "Access to this resource requires a valid payment or active subscription."
	StatusMessageForbidden                    = "You do not have sufficient permissions to access this resource."
	StatusMessageNotFound                     = "The requested resource does not exist or may have been moved."
	StatusMessageMethodNotAllowed             = "The HTTP method used is not allowed for this endpoint."
	StatusMessageNotAcceptable                = "The server cannot produce a response matching the criteria specified in the Accept headers."
	StatusMessageProxyAuthRequired            = "Proxy authentication is required to complete this request."
	StatusMessageRequestTimeout               = "The request took too long to complete. Please try again."
	StatusMessageConflict                     = "The request could not be completed due to a conflict with the current state of the resource."
	StatusMessageGone                         = "The requested resource has been permanently removed and is no longer available."
	StatusMessageLengthRequired               = "The Content-Length header is required but was not provided."
	StatusMessagePreconditionFailed           = "One or more preconditions in the request headers could not be met."
	StatusMessageRequestEntityTooLarge        = "The request payload exceeds the maximum size allowed by the server."
	StatusMessageRequestURITooLong            = "The request URI exceeds the maximum length the server is able to process."
	StatusMessageUnsupportedMediaType         = "The media type of the request is not supported. Please check the Content-Type header."
	StatusMessageRequestedRangeNotSatisfiable = "The requested byte range cannot be satisfied for this resource."
	StatusMessageExpectationFailed            = "The server could not meet the expectation specified in the Expect request header."
	StatusMessageTeapot                       = "I'm a teapot. (RFC 2324 Easter egg — brewing coffee is not supported.)"
	StatusMessageMisdirectedRequest           = "The request was sent to a server that is not configured to produce a response for this resource."
	StatusMessageUnprocessableEntity          = "The submitted entity has validation errors. Please review the fields in the data below."
	StatusMessageLocked                       = "The resource is currently locked and cannot be modified."
	StatusMessageFailedDependency             = "This request could not be completed because a dependent request failed."
	StatusMessageTooEarly                     = "The server is not willing to process this request as it may be replayed."
	StatusMessageUpgradeRequired              = "This request must be made using a more recent protocol version."
	StatusMessagePreconditionRequired         = "This endpoint requires the request to include conditional headers."
	StatusMessageTooManyRequests              = "You have exceeded the allowed request rate. Please wait before retrying."
	StatusMessageRequestHeaderFieldsTooLarge  = "One or more request header fields exceed the maximum allowed size."
	StatusMessageUnavailableForLegalReasons   = "Access to this resource has been restricted for legal reasons."

	// 5xx Server Errors
	StatusMessageInternalServerError           = "An unexpected error occurred on the server. Please try again later."
	StatusMessageNotImplemented                = "The server does not support the functionality required to fulfill this request."
	StatusMessageBadGateway                    = "The server received an invalid or unexpected response from an upstream service."
	StatusMessageServiceUnavailable            = "The server is temporarily unavailable due to maintenance or high load. Please try again later."
	StatusMessageGatewayTimeout                = "The upstream server did not respond in time. Please try again later."
	StatusMessageHTTPVersionNotSupported       = "The HTTP protocol version used in this request is not supported by the server."
	StatusMessageVariantAlsoNegotiates         = "The server has an internal configuration error and cannot complete content negotiation."
	StatusMessageInsufficientStorage           = "The server does not have enough storage to process this request."
	StatusMessageLoopDetected                  = "The server detected an infinite loop while processing the request and aborted."
	StatusMessageNotExtended                   = "The request requires additional extensions that the server does not support."
	StatusMessageNetworkAuthenticationRequired = "Network authentication is required before this request can be processed."
)
