package graph
import (
    "strings"
    "errors"
)
// 
type DeviceLogCollectionTemplateType int

const (
    PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE DeviceLogCollectionTemplateType = iota
)

func (i DeviceLogCollectionTemplateType) String() string {
    return []string{"PREDEFINED"}[i]
}
func ParseDeviceLogCollectionTemplateType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PREDEFINED":
            return PREDEFINED_DEVICELOGCOLLECTIONTEMPLATETYPE, nil
    }
    return 0, errors.New("Unknown DeviceLogCollectionTemplateType value: " + v)
}
func SerializeDeviceLogCollectionTemplateType(values []DeviceLogCollectionTemplateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
