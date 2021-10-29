package graph
import (
    "strings"
    "errors"
)
// 
type WindowsAutopilotEnrollmentType int

const (
    UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE WindowsAutopilotEnrollmentType = iota
    AZUREADJOINEDWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINED_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
)

func (i WindowsAutopilotEnrollmentType) String() string {
    return []string{"UNKNOWN", "AZUREADJOINEDWITHAUTOPILOTPROFILE", "OFFLINEDOMAINJOINED", "AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE", "AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE", "AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE", "AZUREADJOINEDWITHWHITEGLOVE", "OFFLINEDOMAINJOINEDWITHWHITEGLOVE", "OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE"}[i]
}
func ParseWindowsAutopilotEnrollmentType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "AZUREADJOINEDWITHAUTOPILOTPROFILE":
            return AZUREADJOINEDWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "OFFLINEDOMAINJOINED":
            return OFFLINEDOMAINJOINED_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE":
            return AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE":
            return AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE":
            return AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "AZUREADJOINEDWITHWHITEGLOVE":
            return AZUREADJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "OFFLINEDOMAINJOINEDWITHWHITEGLOVE":
            return OFFLINEDOMAINJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
        case "OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE":
            return OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE, nil
    }
    return 0, errors.New("Unknown WindowsAutopilotEnrollmentType value: " + v)
}
func SerializeWindowsAutopilotEnrollmentType(values []WindowsAutopilotEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
