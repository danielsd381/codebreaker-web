package main

var secret = ""
func setSecret(sec string){
	secret = sec;
}

func validate(value string)string{
	result := ""
	for i := 0; i < len(value); i++ {
		for j := 0; j < len(secret); j++ {
			if (i==j)&&(value[i]==secret[j]){
					result = "X"+result;
					break;
			}else if(i!=j) && (value[i]==secret[j]){
					result += "-";
					break;
			}
		}
	}
	return result
}
