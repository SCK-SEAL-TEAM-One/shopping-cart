import React from 'react'
import { Container, Row, Button } from 'react-bootstrap'
import fetch from 'isomorphic-unfetch'

export default class Payment extends React.Component {
  render() {
    return (
      <Container>
        <Row>
          <form>
            <div>
              <input type="radio" value="credit" id="CreditCradPayment" checked />Credit Crad
              <input type="radio" value="credit" id="DebitCradType" disabled />Debit Crad
              <input type="radio" value="credit" id="LinePayType" disabled />Line Pay
            </div>
            <div>
              <label>เลขบัตร: </label>
              <input type="text" id="cardNumber"/>
            </div>
            <div>
              <label>วันหมดอายุ: </label>
              <input type="text" id="expiredMonth"/>/
              <input type="text" id="expiredYear"/>
            </div>
            <div>
              <label>CVV: </label>
              <input type="text" id="cvv"/>
            </div>
            <div>
              <label>ชื่อ: </label>
              <input type="text" id="cardName"/>
            </div>
            <div>
              <label>ยอกชำระ: </label>
              <span id="totalPrice">102.00 USD</span>
            </div>
            <input id="confirmPayment" type="submit" value="ยืนยันการชำระเงิน"></input>
          </form>
        </Row>
      </Container>
    )
  }
}