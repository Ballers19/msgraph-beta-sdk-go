package ediscovery
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type CaseOperationStatus int

const (
    NOTSTARTED_CASEOPERATIONSTATUS CaseOperationStatus = iota
    SUBMISSIONFAILED_CASEOPERATIONSTATUS
    RUNNING_CASEOPERATIONSTATUS
    SUCCEEDED_CASEOPERATIONSTATUS
    PARTIALLYSUCCEEDED_CASEOPERATIONSTATUS
    FAILED_CASEOPERATIONSTATUS
)

func (i CaseOperationStatus) String() string {
    return []string{"notStarted", "submissionFailed", "running", "succeeded", "partiallySucceeded", "failed"}[i]
}
func ParseCaseOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_CASEOPERATIONSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_CASEOPERATIONSTATUS
        case "submissionFailed":
            result = SUBMISSIONFAILED_CASEOPERATIONSTATUS
        case "running":
            result = RUNNING_CASEOPERATIONSTATUS
        case "succeeded":
            result = SUCCEEDED_CASEOPERATIONSTATUS
        case "partiallySucceeded":
            result = PARTIALLYSUCCEEDED_CASEOPERATIONSTATUS
        case "failed":
            result = FAILED_CASEOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown CaseOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeCaseOperationStatus(values []CaseOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
