package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type WindowsInformationProtectionEnforcementLevel int

const (
    // No protection enforcement
    NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL WindowsInformationProtectionEnforcementLevel = iota
    // Encrypt and Audit only
    ENCRYPTANDAUDITONLY_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
    // Encrypt, Audit and Prompt
    ENCRYPTAUDITANDPROMPT_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
    // Encrypt, Audit and Block
    ENCRYPTAUDITANDBLOCK_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
)

func (i WindowsInformationProtectionEnforcementLevel) String() string {
    return []string{"noProtection", "encryptAndAuditOnly", "encryptAuditAndPrompt", "encryptAuditAndBlock"}[i]
}
func ParseWindowsInformationProtectionEnforcementLevel(v string) (interface{}, error) {
    result := NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
    switch v {
        case "noProtection":
            result = NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        case "encryptAndAuditOnly":
            result = ENCRYPTANDAUDITONLY_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        case "encryptAuditAndPrompt":
            result = ENCRYPTAUDITANDPROMPT_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        case "encryptAuditAndBlock":
            result = ENCRYPTAUDITANDBLOCK_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        default:
            return 0, errors.New("Unknown WindowsInformationProtectionEnforcementLevel value: " + v)
    }
    return &result, nil
}
func SerializeWindowsInformationProtectionEnforcementLevel(values []WindowsInformationProtectionEnforcementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
