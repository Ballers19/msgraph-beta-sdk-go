package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ThreatExpectedAssessment int

const (
    BLOCK_THREATEXPECTEDASSESSMENT ThreatExpectedAssessment = iota
    UNBLOCK_THREATEXPECTEDASSESSMENT
)

func (i ThreatExpectedAssessment) String() string {
    return []string{"block", "unblock"}[i]
}
func ParseThreatExpectedAssessment(v string) (interface{}, error) {
    result := BLOCK_THREATEXPECTEDASSESSMENT
    switch v {
        case "block":
            result = BLOCK_THREATEXPECTEDASSESSMENT
        case "unblock":
            result = UNBLOCK_THREATEXPECTEDASSESSMENT
        default:
            return 0, errors.New("Unknown ThreatExpectedAssessment value: " + v)
    }
    return &result, nil
}
func SerializeThreatExpectedAssessment(values []ThreatExpectedAssessment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
