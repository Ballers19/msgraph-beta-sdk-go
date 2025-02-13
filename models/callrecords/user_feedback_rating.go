package callrecords
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type UserFeedbackRating int

const (
    NOTRATED_USERFEEDBACKRATING UserFeedbackRating = iota
    BAD_USERFEEDBACKRATING
    POOR_USERFEEDBACKRATING
    FAIR_USERFEEDBACKRATING
    GOOD_USERFEEDBACKRATING
    EXCELLENT_USERFEEDBACKRATING
    UNKNOWNFUTUREVALUE_USERFEEDBACKRATING
)

func (i UserFeedbackRating) String() string {
    return []string{"notRated", "bad", "poor", "fair", "good", "excellent", "unknownFutureValue"}[i]
}
func ParseUserFeedbackRating(v string) (interface{}, error) {
    result := NOTRATED_USERFEEDBACKRATING
    switch v {
        case "notRated":
            result = NOTRATED_USERFEEDBACKRATING
        case "bad":
            result = BAD_USERFEEDBACKRATING
        case "poor":
            result = POOR_USERFEEDBACKRATING
        case "fair":
            result = FAIR_USERFEEDBACKRATING
        case "good":
            result = GOOD_USERFEEDBACKRATING
        case "excellent":
            result = EXCELLENT_USERFEEDBACKRATING
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERFEEDBACKRATING
        default:
            return 0, errors.New("Unknown UserFeedbackRating value: " + v)
    }
    return &result, nil
}
func SerializeUserFeedbackRating(values []UserFeedbackRating) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
