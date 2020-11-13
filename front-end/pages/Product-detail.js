import React from "react";
import { Container, Button, Form, Row, Col } from "react-bootstrap";
import Route from "next/router";

export default class ProductDetail extends React.Component {
  addToCart() {
    Route.push("/Confirm-order");
  }

  render() {
    return (
      <Container>
        <h1 id="productName">43 Piece dinner Set</h1>
        <Form>
          <Row>
            <Col xs="3">
              <img
                src="https://i.pinimg.com/474x/17/43/2f/17432f12ec88c0d0ea3d0cffc69d25ce.jpg"
                alt="product"
                width="100%"
              />
            </Col>
            <Col xs="9">
              <Row>
                <Col xs="3">Brand</Col>
                <Col xs="3" id="productBrand">
                  CoolKidz
                </Col>
              </Row>
              <Row>
                <Col xs="3">Gender</Col>
                <Col xs="3" id="productGender">
                  UNISEX
                </Col>
              </Row>
              <Row>
                <Col xs="3">Age</Col>
                <Col xs="3" id="productAge">
                  13+
                </Col>
              </Row>
              <Row>
                <Col xs="3">Price</Col>
                <Col xs="3" id="productPrice">
                  12.95 USD
                </Col>
              </Row>
              <Row>
                <Col xs="3">Quantity</Col>
                <Col xs="3">
                  <Form.Control
                    type="number"
                    id="productQuantity"
                    placeholder="0"
                  />
                </Col>
              </Row>
              <Row>
                <Col xs="6">
                  <Button
                    block
                    className="btn-custom"
                    id="addToCart"
                    onClick={() => this.addToCart()}
                  >
                    Add to Cart
                  </Button>
                </Col>
              </Row>
            </Col>
          </Row>
        </Form>
      </Container>
    );
  }
}
