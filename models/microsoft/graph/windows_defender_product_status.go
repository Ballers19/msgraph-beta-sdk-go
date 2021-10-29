package graph
import (
    "strings"
    "errors"
)
// 
type WindowsDefenderProductStatus int

const (
    NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS WindowsDefenderProductStatus = iota
    SERVICENOTRUNNING_WINDOWSDEFENDERPRODUCTSTATUS
    SERVICESTARTEDWITHOUTMALWAREPROTECTION_WINDOWSDEFENDERPRODUCTSTATUS
    PENDINGFULLSCANDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    PENDINGREBOOTDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    PENDINGMANUALSTEPSDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS
    AVSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    ASSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
    NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS
    SYSTEMINITIATEDSCANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    SYSTEMINITIATEDCLEANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    SAMPLESPENDINGSUBMISSION_WINDOWSDEFENDERPRODUCTSTATUS
    PRODUCTRUNNINGINEVALUATIONMODE_WINDOWSDEFENDERPRODUCTSTATUS
    PRODUCTRUNNINGINNONGENUINEMODE_WINDOWSDEFENDERPRODUCTSTATUS
    PRODUCTEXPIRED_WINDOWSDEFENDERPRODUCTSTATUS
    OFFLINESCANREQUIRED_WINDOWSDEFENDERPRODUCTSTATUS
    SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN_WINDOWSDEFENDERPRODUCTSTATUS
    THREATREMEDIATIONFAILEDCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
    THREATREMEDIATIONFAILEDNONCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS
    NOSTATUSFLAGSSET_WINDOWSDEFENDERPRODUCTSTATUS
    PLATFORMOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS
    PLATFORMUPDATEINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS
    PLATFORMABOUTTOBEOUTDATED_WINDOWSDEFENDERPRODUCTSTATUS
    SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING_WINDOWSDEFENDERPRODUCTSTATUS
    WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS
)

func (i WindowsDefenderProductStatus) String() string {
    return []string{"NOSTATUS", "SERVICENOTRUNNING", "SERVICESTARTEDWITHOUTMALWAREPROTECTION", "PENDINGFULLSCANDUETOTHREATACTION", "PENDINGREBOOTDUETOTHREATACTION", "PENDINGMANUALSTEPSDUETOTHREATACTION", "AVSIGNATURESOUTOFDATE", "ASSIGNATURESOUTOFDATE", "NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD", "NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD", "SYSTEMINITIATEDSCANINPROGRESS", "SYSTEMINITIATEDCLEANINPROGRESS", "SAMPLESPENDINGSUBMISSION", "PRODUCTRUNNINGINEVALUATIONMODE", "PRODUCTRUNNINGINNONGENUINEMODE", "PRODUCTEXPIRED", "OFFLINESCANREQUIRED", "SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN", "THREATREMEDIATIONFAILEDCRITICALLY", "THREATREMEDIATIONFAILEDNONCRITICALLY", "NOSTATUSFLAGSSET", "PLATFORMOUTOFDATE", "PLATFORMUPDATEINPROGRESS", "PLATFORMABOUTTOBEOUTDATED", "SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING", "WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL"}[i]
}
func ParseWindowsDefenderProductStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOSTATUS":
            return NOSTATUS_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SERVICENOTRUNNING":
            return SERVICENOTRUNNING_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SERVICESTARTEDWITHOUTMALWAREPROTECTION":
            return SERVICESTARTEDWITHOUTMALWAREPROTECTION_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PENDINGFULLSCANDUETOTHREATACTION":
            return PENDINGFULLSCANDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PENDINGREBOOTDUETOTHREATACTION":
            return PENDINGREBOOTDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PENDINGMANUALSTEPSDUETOTHREATACTION":
            return PENDINGMANUALSTEPSDUETOTHREATACTION_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "AVSIGNATURESOUTOFDATE":
            return AVSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "ASSIGNATURESOUTOFDATE":
            return ASSIGNATURESOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD":
            return NOQUICKSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD":
            return NOFULLSCANHAPPENEDFORSPECIFIEDPERIOD_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SYSTEMINITIATEDSCANINPROGRESS":
            return SYSTEMINITIATEDSCANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SYSTEMINITIATEDCLEANINPROGRESS":
            return SYSTEMINITIATEDCLEANINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SAMPLESPENDINGSUBMISSION":
            return SAMPLESPENDINGSUBMISSION_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PRODUCTRUNNINGINEVALUATIONMODE":
            return PRODUCTRUNNINGINEVALUATIONMODE_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PRODUCTRUNNINGINNONGENUINEMODE":
            return PRODUCTRUNNINGINNONGENUINEMODE_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PRODUCTEXPIRED":
            return PRODUCTEXPIRED_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "OFFLINESCANREQUIRED":
            return OFFLINESCANREQUIRED_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN":
            return SERVICESHUTDOWNASPARTOFSYSTEMSHUTDOWN_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "THREATREMEDIATIONFAILEDCRITICALLY":
            return THREATREMEDIATIONFAILEDCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "THREATREMEDIATIONFAILEDNONCRITICALLY":
            return THREATREMEDIATIONFAILEDNONCRITICALLY_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "NOSTATUSFLAGSSET":
            return NOSTATUSFLAGSSET_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PLATFORMOUTOFDATE":
            return PLATFORMOUTOFDATE_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PLATFORMUPDATEINPROGRESS":
            return PLATFORMUPDATEINPROGRESS_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "PLATFORMABOUTTOBEOUTDATED":
            return PLATFORMABOUTTOBEOUTDATED_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING":
            return SIGNATUREORPLATFORMENDOFLIFEISPASTORISIMPENDING_WINDOWSDEFENDERPRODUCTSTATUS, nil
        case "WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL":
            return WINDOWSSMODESIGNATURESINUSEONNONWIN10SINSTALL_WINDOWSDEFENDERPRODUCTSTATUS, nil
    }
    return 0, errors.New("Unknown WindowsDefenderProductStatus value: " + v)
}
func SerializeWindowsDefenderProductStatus(values []WindowsDefenderProductStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
