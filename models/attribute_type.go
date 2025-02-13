package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type AttributeType int

const (
    STRING_ATTRIBUTETYPE AttributeType = iota
    INTEGER_ATTRIBUTETYPE
    REFERENCE_ATTRIBUTETYPE
    BINARY_ATTRIBUTETYPE
    BOOLEAN_ATTRIBUTETYPE
    DATETIME_ATTRIBUTETYPE
)

func (i AttributeType) String() string {
    return []string{"String", "Integer", "Reference", "Binary", "Boolean", "DateTime"}[i]
}
func ParseAttributeType(v string) (interface{}, error) {
    result := STRING_ATTRIBUTETYPE
    switch v {
        case "String":
            result = STRING_ATTRIBUTETYPE
        case "Integer":
            result = INTEGER_ATTRIBUTETYPE
        case "Reference":
            result = REFERENCE_ATTRIBUTETYPE
        case "Binary":
            result = BINARY_ATTRIBUTETYPE
        case "Boolean":
            result = BOOLEAN_ATTRIBUTETYPE
        case "DateTime":
            result = DATETIME_ATTRIBUTETYPE
        default:
            return 0, errors.New("Unknown AttributeType value: " + v)
    }
    return &result, nil
}
func SerializeAttributeType(values []AttributeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
