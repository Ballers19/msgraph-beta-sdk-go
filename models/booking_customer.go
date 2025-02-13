package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomer 
type BookingCustomer struct {
    BookingPerson
    // Addresses associated with the customer, including home, business and other addresses.
    addresses []PhysicalAddressable
    // Phone numbers associated with the customer, including home, business and mobile numbers.
    phones []Phoneable
}
// NewBookingCustomer instantiates a new BookingCustomer and sets the default values.
func NewBookingCustomer()(*BookingCustomer) {
    m := &BookingCustomer{
        BookingPerson: *NewBookingPerson(),
    }
    return m
}
// CreateBookingCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingCustomer(), nil
}
// GetAddresses gets the addresses property value. Addresses associated with the customer, including home, business and other addresses.
func (m *BookingCustomer) GetAddresses()([]PhysicalAddressable) {
    if m == nil {
        return nil
    } else {
        return m.addresses
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingPerson.GetFieldDeserializers()
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhysicalAddressable, len(val))
            for i, v := range val {
                res[i] = v.(PhysicalAddressable)
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                res[i] = v.(Phoneable)
            }
            m.SetPhones(res)
        }
        return nil
    }
    return res
}
// GetPhones gets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) GetPhones()([]Phoneable) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// Serialize serializes information the current object
func (m *BookingCustomer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingPerson.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddresses sets the addresses property value. Addresses associated with the customer, including home, business and other addresses.
func (m *BookingCustomer) SetAddresses(value []PhysicalAddressable)() {
    if m != nil {
        m.addresses = value
    }
}
// SetPhones sets the phones property value. Phone numbers associated with the customer, including home, business and mobile numbers.
func (m *BookingCustomer) SetPhones(value []Phoneable)() {
    if m != nil {
        m.phones = value
    }
}
