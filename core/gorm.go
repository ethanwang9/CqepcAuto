package core

import (
	"CqepcAuto/global"
	"CqepcAuto/utils"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// name: 数据库
// author: axel wong
// desc: sqlite3 数据库初始化

// GormBySqliteInit 初始化
func GormBySqliteInit() *gorm.DB {
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表名复数
		},
	}

	// 判断是否有数据文件
	if is, _ := utils.PathExists("./db/"); !is {
		utils.CreateDir("./db/")
	}
	is, _ := utils.PathExists("./db/db.db")

	db, err := gorm.Open(sqlite.Open("./db/db.db"), &gormConfig)
	if err != nil {
		global.APP_LOG.Panic("数据库初始化失败", zap.Error(err))
	}

	// 没有数据库文件运行sql语句
	if !is {
		sqlString := `
PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS "log";
CREATE TABLE "log" (
  "uuid" text NOT NULL,
  "uid" text NOT NULL,
  "msg" TEXT NOT NULL,
  "msg_type" TEXT NOT NULL,
  "data" TEXT NOT NULL,
  "cqepc_auto_flag" TEXT NOT NULL,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  PRIMARY KEY ("uuid")
);

-- ----------------------------
-- Table structure for tj
-- ----------------------------
DROP TABLE IF EXISTS "tj";
CREATE TABLE "tj" (
  "uuid" text NOT NULL,
  "uid" text NOT NULL,
  "data" TEXT NOT NULL,
  "cqepc_auto_flag" TEXT NOT NULL,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  PRIMARY KEY ("uuid")
);

-- ----------------------------
-- Table structure for today
-- ----------------------------
DROP TABLE IF EXISTS "today";
CREATE TABLE "today" (
  "uid" text NOT NULL,
  "p_id" text NOT NULL,
  "data" TEXT NOT NULL,
  "pk_data" TEXT,
  "is_pk" text NOT NULL,
  "cqepc_auto_flag" TEXT NOT NULL,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
  "uid" text NOT NULL,
  "s_id" text NOT NULL,
  "s_pass" TEXT NOT NULL,
  "s_openid" text,
  "s_name" TEXT NOT NULL,
  "s_class" TEXT NOT NULL,
  "s_class_code" TEXT NOT NULL,
  "pk_token" TEXT NOT NULL,
  "login_type" TEXT NOT NULL,
  "n_token" TEXT NOT NULL,
  "n_secret" TEXT NOT NULL,
  "n_phone" TEXT NOT NULL,
  "is_stop" TEXT NOT NULL,
  "cqepc_auto_flag" TEXT NOT NULL,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  PRIMARY KEY ("uid")
);

PRAGMA foreign_keys = true;
`
		db.Exec(sqlString)
	}

	return db
}
