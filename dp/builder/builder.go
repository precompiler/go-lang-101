package builder

import (
	"fmt"
)

type connection struct {
}

type ConnectionBuilder struct {
	dburl    string
	username string
	password string
}

func (b *ConnectionBuilder) SetDBUrl(dburl string) *ConnectionBuilder {
	b.dburl = dburl
	return b
}

func (b *ConnectionBuilder) SetUsername(username string) *ConnectionBuilder {
	b.username = username
	return b
}

func (b *ConnectionBuilder) SetPassword(password string) *ConnectionBuilder {
	b.password = password
	return b
}

func (b *ConnectionBuilder) Build() *connection {
	fmt.Println("Creating Database Connection", b.dburl, b.username, b.password)
	return new(connection)
}
