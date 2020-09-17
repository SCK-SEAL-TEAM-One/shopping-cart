import React from "react";
import PropTypes from "prop-types";
import { Button, Card, Col, Row } from "react-bootstrap";

function ProductCard({ item }) {
  const card = item.map(
    ({ productImage, productName, getProductDetail, productPrice }, index) => {
      const key = index + 1;
      return (
        <Col xs={4} key={key}>
          <Card>
            <Card.Body>
              <div className="card-img">
                <img
                  src={productImage}
                  className="card-img"
                  id={`productImage-${key}`}
                />
              </div>
              <h5 id={`productName-${key}`}>{productName}</h5>
              <p id={`productPrice-${key}`}>{productPrice}</p>
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
    productImage: PropTypes.string,
    productName: PropTypes.string,
    productPrice: PropTypes.string,
    getProductDetail: PropTypes.func,
  },
};

ProductCard.defaultProps = {
  item: {},
};

export default ProductCard;
