*** Variable ***
${url}   http://localhost/Confirm-order

*** Keywords ***
ตรวจสอบข้อมูลสินค้า 
    [Arguments]    ${productName}    ${productPrice}    ${productQuantity}    ${productBrand}
    Element Text Should Be       id=productName     ${productName}
    Element Text Should Be       id=productPrice     ${productPrice}
    Element Text Should Be       id=productQuantity     ${productQuantity}
    Element Text Should Be       id=productBrand     ${productBrand}

เพิ่มสินค้าในตะกร้า
    Click Element        id=addCart

กรอกที่อยู่จัดส่ง
    [Arguments]    ${recipientName}    ${shippingAddress}    ${shippingSubDistrict}    ${shippingDistrict}    ${shippingProvince}    ${shippingZipCode}    ${recipientPhoneNumber}
    Input Text     id=recipientName        ${recipientName}
    Input Text     id=shippingAddress        ${shippingAddress}
    Input Text     id=shippingSubDistrict        ${shippingSubDistrict}
    Input Text     id=shippingDistrict        ${shippingDistrict}
    Input Text     id=shippingProvince        ${shippingProvince}
    Input Text     id=shippingZipCode        ${shippingZipCode}
    Input Text     id=recipientPhoneNumber        ${recipientPhoneNumber}

ตรวจสอบคำสั่งซื้อ
    Click Element        id=verifyOrder

ตรวจสอบที่อยู่ในการจัดส่ง
    [Arguments]    ${recipientName}    ${shippingAddress}    ${shippingSubDistrict}    ${shippingDistrict}    ${shippingProvince}    ${shippingZipCode}    ${recipientPhoneNumber}
    Element Text Should Be     id=recipientName        ${recipientName}
    Element Text Should Be     id=shippingAddress        ${shippingAddress}
    Element Text Should Be     id=shippingSubDistrict        ${shippingSubDistrict}
    Element Text Should Be     id=shippingDistrict        ${shippingDistrict}
    Element Text Should Be     id=shippingProvince        ${shippingProvince}
    Element Text Should Be     id=shippingZipCode        ${shippingZipCode}
    Element Text Should Be     id=recipientPhoneNumber        ${recipientPhoneNumber}

ตรวจสอบสรุปรายการสั่งซื้อ
    [Arguments]    ${totalProductPrice}    ${totalShippingCharge}    ${totalAmount}
    Element Text Should Be       id=totalProductPrice     ${totalProductPrice}
    Element Text Should Be       id=totalShippingCharge     ${totalShippingCharge}
    Element Text Should Be       id=totalAmount          ${totalAmount}

ตรวจสอบตะกร้าสินค้า
    [Arguments]    ${productName}    ${productPrice}    ${productQuantity}    ${productBrand}
    Element Text Should Be       id=cartProductName     ${productName}
    Element Text Should Be       id=cartProductPrice     ${productPrice}
    Element Text Should Be       id=cartProductQuantity     ${productQuantity}
    Element Text Should Be       id=cartProductBrand     ${productBrand}

ยืนยันคำสั่งซื้อ
    Click Element        id=confirmPayment

ชำระค่าสินค้า
    [Arguments]    ${cardNumber}    ${expiredMonth}    ${expiredYear}    ${cvv}    ${cardName}    ${totalPrice}
    Click Element    id=paymentType
    Input Text    id=cardNumber    ${cardNumber}
    Input Text    id=expiredMonth    ${expiredMonth}
    Input Text    id=expiredYear    ${expiredYear}
    Input Text    id=cvv    ${cvv}
    Input Text    id=cardName    ${cardName}
    Element Text Should Be    id=totalPrice    ${totalPrice}

ได้รับการแจ้งเตือน
    [Arguments]    ${notify}
    Element Text Should Be    id=notify    ${notify}
    Click Button    id=buttonClose