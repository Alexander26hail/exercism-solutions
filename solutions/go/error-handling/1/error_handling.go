package erratum
//import "errors"
func Use(opener ResourceOpener, input string) (err error) {
    var recurso Resource
	for {
        recurso , err = opener()
        if err == nil{
            break
        }
        _ , isTransit := err.(TransientError)

        if !isTransit{
            return err
        }
    }
    defer recurso.Close()
	defer func(){
      r:=   recover()
      if r !=nil{
          value , ok := r.(FrobError)
          if ok {
              recurso.Defrob(value.defrobTag)
 
          }

          value2, _ := r.(error)
		  err = value2
         
          
      }
    }()
	
	recurso.Frob(input)
    
    return nil
}
