package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type RiskDetectionTimingType int

const (
    NOTDEFINED_RISKDETECTIONTIMINGTYPE RiskDetectionTimingType = iota
    REALTIME_RISKDETECTIONTIMINGTYPE
    NEARREALTIME_RISKDETECTIONTIMINGTYPE
    OFFLINE_RISKDETECTIONTIMINGTYPE
    UNKNOWNFUTUREVALUE_RISKDETECTIONTIMINGTYPE
)

func (i RiskDetectionTimingType) String() string {
    return []string{"notDefined", "realtime", "nearRealtime", "offline", "unknownFutureValue"}[i]
}
func ParseRiskDetectionTimingType(v string) (interface{}, error) {
    result := NOTDEFINED_RISKDETECTIONTIMINGTYPE
    switch v {
        case "notDefined":
            result = NOTDEFINED_RISKDETECTIONTIMINGTYPE
        case "realtime":
            result = REALTIME_RISKDETECTIONTIMINGTYPE
        case "nearRealtime":
            result = NEARREALTIME_RISKDETECTIONTIMINGTYPE
        case "offline":
            result = OFFLINE_RISKDETECTIONTIMINGTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RISKDETECTIONTIMINGTYPE
        default:
            return 0, errors.New("Unknown RiskDetectionTimingType value: " + v)
    }
    return &result, nil
}
func SerializeRiskDetectionTimingType(values []RiskDetectionTimingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
