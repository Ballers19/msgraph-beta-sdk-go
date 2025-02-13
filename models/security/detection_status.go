package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type DetectionStatus int

const (
    DETECTED_DETECTIONSTATUS DetectionStatus = iota
    BLOCKED_DETECTIONSTATUS
    PREVENTED_DETECTIONSTATUS
    UNKNOWNFUTUREVALUE_DETECTIONSTATUS
)

func (i DetectionStatus) String() string {
    return []string{"detected", "blocked", "prevented", "unknownFutureValue"}[i]
}
func ParseDetectionStatus(v string) (interface{}, error) {
    result := DETECTED_DETECTIONSTATUS
    switch v {
        case "detected":
            result = DETECTED_DETECTIONSTATUS
        case "blocked":
            result = BLOCKED_DETECTIONSTATUS
        case "prevented":
            result = PREVENTED_DETECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DETECTIONSTATUS
        default:
            return 0, errors.New("Unknown DetectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeDetectionStatus(values []DetectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
