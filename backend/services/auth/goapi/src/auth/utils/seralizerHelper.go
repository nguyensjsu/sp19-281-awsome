package utils

//
//// go binary encoder
//func ToGOB64(m models.Session) string {
//	b := bytes.Buffer{}
//	e := gob.NewEncoder(&b)
//	err := e.Encode(m)
//	if err != nil { fmt.Println(`failed gob Encode`, err) }
//	return base64.StdEncoding.EncodeToString(b.Bytes())
//}
//
//// go binary decoder
//func FromGOB64(str string) models.Session {
//	m := models.Session{}
//	by, err := base64.StdEncoding.DecodeString(str)
//	if err != nil { fmt.Println(`failed base64 Decode`, err); }
//	b := bytes.Buffer{}
//	b.Write(by)
//	d := gob.NewDecoder(&b)
//	err = d.Decode(&m)
//	if err != nil { fmt.Println(`failed gob Decode`, err); }
//	return m
//}
