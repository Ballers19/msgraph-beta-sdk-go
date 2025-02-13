package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileDetails 
type FileDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The fileName property
    fileName *string
    // The filePath property
    filePath *string
    // The filePublisher property
    filePublisher *string
    // The fileSize property
    fileSize *int64
    // The issuer property
    issuer *string
    // The sha1 property
    sha1 *string
    // The sha256 property
    sha256 *string
    // The signer property
    signer *string
}
// NewFileDetails instantiates a new fileDetails and sets the default values.
func NewFileDetails()(*FileDetails) {
    m := &FileDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFileDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["filePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
        }
        return nil
    }
    res["filePublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePublisher(val)
        }
        return nil
    }
    res["fileSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["sha1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha1(val)
        }
        return nil
    }
    res["sha256"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256(val)
        }
        return nil
    }
    res["signer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigner(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The fileName property
func (m *FileDetails) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetFilePath gets the filePath property value. The filePath property
func (m *FileDetails) GetFilePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePath
    }
}
// GetFilePublisher gets the filePublisher property value. The filePublisher property
func (m *FileDetails) GetFilePublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePublisher
    }
}
// GetFileSize gets the fileSize property value. The fileSize property
func (m *FileDetails) GetFileSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileSize
    }
}
// GetIssuer gets the issuer property value. The issuer property
func (m *FileDetails) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
// GetSha1 gets the sha1 property value. The sha1 property
func (m *FileDetails) GetSha1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sha1
    }
}
// GetSha256 gets the sha256 property value. The sha256 property
func (m *FileDetails) GetSha256()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sha256
    }
}
// GetSigner gets the signer property value. The signer property
func (m *FileDetails) GetSigner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signer
    }
}
// Serialize serializes information the current object
func (m *FileDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePath", m.GetFilePath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePublisher", m.GetFilePublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("fileSize", m.GetFileSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha1", m.GetSha1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha256", m.GetSha256())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signer", m.GetSigner())
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
func (m *FileDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFileName sets the fileName property value. The fileName property
func (m *FileDetails) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetFilePath sets the filePath property value. The filePath property
func (m *FileDetails) SetFilePath(value *string)() {
    if m != nil {
        m.filePath = value
    }
}
// SetFilePublisher sets the filePublisher property value. The filePublisher property
func (m *FileDetails) SetFilePublisher(value *string)() {
    if m != nil {
        m.filePublisher = value
    }
}
// SetFileSize sets the fileSize property value. The fileSize property
func (m *FileDetails) SetFileSize(value *int64)() {
    if m != nil {
        m.fileSize = value
    }
}
// SetIssuer sets the issuer property value. The issuer property
func (m *FileDetails) SetIssuer(value *string)() {
    if m != nil {
        m.issuer = value
    }
}
// SetSha1 sets the sha1 property value. The sha1 property
func (m *FileDetails) SetSha1(value *string)() {
    if m != nil {
        m.sha1 = value
    }
}
// SetSha256 sets the sha256 property value. The sha256 property
func (m *FileDetails) SetSha256(value *string)() {
    if m != nil {
        m.sha256 = value
    }
}
// SetSigner sets the signer property value. The signer property
func (m *FileDetails) SetSigner(value *string)() {
    if m != nil {
        m.signer = value
    }
}
