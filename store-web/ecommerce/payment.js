export default function checkPaymentMethod(cardNumber) {
  if (cardNumber.startsWith('5')) {
    return 'MASTER'
  }

  if (cardNumber.startsWith('4')) {
    return 'VISA'
  }

  return ''
}
