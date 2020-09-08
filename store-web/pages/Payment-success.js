import React from 'react'
import { Container, Button, Image } from 'react-bootstrap'
import Route from 'next/router'


export default class PaymentSucess extends React.Component {
  constructor() {
    super()
    this.message = 'วันเวลาที่ชำระเงิน 1/3/2563 13:30:00 หมายเลขคำสั่งซื้อ 8004359103 คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry ด้วยหมายเลข 1785261900'
  }

  goHome() {
    Route.push('/Product-list')
  }

  render() {
    return (
      <Container>
        <div>
          <h1 id="title">ชำระเงินสำเร็จ</h1>
          <Image src="https://www.pngitem.com/pimgs/m/69-692608_transparent-answer-icon-png-check-pass-icon-png.png" width="8%" />
          <div id="notify">
            {this.message}
          </div>
          <Button id="goHome" onClick={() => this.goHome()}>กลับไปหน้าหลัก</Button>
        </div>
      </Container>
    )
  }
}
