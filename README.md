# netflikss-golang
backend for netflikss

# 99design commands
go run github.com/99designs/gqlgen generate

# example query graphql/playground
mutation buildSeries {
  buildSeriesFromInfo(input:true) 
}

query createInfoJson{
  createInfoJson(input:true)
}

query findSerie {
  	series{
      label,
      seasons{
        number,
        label, 
        directoryName
        directoryName
        episodes{
          label,
        number,
         fileName 
      	}
      }
    }
}


