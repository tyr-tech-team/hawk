package server

// AddScheduleRes -
type AddScheduleRes struct {
	// List -
	List []*AddScheduleDetail
}

// UpdateScheduleRes -
type UpdateScheduleRes struct {
	// List -
	List []*Result
}

// CancelScheduleRes -
type CancelScheduleRes struct {
	// List -
	List []*Result
}

// AddScheduleDetail -
type AddScheduleDetail struct {
	// ScheduleID -
	ScheduleID string

	// Result -
	Result *Result
}

// Result -
type Result struct {
	// Success -
	Success bool

	// ErrMsg -
	ErrMsg error
}

const (
	// ADD_SCHEDULE_REPLY_TOPIC -
	ADD_SCHEDULE_REPLY_TOPIC = "sm-add-schedule-reply"

	// UPDATE_SCHEDULE_REPLY_TOPIC -
	UPDATE_SCHEDULE_REPLY_TOPIC = "sm-update-schedule-reply"

	// CANCEL_SCHEDULE_REPLY_TOPIC -
	CANCEL_SCHEDULE_REPLY_TOPIC = "sm-cancel-schedule-reply"
)
