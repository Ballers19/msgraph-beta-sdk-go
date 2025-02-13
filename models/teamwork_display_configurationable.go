package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDisplayConfigurationable 
type TeamworkDisplayConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguredDisplays()([]TeamworkConfiguredPeripheralable)
    GetDisplayCount()(*int32)
    GetInBuiltDisplayScreenConfiguration()(TeamworkDisplayScreenConfigurationable)
    GetIsContentDuplicationAllowed()(*bool)
    GetIsDualDisplayModeEnabled()(*bool)
    SetConfiguredDisplays(value []TeamworkConfiguredPeripheralable)()
    SetDisplayCount(value *int32)()
    SetInBuiltDisplayScreenConfiguration(value TeamworkDisplayScreenConfigurationable)()
    SetIsContentDuplicationAllowed(value *bool)()
    SetIsDualDisplayModeEnabled(value *bool)()
}
