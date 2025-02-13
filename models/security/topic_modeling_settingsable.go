package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TopicModelingSettingsable 
type TopicModelingSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDynamicallyAdjustTopicCount()(*bool)
    GetIgnoreNumbers()(*bool)
    GetIsEnabled()(*bool)
    GetTopicCount()(*int32)
    SetDynamicallyAdjustTopicCount(value *bool)()
    SetIgnoreNumbers(value *bool)()
    SetIsEnabled(value *bool)()
    SetTopicCount(value *int32)()
}
