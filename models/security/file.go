package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// File provides operations to manage the collection of activityStatistics entities.
type File struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The content property
    content []byte
    // The dateTime property
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The extension property
    extension *string
    // The extractedTextContent property
    extractedTextContent []byte
    // The mediaType property
    mediaType *string
    // The name property
    name *string
    // The otherProperties property
    otherProperties StringValueDictionaryable
    // The processingStatus property
    processingStatus *FileProcessingStatus
    // The senderOrAuthors property
    senderOrAuthors []string
    // The size property
    size *int64
    // The sourceType property
    sourceType *SourceType
    // The subjectTitle property
    subjectTitle *string
}
// NewFile instantiates a new file and sets the default values.
func NewFile()(*File) {
    m := &File{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odatatypeValue := "#microsoft.graph.security.file";
    m.SetType(&odatatypeValue);
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.security.ediscoveryFile":
                        return NewEdiscoveryFile(), nil
                }
            }
        }
    }
    return NewFile(), nil
}
// GetContent gets the content property value. The content property
func (m *File) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetDateTime gets the dateTime property value. The dateTime property
func (m *File) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// GetExtension gets the extension property value. The extension property
func (m *File) GetExtension()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extension
    }
}
// GetExtractedTextContent gets the extractedTextContent property value. The extractedTextContent property
func (m *File) GetExtractedTextContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.extractedTextContent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *File) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["dateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
        }
        return nil
    }
    res["extension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtension(val)
        }
        return nil
    }
    res["extractedTextContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtractedTextContent(val)
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["otherProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStringValueDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherProperties(val.(StringValueDictionaryable))
        }
        return nil
    }
    res["processingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileProcessingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingStatus(val.(*FileProcessingStatus))
        }
        return nil
    }
    res["senderOrAuthors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSenderOrAuthors(res)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceType(val.(*SourceType))
        }
        return nil
    }
    res["subjectTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectTitle(val)
        }
        return nil
    }
    return res
}
// GetMediaType gets the mediaType property value. The mediaType property
func (m *File) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// GetName gets the name property value. The name property
func (m *File) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOtherProperties gets the otherProperties property value. The otherProperties property
func (m *File) GetOtherProperties()(StringValueDictionaryable) {
    if m == nil {
        return nil
    } else {
        return m.otherProperties
    }
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *File) GetProcessingStatus()(*FileProcessingStatus) {
    if m == nil {
        return nil
    } else {
        return m.processingStatus
    }
}
// GetSenderOrAuthors gets the senderOrAuthors property value. The senderOrAuthors property
func (m *File) GetSenderOrAuthors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.senderOrAuthors
    }
}
// GetSize gets the size property value. The size property
func (m *File) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetSourceType gets the sourceType property value. The sourceType property
func (m *File) GetSourceType()(*SourceType) {
    if m == nil {
        return nil
    } else {
        return m.sourceType
    }
}
// GetSubjectTitle gets the subjectTitle property value. The subjectTitle property
func (m *File) GetSubjectTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectTitle
    }
}
// Serialize serializes information the current object
func (m *File) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("extension", m.GetExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("extractedTextContent", m.GetExtractedTextContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("otherProperties", m.GetOtherProperties())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingStatus() != nil {
        cast := (*m.GetProcessingStatus()).String()
        err = writer.WriteStringValue("processingStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSenderOrAuthors() != nil {
        err = writer.WriteCollectionOfStringValues("senderOrAuthors", m.GetSenderOrAuthors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := (*m.GetSourceType()).String()
        err = writer.WriteStringValue("sourceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectTitle", m.GetSubjectTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *File) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetDateTime sets the dateTime property value. The dateTime property
func (m *File) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.dateTime = value
    }
}
// SetExtension sets the extension property value. The extension property
func (m *File) SetExtension(value *string)() {
    if m != nil {
        m.extension = value
    }
}
// SetExtractedTextContent sets the extractedTextContent property value. The extractedTextContent property
func (m *File) SetExtractedTextContent(value []byte)() {
    if m != nil {
        m.extractedTextContent = value
    }
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *File) SetMediaType(value *string)() {
    if m != nil {
        m.mediaType = value
    }
}
// SetName sets the name property value. The name property
func (m *File) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOtherProperties sets the otherProperties property value. The otherProperties property
func (m *File) SetOtherProperties(value StringValueDictionaryable)() {
    if m != nil {
        m.otherProperties = value
    }
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *File) SetProcessingStatus(value *FileProcessingStatus)() {
    if m != nil {
        m.processingStatus = value
    }
}
// SetSenderOrAuthors sets the senderOrAuthors property value. The senderOrAuthors property
func (m *File) SetSenderOrAuthors(value []string)() {
    if m != nil {
        m.senderOrAuthors = value
    }
}
// SetSize sets the size property value. The size property
func (m *File) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
// SetSourceType sets the sourceType property value. The sourceType property
func (m *File) SetSourceType(value *SourceType)() {
    if m != nil {
        m.sourceType = value
    }
}
// SetSubjectTitle sets the subjectTitle property value. The subjectTitle property
func (m *File) SetSubjectTitle(value *string)() {
    if m != nil {
        m.subjectTitle = value
    }
}
