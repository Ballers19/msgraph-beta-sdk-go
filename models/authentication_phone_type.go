package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AuthenticationPhoneType int

const (
    MOBILE_AUTHENTICATIONPHONETYPE AuthenticationPhoneType = iota
    ALTERNATEMOBILE_AUTHENTICATIONPHONETYPE
    OFFICE_AUTHENTICATIONPHONETYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONPHONETYPE
)

func (i AuthenticationPhoneType) String() string {
    return []string{"mobile", "alternateMobile", "office", "unknownFutureValue"}[i]
}
func ParseAuthenticationPhoneType(v string) (interface{}, error) {
    result := MOBILE_AUTHENTICATIONPHONETYPE
    switch v {
        case "mobile":
            result = MOBILE_AUTHENTICATIONPHONETYPE
        case "alternateMobile":
            result = ALTERNATEMOBILE_AUTHENTICATIONPHONETYPE
        case "office":
            result = OFFICE_AUTHENTICATIONPHONETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONPHONETYPE
        default:
            return 0, errors.New("Unknown AuthenticationPhoneType value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationPhoneType(values []AuthenticationPhoneType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
