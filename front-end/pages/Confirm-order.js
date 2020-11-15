import React from 'react'
import { Container, Row, Button } from 'react-bootstrap'
import fetch from 'isomorphic-unfetch'
import Cookies from 'js-cookie'
import Route from 'next/router'
import CartItem from '../components/CartItem'


export default class ConfirmOrder extends React.Component {
  constructor(props) {
    super(props)

    this.submitOrder = this.submitOrder.bind(this)
  }

  createCookies() {
    const cart = [{
      product_id: 1,
      productName: '43 Piece dinner Set',
      productPrice: 12.95,
      productImage: 'https://i.pinimg.com/474x/17/43/2f/17432f12ec88c0d0ea3d0cffc69d25ce.jpg',
      quantity: 1,
    }]
    Cookies.set('cart', JSON.stringify(cart), { expires: 7, path: '' })

    const shipping = {
      shipping_method: 'kerry',
      shipping_address: '405/37 ถ.มหิดล',
      shipping_sub_district: 'ท่าศาลา',
      shipping_district: 'เมือง',
      shipping_province: 'เชียงใหม่',
      shipping_zip_code: '50000',
      recipient_name: 'ณัฐญา ชุติบุตร',
      recipient_phone_number: '0970809292',
    }
    Cookies.set('shipping', JSON.stringify(shipping), { expires: 7, path: '' })
  }


  submitOrder() {
    const cartItems = Cookies.getJSON('cart')
    const cart = cartItems.map(({ product_id, quantity }) => ({ product_id, quantity }))
    const shipping = Cookies.getJSON('shipping')

    fetch('/api/v1/order', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        cart,
        ...shipping,
      }),
    })
      .then((response) => response.json())
      .then((order) => {
        if (order.order_id) {
          Cookies.set('order', JSON.stringify(order), { expires: 7, path: '' })
          Route.push('/Payment')
        }
      })
  }

  render() {
    this.createCookies()
    const productList = Cookies.getJSON('cart')
    return (
      <Container>
        <Row>ยืนยันคำสั่งซื้อ</Row>
        <Row>
          <div>ที่อยู่ในการจัดส่ง:</div>
          <div>
            คุณ
            {' '}
            <span id="receiverName">ณัฐญา ชุติบุตร</span>
            <span id="recevierAddress">405/37 ถ.มหิดล ต.ท่าศาลา อ.เมือง จ.เชียงใหม่ 50000</span>
            <span id="recevierPhonenumber">0970809292</span>
          </div>
        </Row>
        <div>
          <div>รายการชำระเงิน</div>
          <table>
            <tr>
              <td>ค่าสินค้า</td>
              <td id="totalProductPrice">12.95 USD</td>
            </tr>
            <tr>
              <td>ค่าจัดส่ง</td>
              <td id="totalShippingCharge">2.00 USD</td>
            </tr>
            <tr>
              <td>รวมทั้งสิ้น</td>
              <td id="totalAmount">14.95 USD</td>
            </tr>
          </table>
        </div>
        <div>
          <div>รายการสินค้า</div>
          <div>
            {productList && <CartItem item={productList} />}
          </div>
        </div>
        <div>
          <Button id="editAddress">แก้ไขที่อยู่จัดส่ง</Button>
          <Button id="confirmPayment" onClick={() => this.submitOrder()}>ยืนยันคำสั่งซื้อและชำระเงิน</Button>
        </div>
      </Container>
    )
  }
}
