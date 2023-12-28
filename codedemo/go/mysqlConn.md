- 数据库连接(open后需要及时关闭，否则会占用sql连接数)：
```go
func initMysql() {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, passwd, host, port, dbname))
	if err != nil {
		log.E("Connecting to mysql failed, err : %v", err)
	}
}
```

-  数据操作

```go
func (h *httpQueue) upload(db *sql.DB, potid string) {
	log.DBG("upload http response")
	header, body := h.getFeatures()
	if header == nil {
		log.E("parse response failed")
		return
	}
	var headstr = make(map[string]string)
	for key, value := range header {
		headstr[key] = ""
		for _, str := range value {
			headstr[key] += str
		}
	}
	bt, _ := json.Marshal(headstr)
	if len(header) != 0 && len(body) != 0 {
		uploadItem := httpUpload{
			potid,
			h.url,
			bt,
			string(body),
			time.Now(),
		}
		res, err := db.Query("SELECT 1 FROM emulator_response where pot_id=? and url=?", uploadItem.PotType, uploadItem.ReqUrl)
		if err != nil {
			log.E("Error when insert data, %v", err)
			return
		}
		if res.Next() {
			_, err = db.Exec("DELETE FROM emulator_response where pot_id=? and url=?", uploadItem.PotType, uploadItem.ReqUrl)
			if err != nil {
				log.E("Error when insert data, %v", err)
			}
		}
		res.Close()
		_, err = db.Exec("INSERT INTO emulator_response(pot_id, url, headers, httpresponse, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?)",
			uploadItem.PotType, uploadItem.ReqUrl, uploadItem.RespHeader, uploadItem.RespBody, uploadItem.Time, uploadItem.Time)
		if err != nil {
			log.E("Error when insert data, %v", err)
		} else {
			log.DBG("uploaded")
		}

		// log.DBG("sendreq: %s\nsendresp header: %s\nsendresp body: %s\n", uploadItem.ReqUrl, string(uploadItem.RespHeader), string(uploadItem.RespBody))
	} else {
		log.E("Error occured when parsing response package")
	}

}
```
