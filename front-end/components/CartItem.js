import React from "react";
import PropTypes from "prop-types";
import { Row, Col, Form } from "react-bootstrap";

function CartItem({ item }) {
  const cart = item.map(
    ({ productImage, productName, quantity, productPrice }, index) => {
      const key = index + 1;
      return (
        <Row className="items-product" key={key}>
          <Col xs="3">
            <img
              className="product-image"
              id={`productImage-${key}`}
              src={productImage}
            />
          </Col>
          <Col xs="5">
            <Row>
              <Col xs="12" id={`productName-${key}`}>
                {productName}
              </Col>
              <Col xs="6">Gender Unisex</Col>
              <Col xs="6">Age 13+</Col>
              <Col xs="12">In Stock</Col>
              <Col xs="12">Delete</Col>
            </Row>
          </Col>
          <Col xs="2">
            <Form.Group controlId="formBasicEmail">
              <Form.Label>Quantity</Form.Label>
              <Form.Control
                type="number"
                id={`productQuantity-${key}`}
                value={quantity}
              />
            </Form.Group>
          </Col>
          <Col xs="2">
            <Form.Group controlId="formBasicEmail">
              <Form.Label>Price</Form.Label>
              <div id={`productPrice-${key}`}>{productPrice}</div>
            </Form.Group>
          </Col>
        </Row>
      );
    }
  );

  return <div>{cart}</div>;
}
CartItem.propTypes = {
  item: {
    productImage: PropTypes.string,
    productName: PropTypes.string,
    quantity: PropTypes.number,
    productPrice: PropTypes.string,
  },
};

CartItem.defaultProps = {
  item: {},
};

export default CartItem;
