package graph
import (
    "strings"
    "errors"
)
// 
type DeviceGuardVirtualizationBasedSecurityState int

const (
    RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE DeviceGuardVirtualizationBasedSecurityState = iota
    REBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    REQUIRE64BITARCHITECTURE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    NOTLICENSED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    NOTCONFIGURED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    DOESNOTMEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    OTHER_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
)

func (i DeviceGuardVirtualizationBasedSecurityState) String() string {
    return []string{"RUNNING", "REBOOTREQUIRED", "REQUIRE64BITARCHITECTURE", "NOTLICENSED", "NOTCONFIGURED", "DOESNOTMEETHARDWAREREQUIREMENTS", "OTHER"}[i]
}
func ParseDeviceGuardVirtualizationBasedSecurityState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "RUNNING":
            return RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
        case "REBOOTREQUIRED":
            return REBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
        case "REQUIRE64BITARCHITECTURE":
            return REQUIRE64BITARCHITECTURE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
        case "NOTLICENSED":
            return NOTLICENSED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
        case "NOTCONFIGURED":
            return NOTCONFIGURED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
        case "DOESNOTMEETHARDWAREREQUIREMENTS":
            return DOESNOTMEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
        case "OTHER":
            return OTHER_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE, nil
    }
    return 0, errors.New("Unknown DeviceGuardVirtualizationBasedSecurityState value: " + v)
}
func SerializeDeviceGuardVirtualizationBasedSecurityState(values []DeviceGuardVirtualizationBasedSecurityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
