package externalconnectors
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ConnectionOperationStatus int

const (
    UNSPECIFIED_CONNECTIONOPERATIONSTATUS ConnectionOperationStatus = iota
    INPROGRESS_CONNECTIONOPERATIONSTATUS
    COMPLETED_CONNECTIONOPERATIONSTATUS
    FAILED_CONNECTIONOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_CONNECTIONOPERATIONSTATUS
)

func (i ConnectionOperationStatus) String() string {
    return []string{"unspecified", "inprogress", "completed", "failed", "unknownFutureValue"}[i]
}
func ParseConnectionOperationStatus(v string) (interface{}, error) {
    result := UNSPECIFIED_CONNECTIONOPERATIONSTATUS
    switch v {
        case "unspecified":
            result = UNSPECIFIED_CONNECTIONOPERATIONSTATUS
        case "inprogress":
            result = INPROGRESS_CONNECTIONOPERATIONSTATUS
        case "completed":
            result = COMPLETED_CONNECTIONOPERATIONSTATUS
        case "failed":
            result = FAILED_CONNECTIONOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONNECTIONOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown ConnectionOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeConnectionOperationStatus(values []ConnectionOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
