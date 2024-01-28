package cart

import (
	"database/sql"
	"log"
)

type CartService struct {
	CartRepository CartRepository
}

func (cartService CartService) GetCart(uid int) ([]CartDetail, error) {
	carts, err := cartService.CartRepository.GetCartDetail(uid)
	if err != nil {
		log.Printf("CartRepository.GetCartDetail internal error %s", err.Error())
	}
	return carts, err
}

func (cartService CartService) AddCart(uid int, submitedCart SubmitedCart) (string, error) {
	cart, err := cartService.CartRepository.GetCartByProductID(uid, submitedCart.ProductID)
	act := "updated"
	if err == sql.ErrNoRows {
		act = "added"
		cartService.CartRepository.CreateCart(uid, submitedCart.ProductID, submitedCart.Quantity)
		return act, nil
	}
	err = cartService.CartRepository.UpdateCart(uid, submitedCart.ProductID, submitedCart.Quantity+cart.Quantity)
	if err != nil {
		log.Printf("CartRepository.UpdateCart internal error %s", err.Error())
		return "", err
	}
	return act, nil
}

func (cartService CartService) UpdateCart(uid int, submitedCart SubmitedCart) (string, error) {
	act := "updated"
	if submitedCart.Quantity == 0 {
		act = "deleted"
		err := cartService.CartRepository.DeleteCart(uid, submitedCart.ProductID)
		if err != nil {
			log.Printf("CartRepository.DeleteCart internal error %s", err.Error())
			return "", err
		}
	} else {
		err := cartService.CartRepository.UpdateCart(uid, submitedCart.ProductID, submitedCart.Quantity)
		if err != nil {
			log.Printf("CartRepository.UpdateCart internal error %s", err.Error())
			return "", err
		}
	}
	return act, nil

}
