package spinner

import (
	"graphdbcli/internal/data_objects/spinner_status"

	"github.com/enescakir/emoji"
)

var cancelMessage = "Data load have been canceled"

var ReadData = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Data has been accessed successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to read data file",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Reading data...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  emoji.StopSign,
	},
}

var CreateHTTPRequest = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "HTTP request has been created",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to create HTTP request",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Creating HTTP request...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  emoji.StopSign,
	},
}

var SendHTTPRequest = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Data has been loaded",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to send data",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Sending data...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  emoji.StopSign,
	},
}
