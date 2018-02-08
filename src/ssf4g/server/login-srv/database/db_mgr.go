package dbmgr

import (
	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/database/login-dao"
	"taptap-kit/lib/conf"
)

var (
	_market_dao *marketdao.MarketDao

	_stats_dao *statsdao.StatsDao

	_forum_dao *forumdao.ForumDao
)

func init() {
	marketUrl := ""
	statsUrl := ""
	forumUrl := ""

	maxIdleConn, err := conf.GetIniData().Int("db_max_idle_conn")

	if err != nil {
		errMsg := tlog.Error("init database (max idle conn) err (%v).", err)

		tlog.AsyncSend(tlog.NewErrorData(err, errMsg))

		return
	}

	maxOpenConn, err := conf.GetIniData().Int("db_max_open_conn")

	if err != nil {
		errMsg := tlog.Error("init database (max open conn) err (%v).", err)

		tlog.AsyncSend(tlog.NewErrorData(err, errMsg))

		return
	}

	runMode := conf.GetIniData().String("run_mode")

	if runMode == "prod" {
		marketUrl = conf.GetIniData().String("prod::market_db")

		statsUrl = conf.GetIniData().String("prod::stats_db")

		forumUrl = conf.GetIniData().String("prod::forum_db")
	} else {
		marketUrl = conf.GetIniData().String("dev::market_db")

		statsUrl = conf.GetIniData().String("dev::stats_db")

		forumUrl = conf.GetIniData().String("dev::forum_db")
	}

	_market_dao = &marketdao.MarketDao{}

	errData := _market_dao.InitMarketDao(marketUrl, maxIdleConn, maxOpenConn)

	if errData != nil {
		errMsg := tlog.Error("init market dao (%s, %d, %d) err (%v).", marketUrl, maxIdleConn, maxOpenConn, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	_stats_dao = &statsdao.StatsDao{}

	errData = _stats_dao.InitStatsDao(statsUrl, maxIdleConn, maxOpenConn)

	if errData != nil {
		errMsg := tlog.Error("init stats dao (%s, %d, %d) err (%v).", statsUrl, maxIdleConn, maxOpenConn, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	_forum_dao = &forumdao.ForumDao{}

	errData = _forum_dao.InitForumDao(forumUrl, maxIdleConn, maxOpenConn)

	if errData != nil {
		errMsg := tlog.Error("init forum dao (%s, %d, %d) err (%v).", forumUrl, maxIdleConn, maxOpenConn, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	return
}

func GetMarketDao() *marketdao.MarketDao {
	return _market_dao
}

func GetStatsDao() *statsdao.StatsDao {
	return _stats_dao
}

func GetForumDao() *forumdao.ForumDao {
	return _forum_dao
}
