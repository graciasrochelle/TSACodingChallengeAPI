package storage

import (
	"TSACodingChallengeAPI/src/common"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"

	// SQL Driver
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug = flag.Bool("debug", false, "enable debugging")
	db    *sql.DB
)

// Service storage service interface
type Service interface {
	CreateConnectionPool() error
	ReadContacts() (Contacts, error)
	ReadPhoneNumbers(contactID string) (PhoneNumbers, error)
	CreateContact(contact SQLContact) error
	CreatePhoneNumber(phoneNumber SQLPhoneNumber) error
}

type service struct {
	config common.Config
}

// NewService creates new storage service
func NewService(config common.Config) Service {
	return &service{
		config: config,
	}
}

// CreateConnectionPool create connection pool
func (s *service) CreateConnectionPool() (err error) {
	flag.Parse()

	connString := s.config.ConnString
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		e := fmt.Sprintf("Open connection failed: <%s>", err.Error())
		return errors.New(e)
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Printf("Connected to database!\n")

	return nil
}

// ReadContacts reads all contact records
func (s *service) ReadContacts() (contacts Contacts, err error) {
	ctx := context.Background()

	// Check if database is alive.
	if err = db.PingContext(ctx); err != nil {
		return
	}

	tsql := fmt.Sprintf("select * from tblContacts;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		c := &SQLContact{}
		if err = rows.Scan(&c.ContactID, &c.FullName, &c.Email); err != nil {
			return Contacts{}, err
		}
		contacts.Contacts = append(contacts.Contacts, *c)
	}

	return
}

// ReadContacts reads all phone records for a contactID
func (s *service) ReadPhoneNumbers(contactID string) (phoneNumbers PhoneNumbers, err error) {
	ctx := context.Background()

	// Check if database is alive.
	if err = db.PingContext(ctx); err != nil {
		return
	}

	tsql := fmt.Sprintf("select * from tblPhoneNumbers where contactId='%s';", contactID)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		p := &SQLPhoneNumber{}
		if err = rows.Scan(&p.PhoneID, &p.ContactID, &p.PhoneNumber); err != nil {
			return PhoneNumbers{}, err
		}
		phoneNumbers.PhoneNumbers = append(phoneNumbers.PhoneNumbers, *p)
	}

	return
}

// CreateContact inserts an contact record
func (s *service) CreateContact(contact SQLContact) (err error) {
	ctx := context.Background()

	if db == nil {
		err = errors.New("CreateEmployee: db is null")
		return
	}

	// Check if database is alive.
	if err = db.PingContext(ctx); err != nil {
		return
	}

	tsql := "insert into tblContacts (contactId, fullName, email) VALUES (@contactId, @fullName, @email);"

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_ = stmt.QueryRowContext(
		ctx,
		sql.Named("contactId", contact.ContactID),
		sql.Named("fullName", contact.FullName),
		sql.Named("email", contact.Email))

	return
}

// CreatePhoneNumber inserts an phone record for a contactID
func (s *service) CreatePhoneNumber(phoneNumber SQLPhoneNumber) (err error) {
	ctx := context.Background()

	if db == nil {
		err = errors.New("CreateEmployee: db is null")
		return
	}

	// Check if database is alive.
	if err = db.PingContext(ctx); err != nil {
		return
	}

	tsql := "insert into tblPhoneNumbers (phoneId, contactId, phoneNumber) VALUES (@phoneId, @contactId, @phoneNumber);"

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_ = stmt.QueryRowContext(
		ctx,
		sql.Named("phoneId", phoneNumber.PhoneID),
		sql.Named("contactId", phoneNumber.ContactID),
		sql.Named("phoneNumber", phoneNumber.PhoneNumber))

	return
}
