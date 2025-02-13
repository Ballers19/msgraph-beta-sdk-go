package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ScheduleChangeState int

const (
    PENDING_SCHEDULECHANGESTATE ScheduleChangeState = iota
    APPROVED_SCHEDULECHANGESTATE
    DECLINED_SCHEDULECHANGESTATE
    UNKNOWNFUTUREVALUE_SCHEDULECHANGESTATE
)

func (i ScheduleChangeState) String() string {
    return []string{"pending", "approved", "declined", "unknownFutureValue"}[i]
}
func ParseScheduleChangeState(v string) (interface{}, error) {
    result := PENDING_SCHEDULECHANGESTATE
    switch v {
        case "pending":
            result = PENDING_SCHEDULECHANGESTATE
        case "approved":
            result = APPROVED_SCHEDULECHANGESTATE
        case "declined":
            result = DECLINED_SCHEDULECHANGESTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SCHEDULECHANGESTATE
        default:
            return 0, errors.New("Unknown ScheduleChangeState value: " + v)
    }
    return &result, nil
}
func SerializeScheduleChangeState(values []ScheduleChangeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
