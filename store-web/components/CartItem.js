import React from 'react'
import {Container, Row, Button} from 'react-bootstrap'

function CartItem(props) {

  const table = props.item.map((obj,index) => {
    const key = index+1
    return ( 
      <tr>
        <td>{key}</td>
        <td id={`productImage-${key}`}><img src={obj.productImage}/></td>
        <td id={`productName-${key}`}>{obj.productName}</td>
        <td id={`productQuantity-${key}`}>{obj.quantity}</td>
        <td id={`productPrice-${key}`}>{obj.productPrice}</td>
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
export default CartItem