package spinner

import (
	"graphdbcli/internal/data_objects/spinner_status"

	"github.com/enescakir/emoji"
)

var CheckingPrerequisites = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "All prerequisites are fulfilled",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "A prerequisite was not fulfilled",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Prerequisites checking in progress...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: "Prerequisites check cancelled",
		Status:  emoji.StopSign,
	},
}

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
		Message: " Workbench initialization in progress",
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
		Message: " Workbench is starting",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: "Workbench starting cancelled",
		Status:  emoji.StopSign,
	},
}
