import React from 'react'
import {Container, Row, Button} from 'react-bootstrap'
import CartItem from '../components/CartItem'

function ConfirmOrder() {
  const productList = [{
    "id":1,
    "productName":"43 Piecee Dinner Set",
    "productPrice":10.00,
    "productImage":".jpg",
    "quantity":1
  },
  {
    "id":1,
    "productName":"43 Piecee Dinner Set",
    "productPrice":10.00,
    "productImage":".jpg",
    "quantity":1
  }]
  return (
    <Container>
      <Row>ยืนยันคำสั่งซื้อ</Row>
      <Row>
        <div>ที่อยู่ในการจัดส่ง:</div>
        <div>
          คุณ <span id="receiverName">ณัฐญา ชุติบุตร</span> 
          <span id="recevierAddress">405/37 ถ.มหิดล ต.ท่าศาลา อ.เมือง จ.เชียงใหม่ 50000</span> 
          <spcn id="recevierPhonenumber">0970809292</spcn>
        </div>
      </Row>
      <div>
        <div>รายการชำระเงิน</div>
          <table>
            <tr>
              <td>ค่าสินค้า</td>
              <td id="totalProductPrice">100.00 USD</td>
            </tr>
            <tr>
              <td>ค่าจัดส่ง</td>
              <td id="totalShippingCharge">2.00 USD</td>
            </tr>
            <tr>
              <td>รวมทั้งสิ้น</td>
              <td id="totalAmount">102.00 USD</td>
            </tr>
          </table>
      </div>
      <div>
        <div>รายการสินค้า</div>
        <div>
          <CartItem item={productList}/>
        </div>
      </div>
      <div>
        <Button id="editAddress">แก้ไขที่อยู่จัดส่ง</Button>
        <a href="/payment"><Button id="confirmPayment">ยืนยันคำสั่งซื้อและชำระเงิน</Button></a>
      </div>
    </Container>
  )
}  
export default ConfirmOrder