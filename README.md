# filetree-cli
get filtree as json

i.e. 
go run main.go --path ./files
```
{"root":[{"type":"folder","name":"a","files":[{"type":"folder","name":"2","files":[{"type":"file","name":"file","files":[]}]},{"type":"file","name":"file","files":[]}]},{"type":"folder","name":"b","files":[{"type":"folder","name":"1","files":[{"type":"folder","name":"1","files":[{"type":"file","name":"file","files":[]}]}]}]},{"type":"file","name":"file","files":[]}]}
```
or pretty with the `--pretty` flag
```
{
    "root": [
        {
            "type": "folder",
            "name": "a",
            "files": [
                {
                    "type": "folder",
                    "name": "2",
                    "files": [
                        {
                            "type": "file",
                            "name": "file",
                            "files": []
                        }
                    ]
                },
                {
                    "type": "file",
                    "name": "file",
                    "files": []
                }
            ]
        },
        {
            "type": "folder",
            "name": "b",
            "files": [
                {
                    "type": "folder",
                    "name": "1",
                    "files": [
                        {
                            "type": "folder",
                            "name": "1",
                            "files": [
                                {
                                    "type": "file",
                                    "name": "file",
                                    "files": []
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "type": "file",
            "name": "file",
            "files": []
        }
    ]
}
```


