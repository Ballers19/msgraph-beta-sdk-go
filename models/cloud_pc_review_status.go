package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcReviewStatus 
type CloudPcReviewStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
    azureStorageAccountId *string
    // The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
    azureStorageAccountName *string
    // True if the Cloud PC is set to in review by the administrator.
    inReview *bool
    // The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    restorePointDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    reviewStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
    subscriptionId *string
    // The name of the Azure subscription in which the Cloud PC snapshot is being saved.
    subscriptionName *string
    // The userAccessLevel property
    userAccessLevel *CloudPcUserAccessLevel
}
// NewCloudPcReviewStatus instantiates a new cloudPcReviewStatus and sets the default values.
func NewCloudPcReviewStatus()(*CloudPcReviewStatus) {
    m := &CloudPcReviewStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcReviewStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcReviewStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcReviewStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcReviewStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAzureStorageAccountId gets the azureStorageAccountId property value. The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureStorageAccountId
    }
}
// GetAzureStorageAccountName gets the azureStorageAccountName property value. The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureStorageAccountName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcReviewStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["azureStorageAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageAccountId(val)
        }
        return nil
    }
    res["azureStorageAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageAccountName(val)
        }
        return nil
    }
    res["inReview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInReview(val)
        }
        return nil
    }
    res["restorePointDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestorePointDateTime(val)
        }
        return nil
    }
    res["reviewStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStartDateTime(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["subscriptionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionName(val)
        }
        return nil
    }
    res["userAccessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccessLevel(val.(*CloudPcUserAccessLevel))
        }
        return nil
    }
    return res
}
// GetInReview gets the inReview property value. True if the Cloud PC is set to in review by the administrator.
func (m *CloudPcReviewStatus) GetInReview()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inReview
    }
}
// GetRestorePointDateTime gets the restorePointDateTime property value. The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.restorePointDateTime
    }
}
// GetReviewStartDateTime gets the reviewStartDateTime property value. The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) GetReviewStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewStartDateTime
    }
}
// GetSubscriptionId gets the subscriptionId property value. The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
func (m *CloudPcReviewStatus) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
// GetSubscriptionName gets the subscriptionName property value. The name of the Azure subscription in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetSubscriptionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionName
    }
}
// GetUserAccessLevel gets the userAccessLevel property value. The userAccessLevel property
func (m *CloudPcReviewStatus) GetUserAccessLevel()(*CloudPcUserAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.userAccessLevel
    }
}
// Serialize serializes information the current object
func (m *CloudPcReviewStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("azureStorageAccountId", m.GetAzureStorageAccountId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureStorageAccountName", m.GetAzureStorageAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inReview", m.GetInReview())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("restorePointDateTime", m.GetRestorePointDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("reviewStartDateTime", m.GetReviewStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionName", m.GetSubscriptionName())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccessLevel() != nil {
        cast := (*m.GetUserAccessLevel()).String()
        err := writer.WriteStringValue("userAccessLevel", &cast)
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
func (m *CloudPcReviewStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAzureStorageAccountId sets the azureStorageAccountId property value. The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageAccountId(value *string)() {
    if m != nil {
        m.azureStorageAccountId = value
    }
}
// SetAzureStorageAccountName sets the azureStorageAccountName property value. The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageAccountName(value *string)() {
    if m != nil {
        m.azureStorageAccountName = value
    }
}
// SetInReview sets the inReview property value. True if the Cloud PC is set to in review by the administrator.
func (m *CloudPcReviewStatus) SetInReview(value *bool)() {
    if m != nil {
        m.inReview = value
    }
}
// SetRestorePointDateTime sets the restorePointDateTime property value. The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.restorePointDateTime = value
    }
}
// SetReviewStartDateTime sets the reviewStartDateTime property value. The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) SetReviewStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewStartDateTime = value
    }
}
// SetSubscriptionId sets the subscriptionId property value. The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
func (m *CloudPcReviewStatus) SetSubscriptionId(value *string)() {
    if m != nil {
        m.subscriptionId = value
    }
}
// SetSubscriptionName sets the subscriptionName property value. The name of the Azure subscription in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetSubscriptionName(value *string)() {
    if m != nil {
        m.subscriptionName = value
    }
}
// SetUserAccessLevel sets the userAccessLevel property value. The userAccessLevel property
func (m *CloudPcReviewStatus) SetUserAccessLevel(value *CloudPcUserAccessLevel)() {
    if m != nil {
        m.userAccessLevel = value
    }
}
