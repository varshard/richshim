package richshim

import (
	"fmt"

	"github.com/go-kivik/kivik" // Development version of Kivik
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
)

// QueryIterator allows a chaincode to iterate over a set of
// key/value pairs returned by range and execute query.
type QueryIterator struct {
	Rows *kivik.Rows
}

func (q *QueryIterator) Next() (*queryresult.KV, error) {
	err := q.Rows.Err()
	if err != nil {
		return nil, err
	} else {
		var value []byte
		q.Rows.ScanValue(&value)
		fmt.Printf("key: [%s]\n", q.Rows.ID())
		kv := queryresult.KV{
			Key:   q.Rows.ID(),
			Value: value,
		}

		return &kv, nil
	}
}

func (q QueryIterator) Close() error {
	return q.Rows.Close()
}

func (q QueryIterator) HasNext() bool {
	return q.Rows.Next()
}
