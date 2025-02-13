package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type MessageActionFlag int

const (
    ANY_MESSAGEACTIONFLAG MessageActionFlag = iota
    CALL_MESSAGEACTIONFLAG
    DONOTFORWARD_MESSAGEACTIONFLAG
    FOLLOWUP_MESSAGEACTIONFLAG
    FYI_MESSAGEACTIONFLAG
    FORWARD_MESSAGEACTIONFLAG
    NORESPONSENECESSARY_MESSAGEACTIONFLAG
    READ_MESSAGEACTIONFLAG
    REPLY_MESSAGEACTIONFLAG
    REPLYTOALL_MESSAGEACTIONFLAG
    REVIEW_MESSAGEACTIONFLAG
)

func (i MessageActionFlag) String() string {
    return []string{"any", "call", "doNotForward", "followUp", "fyi", "forward", "noResponseNecessary", "read", "reply", "replyToAll", "review"}[i]
}
func ParseMessageActionFlag(v string) (interface{}, error) {
    result := ANY_MESSAGEACTIONFLAG
    switch v {
        case "any":
            result = ANY_MESSAGEACTIONFLAG
        case "call":
            result = CALL_MESSAGEACTIONFLAG
        case "doNotForward":
            result = DONOTFORWARD_MESSAGEACTIONFLAG
        case "followUp":
            result = FOLLOWUP_MESSAGEACTIONFLAG
        case "fyi":
            result = FYI_MESSAGEACTIONFLAG
        case "forward":
            result = FORWARD_MESSAGEACTIONFLAG
        case "noResponseNecessary":
            result = NORESPONSENECESSARY_MESSAGEACTIONFLAG
        case "read":
            result = READ_MESSAGEACTIONFLAG
        case "reply":
            result = REPLY_MESSAGEACTIONFLAG
        case "replyToAll":
            result = REPLYTOALL_MESSAGEACTIONFLAG
        case "review":
            result = REVIEW_MESSAGEACTIONFLAG
        default:
            return 0, errors.New("Unknown MessageActionFlag value: " + v)
    }
    return &result, nil
}
func SerializeMessageActionFlag(values []MessageActionFlag) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
