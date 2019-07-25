package utils

//结构体转map
func Struct2Map(obj interface{}) (data map[interface{}]interface{}, err error) {
	data = make(map[interface{}]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	err = nil
	return
}
