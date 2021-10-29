package graph
import (
    "strings"
    "errors"
)
// 
type AndroidDeviceOwnerEnrollmentMode int

const (
    CORPORATEOWNEDDEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE AndroidDeviceOwnerEnrollmentMode = iota
    CORPORATEOWNEDFULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTMODE
    CORPORATEOWNEDWORKPROFILE_ANDROIDDEVICEOWNERENROLLMENTMODE
    CORPORATEOWNEDAOSPUSERLESSDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
    CORPORATEOWNEDAOSPUSERASSOCIATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE
)

func (i AndroidDeviceOwnerEnrollmentMode) String() string {
    return []string{"CORPORATEOWNEDDEDICATEDDEVICE", "CORPORATEOWNEDFULLYMANAGED", "CORPORATEOWNEDWORKPROFILE", "CORPORATEOWNEDAOSPUSERLESSDEVICE", "CORPORATEOWNEDAOSPUSERASSOCIATEDDEVICE"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentMode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CORPORATEOWNEDDEDICATEDDEVICE":
            return CORPORATEOWNEDDEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE, nil
        case "CORPORATEOWNEDFULLYMANAGED":
            return CORPORATEOWNEDFULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTMODE, nil
        case "CORPORATEOWNEDWORKPROFILE":
            return CORPORATEOWNEDWORKPROFILE_ANDROIDDEVICEOWNERENROLLMENTMODE, nil
        case "CORPORATEOWNEDAOSPUSERLESSDEVICE":
            return CORPORATEOWNEDAOSPUSERLESSDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE, nil
        case "CORPORATEOWNEDAOSPUSERASSOCIATEDDEVICE":
            return CORPORATEOWNEDAOSPUSERASSOCIATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTMODE, nil
    }
    return 0, errors.New("Unknown AndroidDeviceOwnerEnrollmentMode value: " + v)
}
func SerializeAndroidDeviceOwnerEnrollmentMode(values []AndroidDeviceOwnerEnrollmentMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
