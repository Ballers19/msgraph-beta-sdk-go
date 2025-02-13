package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type SimulationAttackTechnique int

const (
    UNKNOWN_SIMULATIONATTACKTECHNIQUE SimulationAttackTechnique = iota
    CREDENTIALHARVESTING_SIMULATIONATTACKTECHNIQUE
    ATTACHMENTMALWARE_SIMULATIONATTACKTECHNIQUE
    DRIVEBYURL_SIMULATIONATTACKTECHNIQUE
    LINKINATTACHMENT_SIMULATIONATTACKTECHNIQUE
    LINKTOMALWAREFILE_SIMULATIONATTACKTECHNIQUE
    UNKNOWNFUTUREVALUE_SIMULATIONATTACKTECHNIQUE
)

func (i SimulationAttackTechnique) String() string {
    return []string{"unknown", "credentialHarvesting", "attachmentMalware", "driveByUrl", "linkInAttachment", "linkToMalwareFile", "unknownFutureValue"}[i]
}
func ParseSimulationAttackTechnique(v string) (interface{}, error) {
    result := UNKNOWN_SIMULATIONATTACKTECHNIQUE
    switch v {
        case "unknown":
            result = UNKNOWN_SIMULATIONATTACKTECHNIQUE
        case "credentialHarvesting":
            result = CREDENTIALHARVESTING_SIMULATIONATTACKTECHNIQUE
        case "attachmentMalware":
            result = ATTACHMENTMALWARE_SIMULATIONATTACKTECHNIQUE
        case "driveByUrl":
            result = DRIVEBYURL_SIMULATIONATTACKTECHNIQUE
        case "linkInAttachment":
            result = LINKINATTACHMENT_SIMULATIONATTACKTECHNIQUE
        case "linkToMalwareFile":
            result = LINKTOMALWAREFILE_SIMULATIONATTACKTECHNIQUE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIMULATIONATTACKTECHNIQUE
        default:
            return 0, errors.New("Unknown SimulationAttackTechnique value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAttackTechnique(values []SimulationAttackTechnique) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
