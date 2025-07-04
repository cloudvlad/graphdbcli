package spinner

import (
	"graphdbcli/internal/data_objects/spinner_status"

	"github.com/enescakir/emoji"
)

var cancelMessage = "Backup cancelled!"
var cancelEmoji = emoji.StopSign

var CreatingBackupStatuses = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Backup request preparations finished",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Backup request preparations failed",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: "Backup request preparations in-progress",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var BackupCreationStatuses = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Backup created successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Backup creation failed",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: "Backup creation in-progress",
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}
