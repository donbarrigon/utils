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

	// 4xx Client Errors - Mensajes descriptivos
	StatusMessageBadRequest                   = "The request could not be understood or was missing required parameters"
	StatusMessageUnauthorized                 = "Authentication is required to access this resource"
	StatusMessagePaymentRequired              = "Payment is required to access this resource"
	StatusMessageForbidden                    = "You don't have permission to access this resource"
	StatusMessageNotFound                     = "The requested resource could not be found"
	StatusMessageMethodNotAllowed             = "The request method is not supported for this resource"
	StatusMessageNotAcceptable                = "The requested resource cannot generate content acceptable according to the Accept headers"
	StatusMessageProxyAuthRequired            = "Authentication with the proxy is required"
	StatusMessageRequestTimeout               = "The server timed out waiting for the request"
	StatusMessageConflict                     = "The request conflicts with the current state of the server"
	StatusMessageGone                         = "The requested resource is no longer available"
	StatusMessageLengthRequired               = "The request did not specify the length of its content"
	StatusMessagePreconditionFailed           = "One or more conditions in the request header fields evaluated to false"
	StatusMessageRequestEntityTooLarge        = "The request is larger than the server is willing or able to process"
	StatusMessageRequestURITooLong            = "The URI provided was too long for the server to process"
	StatusMessageUnsupportedMediaType         = "The request entity has a media type which the server does not support"
	StatusMessageRequestedRangeNotSatisfiable = "The requested range cannot be satisfied"
	StatusMessageExpectationFailed            = "The server cannot meet the requirements of the Expect request header"
	StatusMessageTeapot                       = "I'm a teapot (this is an Easter egg from the HTTP specification)"
	StatusMessageMisdirectedRequest           = "The request was directed at a server that is not able to produce a response"
	StatusMessageUnprocessableEntity          = "The request was well-formed but contains semantic errors"
	StatusMessageLocked                       = "The resource that is being accessed is locked"
	StatusMessageFailedDependency             = "The request failed due to failure of a previous request"
	StatusMessageTooEarly                     = "The server is unwilling to risk processing a request that might be replayed"
	StatusMessageUpgradeRequired              = "The client should switch to a different protocol"
	StatusMessagePreconditionRequired         = "The server requires the request to be conditional"
	StatusMessageTooManyRequests              = "You have sent too many requests in a given amount of time"
	StatusMessageRequestHeaderFieldsTooLarge  = "The request header fields are too large"
	StatusMessageUnavailableForLegalReasons   = "This resource is unavailable for legal reasons"

	// 5xx Server Errors - Mensajes descriptivos
	StatusMessageInternalServerError           = "The server encountered an unexpected condition that prevented it from fulfilling the request"
	StatusMessageNotImplemented                = "The server does not support the functionality required to fulfill the request"
	StatusMessageBadGateway                    = "The server received an invalid response from an upstream server"
	StatusMessageServiceUnavailable            = "The server is currently unable to handle the request due to temporary overload or maintenance"
	StatusMessageGatewayTimeout                = "The server did not receive a timely response from an upstream server"
	StatusMessageHTTPVersionNotSupported       = "The server does not support the HTTP protocol version used in the request"
	StatusMessageVariantAlsoNegotiates         = "The server has an internal configuration error"
	StatusMessageInsufficientStorage           = "The server is unable to store the representation needed to complete the request"
	StatusMessageLoopDetected                  = "The server detected an infinite loop while processing the request"
	StatusMessageNotExtended                   = "Further extensions to the request are required for the server to fulfill it"
	StatusMessageNetworkAuthenticationRequired = "You need to authenticate to gain network access"
)
