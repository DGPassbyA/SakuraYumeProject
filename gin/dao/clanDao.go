package clanDao

import (
	"database/sql"
	"fmt"
	"log"
	"main/conf"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Member struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
}
type History struct {
	Eid int `json:"eid"`
	// Uid   int    `json:"uid"`
	Player string `json:"player"`
	Time   string `json:"time"`
	Round  int    `json:"round"`
	Boss   int    `json:"boss"`
	Dmg    int    `json:"dmg"`
	Flag   int    `json:"flag"`
}

func _connect() (*sql.DB, error) {
	database_path := conf.DatabasePath
	db, err := sql.Open("sqlite3", database_path)
	return db, err
}
func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func GetGidAndCidByClanName(clanName string) (int, int) {
	if clanName == "" {
		return 0, -1
	}
	db, err := _connect()
	checkErr(err)
	defer db.Close()
	sql := "SELECT gid, cid FROM clan WHERE name = ? "
	rows, err := db.Query(sql, clanName)
	checkErr(err)
	defer rows.Close()
	var gid, cid int
	for rows.Next() {
		err := rows.Scan(&gid, &cid)
		checkErr(err)
	}
	return gid, cid
}

//公会名查询公会成员
func GetMembersByClanName(clanName string) []Member {
	gid, _ := GetGidAndCidByClanName(clanName)
	if gid == 0 {
		return nil
	}
	db, err := _connect()
	checkErr(err)
	defer db.Close()
	sql := "SELECT uid,name FROM member WHERE gid = ? "
	rows, err := db.Query(sql, gid)
	checkErr(err)
	defer rows.Close()
	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.Uid, &member.Name)
		checkErr(err)
		members = append(members, member)
	}
	return members
}
func getNameByUid(uid int, members []Member) string {
	for _, m := range members {
		if m.Uid == uid {
			return m.Name
		}
	}
	return "???"
}

//时间+公会名查询公会出刀记录
func GetHistoryByTime(clanName string, time string) []History {
	if clanName == "" {
		return nil
	}
	gid, cid := GetGidAndCidByClanName(clanName)
	if gid == 0 || cid == -1 {
		return nil
	}
	tableName := fmt.Sprintf("battle_%d_%d_%s", gid, cid, time)
	db, err := _connect()
	checkErr(err)
	defer db.Close()
	sql := "SELECT * FROM " + tableName + " WHERE alt = ? "
	rows, err := db.Query(sql, gid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var historyList []History
	members := GetMembersByClanName(clanName)
	for rows.Next() {
		var history History
		var uid, alt int
		err := rows.Scan(&history.Eid, &uid, &alt, &history.Time, &history.Round, &history.Boss, &history.Dmg, &history.Flag)
		checkErr(err)
		if tmp_name := getNameByUid(uid, members); tmp_name != "" {
			history.Player = tmp_name
		}
		historyList = append(historyList, history)
	}
	return historyList
}
func InsertHistory(clanName string, month string, uid, round, boss, dmg, flag uint) bool {
	gid, cid := GetGidAndCidByClanName(clanName)
	if gid == 0 || cid == -1 {
		return false
	}
	tableName := fmt.Sprintf("battle_%d_%d_%s", gid, cid, month)
	//flag：整刀0 尾刀1 补时刀2
	db, err := _connect()
	checkErr(err)
	dt := time.Now()
	sqlStatement, err := db.Prepare("INSERT INTO " + tableName + " (`uid`, `alt`, `time`, `round`, `boss`, `dmg`, `flag`) VALUES (?, ?, ?, ?, ?, ?, ?)")
	res, err := sqlStatement.Exec(uid, gid, dt, round, boss, dmg, flag)
	checkErr(err)
	rowCount, err := res.RowsAffected()
	db.Close()
	if rowCount == 1 {
		return true
	} else {
		return false
	}
}
