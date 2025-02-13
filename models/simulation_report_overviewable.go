package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationReportOverviewable 
type SimulationReportOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecommendedActions()([]RecommendedActionable)
    GetResolvedTargetsCount()(*int32)
    GetSimulationEventsContent()(SimulationEventsContentable)
    GetTrainingEventsContent()(TrainingEventsContentable)
    SetRecommendedActions(value []RecommendedActionable)()
    SetResolvedTargetsCount(value *int32)()
    SetSimulationEventsContent(value SimulationEventsContentable)()
    SetTrainingEventsContent(value TrainingEventsContentable)()
}
