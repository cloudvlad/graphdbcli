package spinner

import (
	"graphdbcli/internal/data_objects/spinner_status"

	"github.com/enescakir/emoji"
)

var InitializingWorkbenchStatuses = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench initialization finished successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench initialization failed",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench initialization in-progress",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench initialization cancelled",
		Status:  emoji.StopSign,
	},
}

var StartingWorkbenchStatuses = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench started successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench starting failed",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench starting in-progress",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench starting cancelled",
		Status:  emoji.StopSign,
	},
}
