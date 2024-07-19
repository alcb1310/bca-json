package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	"github.com/alcb1310/bca-json/internals/server"
)

func TestRegisterCompany(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BCA Suite")
}

var _ = Describe("Company", Ordered, func() {
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

	When("registering a company", func() {
		var testDB database.Service
		var httpServer server.Server

		BeforeAll(func() {
			testDB = database.New(connString)
			Expect(testDB).NotTo(BeNil())

			httpServer = *server.New(testDB, "secret")
			Expect(httpServer).NotTo(BeNil())
		})

		It("should successfully fetch a role from the database", func() {
			r, err := testDB.GetRole("admin")
			Expect(err).NotTo(HaveOccurred())
			Expect(strings.TrimSpace(r.ID)).To(Equal("a"))
			Expect(r.Name).To(Equal("admin"))
		})

		It("should be able to register a company", func() {
			data := make(map[string]string)
			data["ruc"] = "1791838300001"
			data["name"] = "Company Name"
			data["email"] = "a@b.c"
			data["user_name"] = "alcb"
			data["password"] = "albca"
			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(data)
			Expect(err).To(BeNil())

			req, err := http.NewRequest("POST", "/api/v2/companies", &buf)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusCreated))

			data = make(map[string]string)
			data["email"] = "a@b.c"
			data["password"] = "albca"
			err = json.NewEncoder(&buf).Encode(data)
			Expect(err).To(BeNil())

			req, err = http.NewRequest("POST", "/api/v2/login", &buf)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")
			rr = httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusOK))

			token := struct {
				Token string
			}{}
			err = json.Unmarshal(rr.Body.Bytes(), &token)
			Expect(err).To(BeNil())

			Expect(token.Token).NotTo(Equal(""))
		})
	})
})
