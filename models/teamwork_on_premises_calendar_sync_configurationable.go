package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkOnPremisesCalendarSyncConfigurationable 
type TeamworkOnPremisesCalendarSyncConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomain()(*string)
    GetDomainUserName()(*string)
    GetSmtpAddress()(*string)
    SetDomain(value *string)()
    SetDomainUserName(value *string)()
    SetSmtpAddress(value *string)()
}
