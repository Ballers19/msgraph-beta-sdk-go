package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobSubjectable 
type SynchronizationJobSubjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLinks()(SynchronizationLinkedObjectsable)
    GetObjectId()(*string)
    GetObjectTypeName()(*string)
    SetLinks(value SynchronizationLinkedObjectsable)()
    SetObjectId(value *string)()
    SetObjectTypeName(value *string)()
}
