import React from 'react'

class Building extends React.Component {
  render() {
    return (
      <div className='fl w-auto ba pa3'>
        <span className='db'>{this.props.building.name}</span>
        <span className='db'>{this.props.building.id}</span>
      </div>
    )
  }
}

export default Building
