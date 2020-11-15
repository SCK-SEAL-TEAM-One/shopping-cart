import React from "react";
import { Container, Form, Row, Col, Button, CardDeck } from "react-bootstrap";
import Route from "next/router";
import ProductCard from "../components/ProductCard";

export default class ProductList extends React.Component {
  getProductDetail() {
    Route.push("/Product-detail");
  }

  render() {
    const item = [
      {
        productImage:
          "17432f12ec88c0d0ea3d0cffc69d25ce.jpg",
        productName: "43 Piece Dinner Set",
        productPrice: "12.95 USD",
        getProductDetail: this.getProductDetail.bind(this),
      },
      {
        productImage:
          "61uc4bgUPlL._AC_SL1500_.jpg",
        productName: "Balance Training Bicycle",
        productPrice: "119.95 USD",
        getProductDetail: this.getProductDetail.bind(this),
      },
    ];
    return (
      <Container>
        <Form>
          <Form.Row>
            <Form.Group as={Col} controlId="inputAge">
              <Form.Label>Select age</Form.Label>
              <Form.Control
                as="select"
                className="mr-sm-2"
                id="inputAge"
                custom
              >
                <option value="select">Choose</option>
                <option value="1">1-3</option>
                <option value="2">3-10</option>
              </Form.Control>
            </Form.Group>

            <Form.Group as={Col} controlId="inputGender">
              <Form.Label>Select gender</Form.Label>
              <Form.Control
                as="select"
                className="mr-sm-2"
                id="inputGender"
                custom
              >
                <option value="select">Choose</option>
                <option value="girl">girl</option>
                <option value="boy">boy</option>
                <option value="unisex">unisex</option>
              </Form.Control>
            </Form.Group>
            <Form.Group as={Col} controlId="inputGender">
              <Button
                id="search"
                variant="primary"
                type="submit"
                className="btn-custom"
              >
                Search
              </Button>
            </Form.Group>
          </Form.Row>
        </Form>

        <ProductCard item={item} />
      </Container>
    );
  }
}
