package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type SecurityResourceType int

const (
    UNKNOWN_SECURITYRESOURCETYPE SecurityResourceType = iota
    ATTACKED_SECURITYRESOURCETYPE
    RELATED_SECURITYRESOURCETYPE
    UNKNOWNFUTUREVALUE_SECURITYRESOURCETYPE
)

func (i SecurityResourceType) String() string {
    return []string{"unknown", "attacked", "related", "unknownFutureValue"}[i]
}
func ParseSecurityResourceType(v string) (interface{}, error) {
    result := UNKNOWN_SECURITYRESOURCETYPE
    switch v {
        case "unknown":
            result = UNKNOWN_SECURITYRESOURCETYPE
        case "attacked":
            result = ATTACKED_SECURITYRESOURCETYPE
        case "related":
            result = RELATED_SECURITYRESOURCETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYRESOURCETYPE
        default:
            return 0, errors.New("Unknown SecurityResourceType value: " + v)
    }
    return &result, nil
}
func SerializeSecurityResourceType(values []SecurityResourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
