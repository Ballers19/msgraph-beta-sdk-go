package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementConfigurationControlType int

const (
    DEFAULT_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE DeviceManagementConfigurationControlType = iota
    DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
    CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE
)

func (i DeviceManagementConfigurationControlType) String() string {
    return []string{"DEFAULT", "DROPDOWN", "SMALLTEXTBOX", "LARGETEXTBOX", "TOGGLE", "MULTIHEADERGRID", "CONTEXTPANE"}[i]
}
func ParseDeviceManagementConfigurationControlType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
        case "DROPDOWN":
            return DROPDOWN_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
        case "SMALLTEXTBOX":
            return SMALLTEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
        case "LARGETEXTBOX":
            return LARGETEXTBOX_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
        case "TOGGLE":
            return TOGGLE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
        case "MULTIHEADERGRID":
            return MULTIHEADERGRID_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
        case "CONTEXTPANE":
            return CONTEXTPANE_DEVICEMANAGEMENTCONFIGURATIONCONTROLTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationControlType value: " + v)
}
func SerializeDeviceManagementConfigurationControlType(values []DeviceManagementConfigurationControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
