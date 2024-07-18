package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/alcb1310/bca-json/internals/database"
	"github.com/alcb1310/bca-json/internals/server"
	"github.com/alcb1310/bca-json/internals/types"
)

var _ = Describe("Users", Ordered, func() {
	var conntainer testcontainers.Container
	var ctx context.Context
	var connString string
	dbName := "bcatestcompany"
	dbUser := "testuser"
	dbPassword := "testuser"

	BeforeAll(func() {
		ctx = context.Background()
		c, err := postgres.Run(ctx,
			"docker.io/postgres:14-alpine",
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
				filepath.Join("..", "..", "scripts", "u001_user.sql"),
			),
		)
		Expect(err).NotTo(HaveOccurred())
		connString, err = c.ConnectionString(ctx)
		Expect(err).NotTo(HaveOccurred())
		conntainer = c
	})

	AfterAll(func() {
		err := conntainer.Terminate(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	When("managing users", func() {
		var httpServer server.Server
		var token string

		BeforeAll(func() {
			testDB := database.New(connString)
			Expect(testDB).NotTo(BeNil())

			httpServer = *server.New(testDB, "secret")
			Expect(httpServer).NotTo(BeNil())
			var buf bytes.Buffer
			data := make(map[string]string)
			data["email"] = "test@test.com"
			data["password"] = "test"
			err := json.NewEncoder(&buf).Encode(data)
			Expect(err).To(BeNil())

			req, err := http.NewRequest("POST", "/api/v2/login", &buf)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusOK))

			tokenResponse := struct {
				Token string
			}{}
			err = json.Unmarshal(rr.Body.Bytes(), &tokenResponse)
			Expect(err).To(BeNil())
			token = tokenResponse.Token
		})

		It("should get all the users", func() {
			req, err := http.NewRequest("GET", "/api/v2/bca/users", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusOK))

			var usersResponse []types.User
			err = json.Unmarshal(rr.Body.Bytes(), &usersResponse)
			Expect(err).To(BeNil())
			Expect(len(usersResponse)).To(Equal(2))
		})

		It("should get the user by id", func() {
			req, err := http.NewRequest("GET", "/api/v2/bca/users/0cd001ff-2c33-460b-8876-73e51dfb053e", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusOK))

			var user types.User
			err = json.Unmarshal(rr.Body.Bytes(), &user)
			Expect(err).To(BeNil())
			Expect(user.ID).To(Equal(uuid.MustParse("0cd001ff-2c33-460b-8876-73e51dfb053e")))
			Expect(user.Name).To(Equal("Another User"))
			Expect(user.Email).To(Equal("a@b.c"))
		})

		It("should return 404 when user not found", func() {
			req, err := http.NewRequest("GET", "/api/v2/bca/users/4cd001ff-2c33-460b-8876-73e51dfb053e", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusNotFound))

            errorResponse := struct {
                Error string
            }{}
            err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
            Expect(err).To(BeNil())
            Expect(errorResponse.Error).To(Equal("User not found"))
		})

        It("should return the logged in user", func() {
            req, err := http.NewRequest("GET", "/api/v2/bca/users/me", nil)
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            Expect(err).To(BeNil())

            rr := httptest.NewRecorder()
            httpServer.Server.ServeHTTP(rr, req)
            Expect(rr.Code).To(Equal(http.StatusOK))

            var user types.User
            err = json.Unmarshal(rr.Body.Bytes(), &user)
            Expect(err).To(BeNil())
            Expect(user.ID).To(Equal(uuid.MustParse("0cd002ff-2c33-460b-8876-73e51dfb053e")))
            Expect(user.Name).To(Equal("Test User"))
            Expect(user.Email).To(Equal("test@test.com"))
        })

        It("should create a user", func() {
            req, err := http.NewRequest("POST", "/api/v2/bca/users", bytes.NewBuffer([]byte(`{"name": "Testing User", "email": "testing@test.com", "password": "test"}`)))
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            Expect(err).To(BeNil())

            rr := httptest.NewRecorder()
            httpServer.Server.ServeHTTP(rr, req)
            Expect(rr.Code).To(Equal(http.StatusCreated))

            var user types.User
            var userResponse map[string]types.User
            err = json.Unmarshal(rr.Body.Bytes(), &userResponse)
            user = userResponse["user"]
            Expect(err).To(BeNil())
            Expect(user.ID).NotTo(BeNil())
            Expect(user.Name).To(Equal("Testing User"))
            Expect(user.Email).To(Equal("testing@test.com"))
        })

        It("should return 409 when user already exists", func() {
            req, err := http.NewRequest("POST", "/api/v2/bca/users", bytes.NewBuffer([]byte(`{"name": "Test User", "email": "test@test.com", "password": "test"}`)))
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            Expect(err).To(BeNil())

            rr := httptest.NewRecorder()
            httpServer.Server.ServeHTTP(rr, req)
            Expect(rr.Code).To(Equal(http.StatusConflict))
        })
	})
})
