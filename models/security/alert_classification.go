package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AlertClassification int

const (
    UNKNOWN_ALERTCLASSIFICATION AlertClassification = iota
    FALSEPOSITIVE_ALERTCLASSIFICATION
    TRUEPOSITIVE_ALERTCLASSIFICATION
    INFORMATIONALEXPECTEDACTIVITY_ALERTCLASSIFICATION
    UNKNOWNFUTUREVALUE_ALERTCLASSIFICATION
)

func (i AlertClassification) String() string {
    return []string{"unknown", "falsePositive", "truePositive", "informationalExpectedActivity", "unknownFutureValue"}[i]
}
func ParseAlertClassification(v string) (interface{}, error) {
    result := UNKNOWN_ALERTCLASSIFICATION
    switch v {
        case "unknown":
            result = UNKNOWN_ALERTCLASSIFICATION
        case "falsePositive":
            result = FALSEPOSITIVE_ALERTCLASSIFICATION
        case "truePositive":
            result = TRUEPOSITIVE_ALERTCLASSIFICATION
        case "informationalExpectedActivity":
            result = INFORMATIONALEXPECTEDACTIVITY_ALERTCLASSIFICATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTCLASSIFICATION
        default:
            return 0, errors.New("Unknown AlertClassification value: " + v)
    }
    return &result, nil
}
func SerializeAlertClassification(values []AlertClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
