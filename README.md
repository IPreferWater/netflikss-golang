# netflikss-golang

## purpose
Netflikss-golang is a **GraphQL** project to create your own streaming service using **local files**
This back-end project is working well with the [PWA flutter project](https://github.com/IPreferWater/netflikss-flutter "PWA flutter project")


## use it !

##  generate the file
This project use github.com/99designs/gqlgen for graphQL, the schema will generate the files

```
cd 'your-project'
go run github.com/99designs/gqlgen generate
```

## configuration
We use a golang file-server to access local files.
In the configuration.json 
```
 "fileServerPath": "",
```

is the path the golang file-server will use for localhost
if blank, root will be use

```
 "stockPath": "/Movies/example_netflikss",
```

is the path of where you store your local files.
it must be accessible after your fileServerPath

the purpose of this two variable is, in the future, to have multiple 'stockPath'

```
 "serverConfiguration" : {
    "port":"7171",
    "allowedOrigin":"http://localhost:64594"
 }
 ```
allowedOrigin is the url of your front-end application

## start
```
go run main.go
 ```


###  playground 
you can access to a playground via
```
http://localhost:7171/playground
```

The server will create a file info.json in each directory of your **stockPath**
He will try to guess most informations but you might have to complete it

```
mutation createInfoJson{
  createInfoJson(input:true)
}
```
####example of info.json
```
{
 "info": {
  "directory": "tetard",
  "label": "tetard",
  "stockPath": "/Movies/example_netflikss",
  "img": "tetard.jpg",
  "type": "movie"
 },
 "fileName": "tetard.mp4"
}
```

The server will read all info.json available in your stockPath and make them accessible in GraphQL
```
mutation buildSeries {
  buildSeriesFromInfo(input:true) 
}
```

The server will query your available videos
you will likely query the configuration at the same time

```
 netflikss{
    configuration{
      fileServerPath,
      stockPath,
      serverConfiguration{port,allowedOrigin}
    },
      	series{
      label
    },
	movies{
	      label
	}
  }
```

The server will update his configuration
```
mutation updateConfig{
  updateConfig(input:{
    fileServerPath:"newFileServerPath",
    stockPath:"/Movies/example_netflikss",
    port:"7171",
    allowedOrigin:"http://localhost:64594"
  })
}
```

