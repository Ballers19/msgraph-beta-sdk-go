package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type RetentionTrigger int

const (
    DATELABELED_RETENTIONTRIGGER RetentionTrigger = iota
    DATECREATED_RETENTIONTRIGGER
    DATEMODIFIED_RETENTIONTRIGGER
    DATEOFEVENT_RETENTIONTRIGGER
    UNKNOWNFUTUREVALUE_RETENTIONTRIGGER
)

func (i RetentionTrigger) String() string {
    return []string{"dateLabeled", "dateCreated", "dateModified", "dateOfEvent", "unknownFutureValue"}[i]
}
func ParseRetentionTrigger(v string) (interface{}, error) {
    result := DATELABELED_RETENTIONTRIGGER
    switch v {
        case "dateLabeled":
            result = DATELABELED_RETENTIONTRIGGER
        case "dateCreated":
            result = DATECREATED_RETENTIONTRIGGER
        case "dateModified":
            result = DATEMODIFIED_RETENTIONTRIGGER
        case "dateOfEvent":
            result = DATEOFEVENT_RETENTIONTRIGGER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RETENTIONTRIGGER
        default:
            return 0, errors.New("Unknown RetentionTrigger value: " + v)
    }
    return &result, nil
}
func SerializeRetentionTrigger(values []RetentionTrigger) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
