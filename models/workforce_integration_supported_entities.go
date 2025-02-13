package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type WorkforceIntegrationSupportedEntities int

const (
    NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES WorkforceIntegrationSupportedEntities = iota
    SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMECARD_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMEOFFREASON_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMEOFF_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    TIMEOFFREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
)

func (i WorkforceIntegrationSupportedEntities) String() string {
    return []string{"none", "shift", "swapRequest", "userShiftPreferences", "openShift", "openShiftRequest", "offerShiftRequest", "unknownFutureValue", "timeCard", "timeOffReason", "timeOff", "timeOffRequest"}[i]
}
func ParseWorkforceIntegrationSupportedEntities(v string) (interface{}, error) {
    result := NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    switch v {
        case "none":
            result = NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "shift":
            result = SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "swapRequest":
            result = SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "userShiftPreferences":
            result = USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "openShift":
            result = OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "openShiftRequest":
            result = OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "offerShiftRequest":
            result = OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "timeCard":
            result = TIMECARD_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "timeOffReason":
            result = TIMEOFFREASON_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "timeOff":
            result = TIMEOFF_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "timeOffRequest":
            result = TIMEOFFREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        default:
            return 0, errors.New("Unknown WorkforceIntegrationSupportedEntities value: " + v)
    }
    return &result, nil
}
func SerializeWorkforceIntegrationSupportedEntities(values []WorkforceIntegrationSupportedEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
