package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type CloudPcProvisioningPolicyImageType int

const (
    GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE CloudPcProvisioningPolicyImageType = iota
    CUSTOM_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
)

func (i CloudPcProvisioningPolicyImageType) String() string {
    return []string{"gallery", "custom"}[i]
}
func ParseCloudPcProvisioningPolicyImageType(v string) (interface{}, error) {
    result := GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
    switch v {
        case "gallery":
            result = GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
        case "custom":
            result = CUSTOM_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
        default:
            return 0, errors.New("Unknown CloudPcProvisioningPolicyImageType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcProvisioningPolicyImageType(values []CloudPcProvisioningPolicyImageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
