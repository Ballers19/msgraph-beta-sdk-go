package callrecords
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type ClientPlatform int

const (
    UNKNOWN_CLIENTPLATFORM ClientPlatform = iota
    WINDOWS_CLIENTPLATFORM
    MACOS_CLIENTPLATFORM
    IOS_CLIENTPLATFORM
    ANDROID_CLIENTPLATFORM
    WEB_CLIENTPLATFORM
    IPPHONE_CLIENTPLATFORM
    ROOMSYSTEM_CLIENTPLATFORM
    SURFACEHUB_CLIENTPLATFORM
    HOLOLENS_CLIENTPLATFORM
    UNKNOWNFUTUREVALUE_CLIENTPLATFORM
)

func (i ClientPlatform) String() string {
    return []string{"unknown", "windows", "macOS", "iOS", "android", "web", "ipPhone", "roomSystem", "surfaceHub", "holoLens", "unknownFutureValue"}[i]
}
func ParseClientPlatform(v string) (interface{}, error) {
    result := UNKNOWN_CLIENTPLATFORM
    switch v {
        case "unknown":
            result = UNKNOWN_CLIENTPLATFORM
        case "windows":
            result = WINDOWS_CLIENTPLATFORM
        case "macOS":
            result = MACOS_CLIENTPLATFORM
        case "iOS":
            result = IOS_CLIENTPLATFORM
        case "android":
            result = ANDROID_CLIENTPLATFORM
        case "web":
            result = WEB_CLIENTPLATFORM
        case "ipPhone":
            result = IPPHONE_CLIENTPLATFORM
        case "roomSystem":
            result = ROOMSYSTEM_CLIENTPLATFORM
        case "surfaceHub":
            result = SURFACEHUB_CLIENTPLATFORM
        case "holoLens":
            result = HOLOLENS_CLIENTPLATFORM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLIENTPLATFORM
        default:
            return 0, errors.New("Unknown ClientPlatform value: " + v)
    }
    return &result, nil
}
func SerializeClientPlatform(values []ClientPlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
