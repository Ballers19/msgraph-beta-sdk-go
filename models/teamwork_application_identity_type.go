package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type TeamworkApplicationIdentityType int

const (
    AADAPPLICATION_TEAMWORKAPPLICATIONIDENTITYTYPE TeamworkApplicationIdentityType = iota
    BOT_TEAMWORKAPPLICATIONIDENTITYTYPE
    TENANTBOT_TEAMWORKAPPLICATIONIDENTITYTYPE
    OFFICE365CONNECTOR_TEAMWORKAPPLICATIONIDENTITYTYPE
    OUTGOINGWEBHOOK_TEAMWORKAPPLICATIONIDENTITYTYPE
    UNKNOWNFUTUREVALUE_TEAMWORKAPPLICATIONIDENTITYTYPE
)

func (i TeamworkApplicationIdentityType) String() string {
    return []string{"aadApplication", "bot", "tenantBot", "office365Connector", "outgoingWebhook", "unknownFutureValue"}[i]
}
func ParseTeamworkApplicationIdentityType(v string) (interface{}, error) {
    result := AADAPPLICATION_TEAMWORKAPPLICATIONIDENTITYTYPE
    switch v {
        case "aadApplication":
            result = AADAPPLICATION_TEAMWORKAPPLICATIONIDENTITYTYPE
        case "bot":
            result = BOT_TEAMWORKAPPLICATIONIDENTITYTYPE
        case "tenantBot":
            result = TENANTBOT_TEAMWORKAPPLICATIONIDENTITYTYPE
        case "office365Connector":
            result = OFFICE365CONNECTOR_TEAMWORKAPPLICATIONIDENTITYTYPE
        case "outgoingWebhook":
            result = OUTGOINGWEBHOOK_TEAMWORKAPPLICATIONIDENTITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKAPPLICATIONIDENTITYTYPE
        default:
            return 0, errors.New("Unknown TeamworkApplicationIdentityType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkApplicationIdentityType(values []TeamworkApplicationIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
