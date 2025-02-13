package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleSuccessStatistics 
type RoleSuccessStatistics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The permanentFail property
    permanentFail *int64
    // The permanentSuccess property
    permanentSuccess *int64
    // The removeFail property
    removeFail *int64
    // The removeSuccess property
    removeSuccess *int64
    // The roleId property
    roleId *string
    // The roleName property
    roleName *string
    // The temporaryFail property
    temporaryFail *int64
    // The temporarySuccess property
    temporarySuccess *int64
    // The unknownFail property
    unknownFail *int64
}
// NewRoleSuccessStatistics instantiates a new roleSuccessStatistics and sets the default values.
func NewRoleSuccessStatistics()(*RoleSuccessStatistics) {
    m := &RoleSuccessStatistics{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleSuccessStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleSuccessStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleSuccessStatistics(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleSuccessStatistics) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleSuccessStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["permanentFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentFail(val)
        }
        return nil
    }
    res["permanentSuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentSuccess(val)
        }
        return nil
    }
    res["removeFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveFail(val)
        }
        return nil
    }
    res["removeSuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveSuccess(val)
        }
        return nil
    }
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["temporaryFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporaryFail(val)
        }
        return nil
    }
    res["temporarySuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporarySuccess(val)
        }
        return nil
    }
    res["unknownFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownFail(val)
        }
        return nil
    }
    return res
}
// GetPermanentFail gets the permanentFail property value. The permanentFail property
func (m *RoleSuccessStatistics) GetPermanentFail()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.permanentFail
    }
}
// GetPermanentSuccess gets the permanentSuccess property value. The permanentSuccess property
func (m *RoleSuccessStatistics) GetPermanentSuccess()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.permanentSuccess
    }
}
// GetRemoveFail gets the removeFail property value. The removeFail property
func (m *RoleSuccessStatistics) GetRemoveFail()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.removeFail
    }
}
// GetRemoveSuccess gets the removeSuccess property value. The removeSuccess property
func (m *RoleSuccessStatistics) GetRemoveSuccess()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.removeSuccess
    }
}
// GetRoleId gets the roleId property value. The roleId property
func (m *RoleSuccessStatistics) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// GetRoleName gets the roleName property value. The roleName property
func (m *RoleSuccessStatistics) GetRoleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleName
    }
}
// GetTemporaryFail gets the temporaryFail property value. The temporaryFail property
func (m *RoleSuccessStatistics) GetTemporaryFail()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.temporaryFail
    }
}
// GetTemporarySuccess gets the temporarySuccess property value. The temporarySuccess property
func (m *RoleSuccessStatistics) GetTemporarySuccess()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.temporarySuccess
    }
}
// GetUnknownFail gets the unknownFail property value. The unknownFail property
func (m *RoleSuccessStatistics) GetUnknownFail()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unknownFail
    }
}
// Serialize serializes information the current object
func (m *RoleSuccessStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("permanentFail", m.GetPermanentFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("permanentSuccess", m.GetPermanentSuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("removeFail", m.GetRemoveFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("removeSuccess", m.GetRemoveSuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roleName", m.GetRoleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("temporaryFail", m.GetTemporaryFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("temporarySuccess", m.GetTemporarySuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("unknownFail", m.GetUnknownFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleSuccessStatistics) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPermanentFail sets the permanentFail property value. The permanentFail property
func (m *RoleSuccessStatistics) SetPermanentFail(value *int64)() {
    if m != nil {
        m.permanentFail = value
    }
}
// SetPermanentSuccess sets the permanentSuccess property value. The permanentSuccess property
func (m *RoleSuccessStatistics) SetPermanentSuccess(value *int64)() {
    if m != nil {
        m.permanentSuccess = value
    }
}
// SetRemoveFail sets the removeFail property value. The removeFail property
func (m *RoleSuccessStatistics) SetRemoveFail(value *int64)() {
    if m != nil {
        m.removeFail = value
    }
}
// SetRemoveSuccess sets the removeSuccess property value. The removeSuccess property
func (m *RoleSuccessStatistics) SetRemoveSuccess(value *int64)() {
    if m != nil {
        m.removeSuccess = value
    }
}
// SetRoleId sets the roleId property value. The roleId property
func (m *RoleSuccessStatistics) SetRoleId(value *string)() {
    if m != nil {
        m.roleId = value
    }
}
// SetRoleName sets the roleName property value. The roleName property
func (m *RoleSuccessStatistics) SetRoleName(value *string)() {
    if m != nil {
        m.roleName = value
    }
}
// SetTemporaryFail sets the temporaryFail property value. The temporaryFail property
func (m *RoleSuccessStatistics) SetTemporaryFail(value *int64)() {
    if m != nil {
        m.temporaryFail = value
    }
}
// SetTemporarySuccess sets the temporarySuccess property value. The temporarySuccess property
func (m *RoleSuccessStatistics) SetTemporarySuccess(value *int64)() {
    if m != nil {
        m.temporarySuccess = value
    }
}
// SetUnknownFail sets the unknownFail property value. The unknownFail property
func (m *RoleSuccessStatistics) SetUnknownFail(value *int64)() {
    if m != nil {
        m.unknownFail = value
    }
}
