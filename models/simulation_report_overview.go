package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationReportOverview 
type SimulationReportOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
    recommendedActions []RecommendedActionable
    // Number of valid users in the attack simulation and training campaign.
    resolvedTargetsCount *int32
    // Summary of simulation events in the attack simulation and training campaign.
    simulationEventsContent SimulationEventsContentable
    // Summary of assigned trainings in the attack simulation and training campaign.
    trainingEventsContent TrainingEventsContentable
}
// NewSimulationReportOverview instantiates a new simulationReportOverview and sets the default values.
func NewSimulationReportOverview()(*SimulationReportOverview) {
    m := &SimulationReportOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSimulationReportOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationReportOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationReportOverview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReportOverview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationReportOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recommendedActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecommendedActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendedActionable, len(val))
            for i, v := range val {
                res[i] = v.(RecommendedActionable)
            }
            m.SetRecommendedActions(res)
        }
        return nil
    }
    res["resolvedTargetsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedTargetsCount(val)
        }
        return nil
    }
    res["simulationEventsContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimulationEventsContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationEventsContent(val.(SimulationEventsContentable))
        }
        return nil
    }
    res["trainingEventsContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingEventsContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingEventsContent(val.(TrainingEventsContentable))
        }
        return nil
    }
    return res
}
// GetRecommendedActions gets the recommendedActions property value. List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
func (m *SimulationReportOverview) GetRecommendedActions()([]RecommendedActionable) {
    if m == nil {
        return nil
    } else {
        return m.recommendedActions
    }
}
// GetResolvedTargetsCount gets the resolvedTargetsCount property value. Number of valid users in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetResolvedTargetsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.resolvedTargetsCount
    }
}
// GetSimulationEventsContent gets the simulationEventsContent property value. Summary of simulation events in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetSimulationEventsContent()(SimulationEventsContentable) {
    if m == nil {
        return nil
    } else {
        return m.simulationEventsContent
    }
}
// GetTrainingEventsContent gets the trainingEventsContent property value. Summary of assigned trainings in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetTrainingEventsContent()(TrainingEventsContentable) {
    if m == nil {
        return nil
    } else {
        return m.trainingEventsContent
    }
}
// Serialize serializes information the current object
func (m *SimulationReportOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRecommendedActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecommendedActions()))
        for i, v := range m.GetRecommendedActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recommendedActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("resolvedTargetsCount", m.GetResolvedTargetsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("simulationEventsContent", m.GetSimulationEventsContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingEventsContent", m.GetTrainingEventsContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReportOverview) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRecommendedActions sets the recommendedActions property value. List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
func (m *SimulationReportOverview) SetRecommendedActions(value []RecommendedActionable)() {
    if m != nil {
        m.recommendedActions = value
    }
}
// SetResolvedTargetsCount sets the resolvedTargetsCount property value. Number of valid users in the attack simulation and training campaign.
func (m *SimulationReportOverview) SetResolvedTargetsCount(value *int32)() {
    if m != nil {
        m.resolvedTargetsCount = value
    }
}
// SetSimulationEventsContent sets the simulationEventsContent property value. Summary of simulation events in the attack simulation and training campaign.
func (m *SimulationReportOverview) SetSimulationEventsContent(value SimulationEventsContentable)() {
    if m != nil {
        m.simulationEventsContent = value
    }
}
// SetTrainingEventsContent sets the trainingEventsContent property value. Summary of assigned trainings in the attack simulation and training campaign.
func (m *SimulationReportOverview) SetTrainingEventsContent(value TrainingEventsContentable)() {
    if m != nil {
        m.trainingEventsContent = value
    }
}
