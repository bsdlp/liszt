import React from 'react'
import Building from './Building'

class ListBuilding extends React.Component {
  render() {
    return (
      <div className='w-100 flex justify-center'>
        <div className='w-100' style={{ maxWidth: 400 }}>
          {this.props.query.Buildings.map(({node}) =>
            <Building key={node.id} post={node} />
          )}
        </div>
      </div>
    );
  }
}

export default ListBuilding
