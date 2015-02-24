package sunday

import ()

/*
    ID      string
    Values  map[interface{}]interface{}
    Options *Options
    IsNew   bool
    */

type Session interface {
    ID() string  
    
}