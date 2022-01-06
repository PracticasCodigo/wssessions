package settings

var FAKE_API = true


func IsFakeAPI() bool{
	if FAKE_API{
		return true
	}else{
		return false
	}
}
