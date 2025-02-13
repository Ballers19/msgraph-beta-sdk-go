package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type BehaviorDuringRetentionPeriod int

const (
    DONOTRETAIN_BEHAVIORDURINGRETENTIONPERIOD BehaviorDuringRetentionPeriod = iota
    RETAIN_BEHAVIORDURINGRETENTIONPERIOD
    RETAINASRECORD_BEHAVIORDURINGRETENTIONPERIOD
    RETAINASREGULATORYRECORD_BEHAVIORDURINGRETENTIONPERIOD
    UNKNOWNFUTUREVALUE_BEHAVIORDURINGRETENTIONPERIOD
)

func (i BehaviorDuringRetentionPeriod) String() string {
    return []string{"doNotRetain", "retain", "retainAsRecord", "retainAsRegulatoryRecord", "unknownFutureValue"}[i]
}
func ParseBehaviorDuringRetentionPeriod(v string) (interface{}, error) {
    result := DONOTRETAIN_BEHAVIORDURINGRETENTIONPERIOD
    switch v {
        case "doNotRetain":
            result = DONOTRETAIN_BEHAVIORDURINGRETENTIONPERIOD
        case "retain":
            result = RETAIN_BEHAVIORDURINGRETENTIONPERIOD
        case "retainAsRecord":
            result = RETAINASRECORD_BEHAVIORDURINGRETENTIONPERIOD
        case "retainAsRegulatoryRecord":
            result = RETAINASREGULATORYRECORD_BEHAVIORDURINGRETENTIONPERIOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BEHAVIORDURINGRETENTIONPERIOD
        default:
            return 0, errors.New("Unknown BehaviorDuringRetentionPeriod value: " + v)
    }
    return &result, nil
}
func SerializeBehaviorDuringRetentionPeriod(values []BehaviorDuringRetentionPeriod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
