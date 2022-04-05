package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the commsApplication singleton.
type OnlineMeetingRole int

const (
    ATTENDEE_ONLINEMEETINGROLE OnlineMeetingRole = iota
    PRESENTER_ONLINEMEETINGROLE
    UNKNOWNFUTUREVALUE_ONLINEMEETINGROLE
    PRODUCER_ONLINEMEETINGROLE
    COORGANIZER_ONLINEMEETINGROLE
)

func (i OnlineMeetingRole) String() string {
    return []string{"ATTENDEE", "PRESENTER", "UNKNOWNFUTUREVALUE", "PRODUCER", "COORGANIZER"}[i]
}
func ParseOnlineMeetingRole(v string) (interface{}, error) {
    result := ATTENDEE_ONLINEMEETINGROLE
    switch strings.ToUpper(v) {
        case "ATTENDEE":
            result = ATTENDEE_ONLINEMEETINGROLE
        case "PRESENTER":
            result = PRESENTER_ONLINEMEETINGROLE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ONLINEMEETINGROLE
        case "PRODUCER":
            result = PRODUCER_ONLINEMEETINGROLE
        case "COORGANIZER":
            result = COORGANIZER_ONLINEMEETINGROLE
        default:
            return 0, errors.New("Unknown OnlineMeetingRole value: " + v)
    }
    return &result, nil
}
func SerializeOnlineMeetingRole(values []OnlineMeetingRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
