package jsEditor

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)
import "github.com/robertkrimen/otto"

// THIS IS ADDED
// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-helloworld")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// THIS HAS CHANGED
// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	vm := otto.New()
	// vm.Run(`
	// 	abc = 2 + 2;
	// 	console.log("The value of abc is " + abc); // 4
	// `)

	args := context.GetInput("args")

	vm.Set("args", args)
	vm.Set("context", context)

	log.Debugf("The value of k is %s", context.ActivityHost().Name()) // 4

	if _, err := vm.Run(context.GetInput("jscode")); err == nil {
		if value, err := vm.Get("args"); err == nil {
			if argsOut1, err := value.ToString(); err == nil {
				log.Debugf("The value of i is %s", argsOut1) // 4
				// context.SetOutput("args_out", argsOut)
				context.SetOutput("args_out", argsOut1)

				log.Debugf("The value of j is %s", argsOut1) // 4
			}
		}
	} else {
		context.SetOutput("args_out", args)
	}

	// 	if argsOut1, err := vm.Get("args"); err == nil {
	// 		log.Debugf("The value of i is %s", argsOut1) // 4
	// 		// context.SetOutput("args_out", argsOut)
	// 		context.SetOutput("args_out", argsOut1.(string))

	// 		log.Debugf("The value of j is %s", argsOut1) // 4

	// 	}
	// }
	// Get the activity data from the context
	// name := context.GetInput("name").(string)
	// salutation := context.GetInput("salutation").(string)

	// Use the log object to log the greeting
	//log.Debugf("The Flogo engine says [%s] to [%s]", salutation, name)

	// Set the result as part of the context
	// if value, err := vm.Get("abc"); err == nil {

	// }

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
