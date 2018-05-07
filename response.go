package lapns




var(
	PushResponseErrCode = map[uint8]string{
		0 : "No errors encountered",
		1 : "Processing error",
		2 : "Missing device token",
		3 : "Missing topic",
		4 : "Missing payload",
		5 : "Invalid token size",
		6 : "Invalid topic size",
		7 : "Invalid payload size",
		8 : "Invalid token",
		10 : "Shutdown",
		128 : "Protocol error (APNs could not parse the notification)",
		255 : "None (unknown)",
	}
)

type Response struct {
	Sucess bool
	Err error
}
