package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type DeviceComplianceActionType int

const (
    // No Action
    NOACTION_DEVICECOMPLIANCEACTIONTYPE DeviceComplianceActionType = iota
    // Send Notification
    NOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
    // Block the device in AAD
    BLOCK_DEVICECOMPLIANCEACTIONTYPE
    // Retire the device
    RETIRE_DEVICECOMPLIANCEACTIONTYPE
    // Wipe the device
    WIPE_DEVICECOMPLIANCEACTIONTYPE
    // Remove Resource Access Profiles from the device
    REMOVERESOURCEACCESSPROFILES_DEVICECOMPLIANCEACTIONTYPE
    // Send push notification to device
    PUSHNOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
    // Remotely lock the device
    REMOTELOCK_DEVICECOMPLIANCEACTIONTYPE
)

func (i DeviceComplianceActionType) String() string {
    return []string{"noAction", "notification", "block", "retire", "wipe", "removeResourceAccessProfiles", "pushNotification", "remoteLock"}[i]
}
func ParseDeviceComplianceActionType(v string) (interface{}, error) {
    result := NOACTION_DEVICECOMPLIANCEACTIONTYPE
    switch v {
        case "noAction":
            result = NOACTION_DEVICECOMPLIANCEACTIONTYPE
        case "notification":
            result = NOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
        case "block":
            result = BLOCK_DEVICECOMPLIANCEACTIONTYPE
        case "retire":
            result = RETIRE_DEVICECOMPLIANCEACTIONTYPE
        case "wipe":
            result = WIPE_DEVICECOMPLIANCEACTIONTYPE
        case "removeResourceAccessProfiles":
            result = REMOVERESOURCEACCESSPROFILES_DEVICECOMPLIANCEACTIONTYPE
        case "pushNotification":
            result = PUSHNOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
        case "remoteLock":
            result = REMOTELOCK_DEVICECOMPLIANCEACTIONTYPE
        default:
            return 0, errors.New("Unknown DeviceComplianceActionType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceComplianceActionType(values []DeviceComplianceActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
