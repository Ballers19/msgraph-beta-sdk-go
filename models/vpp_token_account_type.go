package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type VppTokenAccountType int

const (
    // Apple Volume Purchase Program token associated with an business program.
    BUSINESS_VPPTOKENACCOUNTTYPE VppTokenAccountType = iota
    // Apple Volume Purchase Program token associated with an education program.
    EDUCATION_VPPTOKENACCOUNTTYPE
)

func (i VppTokenAccountType) String() string {
    return []string{"business", "education"}[i]
}
func ParseVppTokenAccountType(v string) (interface{}, error) {
    result := BUSINESS_VPPTOKENACCOUNTTYPE
    switch v {
        case "business":
            result = BUSINESS_VPPTOKENACCOUNTTYPE
        case "education":
            result = EDUCATION_VPPTOKENACCOUNTTYPE
        default:
            return 0, errors.New("Unknown VppTokenAccountType value: " + v)
    }
    return &result, nil
}
func SerializeVppTokenAccountType(values []VppTokenAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
