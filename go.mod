module github.com/rudderlabs/rudder-server

go 1.16

require (
	cloud.google.com/go/bigquery v1.8.0
	cloud.google.com/go/pubsub v1.3.1
	cloud.google.com/go/storage v1.10.0
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/azure-storage-blob-go v0.8.0
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.8.1 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/ClickHouse/clickhouse-go v1.4.8
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/EagleChen/mapmutex v0.0.0-20180418073615-e1a5ae258d8d // indirect
	github.com/EagleChen/restrictor v0.0.0-20180420073700-9b81bbf8df1d
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/Shopify/sarama v1.26.1
	github.com/araddon/dateparse v0.0.0-20190622164848-0fb0a474d195
	github.com/aws/aws-sdk-go v1.37.23
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bugsnag/bugsnag-go v1.5.3
	github.com/bugsnag/panicwrap v1.2.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/containerd/continuity v0.1.0 // indirect
	github.com/denisenkom/go-mssqldb v0.10.0
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/fsnotify/fsnotify v1.5.1
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/go-ini/ini v1.63.2 // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.0
	github.com/gomodule/redigo v1.8.5
	github.com/gorilla/mux v1.7.4
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/hashicorp/go-retryablehttp v0.6.6
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce
	github.com/iancoleman/strcase v0.1.3
	github.com/jeremywohl/flatten v1.0.1
	github.com/jinzhu/copier v0.3.2
	github.com/joho/godotenv v1.3.0
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.3.0
	github.com/mattn/go-ieproxy v0.0.0-20200203040449-2dbc853185d9 // indirect
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/minio/minio-go/v6 v6.0.49
	github.com/mkmik/multierror v0.3.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.10.1
	github.com/opencontainers/runc v1.0.1 // indirect
	github.com/ory/dockertest v3.3.5+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4 // indirect
	github.com/rs/cors v1.7.0
	github.com/rudderlabs/analytics-go v3.3.1+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/segmentio/backo-go v0.0.0-20160424052352-204274ad699c // indirect
	github.com/segmentio/ksuid v1.0.2
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/snowflakedb/gosnowflake v1.3.4
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.5.0
	github.com/tidwall/gjson v1.5.0
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/tidwall/sjson v1.0.4
	github.com/xdg/scram v1.0.3
	github.com/xitongsys/parquet-go v1.6.1-0.20210531003158-8ed615220b7d
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.14.0
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools v0.1.6 // indirect
	google.golang.org/api v0.39.0
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/alexcesaro/statsd.v2 v2.0.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gotest.tools v2.2.0+incompatible // indirect
)
