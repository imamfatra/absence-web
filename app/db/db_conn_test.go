package db

import "testing"

func TestDBConnection(t *testing.T) {
    db := NewDB()
    err := db.Ping()
    if err != nil {
        t.Fatalf("Connection failed, err: %v", err)
    }
    t.Log("Connection db succesfull")
}

