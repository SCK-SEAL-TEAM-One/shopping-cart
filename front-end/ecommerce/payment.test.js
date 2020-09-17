import checkPaymentMethod from './payment'

describe('CheckPaymentMethod', () => {
  it('Input 52477603 Should Be MASTER', () => {
    const expectedMethodPayment = 'MASTER'

    const cardNumber = '52477603'

    const actualPaymentMethod = checkPaymentMethod(cardNumber)

    expect(actualPaymentMethod).toBe(expectedMethodPayment)
  })

  it('Input 47197005 Should Be VISA', () => {
    const expectedMethodPayment = 'VISA'

    const cardNumber = '47197005'

    const actualPaymentMethod = checkPaymentMethod(cardNumber)

    expect(actualPaymentMethod).toBe(expectedMethodPayment)
  })
})
