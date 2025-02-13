package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type CloudAppSecuritySessionControlType int

const (
    MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE CloudAppSecuritySessionControlType = iota
    MONITORONLY_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    BLOCKDOWNLOADS_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    UNKNOWNFUTUREVALUE_CLOUDAPPSECURITYSESSIONCONTROLTYPE
)

func (i CloudAppSecuritySessionControlType) String() string {
    return []string{"mcasConfigured", "monitorOnly", "blockDownloads", "unknownFutureValue"}[i]
}
func ParseCloudAppSecuritySessionControlType(v string) (interface{}, error) {
    result := MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    switch v {
        case "mcasConfigured":
            result = MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        case "monitorOnly":
            result = MONITORONLY_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        case "blockDownloads":
            result = BLOCKDOWNLOADS_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDAPPSECURITYSESSIONCONTROLTYPE
        default:
            return 0, errors.New("Unknown CloudAppSecuritySessionControlType value: " + v)
    }
    return &result, nil
}
func SerializeCloudAppSecuritySessionControlType(values []CloudAppSecuritySessionControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
