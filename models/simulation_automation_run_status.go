package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type SimulationAutomationRunStatus int

const (
    UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS SimulationAutomationRunStatus = iota
    RUNNING_SIMULATIONAUTOMATIONRUNSTATUS
    SUCCEEDED_SIMULATIONAUTOMATIONRUNSTATUS
    FAILED_SIMULATIONAUTOMATIONRUNSTATUS
    SKIPPED_SIMULATIONAUTOMATIONRUNSTATUS
    UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONRUNSTATUS
)

func (i SimulationAutomationRunStatus) String() string {
    return []string{"unknown", "running", "succeeded", "failed", "skipped", "unknownFutureValue"}[i]
}
func ParseSimulationAutomationRunStatus(v string) (interface{}, error) {
    result := UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_SIMULATIONAUTOMATIONRUNSTATUS
        case "running":
            result = RUNNING_SIMULATIONAUTOMATIONRUNSTATUS
        case "succeeded":
            result = SUCCEEDED_SIMULATIONAUTOMATIONRUNSTATUS
        case "failed":
            result = FAILED_SIMULATIONAUTOMATIONRUNSTATUS
        case "skipped":
            result = SKIPPED_SIMULATIONAUTOMATIONRUNSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONRUNSTATUS
        default:
            return 0, errors.New("Unknown SimulationAutomationRunStatus value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAutomationRunStatus(values []SimulationAutomationRunStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
