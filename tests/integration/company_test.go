package integration

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/alcb1310/bca-json/internals/database"
)

func TestRegisterCompany(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Company Suite")
}

var _ = Describe("Client", Ordered, func() {
	var container testcontainers.Container
	var ctx context.Context
	var connString string
	dbName := "bcatestcompany"
	dbUser := "testuser"
	dbPassword := "testpassword"

	BeforeAll(func() {
		ctx = context.Background()

		c, err := postgres.Run(ctx,
			"docker.io/postgres:16-alpine",
			postgres.WithDatabase(dbName),
			postgres.WithUsername(dbUser),
			postgres.WithPassword(dbPassword),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(5*time.Second),
			),

			postgres.WithInitScripts(
				filepath.Join("..", "..", "scripts", "tables.sql"),
				filepath.Join("..", "..", "scripts", "u000_role.sql"),
			),
		)
		Expect(err).NotTo(HaveOccurred())

		connString, err = c.ConnectionString(ctx)
		Expect(err).NotTo(HaveOccurred())

		container = c
	})

	AfterAll(func() {
		err := container.Terminate(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	When("fetching from the database", func() {
		var testDB database.Service
		BeforeAll(func() {
			testDB = database.New(connString)
			Expect(testDB).NotTo(BeNil())
		})

		It("should successfully connect to the database", func() {
			Expect(testDB).NotTo(BeNil())
		})

		It("should successfully fetch a role from the database", func() {
			r, err := testDB.GetRole("admin")
			Expect(err).NotTo(HaveOccurred())
            Expect(strings.TrimSpace(r.ID)).To(Equal("a"))
            Expect(r.Name).To(Equal("admin"))
		})
	})
})
