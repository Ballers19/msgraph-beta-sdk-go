package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileEvidence 
type FileEvidence struct {
    AlertEvidence
    // The detectionStatus property
    detectionStatus *DetectionStatus
    // The fileDetails property
    fileDetails FileDetailsable
    // The mdeDeviceId property
    mdeDeviceId *string
}
// NewFileEvidence instantiates a new FileEvidence and sets the default values.
func NewFileEvidence()(*FileEvidence) {
    m := &FileEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateFileEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileEvidence(), nil
}
// GetDetectionStatus gets the detectionStatus property value. The detectionStatus property
func (m *FileEvidence) GetDetectionStatus()(*DetectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.detectionStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["detectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionStatus(val.(*DetectionStatus))
        }
        return nil
    }
    res["fileDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileDetails(val.(FileDetailsable))
        }
        return nil
    }
    res["mdeDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdeDeviceId(val)
        }
        return nil
    }
    return res
}
// GetFileDetails gets the fileDetails property value. The fileDetails property
func (m *FileEvidence) GetFileDetails()(FileDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.fileDetails
    }
}
// GetMdeDeviceId gets the mdeDeviceId property value. The mdeDeviceId property
func (m *FileEvidence) GetMdeDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdeDeviceId
    }
}
// Serialize serializes information the current object
func (m *FileEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetectionStatus() != nil {
        cast := (*m.GetDetectionStatus()).String()
        err = writer.WriteStringValue("detectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fileDetails", m.GetFileDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdeDeviceId", m.GetMdeDeviceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionStatus sets the detectionStatus property value. The detectionStatus property
func (m *FileEvidence) SetDetectionStatus(value *DetectionStatus)() {
    if m != nil {
        m.detectionStatus = value
    }
}
// SetFileDetails sets the fileDetails property value. The fileDetails property
func (m *FileEvidence) SetFileDetails(value FileDetailsable)() {
    if m != nil {
        m.fileDetails = value
    }
}
// SetMdeDeviceId sets the mdeDeviceId property value. The mdeDeviceId property
func (m *FileEvidence) SetMdeDeviceId(value *string)() {
    if m != nil {
        m.mdeDeviceId = value
    }
}
