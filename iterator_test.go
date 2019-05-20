package richshim

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/go-kivik/kivik" // Development version of Kivik
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var client *kivik.Client
var db *kivik.DB

const DBName = "richshim"
const ConnectionString = "http://localhost:5984"

var _ = Describe("QueryIterator", func() {
	Describe("HasNext", func() {
		It("returns true if key exist", func() {
			value := map[string]string{
				"name":   "momo",
				"gender": "female",
			}

			id := uuid.New()
			idString := id.String()
			rev, err := db.Put(context.TODO(), idString, value)
			Expect(err).NotTo(HaveOccurred())
			Expect(rev).NotTo(BeEmpty())

			query := map[string]interface{}{
				"selector": map[string]interface{}{
					"_id": map[string]string{"$eq": idString},
				},
			}
			rows, err := db.Find(context.TODO(), query)
			Expect(err).NotTo(HaveOccurred())

			iterator := QueryIterator{
				Rows: rows,
			}
			defer iterator.Close()

			Expect(iterator.HasNext()).To(Equal(true))

			// No more item
			Expect(iterator.HasNext()).To(Equal(false))
		})
	})

	Describe("Next", func() {
		It("returns item", func() {
			value := map[string]interface{}{
				"foo": "bar",
				"data": map[string]string{
					"AAA": "123",
				},
			}

			id := uuid.New()
			idString := id.String()

			_, err := db.Put(context.TODO(), idString, value)
			Expect(err).NotTo(HaveOccurred())

			_, err = db.Put(context.TODO(), uuid.New().String(), map[string]interface{}{
				"foo":  "boo",
				"data": map[string]string{"AAA": "234"},
			})
			Expect(err).NotTo(HaveOccurred())

			query := map[string]interface{}{
				"selector": map[string]interface{}{
					"_id": map[string]string{
						"$eq": idString,
					},
				},
			}

			rows, err := db.Find(context.TODO(), query)
			Expect(err).NotTo(HaveOccurred())

			iterator := QueryIterator{
				Rows: rows,
			}
			defer iterator.Close()

			Expect(iterator.HasNext()).To(Equal(true))
			kv, err := iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(kv.Key).To(Equal(idString))

			var convertedValue map[string]interface{}
			Expect(json.Unmarshal(kv.Value, &convertedValue)).To(Succeed())

			Expect(convertedValue["foo"]).To(Equal("bar"))
		})
	})
})

var _ = BeforeSuite(func() {
	client, err := kivik.New("couch", ConnectionString)
	ctx := context.TODO()

	exist, err := client.DBExists(ctx, DBName)
	Expect(err).NotTo(HaveOccurred())

	if !exist {
		err = client.CreateDB(ctx, DBName)
		Expect(err).NotTo(HaveOccurred())
	}

	db = client.DB(ctx, DBName)
	Expect(db.Err()).NotTo(HaveOccurred())
})

// var _ = AfterSuite(func() {
// 	client.DestroyDB(context.TODO(), DBName)
// 	Expect(client.Close(context.TODO())).NotTo(HaveOccurred())
// })
