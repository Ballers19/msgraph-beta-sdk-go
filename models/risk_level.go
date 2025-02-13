package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type RiskLevel int

const (
    LOW_RISKLEVEL RiskLevel = iota
    MEDIUM_RISKLEVEL
    HIGH_RISKLEVEL
    HIDDEN_RISKLEVEL
    NONE_RISKLEVEL
    UNKNOWNFUTUREVALUE_RISKLEVEL
)

func (i RiskLevel) String() string {
    return []string{"low", "medium", "high", "hidden", "none", "unknownFutureValue"}[i]
}
func ParseRiskLevel(v string) (interface{}, error) {
    result := LOW_RISKLEVEL
    switch v {
        case "low":
            result = LOW_RISKLEVEL
        case "medium":
            result = MEDIUM_RISKLEVEL
        case "high":
            result = HIGH_RISKLEVEL
        case "hidden":
            result = HIDDEN_RISKLEVEL
        case "none":
            result = NONE_RISKLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RISKLEVEL
        default:
            return 0, errors.New("Unknown RiskLevel value: " + v)
    }
    return &result, nil
}
func SerializeRiskLevel(values []RiskLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
