// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/danielqsj/danielqsjblog/app"
	controllers "github.com/danielqsj/danielqsjblog/app/controllers"
	models "github.com/danielqsj/danielqsjblog/app/models"
	tests "github.com/danielqsj/danielqsjblog/tests"
	controllers0 "github.com/revel/revel/modules/static/app/controllers"
	_ "github.com/revel/revel/modules/testrunner/app"
	controllers1 "github.com/revel/revel/modules/testrunner/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
						"blogs",
						"recentCnt",
					},
				},
			},
			&revel.MethodType{
				Name: "WBlog",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					26: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "BlogInfor",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "rcnt", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
						"blog",
						"rcnt",
						"comments",
					},
				},
			},
			&revel.MethodType{
				Name: "Message",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					59: []string{ 
						"messages",
					},
				},
			},
			&revel.MethodType{
				Name: "History",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					73: []string{ 
						"historys",
					},
				},
			},
			&revel.MethodType{
				Name: "Emails",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					83: []string{ 
						"emails",
					},
				},
			},
			&revel.MethodType{
				Name: "About",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					86: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"error",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.WBlog)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Putup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "blog", Type: reflect.TypeOf((**models.Blog)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.WComment)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Docomment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "rcnt", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "comment", Type: reflect.TypeOf((**models.Comment)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.WMessage)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Putup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "message", Type: reflect.TypeOf((**models.Message)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/danielqsj/danielqsjblog/app/models.(*Blog).Validate": { 
			25: "blog.Title",
			29: "blog.Email",
			30: "blog.Email",
			34: "blog.Subject",
		},
		"github.com/danielqsj/danielqsjblog/app/models.(*Comment).Validate": { 
			17: "comment.Email",
			18: "comment.Email",
			23: "comment.Content",
		},
		"github.com/danielqsj/danielqsjblog/app/models.(*Message).Validate": { 
			18: "message.Email",
			19: "message.Email",
			22: "message.QQ",
			25: "message.Url",
			30: "message.Content",
		},
	}
	revel.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
