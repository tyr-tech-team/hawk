package server

type AddScheduleRes struct {
	List []*AddScheduleDetail
}

type UpdateScheduleRes struct {
	List []*Result
}

type CancelScheduleRes struct {
	List []*Result
}

type AddScheduleDetail struct {
	ScheduleID string
	Result     *Result
}

type Result struct {
	Success bool
	ErrMsg  error
}

const (
	ADD_SCHEDULE_REPLY_TOPIC    = "sm-add-schedule-reply"
	UPDATE_SCHEDULE_REPLY_TOPIC = "sm-update-schedule-reply"
	CANCEL_SCHEDULE_REPLY_TOPIC = "sm-cancel-schedule-reply"
)
