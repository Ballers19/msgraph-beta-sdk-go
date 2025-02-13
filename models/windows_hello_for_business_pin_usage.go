package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type WindowsHelloForBusinessPinUsage int

const (
    // Allowed the usage of certain pin rule
    ALLOWED_WINDOWSHELLOFORBUSINESSPINUSAGE WindowsHelloForBusinessPinUsage = iota
    // Enforce the usage of certain pin rule
    REQUIRED_WINDOWSHELLOFORBUSINESSPINUSAGE
    // Forbit the usage of certain pin rule
    DISALLOWED_WINDOWSHELLOFORBUSINESSPINUSAGE
)

func (i WindowsHelloForBusinessPinUsage) String() string {
    return []string{"allowed", "required", "disallowed"}[i]
}
func ParseWindowsHelloForBusinessPinUsage(v string) (interface{}, error) {
    result := ALLOWED_WINDOWSHELLOFORBUSINESSPINUSAGE
    switch v {
        case "allowed":
            result = ALLOWED_WINDOWSHELLOFORBUSINESSPINUSAGE
        case "required":
            result = REQUIRED_WINDOWSHELLOFORBUSINESSPINUSAGE
        case "disallowed":
            result = DISALLOWED_WINDOWSHELLOFORBUSINESSPINUSAGE
        default:
            return 0, errors.New("Unknown WindowsHelloForBusinessPinUsage value: " + v)
    }
    return &result, nil
}
func SerializeWindowsHelloForBusinessPinUsage(values []WindowsHelloForBusinessPinUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
