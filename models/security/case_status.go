package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type CaseStatus int

const (
    UNKNOWN_CASESTATUS CaseStatus = iota
    ACTIVE_CASESTATUS
    PENDINGDELETE_CASESTATUS
    CLOSING_CASESTATUS
    CLOSED_CASESTATUS
    CLOSEDWITHERROR_CASESTATUS
    UNKNOWNFUTUREVALUE_CASESTATUS
)

func (i CaseStatus) String() string {
    return []string{"unknown", "active", "pendingDelete", "closing", "closed", "closedWithError", "unknownFutureValue"}[i]
}
func ParseCaseStatus(v string) (interface{}, error) {
    result := UNKNOWN_CASESTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CASESTATUS
        case "active":
            result = ACTIVE_CASESTATUS
        case "pendingDelete":
            result = PENDINGDELETE_CASESTATUS
        case "closing":
            result = CLOSING_CASESTATUS
        case "closed":
            result = CLOSED_CASESTATUS
        case "closedWithError":
            result = CLOSEDWITHERROR_CASESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CASESTATUS
        default:
            return 0, errors.New("Unknown CaseStatus value: " + v)
    }
    return &result, nil
}
func SerializeCaseStatus(values []CaseStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
