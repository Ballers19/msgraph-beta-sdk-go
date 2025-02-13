package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveContentLocationable 
type SensitiveContentLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfidence()(*int32)
    GetEvidences()([]SensitiveContentEvidenceable)
    GetIdMatch()(*string)
    GetLength()(*int32)
    GetOffset()(*int32)
    SetConfidence(value *int32)()
    SetEvidences(value []SensitiveContentEvidenceable)()
    SetIdMatch(value *string)()
    SetLength(value *int32)()
    SetOffset(value *int32)()
}
