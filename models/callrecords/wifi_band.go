package callrecords
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type WifiBand int

const (
    UNKNOWN_WIFIBAND WifiBand = iota
    FREQUENCY24GHZ_WIFIBAND
    FREQUENCY50GHZ_WIFIBAND
    FREQUENCY60GHZ_WIFIBAND
    UNKNOWNFUTUREVALUE_WIFIBAND
)

func (i WifiBand) String() string {
    return []string{"unknown", "frequency24GHz", "frequency50GHz", "frequency60GHz", "unknownFutureValue"}[i]
}
func ParseWifiBand(v string) (interface{}, error) {
    result := UNKNOWN_WIFIBAND
    switch v {
        case "unknown":
            result = UNKNOWN_WIFIBAND
        case "frequency24GHz":
            result = FREQUENCY24GHZ_WIFIBAND
        case "frequency50GHz":
            result = FREQUENCY50GHZ_WIFIBAND
        case "frequency60GHz":
            result = FREQUENCY60GHZ_WIFIBAND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WIFIBAND
        default:
            return 0, errors.New("Unknown WifiBand value: " + v)
    }
    return &result, nil
}
func SerializeWifiBand(values []WifiBand) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
