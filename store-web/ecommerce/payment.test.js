import checkPaymentMethod from "./payment"

describe('CheckPaymentMethod', () => {
    it('Input 52477603 Should Be MASTER', () => {
        const expectedMethodPayment = "MASTER"

        const cardNumber = "52477603"

        const actualPaymentMethod = checkPaymentMethod(cardNumber)

        expect(actualPaymentMethod).toBe(expectedMethodPayment)

    });
});