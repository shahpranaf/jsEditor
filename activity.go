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

	returnVar := true

	arg1 := context.GetInput("arg1")
	arg2 := context.GetInput("arg2")
	arg3 := context.GetInput("arg3")

	vm.Set("arg1", arg1)
	vm.Set("arg2", arg2)
	vm.Set("arg3", arg3)
	vm.Set("returnVar", true)

	// log.Debugf("The value of k is %s", context.ActivityHost().Name()) // 4

	if _, err := vm.Run(context.GetInput("jscode")); err == nil {
		if value1, err := vm.Get("arg1"); err == nil {
			if argsOut1, err := value1.Export(); err == nil {
				context.SetOutput("arg1_out", argsOut1)
				log.Debugf("The value of j is %s", argsOut1) // 4
			}
		}
		if value2, err := vm.Get("arg2"); err == nil {
			if argsOut2, err := value2.Export(); err == nil {
				context.SetOutput("arg2_out", argsOut2)
				log.Debugf("The value of j is %s", argsOut2) // 4
			}
		}
		if value3, err := vm.Get("arg3"); err == nil {
			if argsOut3, err := value3.Export(); err == nil {
				context.SetOutput("arg3_out", argsOut3)
				log.Debugf("The value of j is %s", argsOut3) // 4
			}
		}

		if returnVar1, err := vm.Get("returnVar"); err == nil {
			returnVar, _ = returnVar1.ToBoolean()
			log.Debugf("The value of reutrn is %s", returnVar) // 4
		}
	} else {
		context.SetOutput("arg1_out", arg1)
		context.SetOutput("arg2_out", arg2)
		context.SetOutput("arg3_out", arg3)
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
	return returnVar, nil
}
