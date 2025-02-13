package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewRecurrenceSettingsable 
type AccessReviewRecurrenceSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDurationInDays()(*int32)
    GetRecurrenceCount()(*int32)
    GetRecurrenceEndType()(*string)
    GetRecurrenceType()(*string)
    SetDurationInDays(value *int32)()
    SetRecurrenceCount(value *int32)()
    SetRecurrenceEndType(value *string)()
    SetRecurrenceType(value *string)()
}
