package tools
import "log"

func errNotify(err error, notify string)  error{
	if err != nil{
		log.Println(notify,": ",  err)
		return err
	}
	return nil
}
