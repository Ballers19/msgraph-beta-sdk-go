package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessApplicationsable 
type ConditionalAccessApplicationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeApplications()([]string)
    GetIncludeApplications()([]string)
    GetIncludeAuthenticationContextClassReferences()([]string)
    GetIncludeUserActions()([]string)
    SetExcludeApplications(value []string)()
    SetIncludeApplications(value []string)()
    SetIncludeAuthenticationContextClassReferences(value []string)()
    SetIncludeUserActions(value []string)()
}
