package graph
import (
    "strings"
    "errors"
)
// 
type DiscoverySource int

const (
    UNKNOWN_DISCOVERYSOURCE DiscoverySource = iota
    ADMINIMPORT_DISCOVERYSOURCE
    DEVICEENROLLMENTPROGRAM_DISCOVERYSOURCE
)

func (i DiscoverySource) String() string {
    return []string{"UNKNOWN", "ADMINIMPORT", "DEVICEENROLLMENTPROGRAM"}[i]
}
func ParseDiscoverySource(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DISCOVERYSOURCE, nil
        case "ADMINIMPORT":
            return ADMINIMPORT_DISCOVERYSOURCE, nil
        case "DEVICEENROLLMENTPROGRAM":
            return DEVICEENROLLMENTPROGRAM_DISCOVERYSOURCE, nil
    }
    return 0, errors.New("Unknown DiscoverySource value: " + v)
}
func SerializeDiscoverySource(values []DiscoverySource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
