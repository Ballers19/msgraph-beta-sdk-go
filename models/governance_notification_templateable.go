package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceNotificationTemplateable 
type GovernanceNotificationTemplateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCulture()(*string)
    GetId()(*string)
    GetSource()(*string)
    GetType()(*string)
    GetVersion()(*string)
    SetCulture(value *string)()
    SetId(value *string)()
    SetSource(value *string)()
    SetType(value *string)()
    SetVersion(value *string)()
}
