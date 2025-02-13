package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AuthenticationMethodTargetType int

const (
    USER_AUTHENTICATIONMETHODTARGETTYPE AuthenticationMethodTargetType = iota
    GROUP_AUTHENTICATIONMETHODTARGETTYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODTARGETTYPE
)

func (i AuthenticationMethodTargetType) String() string {
    return []string{"user", "group", "unknownFutureValue"}[i]
}
func ParseAuthenticationMethodTargetType(v string) (interface{}, error) {
    result := USER_AUTHENTICATIONMETHODTARGETTYPE
    switch v {
        case "user":
            result = USER_AUTHENTICATIONMETHODTARGETTYPE
        case "group":
            result = GROUP_AUTHENTICATIONMETHODTARGETTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODTARGETTYPE
        default:
            return 0, errors.New("Unknown AuthenticationMethodTargetType value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodTargetType(values []AuthenticationMethodTargetType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
