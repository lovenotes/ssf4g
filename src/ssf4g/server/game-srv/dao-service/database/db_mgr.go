package dbmgr

import (
	"ssf4g/common/tlog"
	"ssf4g/server/game-srv/common/srv-config"
	"ssf4g/server/game-srv/dao-service/database/game-dao"
)

var (
	_game_dao *gamedao.GameDao
)

func init() {
	maxIdleConn := srvconfig.GetConfig().DBMaxIdleConn
	maxOpenConn := srvconfig.GetConfig().DBMaxOpenConn

	gameUrl := srvconfig.GetConfig().GameDB

	_game_dao = &gamedao.GameDao{}

	errData := _game_dao.InitGameDao(gameUrl, maxIdleConn, maxOpenConn)

	if errData != nil {
		errMsg := tlog.Error("init game dao (%s, %d, %d) err (%v).", gameUrl, maxIdleConn, maxOpenConn, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	return
}

func GetGameDao() *gamedao.GameDao {
	return _game_dao
}
