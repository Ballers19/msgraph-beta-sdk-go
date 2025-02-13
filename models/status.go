package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type Status int

const (
    ACTIVE_STATUS Status = iota
    UPDATED_STATUS
    DELETED_STATUS
    IGNORED_STATUS
    UNKNOWNFUTUREVALUE_STATUS
)

func (i Status) String() string {
    return []string{"active", "updated", "deleted", "ignored", "unknownFutureValue"}[i]
}
func ParseStatus(v string) (interface{}, error) {
    result := ACTIVE_STATUS
    switch v {
        case "active":
            result = ACTIVE_STATUS
        case "updated":
            result = UPDATED_STATUS
        case "deleted":
            result = DELETED_STATUS
        case "ignored":
            result = IGNORED_STATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_STATUS
        default:
            return 0, errors.New("Unknown Status value: " + v)
    }
    return &result, nil
}
func SerializeStatus(values []Status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
