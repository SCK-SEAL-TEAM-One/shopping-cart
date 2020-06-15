import React from 'react'
import fetch from 'isomorphic-unfetch'
import { Container, Row, Button } from 'react-bootstrap'
import Route from 'next/router'


export default class ProductList extends React.Component {
  render() {
    return (
      <Container>
        <table>
            <tr>
              <dev>
              <img src="https://i.pinimg.com/736x/6c/38/ab/6c38abb1475ac373ee2d9d78da609cb3--fine-porcelain-dinner-sets.jpg" width="20%"></img>
              </dev>
              <div>
                <h3 id="productName-1">43 Piecee Dinner Set</h3>
                <h5 id="productPrice-1">12.95 USD</h5>
              </div>
            </tr>
            <tr>
            <dev>
              <img src="https://images-na.ssl-images-amazon.com/images/I/61uc4bgUPlL._AC_SL1500_.jpg" width="20%"></img>
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