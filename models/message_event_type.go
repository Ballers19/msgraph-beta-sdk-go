package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of messageEvent entities.
type MessageEventType int

const (
    RECEIVED_MESSAGEEVENTTYPE MessageEventType = iota
    SENT_MESSAGEEVENTTYPE
    DELIVERED_MESSAGEEVENTTYPE
    FAILED_MESSAGEEVENTTYPE
    PROCESSINGFAILED_MESSAGEEVENTTYPE
    DISTRIBUTIONGROUPEXPANDED_MESSAGEEVENTTYPE
    SUBMITTED_MESSAGEEVENTTYPE
    DELAYED_MESSAGEEVENTTYPE
    REDIRECTED_MESSAGEEVENTTYPE
    RESOLVED_MESSAGEEVENTTYPE
    DROPPED_MESSAGEEVENTTYPE
    RECIPIENTSADDED_MESSAGEEVENTTYPE
    MALWAREDETECTED_MESSAGEEVENTTYPE
    MALWAREDETECTEDINMESSAGE_MESSAGEEVENTTYPE
    MALWAREDETECTEDINATTACHMENT_MESSAGEEVENTTYPE
    TTZAPPED_MESSAGEEVENTTYPE
    TTDELIVERED_MESSAGEEVENTTYPE
    SPAMDETECTED_MESSAGEEVENTTYPE
    TRANSPORTRULETRIGGERED_MESSAGEEVENTTYPE
    DLPRULETRIGGERED_MESSAGEEVENTTYPE
    JOURNALED_MESSAGEEVENTTYPE
    UNKNOWNFUTUREVALUE_MESSAGEEVENTTYPE
)

func (i MessageEventType) String() string {
    return []string{"RECEIVED", "SENT", "DELIVERED", "FAILED", "PROCESSINGFAILED", "DISTRIBUTIONGROUPEXPANDED", "SUBMITTED", "DELAYED", "REDIRECTED", "RESOLVED", "DROPPED", "RECIPIENTSADDED", "MALWAREDETECTED", "MALWAREDETECTEDINMESSAGE", "MALWAREDETECTEDINATTACHMENT", "TTZAPPED", "TTDELIVERED", "SPAMDETECTED", "TRANSPORTRULETRIGGERED", "DLPRULETRIGGERED", "JOURNALED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMessageEventType(v string) (interface{}, error) {
    result := RECEIVED_MESSAGEEVENTTYPE
    switch strings.ToUpper(v) {
        case "RECEIVED":
            result = RECEIVED_MESSAGEEVENTTYPE
        case "SENT":
            result = SENT_MESSAGEEVENTTYPE
        case "DELIVERED":
            result = DELIVERED_MESSAGEEVENTTYPE
        case "FAILED":
            result = FAILED_MESSAGEEVENTTYPE
        case "PROCESSINGFAILED":
            result = PROCESSINGFAILED_MESSAGEEVENTTYPE
        case "DISTRIBUTIONGROUPEXPANDED":
            result = DISTRIBUTIONGROUPEXPANDED_MESSAGEEVENTTYPE
        case "SUBMITTED":
            result = SUBMITTED_MESSAGEEVENTTYPE
        case "DELAYED":
            result = DELAYED_MESSAGEEVENTTYPE
        case "REDIRECTED":
            result = REDIRECTED_MESSAGEEVENTTYPE
        case "RESOLVED":
            result = RESOLVED_MESSAGEEVENTTYPE
        case "DROPPED":
            result = DROPPED_MESSAGEEVENTTYPE
        case "RECIPIENTSADDED":
            result = RECIPIENTSADDED_MESSAGEEVENTTYPE
        case "MALWAREDETECTED":
            result = MALWAREDETECTED_MESSAGEEVENTTYPE
        case "MALWAREDETECTEDINMESSAGE":
            result = MALWAREDETECTEDINMESSAGE_MESSAGEEVENTTYPE
        case "MALWAREDETECTEDINATTACHMENT":
            result = MALWAREDETECTEDINATTACHMENT_MESSAGEEVENTTYPE
        case "TTZAPPED":
            result = TTZAPPED_MESSAGEEVENTTYPE
        case "TTDELIVERED":
            result = TTDELIVERED_MESSAGEEVENTTYPE
        case "SPAMDETECTED":
            result = SPAMDETECTED_MESSAGEEVENTTYPE
        case "TRANSPORTRULETRIGGERED":
            result = TRANSPORTRULETRIGGERED_MESSAGEEVENTTYPE
        case "DLPRULETRIGGERED":
            result = DLPRULETRIGGERED_MESSAGEEVENTTYPE
        case "JOURNALED":
            result = JOURNALED_MESSAGEEVENTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MESSAGEEVENTTYPE
        default:
            return 0, errors.New("Unknown MessageEventType value: " + v)
    }
    return &result, nil
}
func SerializeMessageEventType(values []MessageEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
