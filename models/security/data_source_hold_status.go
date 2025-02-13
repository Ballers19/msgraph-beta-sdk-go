package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type DataSourceHoldStatus int

const (
    NOTAPPLIED_DATASOURCEHOLDSTATUS DataSourceHoldStatus = iota
    APPLIED_DATASOURCEHOLDSTATUS
    APPLYING_DATASOURCEHOLDSTATUS
    REMOVING_DATASOURCEHOLDSTATUS
    PARTIAL_DATASOURCEHOLDSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCEHOLDSTATUS
)

func (i DataSourceHoldStatus) String() string {
    return []string{"notApplied", "applied", "applying", "removing", "partial", "unknownFutureValue"}[i]
}
func ParseDataSourceHoldStatus(v string) (interface{}, error) {
    result := NOTAPPLIED_DATASOURCEHOLDSTATUS
    switch v {
        case "notApplied":
            result = NOTAPPLIED_DATASOURCEHOLDSTATUS
        case "applied":
            result = APPLIED_DATASOURCEHOLDSTATUS
        case "applying":
            result = APPLYING_DATASOURCEHOLDSTATUS
        case "removing":
            result = REMOVING_DATASOURCEHOLDSTATUS
        case "partial":
            result = PARTIAL_DATASOURCEHOLDSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATASOURCEHOLDSTATUS
        default:
            return 0, errors.New("Unknown DataSourceHoldStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceHoldStatus(values []DataSourceHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
