package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type ImportedWindowsAutopilotDeviceIdentityUploadStatus int

const (
    // No upload status.
    NOUPLOAD_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS ImportedWindowsAutopilotDeviceIdentityUploadStatus = iota
    // Pending status.
    PENDING_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
    // Complete status.
    COMPLETE_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
    // Error status.
    ERROR_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
)

func (i ImportedWindowsAutopilotDeviceIdentityUploadStatus) String() string {
    return []string{"noUpload", "pending", "complete", "error"}[i]
}
func ParseImportedWindowsAutopilotDeviceIdentityUploadStatus(v string) (interface{}, error) {
    result := NOUPLOAD_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
    switch v {
        case "noUpload":
            result = NOUPLOAD_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
        case "pending":
            result = PENDING_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
        case "complete":
            result = COMPLETE_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
        case "error":
            result = ERROR_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYUPLOADSTATUS
        default:
            return 0, errors.New("Unknown ImportedWindowsAutopilotDeviceIdentityUploadStatus value: " + v)
    }
    return &result, nil
}
func SerializeImportedWindowsAutopilotDeviceIdentityUploadStatus(values []ImportedWindowsAutopilotDeviceIdentityUploadStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
