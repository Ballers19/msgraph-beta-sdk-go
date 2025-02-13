package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type OnboardingStatus int

const (
    INSUFFICIENTINFO_ONBOARDINGSTATUS OnboardingStatus = iota
    ONBOARDED_ONBOARDINGSTATUS
    CANBEONBOARDED_ONBOARDINGSTATUS
    UNSUPPORTED_ONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
)

func (i OnboardingStatus) String() string {
    return []string{"insufficientInfo", "onboarded", "canBeOnboarded", "unsupported", "unknownFutureValue"}[i]
}
func ParseOnboardingStatus(v string) (interface{}, error) {
    result := INSUFFICIENTINFO_ONBOARDINGSTATUS
    switch v {
        case "insufficientInfo":
            result = INSUFFICIENTINFO_ONBOARDINGSTATUS
        case "onboarded":
            result = ONBOARDED_ONBOARDINGSTATUS
        case "canBeOnboarded":
            result = CANBEONBOARDED_ONBOARDINGSTATUS
        case "unsupported":
            result = UNSUPPORTED_ONBOARDINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown OnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeOnboardingStatus(values []OnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
