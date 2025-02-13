package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ComplianceStatus int

const (
    UNKNOWN_COMPLIANCESTATUS ComplianceStatus = iota
    NOTAPPLICABLE_COMPLIANCESTATUS
    COMPLIANT_COMPLIANCESTATUS
    REMEDIATED_COMPLIANCESTATUS
    NONCOMPLIANT_COMPLIANCESTATUS
    ERROR_COMPLIANCESTATUS
    CONFLICT_COMPLIANCESTATUS
    NOTASSIGNED_COMPLIANCESTATUS
)

func (i ComplianceStatus) String() string {
    return []string{"unknown", "notApplicable", "compliant", "remediated", "nonCompliant", "error", "conflict", "notAssigned"}[i]
}
func ParseComplianceStatus(v string) (interface{}, error) {
    result := UNKNOWN_COMPLIANCESTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_COMPLIANCESTATUS
        case "notApplicable":
            result = NOTAPPLICABLE_COMPLIANCESTATUS
        case "compliant":
            result = COMPLIANT_COMPLIANCESTATUS
        case "remediated":
            result = REMEDIATED_COMPLIANCESTATUS
        case "nonCompliant":
            result = NONCOMPLIANT_COMPLIANCESTATUS
        case "error":
            result = ERROR_COMPLIANCESTATUS
        case "conflict":
            result = CONFLICT_COMPLIANCESTATUS
        case "notAssigned":
            result = NOTASSIGNED_COMPLIANCESTATUS
        default:
            return 0, errors.New("Unknown ComplianceStatus value: " + v)
    }
    return &result, nil
}
func SerializeComplianceStatus(values []ComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
