package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ManagedBrowserType int

const (
    // Not configured
    NOTCONFIGURED_MANAGEDBROWSERTYPE ManagedBrowserType = iota
    // Microsoft Edge
    MICROSOFTEDGE_MANAGEDBROWSERTYPE
)

func (i ManagedBrowserType) String() string {
    return []string{"notConfigured", "microsoftEdge"}[i]
}
func ParseManagedBrowserType(v string) (interface{}, error) {
    result := NOTCONFIGURED_MANAGEDBROWSERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MANAGEDBROWSERTYPE
        case "microsoftEdge":
            result = MICROSOFTEDGE_MANAGEDBROWSERTYPE
        default:
            return 0, errors.New("Unknown ManagedBrowserType value: " + v)
    }
    return &result, nil
}
func SerializeManagedBrowserType(values []ManagedBrowserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
