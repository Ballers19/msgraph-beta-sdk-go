package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type PrintOperationProcessingState int

const (
    NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE PrintOperationProcessingState = iota
    RUNNING_PRINTOPERATIONPROCESSINGSTATE
    SUCCEEDED_PRINTOPERATIONPROCESSINGSTATE
    FAILED_PRINTOPERATIONPROCESSINGSTATE
    UNKNOWNFUTUREVALUE_PRINTOPERATIONPROCESSINGSTATE
)

func (i PrintOperationProcessingState) String() string {
    return []string{"notStarted", "running", "succeeded", "failed", "unknownFutureValue"}[i]
}
func ParsePrintOperationProcessingState(v string) (interface{}, error) {
    result := NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE
    switch v {
        case "notStarted":
            result = NOTSTARTED_PRINTOPERATIONPROCESSINGSTATE
        case "running":
            result = RUNNING_PRINTOPERATIONPROCESSINGSTATE
        case "succeeded":
            result = SUCCEEDED_PRINTOPERATIONPROCESSINGSTATE
        case "failed":
            result = FAILED_PRINTOPERATIONPROCESSINGSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRINTOPERATIONPROCESSINGSTATE
        default:
            return 0, errors.New("Unknown PrintOperationProcessingState value: " + v)
    }
    return &result, nil
}
func SerializePrintOperationProcessingState(values []PrintOperationProcessingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
