package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"go-zero-study/core/stores/cache"
	"go-zero-study/core/stores/sqlc"
	"go-zero-study/core/stores/sqlx"
	"go-zero-study/core/stringx"
	"go-zero-study/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"strings"
	{{if .time}}"time"{{end}}

	"go-zero-study/core/stores/sqlc"
	"go-zero-study/core/stores/sqlx"
	"go-zero-study/core/stringx"
	"go-zero-study/tools/goctl/model/sql/builderx"
)
`
)
