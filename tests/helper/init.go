package helper

import (
	"log"
	"os"

	"github.com/obulpathi/boltcoin/ethutil"
	logpkg "github.com/obulpathi/boltcoin/logger"
)

var Logger logpkg.LogSystem
var Log = logpkg.NewLogger("TEST")

func init() {
	Logger = logpkg.NewStdLogSystem(os.Stdout, log.LstdFlags, logpkg.InfoLevel)
	logpkg.AddLogSystem(Logger)

	ethutil.ReadConfig(".ethtest", "/tmp/ethtest", "")
	ethutil.Config.Db, _ = NewMemDatabase()
}
