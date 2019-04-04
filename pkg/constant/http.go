package constant

const (
	HTTPScheme = "http"
	TCPScheme  = "tcp"
	UDPScheme  = "udp"

	HTTPResultStatus  = "status"
	HTTPResultError   = "err"
	HTTPResultData    = "data"
	HTTPResultMessage = "message"

	HTTPResultStatusOK    = "OK"
	HTTPResultStatusError = "Error"

	HeaderAuthorization   = "Authorization"
	HeaderContentType     = "Content-Type"
	HeaderAcceptRanges    = "Accept-Ranges"
	HeaderContentRange    = "Content-Range"
	HeaderContentEncoding = "Content-Encoding"
	HeaderContentLength   = "Content-Length"

	LoginURL = "/login" //post
	RootURL  = "/"

	HealthzURL = "/healthz"
	StatusURL  = "/status"

	MetricsURL = "/metrics"
	ProbeURL   = "/probe"
)
