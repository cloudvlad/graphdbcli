package create

import (
	"graphdbcli/internal/data_objects/spinner_status"

	"github.com/enescakir/emoji"
)

var cancelMessage = "Instance destruction cancelled!"
var cancelEmoji = emoji.StopSign

var CheckingForPresentInstance = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "The specified instance was found",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "The specified instance was not found",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Checking for the specified instance in-progress...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var FetchingInstancePID = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "The instance process ID was fetched successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "There was problem fetching the instance process ID",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Fetching instance proced ID in progress...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var CheckingIsInstanceRunning = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "The instance is running",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "The instance is not running. Skipping...",
		Status:  emoji.RightArrow,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Determining whether the instance is running...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var StoppingInstance = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "The instance has been stopped",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "There was problem whilst stopping the instance...",
		Status:  emoji.RightArrow,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Stopping the instance...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var PurgeInstanceSpace = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "The instance has been purged",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "The instance is not running. Skipping",
		Status:  emoji.RightArrow,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Determining whether the instance is running...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}
