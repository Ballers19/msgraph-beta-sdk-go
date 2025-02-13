package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type EducationFeedbackResourceOutcomeStatus int

const (
    NOTPUBLISHED_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS EducationFeedbackResourceOutcomeStatus = iota
    PENDINGPUBLISH_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
    PUBLISHED_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
    FAILEDPUBLISH_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
    UNKNOWNFUTUREVALUE_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
)

func (i EducationFeedbackResourceOutcomeStatus) String() string {
    return []string{"notPublished", "pendingPublish", "published", "failedPublish", "unknownFutureValue"}[i]
}
func ParseEducationFeedbackResourceOutcomeStatus(v string) (interface{}, error) {
    result := NOTPUBLISHED_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
    switch v {
        case "notPublished":
            result = NOTPUBLISHED_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
        case "pendingPublish":
            result = PENDINGPUBLISH_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
        case "published":
            result = PUBLISHED_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
        case "failedPublish":
            result = FAILEDPUBLISH_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONFEEDBACKRESOURCEOUTCOMESTATUS
        default:
            return 0, errors.New("Unknown EducationFeedbackResourceOutcomeStatus value: " + v)
    }
    return &result, nil
}
func SerializeEducationFeedbackResourceOutcomeStatus(values []EducationFeedbackResourceOutcomeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
