
# Installation 
```
go get -u github.com/fatih/color
go get -u github.com/joaobaiense/logger
```

# Example go file 

```
package main 
import (
    log "github.com/joaobaiense/logger"
)
func main(){
    log.Init(os.Stdout)
    log.Info("Hey, this is in color")
}
```
