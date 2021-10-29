package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OutlookUser struct {
    Entity
    // A list of categories defined for the user.
    masterCategories []OutlookCategory;
    // 
    taskFolders []OutlookTaskFolder;
    // 
    taskGroups []OutlookTaskGroup;
    // 
    tasks []OutlookTask;
}
// Instantiates a new outlookUser and sets the default values.
func NewOutlookUser()(*OutlookUser) {
    m := &OutlookUser{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the masterCategories property value. A list of categories defined for the user.
func (m *OutlookUser) GetMasterCategories()([]OutlookCategory) {
    if m == nil {
        return nil
    } else {
        return m.masterCategories
    }
}
// Gets the taskFolders property value. 
func (m *OutlookUser) GetTaskFolders()([]OutlookTaskFolder) {
    if m == nil {
        return nil
    } else {
        return m.taskFolders
    }
}
// Gets the taskGroups property value. 
func (m *OutlookUser) GetTaskGroups()([]OutlookTaskGroup) {
    if m == nil {
        return nil
    } else {
        return m.taskGroups
    }
}
// Gets the tasks property value. 
func (m *OutlookUser) GetTasks()([]OutlookTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// The deserialization information for the current model
func (m *OutlookUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["masterCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookCategory() })
        if err != nil {
            return err
        }
        res := make([]OutlookCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*OutlookCategory))
        }
        m.SetMasterCategories(res)
        return nil
    }
    res["taskFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTaskFolder() })
        if err != nil {
            return err
        }
        res := make([]OutlookTaskFolder, len(val))
        for i, v := range val {
            res[i] = *(v.(*OutlookTaskFolder))
        }
        m.SetTaskFolders(res)
        return nil
    }
    res["taskGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTaskGroup() })
        if err != nil {
            return err
        }
        res := make([]OutlookTaskGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*OutlookTaskGroup))
        }
        m.SetTaskGroups(res)
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookTask() })
        if err != nil {
            return err
        }
        res := make([]OutlookTask, len(val))
        for i, v := range val {
            res[i] = *(v.(*OutlookTask))
        }
        m.SetTasks(res)
        return nil
    }
    return res
}
func (m *OutlookUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OutlookUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMasterCategories()))
        for i, v := range m.GetMasterCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("masterCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskFolders()))
        for i, v := range m.GetTaskFolders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taskFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskGroups()))
        for i, v := range m.GetTaskGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taskGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the masterCategories property value. A list of categories defined for the user.
// Parameters:
//  - value : Value to set for the masterCategories property.
func (m *OutlookUser) SetMasterCategories(value []OutlookCategory)() {
    m.masterCategories = value
}
// Sets the taskFolders property value. 
// Parameters:
//  - value : Value to set for the taskFolders property.
func (m *OutlookUser) SetTaskFolders(value []OutlookTaskFolder)() {
    m.taskFolders = value
}
// Sets the taskGroups property value. 
// Parameters:
//  - value : Value to set for the taskGroups property.
func (m *OutlookUser) SetTaskGroups(value []OutlookTaskGroup)() {
    m.taskGroups = value
}
// Sets the tasks property value. 
// Parameters:
//  - value : Value to set for the tasks property.
func (m *OutlookUser) SetTasks(value []OutlookTask)() {
    m.tasks = value
}
