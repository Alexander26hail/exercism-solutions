package phonenumber
import "strings"
import "fmt"

func Number(phoneNumber string) (string, error) {
    r := strings.NewReplacer("-", "", "+", "", " ", "", ".", "", "(", "", ")", "","@:!-","")
    clear := r.Replace(phoneNumber)
    
    if len(clear) == 11 && clear[0] == '1' {
        clear = clear[1:]
    }
    if len(clear) != 10 {
        return "", fmt.Errorf("invalid number")
    }
    for i:=0; i<len(clear); i ++{
        if clear[i]<'0' || clear[i]>'9'{
            return "", fmt.Errorf("invalid number") 
        }
    }
    if clear[0] < '2' || clear[3] < '2' {
    	return "", fmt.Errorf("invalid number")
	}
    return clear, nil
}

func AreaCode(phoneNumber string) (string, error) {
    clear, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return clear[:3], nil
}

func Format(phoneNumber string) (string, error) {
    clear, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("(%s) %s-%s", clear[:3], clear[3:6], clear[6:]), nil
}
