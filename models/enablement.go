package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type Enablement int

const (
    // Device default value, no intent.
    NOTCONFIGURED_ENABLEMENT Enablement = iota
    // Enables the setting on the device.
    ENABLED_ENABLEMENT
    // Disables the setting on the device.
    DISABLED_ENABLEMENT
)

func (i Enablement) String() string {
    return []string{"notConfigured", "enabled", "disabled"}[i]
}
func ParseEnablement(v string) (interface{}, error) {
    result := NOTCONFIGURED_ENABLEMENT
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ENABLEMENT
        case "enabled":
            result = ENABLED_ENABLEMENT
        case "disabled":
            result = DISABLED_ENABLEMENT
        default:
            return 0, errors.New("Unknown Enablement value: " + v)
    }
    return &result, nil
}
func SerializeEnablement(values []Enablement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
