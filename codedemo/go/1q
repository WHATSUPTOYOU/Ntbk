使用库：
https://github.com/segmentio/kafka-go 或
https://github.com/IBM/sarama

// 函数功能：创建指定消息发送至kafka，KafkaAddr为kafka连接地址，在docker容器中时需要在/etc/hosts中创建域名映射，并通过域名访问
```go
func CreateSpcialPotInfo(serverPort int, srcIP string, srcPort int, proxyIP string, proxyPort int, potname string) {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(KafkaAddr),
		Topic:                  fmt.Sprintf("pot_%s", potname),
		AllowAutoTopicCreation: true,
	}

	recordSrc, err := DB.City(net.ParseIP(srcIP))
	if err != nil {
		log.E(err.Error())
	}

	recordDst, err := DB.City(net.ParseIP(proxyIP))
	if err != nil {
		log.E(err.Error())
	}

	var srcregion, dstregion, srccode, dstcode string

	if len(recordSrc.Subdivisions) == 0 {
		srcregion = ""
		srccode = ""
	} else {
		srcregion = recordSrc.Subdivisions[0].Names["en"]
		srccode = recordSrc.Subdivisions[0].IsoCode
	}

	if len(recordDst.Subdivisions) == 0 {
		dstregion = ""
		dstcode = ""
	} else {
		dstregion = recordDst.Subdivisions[0].Names["en"]
		dstcode = recordDst.Subdivisions[0].IsoCode
	}

	value := KafkaMsg{
		ServerPort: serverPort,
		SrcInfo: PositionInfo{
			IP:            srcIP,
			Port:          srcPort,
			CountryName:   recordSrc.Country.Names["en"],
			RegionName:    srcregion,
			CityName:      recordSrc.City.Names["en"],
			Latigude:      recordSrc.Location.Latitude,
			Longitude:     recordSrc.Location.Longitude,
			RegionCode:    srccode,
			CountryCode2:  recordSrc.Country.IsoCode,
			ContitentCode: recordSrc.Continent.Code,
			Timezone:      recordSrc.Location.TimeZone,
		},
		DstInfo: PositionInfo{
			IP:            proxyIP,
			Port:          proxyPort,
			CountryName:   recordDst.Country.Names["en"],
			RegionName:    dstregion,
			CityName:      recordDst.City.Names["en"],
			Latigude:      recordDst.Location.Latitude,
			Longitude:     recordDst.Location.Longitude,
			RegionCode:    dstcode,
			CountryCode2:  recordDst.Country.IsoCode,
			ContitentCode: recordDst.Continent.Code,
			Timezone:      recordDst.Location.TimeZone,
		},
	}
	log.DBG("kafkaMsg: %v", value)
	buf, err := packet.JSON(value)
	if err != nil {
		log.E("parse kafka message error. %v.", err)
	}
	msgbytes := buf.Bytes()
	// binary.Write(buf, binary.LittleEndian, value)
	messages := []kafka.Message{
		{
			Key:   []byte("specinfo"),
			Value: msgbytes,
		},
	}

	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// attempt to create topic prior to publishing the message
		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.E("unexpected error %v", err)
		}
		break
	}

	if err := w.Close(); err != nil {
		log.E("failed to close writer:", err)
	}
}```
