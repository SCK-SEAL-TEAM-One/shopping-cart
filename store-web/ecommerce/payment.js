export default function checkPaymentMethod(cardNumber) {
    if (cardNumber.startsWith('5')) {
        return "MASTER"
    }
}