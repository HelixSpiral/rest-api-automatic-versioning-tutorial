package database

import "io/ioutil"

// Database struct holds the "host" for our database
// Since this is a local db for testing, the Host is just a file.
type Database struct {
	Host string
}

// New returns a new database
func New(host string) *Database {
	return &Database{
		Host: host,
	}
}

// ReadFullDB reads from a file, returns the file contents or error
func (d *Database) ReadFullDB() ([]byte, error) {
	fileContent, err := ioutil.ReadFile(d.Host)
	if err != nil {
		return []byte{}, err
	}

	return fileContent, nil
}
