package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type UserFlowType int

const (
    SIGNUP_USERFLOWTYPE UserFlowType = iota
    SIGNIN_USERFLOWTYPE
    SIGNUPORSIGNIN_USERFLOWTYPE
    PASSWORDRESET_USERFLOWTYPE
    PROFILEUPDATE_USERFLOWTYPE
    RESOURCEOWNER_USERFLOWTYPE
    UNKNOWNFUTUREVALUE_USERFLOWTYPE
)

func (i UserFlowType) String() string {
    return []string{"signUp", "signIn", "signUpOrSignIn", "passwordReset", "profileUpdate", "resourceOwner", "unknownFutureValue"}[i]
}
func ParseUserFlowType(v string) (interface{}, error) {
    result := SIGNUP_USERFLOWTYPE
    switch v {
        case "signUp":
            result = SIGNUP_USERFLOWTYPE
        case "signIn":
            result = SIGNIN_USERFLOWTYPE
        case "signUpOrSignIn":
            result = SIGNUPORSIGNIN_USERFLOWTYPE
        case "passwordReset":
            result = PASSWORDRESET_USERFLOWTYPE
        case "profileUpdate":
            result = PROFILEUPDATE_USERFLOWTYPE
        case "resourceOwner":
            result = RESOURCEOWNER_USERFLOWTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERFLOWTYPE
        default:
            return 0, errors.New("Unknown UserFlowType value: " + v)
    }
    return &result, nil
}
func SerializeUserFlowType(values []UserFlowType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
