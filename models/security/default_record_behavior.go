package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type DefaultRecordBehavior int

const (
    STARTLOCKED_DEFAULTRECORDBEHAVIOR DefaultRecordBehavior = iota
    STARTUNLOCKED_DEFAULTRECORDBEHAVIOR
    UNKNOWNFUTUREVALUE_DEFAULTRECORDBEHAVIOR
)

func (i DefaultRecordBehavior) String() string {
    return []string{"startLocked", "startUnlocked", "unknownFutureValue"}[i]
}
func ParseDefaultRecordBehavior(v string) (interface{}, error) {
    result := STARTLOCKED_DEFAULTRECORDBEHAVIOR
    switch v {
        case "startLocked":
            result = STARTLOCKED_DEFAULTRECORDBEHAVIOR
        case "startUnlocked":
            result = STARTUNLOCKED_DEFAULTRECORDBEHAVIOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEFAULTRECORDBEHAVIOR
        default:
            return 0, errors.New("Unknown DefaultRecordBehavior value: " + v)
    }
    return &result, nil
}
func SerializeDefaultRecordBehavior(values []DefaultRecordBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
