package main

func parse() {
	///  ここでパースしながらisStruing?にチャネルを渡す
	for {
		if !isString(sssssss) {
			continue
		}
		if !isSql(sssssss) {
			continue
		}
		err := format(ssssss)
		if err != nil {
			return err
		}
	}
}

func isString(input interface{}) bool {
	_, ok := input.(string)
	return ok
}

func isSQL(input interface{}) bool {
	input, _ = input.(string)

	token, err := breakDowonToToken(input)
	if err != nil {
		return
	}
}

func breakDowonToToken(input string) {

}

//どうやってSQLか判断するの？
//与えれた文字列をtokenに分割して、そこにSELECが含まれてるばあいはtrueにして、その文章をSQlと判断する
