package graph
import (
    "strings"
    "errors"
)
// 
type MlClassificationMatchTolerance int

const (
    EXACT_MLCLASSIFICATIONMATCHTOLERANCE MlClassificationMatchTolerance = iota
    NEAR_MLCLASSIFICATIONMATCHTOLERANCE
)

func (i MlClassificationMatchTolerance) String() string {
    return []string{"EXACT", "NEAR"}[i]
}
func ParseMlClassificationMatchTolerance(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "EXACT":
            return EXACT_MLCLASSIFICATIONMATCHTOLERANCE, nil
        case "NEAR":
            return NEAR_MLCLASSIFICATIONMATCHTOLERANCE, nil
    }
    return 0, errors.New("Unknown MlClassificationMatchTolerance value: " + v)
}
func SerializeMlClassificationMatchTolerance(values []MlClassificationMatchTolerance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
