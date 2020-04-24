# netflikss-golang
backend for netflikss

# 99design commands
go run github.com/99designs/gqlgen generate

# example query graphql/playground

mutation buildSeries {
  buildSeriesFromInfo(input:true) 
}

mutation createInfoJson{
  createInfoJson(input:true)
}

mutation updateConfig{
  updateConfig(input:{
    fileServerPath:"newFileServerPath",
    stockPath:"/Movies/example_netflikss",
    port:"7171",
    allowedOrigin:"http://localhost:64594"
  })
}

 netflikss{
    configuration{
      fileServerPath,
      stockPath,
      serverConfiguration{port,allowedOrigin}
    },
      	series{
      label,
          stockPath,
          img
      seasons{
        number,
        label, 
        directoryName
        episodes{
          label,
        number,
         fileName 
      	}
      }
    }
  }