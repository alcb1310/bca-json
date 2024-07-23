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

var companyID = uuid.MustParse("4b114c26-b038-4cfa-ae6e-ad46c73ef59d")

var _ = Describe("Projects", Ordered, func() {
	var container testcontainers.Container
	var ctx context.Context
	var connString string
	dbName := "bcatestcompany"
	dbUser := "testuser"
	dbPassword := "testproject"

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
				filepath.Join("..", "..", "scripts", "u002_project.sql"),
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

	When("managing projects", func() {
		var httpServer server.Server
		var token string

		BeforeEach(func() {
			testDB := database.New(connString)
			Expect(testDB).NotTo(BeNil())

			httpServer = *server.New(testDB, "secret")
			Expect(httpServer).NotTo(BeNil())

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

        It("should get all the projects", func() {
            req, err := http.NewRequest("GET", "/api/v2/bca/projects", nil)
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            Expect(err).To(BeNil())

            rr := httptest.NewRecorder()
            httpServer.Server.ServeHTTP(rr, req)
            Expect(rr.Code).To(Equal(http.StatusOK))

            var projectsResponse []types.Project
            err = json.Unmarshal(rr.Body.Bytes(), &projectsResponse)
            Expect(err).To(BeNil())
            Expect(len(projectsResponse)).To(Equal(3))
        })

        It("should get a project", func() {
            req, err := http.NewRequest("GET", "/api/v2/bca/projects/bcf847df-f415-469e-ae67-dfac273b0e82", nil)
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            Expect(err).To(BeNil())

            rr := httptest.NewRecorder()
            httpServer.Server.ServeHTTP(rr, req)
            Expect(rr.Code).To(Equal(http.StatusOK))

            var projectResponse map[string]types.Project
            err = json.Unmarshal(rr.Body.Bytes(), &projectResponse)
            Expect(err).To(BeNil())
            Expect(projectResponse["project"].ID).NotTo(BeNil())
            Expect(projectResponse["project"].Name).To(Equal("Project 1"))
            Expect(projectResponse["project"].GrossArea).To(Equal(100.50))
            Expect(projectResponse["project"].NetArea).To(Equal(50.25))
            Expect(projectResponse["project"].CompanyID).To(Equal(companyID))
        })

        It("should return not found when getting a project that does not exist", func() {
            req, err := http.NewRequest("GET", "/api/v2/bca/projects/00000000-0000-0000-0000-000000000000", nil)
            req.Header.Set("Content-Type", "application/json")
            req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
            Expect(err).To(BeNil())

            rr := httptest.NewRecorder()
            httpServer.Server.ServeHTTP(rr, req)
            Expect(rr.Code).To(Equal(http.StatusNotFound))
        })

		It("should create a project", func() {
			var buf bytes.Buffer
			data := make(map[string]interface{})
			data["name"] = "Test Project"
			data["gross_area"] = 100
			data["net_area"] = 100

			err := json.NewEncoder(&buf).Encode(data)
			Expect(err).To(BeNil())
			req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusCreated))

			var projectResponse map[string]types.Project
			err = json.Unmarshal(rr.Body.Bytes(), &projectResponse)
			Expect(err).To(BeNil())
			Expect(projectResponse["project"].ID).NotTo(BeNil())
			Expect(projectResponse["project"].GrossArea).To(Equal(100.0))
			Expect(projectResponse["project"].NetArea).To(Equal(100.0))
			Expect(projectResponse["project"].LastClosure).To(BeNil())
			Expect(projectResponse["project"].Name).To(Equal("Test Project"))
		})

		It("should conflict if project already exists", func() {
			var buf bytes.Buffer
			data := make(map[string]interface{})
			data["name"] = "Test Project"
			data["gross_area"] = 100
			data["net_area"] = 100

			err := json.NewEncoder(&buf).Encode(data)
			Expect(err).To(BeNil())
			req, err := http.NewRequest("POST", "/api/v2/bca/projects", &buf)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			Expect(err).To(BeNil())

			rr := httptest.NewRecorder()
			httpServer.Server.ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusConflict))
		})
	})
})
