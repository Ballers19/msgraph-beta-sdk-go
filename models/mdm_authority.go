package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type MdmAuthority int

const (
    // Unknown
    UNKNOWN_MDMAUTHORITY MdmAuthority = iota
    // Intune
    INTUNE_MDMAUTHORITY
    // SCCM
    SCCM_MDMAUTHORITY
    // Office365
    OFFICE365_MDMAUTHORITY
)

func (i MdmAuthority) String() string {
    return []string{"unknown", "intune", "sccm", "office365"}[i]
}
func ParseMdmAuthority(v string) (interface{}, error) {
    result := UNKNOWN_MDMAUTHORITY
    switch v {
        case "unknown":
            result = UNKNOWN_MDMAUTHORITY
        case "intune":
            result = INTUNE_MDMAUTHORITY
        case "sccm":
            result = SCCM_MDMAUTHORITY
        case "office365":
            result = OFFICE365_MDMAUTHORITY
        default:
            return 0, errors.New("Unknown MdmAuthority value: " + v)
    }
    return &result, nil
}
func SerializeMdmAuthority(values []MdmAuthority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
