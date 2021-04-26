module go-zero-study

go 1.13

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/ClickHouse/clickhouse-go v1.4.3
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dsymonds/gotoc v0.0.0-20160928043926-5aebcfc91819
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/builder v0.3.4
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.2 // indirect
	github.com/iancoleman/strcase v0.1.2
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/justinas/alice v1.2.0
	github.com/lib/pq v1.8.0
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/olekukonko/tablewriter v0.0.4
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/spaolacci/murmur3 v1.1.0
	github.com/stretchr/testify v1.6.1
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/urfave/cli v1.22.4
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/xwb1989/sqlparser v0.0.0-20180606152119-120387863bf2
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/automaxprocs v1.3.0
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/grpc v1.32.0
	gopkg.in/h2non/gock.v1 v1.0.15
	gopkg.in/yaml.v2 v2.3.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)
