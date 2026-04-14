package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
	  "message": "pong",
	})
}

func GetHome(c *gin.Context){
	c.HTML(http.StatusOK, "pages/home.html", gin.H{
		"Active": "home",
		"Title": "Stager Generator App",
		"Message": "Welcome to my app.",
	})
}

type ObfuscationType struct {
    ID    int
    Label string
    Description string
}

type FieldType string

const (
    FieldText   FieldType = "text"
    FieldFile   FieldType = "file"
)

type Field struct {
    ID          string
    Label       string
    Type        FieldType
    Placeholder string
    Optional	bool
}

type LoaderType struct {
    ID          int
    Label       string
    Description string
    Fields      []Field
}

func GetObfuscationPage(c *gin.Context){
	c.HTML(http.StatusOK, "pages/obfuscation_page.html", gin.H{
		"Active": "obfuscation_page",
		"Types": []ObfuscationType{
			{
				ID: 1,
				Label: "JScript Obfuscation 1",
				Description: `Description:<br>

				- Inserts comamnd into "template.js" JScript file.<br>
				- "template.js" content is embedded into loader.js as comments (Prefix: "////")<br>
				- "loader.js" will find comments start with "////", regroup it and execute as JScript.`,
			},
			{
				ID: 2,
				Label: "JScript Obfuscation 2",
				Description: `Sample Reference:<br>
				<a href="https://www.swisspost-cybersecurity.ch/news/purelogs-infostealer-analysis-dont-judge-a-png-by-its-header">https://www.swisspost-cybersecurity.ch/news/purelogs-infostealer-analysis-dont-judge-a-png-by-its-header</a><br><br>

				Description:<br>
				- Inserts command into "template.ps1" powershell file<br>
				- Encode powershell into base64 (UTF-16)<br>
				- Insert junk content into encoded powershell<br>
				- Place obfuscated base64 powershell into "loader.js" along with the junk to be removed at runtime.<br>

				`,
			},
		},
	})
}

func GetPsLoaderPage(c *gin.Context){
	c.HTML(http.StatusOK, "pages/ps_loader_page.html", gin.H{
		"Active": "ps_loader_page",
		"Title": "Powershell Loader",
		"Types": []LoaderType{
	        {
	            ID:          1,
	            Label:       "Powershell Loader 1",
	            Description: `Sample Reference:<br>
	            <a href="https://bazaar.abuse.ch/sample/a755759a2efb1f49d639af3f8166cb334e7fd537c3baf454e561f7ad6d07838f/">https://bazaar.abuse.ch/sample/a755759a2efb1f49d639af3f8166cb334e7fd537c3baf454e561f7ad6d07838f/</a><br>

	            Loader Description:<br>
	            - Downloads a base64 encoded and modified .NET executable to the local system<br>
	            - Decodes the .NET executable<br>
	            - Creates a second powershell loader script in C:\ProgramData<br>
	            - The second loader will invoke a static method in the .NET executable<br>
	            `,
	            Fields: []Field{
	            	{ID: "url", Label: "URL where the encoded .NET file will be hosted", Type: FieldText, Placeholder: "e.g. http://localhost/payload.txt", Optional: false},
	                {ID: "localFilepath", Label: "Windows Local File Path to encoded .NET file", Type: FieldText, Placeholder: "e.g. C:\\Windows\\Temp\\output.txt", Optional: false},
	                {ID: "uploadFile", Label: "Upload File", Type: FieldFile},
	                {ID: "getType", Label: "GetType", Type: FieldText, Placeholder: "e.g. ClassLibrary3.Class1", Optional: false},
	                {ID: "getMethod", Label: "GetMethod", Type: FieldText, Placeholder: "e.g. MethodA", Optional: false},
	                // TODO: add support for passing arguments into .NET static method
	                //{ID: "args", Label: "Arguments (blank if none)", Type: FieldText, Placeholder: "e.g. 'arg1', 'arg2'", Optional: true},
	            },
	        },
	        /*
	        {
	            ID:          2,
	            Label:       "Option 2",
	            Description: "Provide a <strong>source path</strong> and a <strong>destination path</strong>.",
	            Fields: []Field{
	                {ID: "sourcePath",      Label: "Source Path",      Type: FieldText, Placeholder: "e.g. C:\\source\\file.txt"},
	                {ID: "destinationPath", Label: "Destination Path", Type: FieldText, Placeholder: "e.g. C:\\dest\\file.txt"},
	            },
	        },
	        {
	            ID:          3,
	            Label:       "Option 3",
	            Description: "Provide a <strong>host</strong>, <strong>port</strong>, and select a <strong>payload file</strong>.",
	            Fields: []Field{
	                {ID: "host",        Label: "Host",         Type: FieldText, Placeholder: "e.g. 192.168.1.1"},
	                {ID: "port",        Label: "Port",         Type: FieldText, Placeholder: "e.g. 8080"},
	                {ID: "payloadFile", Label: "Payload File", Type: FieldFile},
	            },
	        },
	        */
	    },
	})
}