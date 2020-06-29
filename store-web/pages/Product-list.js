import React from 'react'
import fetch from 'isomorphic-unfetch'
import { Container, Row, Button } from 'react-bootstrap'
import Route from 'next/router'


export default class ProductList extends React.Component {
  getProductDetail() {
    Route.push('/Product-detail')
  }

  render() {
    return (
      <Container>
        <table>
          <tr>
            <dev>
              <img src="https://i.pinimg.com/474x/17/43/2f/17432f12ec88c0d0ea3d0cffc69d25ce.jpg" width="20%" />
            </dev>
            <div onClick={() => this.getProductDetail()}>
              <h3 id="productName-1">43 Piecee Dinner Set</h3>
              <h5 id="productPrice-1">12.95 USD</h5>
            </div>
          </tr>
          <tr>
            <dev>
              <img src="https://images-na.ssl-images-amazon.com/images/I/61uc4bgUPlL._AC_SL1500_.jpg" width="20%" />
            </dev>
            <div>
              <h3 id="productName-2">Balance Training Bicycle</h3>
              <h5 id="productPrice-2">119.95 USD</h5>
            </div>
          </tr>
        </table>
      </Container>
    )
  }
}
