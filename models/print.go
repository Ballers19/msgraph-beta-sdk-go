package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Print 
type Print struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of available print connectors.
    connectors []PrintConnectorable
    // The list of print long running operations.
    operations []PrintOperationable
    // The list of printers registered in the tenant.
    printers []Printerable
    // The printerShares property
    printerShares []PrinterShareable
    // The reports property
    reports ReportRootable
    // The list of available Universal Print service endpoints.
    services []PrintServiceable
    // Tenant-wide settings for the Universal Print service.
    settings PrintSettingsable
    // The list of printer shares registered in the tenant.
    shares []PrinterShareable
    // List of abstract definition for a task that can be triggered when various events occur within Universal Print.
    taskDefinitions []PrintTaskDefinitionable
}
// NewPrint instantiates a new Print and sets the default values.
func NewPrint()(*Print) {
    m := &Print{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Print) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConnectors gets the connectors property value. The list of available print connectors.
func (m *Print) GetConnectors()([]PrintConnectorable) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Print) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(PrintConnectorable)
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintOperationable, len(val))
            for i, v := range val {
                res[i] = v.(PrintOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["printers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Printerable, len(val))
            for i, v := range val {
                res[i] = v.(Printerable)
            }
            m.SetPrinters(res)
        }
        return nil
    }
    res["printerShares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrinterShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterShareable, len(val))
            for i, v := range val {
                res[i] = v.(PrinterShareable)
            }
            m.SetPrinterShares(res)
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReportRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(ReportRootable))
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintServiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintServiceable, len(val))
            for i, v := range val {
                res[i] = v.(PrintServiceable)
            }
            m.SetServices(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(PrintSettingsable))
        }
        return nil
    }
    res["shares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrinterShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterShareable, len(val))
            for i, v := range val {
                res[i] = v.(PrinterShareable)
            }
            m.SetShares(res)
        }
        return nil
    }
    res["taskDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintTaskDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(PrintTaskDefinitionable)
            }
            m.SetTaskDefinitions(res)
        }
        return nil
    }
    return res
}
// GetOperations gets the operations property value. The list of print long running operations.
func (m *Print) GetOperations()([]PrintOperationable) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPrinters gets the printers property value. The list of printers registered in the tenant.
func (m *Print) GetPrinters()([]Printerable) {
    if m == nil {
        return nil
    } else {
        return m.printers
    }
}
// GetPrinterShares gets the printerShares property value. The printerShares property
func (m *Print) GetPrinterShares()([]PrinterShareable) {
    if m == nil {
        return nil
    } else {
        return m.printerShares
    }
}
// GetReports gets the reports property value. The reports property
func (m *Print) GetReports()(ReportRootable) {
    if m == nil {
        return nil
    } else {
        return m.reports
    }
}
// GetServices gets the services property value. The list of available Universal Print service endpoints.
func (m *Print) GetServices()([]PrintServiceable) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
// GetSettings gets the settings property value. Tenant-wide settings for the Universal Print service.
func (m *Print) GetSettings()(PrintSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetShares gets the shares property value. The list of printer shares registered in the tenant.
func (m *Print) GetShares()([]PrinterShareable) {
    if m == nil {
        return nil
    } else {
        return m.shares
    }
}
// GetTaskDefinitions gets the taskDefinitions property value. List of abstract definition for a task that can be triggered when various events occur within Universal Print.
func (m *Print) GetTaskDefinitions()([]PrintTaskDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.taskDefinitions
    }
}
// Serialize serializes information the current object
func (m *Print) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("connectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPrinters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrinters()))
        for i, v := range m.GetPrinters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("printers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPrinterShares() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrinterShares()))
        for i, v := range m.GetPrinterShares() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("printerShares", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    if m.GetShares() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShares()))
        for i, v := range m.GetShares() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("shares", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskDefinitions()))
        for i, v := range m.GetTaskDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("taskDefinitions", cast)
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
func (m *Print) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConnectors sets the connectors property value. The list of available print connectors.
func (m *Print) SetConnectors(value []PrintConnectorable)() {
    if m != nil {
        m.connectors = value
    }
}
// SetOperations sets the operations property value. The list of print long running operations.
func (m *Print) SetOperations(value []PrintOperationable)() {
    if m != nil {
        m.operations = value
    }
}
// SetPrinters sets the printers property value. The list of printers registered in the tenant.
func (m *Print) SetPrinters(value []Printerable)() {
    if m != nil {
        m.printers = value
    }
}
// SetPrinterShares sets the printerShares property value. The printerShares property
func (m *Print) SetPrinterShares(value []PrinterShareable)() {
    if m != nil {
        m.printerShares = value
    }
}
// SetReports sets the reports property value. The reports property
func (m *Print) SetReports(value ReportRootable)() {
    if m != nil {
        m.reports = value
    }
}
// SetServices sets the services property value. The list of available Universal Print service endpoints.
func (m *Print) SetServices(value []PrintServiceable)() {
    if m != nil {
        m.services = value
    }
}
// SetSettings sets the settings property value. Tenant-wide settings for the Universal Print service.
func (m *Print) SetSettings(value PrintSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
// SetShares sets the shares property value. The list of printer shares registered in the tenant.
func (m *Print) SetShares(value []PrinterShareable)() {
    if m != nil {
        m.shares = value
    }
}
// SetTaskDefinitions sets the taskDefinitions property value. List of abstract definition for a task that can be triggered when various events occur within Universal Print.
func (m *Print) SetTaskDefinitions(value []PrintTaskDefinitionable)() {
    if m != nil {
        m.taskDefinitions = value
    }
}
