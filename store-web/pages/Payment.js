import React from 'react'
import { Container, Row } from 'react-bootstrap'
import fetch from 'isomorphic-unfetch'
import Cookies from 'js-cookie'
import checkPaymentMethod from '../ecommerce/payment'

export default class Payment extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: '',
      cardNumber: '',
      expiredMonth: '',
      expiredYear: '',
      cvv: '',
      cardName: '',
    };
    this.confrimPayment = this.confrimPayment.bind(this)
    this.handleChangeCardNumber = this.handleChangeCardNumber.bind(this)
    this.handleChangeExpiredMonth = this.handleChangeExpiredMonth.bind(this)
    this.handleChangeExpiredYear = this.handleChangeExpiredYear.bind(this)
    this.handleChangeCVV = this.handleChangeCVV.bind(this)
    this.handleChangeCardName = this.handleChangeCardName.bind(this)
  }

  confrimPayment() {
    const cradType = checkPaymentMethod(this.state.cardNumber)

    const mockRequest = {
      payment_type: this.state.payment_type,
      type: cradType,
      card_number: this.state.cardNumber,
      cvv: this.state.cvv,
      expired_month: parseInt(this.state.expiredMonth),
      expired_year: parseInt(this.state.expiredYear),
      card_name: this.state.cardNumber,
      total_price: 14.95,
    }
    fetch('/api/v1/confirmPayment', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        ...mockRequest,
      }),
    })
      .then((r) => r.json())
      .then((data) => {
        console.log(data)
      })
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  handleChangeCardNumber(event) {
    this.setState({ cardNumber: event.target.value });
  }
  handleChangeExpiredMonth(event) {
    this.setState({ expiredMonth: event.target.value });
  }
  handleChangeExpiredYear(event) {
    this.setState({ expiredYear: event.target.value });
  }
  handleChangeCVV(event) {
    this.setState({ cvv: event.target.value });
  }
  handleChangeCardName(event) {
    this.setState({ cardName: event.target.value });
  }
  render() {
    const { cardNumber, expiredMonth, expiredYear, cvv, cardName } = this.state
    const order = Cookies.getJSON('order')
    return (
      <Container>
        <Row>
          <form onSubmit={this.confrimPayment}>
            <div>
              <input type="radio" value="credit" id="CreditCradPayment" checked="checked" defaultChecked={true} onChange={this.handleChange} />Credit Crad
              <input type="radio" value="credit" id="DebitCradType" disabled />Debit Crad
              <input type="radio" value="credit" id="LinePayType" disabled />Line Pay
            </div>
            <div>
              <label>เลขบัตร: </label>
              <input type="text" id="cardNumber" onChange={this.handleChangeCardNumber} value={cardNumber} />
            </div>
            <div>
              <label>วันหมดอายุ: </label>
              <input type="text" id="expiredMonth" onChange={this.handleChangeExpiredMonth} value={expiredMonth} />/
              <input type="text" id="expiredYear" onChange={this.handleChangeExpiredYear} value={expiredYear} />
            </div>
            <div>
              <label>CVV: </label>
              <input type="text" id="cvv" onChange={this.handleChangeCVV} value={cvv} />
            </div>
            <div>
              <label>ชื่อ: </label>
              <input type="text" id="cardName" onChange={this.handleChangeCardName} value={cardName} />
            </div>
            <div>
              <label>ยอกชำระ: </label>
              <span id="totalPrice">{order? order.total_price:""} USD</span>
            </div>
            <input id="confirmPayment" type="submit" value="ยืนยันการชำระเงิน"></input>
          </form>
        </Row>
      </Container>
    )
  }
}