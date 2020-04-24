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