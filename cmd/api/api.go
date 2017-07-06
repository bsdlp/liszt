package main

import (
	"context"
	"net/http"
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
	"github.com/kelseyhightower/envconfig"
	"github.com/liszt-code/liszt/cmd/api/internal"
	"github.com/liszt-code/liszt/pkg/registry"
	"github.com/liszt-code/liszt/pkg/registry/resolver"
	"github.com/liszt-code/liszt/pkg/registry/schema"
	graphql "github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/sirupsen/logrus"
)

// Config holds configuration options
type Config struct {
	AWSRegion string `envconfig:"aws_region" default:"us-west-2"`

	BindAddress string `evconfig:"bind_address" default:":8080"`

	BuildingTableName string `envconfig:"building_table_name" default:"liszt-buildings-dev"`
	UnitTableName     string `envconfig:"unit_table_name" default:"liszt-units-dev"`
	ResidentTableName string `envconfig:"resident_table_name" default:"liszt-residents-dev"`
}

type panicLogger struct {
	logger logrus.FieldLogger
}

func (pl panicLogger) LogPanic(ctx context.Context, value interface{}) {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]

	pl.logger.WithFields(logrus.Fields{
		"stack": string(buf),
		"error": "graphql: panic occured",
	}).Error(value)
}

func main() {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	cfg := new(Config)
	err := envconfig.Process("liszt", cfg)
	if err != nil {
		logger.Fatal(err)
	}

	sess := session.New(aws.NewConfig().WithRegion(cfg.AWSRegion))
	registrar := &registry.DynamoRegistrar{
		DB: dynamodb.New(sess),
		Config: &registry.DynamoConfig{
			BuildingTableName: cfg.BuildingTableName,
			UnitTableName:     cfg.UnitTableName,
			ResidentTableName: cfg.ResidentTableName,
		},
	}

	gqlSchema, err := schema.Build()
	if err != nil {
		logger.Fatal(err)
	}

	res := &resolver.Resolver{
		Registrar: registrar,
		Logger:    logger,
	}

	scheme, err := graphql.ParseSchema(gqlSchema, res, graphql.Logger(panicLogger{logger: logger}))
	if err != nil {
		logger.Fatal(err)
	}

	mux := chi.NewMux()
	mux.Mount("/v1", internal.NewCRUDService(registrar))
	mux.Handle("/query", &relay.Handler{Schema: scheme})

	mux.Get("/ide", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(gqlIDEPage)
		if err != nil {
			logger.Error(err)
		}
	})

	logger.Fatal(http.ListenAndServe(cfg.BindAddress, mux))
}

var gqlIDEPage = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.3.2/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.3.2/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				graphQLParams.variables = graphQLParams.variables ? JSON.parse(graphQLParams.variables) : null;
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
