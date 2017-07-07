import React from 'react'
import Building from './Building'

class ListBuilding extends React.Component {
  render() {
    return (
      <div className='w-100 flex'>
        <div className='w-100' style={{ maxWidth: 400 }}>
          {this.props.buildings.map((item) =>
            <Building key={item.id} building={item} />
          )}
        </div>
      </div>
    )
  }
}

export default ListBuilding

ListBuilding.defaultProps = {
  buildings: []
}
