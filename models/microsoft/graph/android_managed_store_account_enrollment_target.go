package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedStoreAccountEnrollmentTarget int

const (
    NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET AndroidManagedStoreAccountEnrollmentTarget = iota
    ALL_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    TARGETED_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
)

func (i AndroidManagedStoreAccountEnrollmentTarget) String() string {
    return []string{"NONE", "ALL", "TARGETED", "TARGETEDASENROLLMENTRESTRICTIONS"}[i]
}
func ParseAndroidManagedStoreAccountEnrollmentTarget(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET, nil
        case "ALL":
            return ALL_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET, nil
        case "TARGETED":
            return TARGETED_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET, nil
        case "TARGETEDASENROLLMENTRESTRICTIONS":
            return TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET, nil
    }
    return 0, errors.New("Unknown AndroidManagedStoreAccountEnrollmentTarget value: " + v)
}
func SerializeAndroidManagedStoreAccountEnrollmentTarget(values []AndroidManagedStoreAccountEnrollmentTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
