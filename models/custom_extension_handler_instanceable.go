package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionHandlerInstanceable 
type CustomExtensionHandlerInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomExtensionId()(*string)
    GetExternalCorrelationId()(*string)
    GetStage()(*AccessPackageCustomExtensionStage)
    GetStatus()(*AccessPackageCustomExtensionHandlerStatus)
    SetCustomExtensionId(value *string)()
    SetExternalCorrelationId(value *string)()
    SetStage(value *AccessPackageCustomExtensionStage)()
    SetStatus(value *AccessPackageCustomExtensionHandlerStatus)()
}
