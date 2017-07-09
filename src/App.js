import React, { Component } from 'react';
import {
  QueryRenderer,
  graphql
} from 'react-relay'
import environment from './Environment'
import ListBuilding from './ListBuilding'

const AppAllBuildingQuery = graphql`
  query AppAllBuildingQuery {
    Buildings {
      id
      name
    }
  }
`

class App extends Component {
  render() {
    return (
      <QueryRenderer
        environment={environment}
        query={AppAllBuildingQuery}
        render={({error, props}) => {
          if (error) {
            console.log(error)
            return <div>{error.message}</div>
          } else if (props) {
            return <ListBuilding buildings={props.Buildings} />
          }
          return null
        }}
      />
    )
  }
}

export default App;
