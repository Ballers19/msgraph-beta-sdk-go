package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ThreatAssessmentResultType int

const (
    CHECKPOLICY_THREATASSESSMENTRESULTTYPE ThreatAssessmentResultType = iota
    RESCAN_THREATASSESSMENTRESULTTYPE
    UNKNOWNFUTUREVALUE_THREATASSESSMENTRESULTTYPE
)

func (i ThreatAssessmentResultType) String() string {
    return []string{"checkPolicy", "rescan", "unknownFutureValue"}[i]
}
func ParseThreatAssessmentResultType(v string) (interface{}, error) {
    result := CHECKPOLICY_THREATASSESSMENTRESULTTYPE
    switch v {
        case "checkPolicy":
            result = CHECKPOLICY_THREATASSESSMENTRESULTTYPE
        case "rescan":
            result = RESCAN_THREATASSESSMENTRESULTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_THREATASSESSMENTRESULTTYPE
        default:
            return 0, errors.New("Unknown ThreatAssessmentResultType value: " + v)
    }
    return &result, nil
}
func SerializeThreatAssessmentResultType(values []ThreatAssessmentResultType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
