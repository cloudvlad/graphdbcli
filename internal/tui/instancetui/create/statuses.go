package create

import (
	"graphdbcli/internal/data_objects/spinner_status"

	"github.com/enescakir/emoji"
)

var cancelMessage = "Instance creation cancelled!"
var cancelEmoji = emoji.StopSign

var CreatingInstanceStructure = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Instance structure created successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to create instance structure",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Instance structure setup is in progress...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var SettingUpInstancePort = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Instance port was configured",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to set up instance port",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Setting up instance port...",
		Status:  emoji.Emoji(""),
	},
}

var SettingUpProperties = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Instance properties set successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to set up properties",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Instance properties setup is in progress...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var SettingUpLicense = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Instance license set successfully",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to set instance license",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Setting up instance license...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var CheckingIsInstanceAccessible = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "The GraphDB instance is accessible",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "The GraphDB instance is not accessible",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Checking if the instance is accessible",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var StartingGraphDBInstance = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Successfully started GraphDB instance",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to start GraphDB instance",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " GraphDB is starting...",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}

var CleaningUpGraphDBInstance = spinner_status.SpinnerStatuses{
	SuccessMessage: spinner_status.SpinnerStatusMessage{
		Message: "Successfully reverted",
		Status:  emoji.CheckMark,
	},
	FailureMessage: spinner_status.SpinnerStatusMessage{
		Message: "Failed to revert the creation of the instance",
		Status:  emoji.CrossMark,
	},
	InProgressMessage: spinner_status.SpinnerStatusMessage{
		Message: " Reverting in progress",
		Status:  emoji.Emoji(""),
	},
	CancelledMessage: spinner_status.SpinnerStatusMessage{
		Message: cancelMessage,
		Status:  cancelEmoji,
	},
}
