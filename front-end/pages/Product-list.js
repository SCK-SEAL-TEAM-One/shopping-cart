import React from "react";
import { Container, Form, Row, Col, Button, CardDeck } from "react-bootstrap";
import ProductCard from "../components/ProductCard";

export default class ProductList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      item: [],
    };
  }

  componentDidMount() {
    this.getProductList();
  }

  getProductList() {
    fetch("/api/v1/product", {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    })
      .then((response) => response.json())
      .then((products) => {
        this.setState({ item: products.products });
      });
  }

  render() {
    const item = this.state.item;
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
