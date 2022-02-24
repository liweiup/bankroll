package model

type Datas []struct{
	Torrents
}   //自定义类型
type Torrents map[string]interface{}  //结构体

type WenCaiDetail struct {
	StatusCode int `json:"status_code"`
	Data struct {
		//ParserData struct {
		//	Domain string `json:"domain"`
		//	Sugs []interface{} `json:"sugs"`
		//	QueryType string `json:"query_type"`
		//	ConditionID string `json:"condition_id"`
		//} `json:"parser_data"`
		//SessionCode string `json:"sessionCode"`
		//ShowSwitch bool `json:"show_switch"`
		//StatusCode int `json:"status_code"`
		//Question string `json:"question"`
		//QuestionID string `json:"question_id"`
		//Version string `json:"version"`
		Answer []struct {
			//SrcURL string `json:"src_url"`
			//UrpFlag int `json:"urp_flag"`
			//Img []interface{} `json:"img"`
			//Sug []interface{} `json:"sug"`
			//ShowDelayTime int `json:"show_delay_time"`
			//Source string `json:"source"`
			//Video []interface{} `json:"video"`
			//HisTime int64 `json:"his_time"`
			//ZhenguDuolun int `json:"zhengu_duolun"`
			//Score float64 `json:"score"`
			//DuolunSignal int `json:"duolun_signal"`
			//SecondarySource string `json:"secondary_source"`
			//ID int `json:"id"`
			//TopicID string `json:"topic_id"`
			//AnsType string `json:"ans_type"`
			//Table []interface{} `json:"table"`
			//TextAnswer string `json:"text_answer"`
			//Question string `json:"question"`
			//VoiceTxt string `json:"voice_txt"`
			//ActiveTime int `json:"activeTime"`
			//OutsideURL string `json:"outside_url"`
			//SugTitle string `json:"sug_title"`
			//Intent string `json:"intent"`
			//Diagram []interface{} `json:"diagram"`
			Txt []struct {
				//Color string `json:"color"`
				//AnsDelayTime int `json:"ansDelayTime"`
				//LinkURL string `json:"link_url"`
				//Disambiguationed string `json:"disambiguationed"`
				//Style string `json:"style"`
				//DivNo int `json:"div_no"`
				//Type string `json:"type"`
				Content struct {
					Components []struct {
						//ShowType string `json:"show_type"`
						//TitleConfig struct {
						//	ShowType string `json:"show_type"`
						//	Data struct {
						//		H1 string `json:"h1"`
						//		H2 string `json:"h2"`
						//	} `json:"data"`
						//	Config struct {
						//		Display bool `json:"display"`
						//		Type string `json:"type"`
						//		URL string `json:"url"`
						//	} `json:"config"`
						//} `json:"title_config"`
						Data struct {
							//Columns []struct {
							//	Unit string `json:"unit"`
							//	Domain string `json:"domain"`
							//	Source string `json:"source"`
							//	Type string `json:"type"`
							//	IndexName string `json:"index_name"`
							//	Key string `json:"key"`
							//	IndexID string `json:"indexId,omitempty"`
							//	Timestamp string `json:"timestamp,omitempty"`
							//	SortInfo string `json:"sort_info,omitempty"`
							//	FilterValue string `json:"filter_value,omitempty"`
							//	FilterOper string `json:"filter_oper,omitempty"`
							//	FilterKeys string `json:"filter_keys,omitempty"`
							//} `json:"columns"`
							Datas `json:"datas"`
							//Meta struct {
							//	Ret string `json:"ret"`
							//	Addcondition string `json:"addcondition"`
							//	IsCache string `json:"is_cache"`
							//	Sessionid string `json:"sessionid"`
							//	Source string `json:"source"`
							//	Qid string `json:"qid"`
							//	UrpUseSort string `json:"urp_use_sort"`
							//	Userid string `json:"userid"`
							//	Q string `json:"q"`
							//	Domain string `json:"domain"`
							//	Limit string `json:"limit"`
							//	UrpSortIndex string `json:"urp_sort_index"`
							//	Page string `json:"page"`
							//	UrpSortWay string `json:"urp_sort_way"`
							//	AddIndex string `json:"add_index"`
							//	Uuids string `json:"uuids"`
							//} `json:"meta"`
						} `json:"data"`
						//Puuid int `json:"puuid"`
						//Config struct {
						//	OtherInfo struct {
						//		FooterInfo struct {
						//			Title string `json:"title"`
						//			URL string `json:"url"`
						//		} `json:"footer_info"`
						//	} `json:"other_info"`
						//	Render struct {
						//		LeftNum string `json:"leftNum"`
						//		Merge struct {
						//			NAMING_FAILED struct {
						//				Keys []string `json:"keys"`
						//				Title string `json:"title"`
						//			} `json:"股票简称"`
						//		} `json:"merge"`
						//		Exclude []interface{} `json:"exclude"`
						//	} `json:"render"`
						//} `json:"config"`
						//UUID string `json:"uuid"`
						//OriIndex int `json:"ori_index"`
						//Cid int `json:"cid"`
					} `json:"components"`
					//RequestParams string `json:"request_params"`
					//Global struct {
					//	RequestTime string `json:"requestTime"`
					//	AppModule string `json:"appModule"`
					//	AppVersion string `json:"appVersion"`
					//	QueryType string `json:"queryType"`
					//} `json:"global"`
					//Page struct {
					//	Layout struct {
					//		LayoutData string `json:"layout_data"`
					//		LayoutMode string `json:"layout_mode"`
					//	} `json:"layout"`
					//	CacheTime int `json:"cache_time"`
					//	OriID int `json:"ori_id"`
					//	More struct {
					//		Q string `json:"q"`
					//		Codes []interface{} `json:"codes"`
					//		Isview bool `json:"isview"`
					//		DisplayBoard bool `json:"display_board"`
					//		QueryType string `json:"query_type"`
					//	} `json:"more"`
					//	Tag string `json:"tag"`
					//	OriScene int `json:"ori_scene"`
					//	Uuids []int `json:"uuids"`
					//} `json:"page"`
					ID int `json:"id"`
				} `json:"content"`
				ConditionID string `json:"condition_id"`
			} `json:"txt"`
			Sugs []interface{} `json:"sugs"`
			Idx []struct {
				Pos []int `json:"pos"`
				Type string `json:"type"`
			} `json:"idx"`
			Lenovo []interface{} `json:"lenovo"`
		} `json:"answer"`
		UserID string `json:"user_id"`
		RewriteQuestion string `json:"rewriteQuestion"`
		StatusMsg string `json:"status_msg"`
		Logid string `json:"logid"`
		Logs []interface{} `json:"logs"`
	} `json:"data"`
	StatusMsg string `json:"status_msg"`
	Logid string `json:"logid"`
}