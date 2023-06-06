/*
 Navicat Premium Data Transfer

 Source Server         : cqepc auto
 Source Server Type    : SQLite
 Source Server Version : 3035005
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3035005
 File Encoding         : 65001

 Date: 21/04/2022 00:14:07
*/

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
