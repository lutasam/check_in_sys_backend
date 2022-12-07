package utils

import "github.com/bwmarrin/snowflake"

var (
	userIDGenerator   *snowflake.Node
	recordIDGenerator *snowflake.Node
	fileIDGenerator   *snowflake.Node
	noticeIDGenerator *snowflake.Node
)

func init() {
	var err error
	userIDGenerator, err = snowflake.NewNode(100)
	if err != nil {
		panic(err)
	}

	recordIDGenerator, err = snowflake.NewNode(200)
	if err != nil {
		panic(err)
	}

	fileIDGenerator, err = snowflake.NewNode(300)
	if err != nil {
		panic(err)
	}

	noticeIDGenerator, err = snowflake.NewNode(400)
	if err != nil {
		panic(err)
	}
}

func GenerateUserID() uint64 {
	return uint64(userIDGenerator.Generate().Int64())
}

func GenerateRecordID() uint64 {
	return uint64(recordIDGenerator.Generate().Int64())
}

func GenerateFileID() uint64 {
	return uint64(fileIDGenerator.Generate().Int64())
}

func GenerateNoticeID() uint64 {
	return uint64(noticeIDGenerator.Generate().Int64())
}
