package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ThreatAssessmentContentType int

const (
    MAIL_THREATASSESSMENTCONTENTTYPE ThreatAssessmentContentType = iota
    URL_THREATASSESSMENTCONTENTTYPE
    FILE_THREATASSESSMENTCONTENTTYPE
)

func (i ThreatAssessmentContentType) String() string {
    return []string{"mail", "url", "file"}[i]
}
func ParseThreatAssessmentContentType(v string) (interface{}, error) {
    result := MAIL_THREATASSESSMENTCONTENTTYPE
    switch v {
        case "mail":
            result = MAIL_THREATASSESSMENTCONTENTTYPE
        case "url":
            result = URL_THREATASSESSMENTCONTENTTYPE
        case "file":
            result = FILE_THREATASSESSMENTCONTENTTYPE
        default:
            return 0, errors.New("Unknown ThreatAssessmentContentType value: " + v)
    }
    return &result, nil
}
func SerializeThreatAssessmentContentType(values []ThreatAssessmentContentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
