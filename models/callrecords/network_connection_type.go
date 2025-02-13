package callrecords
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type NetworkConnectionType int

const (
    UNKNOWN_NETWORKCONNECTIONTYPE NetworkConnectionType = iota
    WIRED_NETWORKCONNECTIONTYPE
    WIFI_NETWORKCONNECTIONTYPE
    MOBILE_NETWORKCONNECTIONTYPE
    TUNNEL_NETWORKCONNECTIONTYPE
    UNKNOWNFUTUREVALUE_NETWORKCONNECTIONTYPE
)

func (i NetworkConnectionType) String() string {
    return []string{"unknown", "wired", "wifi", "mobile", "tunnel", "unknownFutureValue"}[i]
}
func ParseNetworkConnectionType(v string) (interface{}, error) {
    result := UNKNOWN_NETWORKCONNECTIONTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_NETWORKCONNECTIONTYPE
        case "wired":
            result = WIRED_NETWORKCONNECTIONTYPE
        case "wifi":
            result = WIFI_NETWORKCONNECTIONTYPE
        case "mobile":
            result = MOBILE_NETWORKCONNECTIONTYPE
        case "tunnel":
            result = TUNNEL_NETWORKCONNECTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NETWORKCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown NetworkConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeNetworkConnectionType(values []NetworkConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
