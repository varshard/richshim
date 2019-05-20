package richshim

import (
	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"     // Development version of Kivik
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type RichMockStub struct {
	shim.MockStub
	DB *kivik.DB
}

// GetQueryResult function can be invoked by a chaincode to perform a
// rich query against state database.  Only supported by state database implementations
// that support rich query.  The query string is in the syntax of the underlying
// state database. An iterator is returned which can be used to iterate (next) over
// the query result set
// func (stub *RichMockStub) GetQueryResult(query string) (shim.StateQueryIteratorInterface, error) {
// 	// Access public data by setting the collection to empty string
// 	collection := ""
// 	// ignore QueryResponseMetadata as it is not applicable for a rich query without pagination

// 	// raw, err := DB.QueryJSON(query)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	rows, err := stub.DB.Find(context.TODO(), query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO: turn raw into StateQueryIteratorInterface
// 	// iterator, _, err := stub.handleGetQueryResult(collection, query, nil)

// 	return iterator, err
// }
