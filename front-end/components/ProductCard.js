import React from "react";
import Route from 'next/router'
import PropTypes from "prop-types";
import { Button, Card, Col, Row } from "react-bootstrap";

function ProductCard({ item }) {
  const getProductDetail = (e) => {
    e.preventDefault()
    Route.push("/Product-detail");
  }
  const card = item.map(
    ({ product_image, product_name, product_price }, index) => {
      const key = index + 1;
      return (
        <Col xs={4} key={key}>
          <Card>
            <Card.Body>
              <div className="card-img">
                <img
                  src={product_image}
                  className="card-img"
                  id={`productImage-${key}`}
                />
              </div>
              <h5 id={`productName-${key}`}>{product_name}</h5>
              <p id={`productPrice-${key}`}>{product_price}</p>
              <Button
                id={`viewMore-${key}`}
                variant="primary"
                onClick={getProductDetail}
              >
                View more
              </Button>
            </Card.Body>
          </Card>
        </Col>
      );
    }
  );
  return <Row>{card}</Row>;
}

ProductCard.propTypes = {
  item: {
    product_image: PropTypes.string,
    product_name: PropTypes.string,
    product_price: PropTypes.string,
  },
};

ProductCard.defaultProps = {
  item: {},
};

export default ProductCard;
