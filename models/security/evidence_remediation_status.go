package security
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type EvidenceRemediationStatus int

const (
    NONE_EVIDENCEREMEDIATIONSTATUS EvidenceRemediationStatus = iota
    REMEDIATED_EVIDENCEREMEDIATIONSTATUS
    PREVENTED_EVIDENCEREMEDIATIONSTATUS
    BLOCKED_EVIDENCEREMEDIATIONSTATUS
    NOTFOUND_EVIDENCEREMEDIATIONSTATUS
    ACTIVE_EVIDENCEREMEDIATIONSTATUS
    PENDINGAPPROVAL_EVIDENCEREMEDIATIONSTATUS
    DECLINED_EVIDENCEREMEDIATIONSTATUS
    NOTREMEDIATED_EVIDENCEREMEDIATIONSTATUS
    RUNNING_EVIDENCEREMEDIATIONSTATUS
    UNKNOWNFUTUREVALUE_EVIDENCEREMEDIATIONSTATUS
)

func (i EvidenceRemediationStatus) String() string {
    return []string{"none", "remediated", "prevented", "blocked", "notFound", "active", "pendingApproval", "declined", "notRemediated", "running", "unknownFutureValue"}[i]
}
func ParseEvidenceRemediationStatus(v string) (interface{}, error) {
    result := NONE_EVIDENCEREMEDIATIONSTATUS
    switch v {
        case "none":
            result = NONE_EVIDENCEREMEDIATIONSTATUS
        case "remediated":
            result = REMEDIATED_EVIDENCEREMEDIATIONSTATUS
        case "prevented":
            result = PREVENTED_EVIDENCEREMEDIATIONSTATUS
        case "blocked":
            result = BLOCKED_EVIDENCEREMEDIATIONSTATUS
        case "notFound":
            result = NOTFOUND_EVIDENCEREMEDIATIONSTATUS
        case "active":
            result = ACTIVE_EVIDENCEREMEDIATIONSTATUS
        case "pendingApproval":
            result = PENDINGAPPROVAL_EVIDENCEREMEDIATIONSTATUS
        case "declined":
            result = DECLINED_EVIDENCEREMEDIATIONSTATUS
        case "notRemediated":
            result = NOTREMEDIATED_EVIDENCEREMEDIATIONSTATUS
        case "running":
            result = RUNNING_EVIDENCEREMEDIATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EVIDENCEREMEDIATIONSTATUS
        default:
            return 0, errors.New("Unknown EvidenceRemediationStatus value: " + v)
    }
    return &result, nil
}
func SerializeEvidenceRemediationStatus(values []EvidenceRemediationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
