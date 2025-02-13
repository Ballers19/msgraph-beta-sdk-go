package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type DataSubjectType int

const (
    CUSTOMER_DATASUBJECTTYPE DataSubjectType = iota
    CURRENTEMPLOYEE_DATASUBJECTTYPE
    FORMEREMPLOYEE_DATASUBJECTTYPE
    PROSPECTIVEEMPLOYEE_DATASUBJECTTYPE
    STUDENT_DATASUBJECTTYPE
    TEACHER_DATASUBJECTTYPE
    FACULTY_DATASUBJECTTYPE
    OTHER_DATASUBJECTTYPE
    UNKNOWNFUTUREVALUE_DATASUBJECTTYPE
)

func (i DataSubjectType) String() string {
    return []string{"customer", "currentEmployee", "formerEmployee", "prospectiveEmployee", "student", "teacher", "faculty", "other", "unknownFutureValue"}[i]
}
func ParseDataSubjectType(v string) (interface{}, error) {
    result := CUSTOMER_DATASUBJECTTYPE
    switch v {
        case "customer":
            result = CUSTOMER_DATASUBJECTTYPE
        case "currentEmployee":
            result = CURRENTEMPLOYEE_DATASUBJECTTYPE
        case "formerEmployee":
            result = FORMEREMPLOYEE_DATASUBJECTTYPE
        case "prospectiveEmployee":
            result = PROSPECTIVEEMPLOYEE_DATASUBJECTTYPE
        case "student":
            result = STUDENT_DATASUBJECTTYPE
        case "teacher":
            result = TEACHER_DATASUBJECTTYPE
        case "faculty":
            result = FACULTY_DATASUBJECTTYPE
        case "other":
            result = OTHER_DATASUBJECTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATASUBJECTTYPE
        default:
            return 0, errors.New("Unknown DataSubjectType value: " + v)
    }
    return &result, nil
}
func SerializeDataSubjectType(values []DataSubjectType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
