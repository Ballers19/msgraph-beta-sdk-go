package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AccessReviewHistoryStatus int

const (
    DONE_ACCESSREVIEWHISTORYSTATUS AccessReviewHistoryStatus = iota
    INPROGRESS_ACCESSREVIEWHISTORYSTATUS
    ERROR_ACCESSREVIEWHISTORYSTATUS
    REQUESTED_ACCESSREVIEWHISTORYSTATUS
    UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYSTATUS
)

func (i AccessReviewHistoryStatus) String() string {
    return []string{"done", "inprogress", "error", "requested", "unknownFutureValue"}[i]
}
func ParseAccessReviewHistoryStatus(v string) (interface{}, error) {
    result := DONE_ACCESSREVIEWHISTORYSTATUS
    switch v {
        case "done":
            result = DONE_ACCESSREVIEWHISTORYSTATUS
        case "inprogress":
            result = INPROGRESS_ACCESSREVIEWHISTORYSTATUS
        case "error":
            result = ERROR_ACCESSREVIEWHISTORYSTATUS
        case "requested":
            result = REQUESTED_ACCESSREVIEWHISTORYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYSTATUS
        default:
            return 0, errors.New("Unknown AccessReviewHistoryStatus value: " + v)
    }
    return &result, nil
}
func SerializeAccessReviewHistoryStatus(values []AccessReviewHistoryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
