import React from 'react'

class Building extends React.Component {
  render() {
    return (
      <div className='w-100 ba pa3'>
        {this.props.building.name}
      </div>
    )
  }
}

export default Building
