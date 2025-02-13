package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type FederatedIdpMfaBehavior int

const (
    ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR FederatedIdpMfaBehavior = iota
    ENFORCEMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
    REJECTMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
    UNKNOWNFUTUREVALUE_FEDERATEDIDPMFABEHAVIOR
)

func (i FederatedIdpMfaBehavior) String() string {
    return []string{"acceptIfMfaDoneByFederatedIdp", "enforceMfaByFederatedIdp", "rejectMfaByFederatedIdp", "unknownFutureValue"}[i]
}
func ParseFederatedIdpMfaBehavior(v string) (interface{}, error) {
    result := ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
    switch v {
        case "acceptIfMfaDoneByFederatedIdp":
            result = ACCEPTIFMFADONEBYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
        case "enforceMfaByFederatedIdp":
            result = ENFORCEMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
        case "rejectMfaByFederatedIdp":
            result = REJECTMFABYFEDERATEDIDP_FEDERATEDIDPMFABEHAVIOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FEDERATEDIDPMFABEHAVIOR
        default:
            return 0, errors.New("Unknown FederatedIdpMfaBehavior value: " + v)
    }
    return &result, nil
}
func SerializeFederatedIdpMfaBehavior(values []FederatedIdpMfaBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
