package pkg

import (
	"github.com/digitalcircle-com-br/restdb/pkg/api"
	"github.com/digitalcircle-com-br/restdb/pkg/infra/db"
	"github.com/digitalcircle-com-br/restdb/pkg/util"
)

func Setup() error {
	return util.Series(db.Setup,api.Setup)
}
