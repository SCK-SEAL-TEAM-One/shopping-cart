import React from 'react'
import { Container, Row, Button } from 'react-bootstrap'
import fetch from 'isomorphic-unfetch'
import Cookies from 'js-cookie'
import Route from 'next/router'
import checkPaymentMethod from '../ecommerce/payment'

export default class Payment extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      paymentType: '',
      cardNumber: '',
      expiredMonth: '',
      expiredYear: '',
      cvv: '',
      cardName: '',
    }
    this.handleChangePaymentType = this.handleChangePaymentType.bind(this)
    this.handleChangeCardNumber = this.handleChangeCardNumber.bind(this)
    this.handleChangeExpiredMonth = this.handleChangeExpiredMonth.bind(this)
    this.handleChangeExpiredYear = this.handleChangeExpiredYear.bind(this)
    this.handleChangeCVV = this.handleChangeCVV.bind(this)
    this.handleChangeCardName = this.handleChangeCardName.bind(this)
  }

  confrimPayment() {
    const {
      paymentType, cardNumber, expiredMonth, expiredYear, cvv, cardName,
    } = this.state
    const cradType = checkPaymentMethod(cardNumber)
    const order = Cookies.getJSON('order')
    const totalPrice = order ? order.total_price : 0
    const orderId = order ? order.order_id : 0

    const request = {
      payment_type: paymentType,
      type: cradType,
      card_number: cardNumber,
      cvv,
      expired_month: parseInt(expiredMonth, 2),
      expired_year: parseInt(expiredYear, 2),
      card_name: cardName,
      total_price: totalPrice,
      order_id: orderId,
    }
    fetch('/api/v1/confirmPayment', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        ...request,
      }),
    })
      .then((r) => r.json())
  }

  handleChangePaymentType(event) {
    this.setState({ paymentType: event.target.value })
  }

  handleChangeCardNumber(event) {
    this.setState({ cardNumber: event.target.value })
  }

  handleChangeExpiredMonth(event) {
    this.setState({ expiredMonth: event.target.value })
  }

  handleChangeExpiredYear(event) {
    this.setState({ expiredYear: event.target.value })
  }

  handleChangeCVV(event) {
    this.setState({ cvv: event.target.value })
  }

  handleChangeCardName(event) {
    this.setState({ cardName: event.target.value })
  }

  payment() {
    Route.push('/Payment-success')
  }

  render() {
    const {
      cardNumber, expiredMonth, expiredYear, cvv, cardName,
    } = this.state
    const order = Cookies.getJSON('order')
    return (
      <Container>
        <Row>
          <form onSubmit={this.confrimPayment}>
            <div>
              <input type="radio" value="credit" id="CreditCradPayment" checked="checked" defaultChecked onChange={this.handleChangePaymentType} />
              Credit Crad
              <input type="radio" value="debit" id="DebitCradType" disabled />
              Debit Crad
              <input type="radio" value="linePay" id="LinePayType" disabled />
              Line Pay
            </div>
            <div>
              <label htmlFor="cardNumber" id="labelCardNumber">
                เลขบัตร:
                <input type="text" id="cardNumber" onChange={this.handleChangeCardNumber} value={cardNumber} />
              </label>
            </div>
            <div>
              <label htmlFor="expiredMonth" id="labelExpiredMonth">
                วันหมดอายุ:
                <input type="text" id="expiredMonth" onChange={this.handleChangeExpiredMonth} value={expiredMonth} />
              </label>
              /
              <input type="text" id="expiredYear" onChange={this.handleChangeExpiredYear} value={expiredYear} />
            </div>
            <div>
              <label htmlFor="cvv" id="labelCvv">
                CVV:
                <input type="text" id="cvv" onChange={this.handleChangeCVV} value={cvv} />
              </label>
            </div>
            <div>
              <label htmlFor="cardName" id="labelCardName">
                ชื่อ:
                <input type="text" id="cardName" onChange={this.handleChangeCardName} value={cardName} />
              </label>
            </div>
            <div>
              <label htmlFor="totalPrice" id="labelTotalPrice">
                ยอกชำระ:
                <span id="totalPrice">
                  {order ? order.total_price : ''}
                  {' '}
                  USD
                </span>
              </label>
            </div>
            <Button id="payment" onClick={() => { this.payment() }}> ยืนยันการชำระเงิน</Button>
          </form>
        </Row>
      </Container>
    )
  }
}
