package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type DeviceManagementReportFileFormat int

const (
    // Comma-separated values
    CSV_DEVICEMANAGEMENTREPORTFILEFORMAT DeviceManagementReportFileFormat = iota
    // Portable Document Format
    PDF_DEVICEMANAGEMENTREPORTFILEFORMAT
)

func (i DeviceManagementReportFileFormat) String() string {
    return []string{"csv", "pdf"}[i]
}
func ParseDeviceManagementReportFileFormat(v string) (interface{}, error) {
    result := CSV_DEVICEMANAGEMENTREPORTFILEFORMAT
    switch v {
        case "csv":
            result = CSV_DEVICEMANAGEMENTREPORTFILEFORMAT
        case "pdf":
            result = PDF_DEVICEMANAGEMENTREPORTFILEFORMAT
        default:
            return 0, errors.New("Unknown DeviceManagementReportFileFormat value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementReportFileFormat(values []DeviceManagementReportFileFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
