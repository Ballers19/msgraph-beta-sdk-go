package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AuthenticationMethodState int

const (
    ENABLED_AUTHENTICATIONMETHODSTATE AuthenticationMethodState = iota
    DISABLED_AUTHENTICATIONMETHODSTATE
)

func (i AuthenticationMethodState) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseAuthenticationMethodState(v string) (interface{}, error) {
    result := ENABLED_AUTHENTICATIONMETHODSTATE
    switch v {
        case "enabled":
            result = ENABLED_AUTHENTICATIONMETHODSTATE
        case "disabled":
            result = DISABLED_AUTHENTICATIONMETHODSTATE
        default:
            return 0, errors.New("Unknown AuthenticationMethodState value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodState(values []AuthenticationMethodState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
