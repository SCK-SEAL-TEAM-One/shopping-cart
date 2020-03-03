import React from 'react'
import PropTypes from 'prop-types'


function CartItem({ item }) {
  const table = item.map(({
    productImage, productName, quantity, productPrice,
  }, index) => {
    const key = index + 1
    return (
      <tr key={key}>
        <td>{key}</td>
        <td id={`productImage-${key}`}><img src={productImage} alt="" /></td>
        <td id={`productName-${key}`}>{productName}</td>
        <td id={`productQuantity-${key}`}>{quantity}</td>
        <td id={`productPrice-${key}`}>{productPrice}</td>
      </tr>
    )
  })

  return (
    <div>
      <table>
        {table}
      </table>
    </div>
  )
}
CartItem.propTypes = {
  item: {
    productImage: PropTypes.string,
    productName: PropTypes.string,
    quantity: PropTypes.number,
    productPrice: PropTypes.string,
  },
}

CartItem.defaultProps = {
  item: {},
}

export default CartItem
