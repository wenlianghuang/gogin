package main

<<<<<<< HEAD
import (
	cookieex "github.com/Go/GoGin/CookieEx"
)
=======
import multipleservice "github.com/Go/GoGin/MultipleService"
>>>>>>> 20211109 modified file

func main() {
	//For "Example" folder
	//example.Example()

	//Parameter in Path
	//parameterinpath.ParameterInPath()

	//Quertystring parameters
	//querystringparameters.QueryStringParameters()

	//Multipart/Urlencoded Form
	//urlencodeform.UrlencodeForm()

	//Custom Recovery behavior
	//customrecovery.CustomRecovery()

	//Model binding and validation
	//bindingvalidation.BindingValidation()

	//Bind Query String or Post Data
	//bindquerypost.BindQueryPost()

	//XML, JSON, YAML and ProtoBuf rendering (error)
	//xmljsonyamlprotobuf.XMLJSONYAMLProtoBuf()

	//SecureJSON
	//securejson.SecureJSON()

	//Using BasicAuth() middleware
	//basicauth.BasicAuth()

	//Set and get a cookie
	cookieex.CookieEx()

	//Goroutines inside a middleware
	//goroutinemiddleware.GoroutineMiddleware()

	//Run multiple service using Gin
	multipleservice.MultipleService()
}
