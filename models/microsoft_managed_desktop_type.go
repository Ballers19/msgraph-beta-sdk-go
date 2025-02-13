package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type MicrosoftManagedDesktopType int

const (
    NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE MicrosoftManagedDesktopType = iota
    PREMIUMMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    STANDARDMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    STARTERMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    UNKNOWNFUTUREVALUE_MICROSOFTMANAGEDDESKTOPTYPE
)

func (i MicrosoftManagedDesktopType) String() string {
    return []string{"notManaged", "premiumManaged", "standardManaged", "starterManaged", "unknownFutureValue"}[i]
}
func ParseMicrosoftManagedDesktopType(v string) (interface{}, error) {
    result := NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
    switch v {
        case "notManaged":
            result = NOTMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "premiumManaged":
            result = PREMIUMMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "standardManaged":
            result = STANDARDMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "starterManaged":
            result = STARTERMANAGED_MICROSOFTMANAGEDDESKTOPTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MICROSOFTMANAGEDDESKTOPTYPE
        default:
            return 0, errors.New("Unknown MicrosoftManagedDesktopType value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftManagedDesktopType(values []MicrosoftManagedDesktopType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
