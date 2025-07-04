package spinner_status

import "github.com/enescakir/emoji"

// SpinnerStatuses provides the structure that is unified for the common spinner
// It provides the 4 possible statuses:
// SuccessMessage specifies the message that will be displayed when the operation is executed successfully
// FailureMessage specifies the message that will be displayed when the operation failed
// InProgressMessage specifies the message that will be displayed whilst the operation is still in progress
// CancelledMessage specifies the message that will be displayed when the operation is canceled
type SpinnerStatuses struct {
	SuccessMessage    SpinnerStatusMessage
	FailureMessage    SpinnerStatusMessage
	InProgressMessage SpinnerStatusMessage
	CancelledMessage  SpinnerStatusMessage
}

// SpinnerStatusMessage provide the base structure for storing the spinner status
type SpinnerStatusMessage struct {
	Message string
	Status  emoji.Emoji
}
