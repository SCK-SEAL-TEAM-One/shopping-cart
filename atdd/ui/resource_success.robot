*** Variable ***
${url}   http://localhost/Product-list

*** Keywords ***
ดูรายละเอียดสินค้า
    Click Element    id=productName-1
ตรวจสอบข้อมูลสินค้า 
    [Arguments]    ${productName}    ${productPrice}    ${productQuantity}    ${productBrand}
    Sleep   0.3
    Element Text Should Be       id=productName-1     ${productName}
    Element Text Should Be       id=productPrice-1     ${productPrice}
    Input Text      id=productQuantity     ${productQuantity}
    Element Text Should Be       id=productBrand     ${productBrand}

เพิ่มสินค้าในตะกร้า
    Click Element        id=addToCart

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
    [Arguments]    ${recipientName}    ${shippingAddress}    ${recipientPhoneNumber}
    Sleep   0.3
    Element Text Should Be     id=receiverName        ${recipientName}
    Element Text Should Be     id=recevierAddress        ${shippingAddress}
    Element Text Should Be     id=recevierPhonenumber        ${recipientPhoneNumber}

ตรวจสอบสรุปรายการสั่งซื้อ
    [Arguments]    ${totalProductPrice}    ${totalShippingCharge}    ${totalAmount}
    Element Text Should Be       id=totalProductPrice     ${totalProductPrice}
    Element Text Should Be       id=totalShippingCharge     ${totalShippingCharge}
    Element Text Should Be       id=totalAmount          ${totalAmount}

ตรวจสอบตะกร้าสินค้า
    [Arguments]    ${productName}    ${productPrice}    ${productQuantity}    ${productBrand}
    Element Text Should Be       id=productName-1     ${productName}
    Element Text Should Be       id=productPrice-1     ${productPrice}
    Element Text Should Be       id=productQuantity-1     ${productQuantity}

ยืนยันคำสั่งซื้อ
    Click Element        id=confirmPayment

ชำระค่าสินค้า
    [Arguments]    ${cardNumber}    ${expiredMonth}    ${expiredYear}    ${cvv}    ${cardName}    ${totalPrice}
    Sleep    0.5
    Input Text    id=cardNumber    ${cardNumber} 
    Input Text    id=expiredMonth    ${expiredMonth}
    Input Text    id=expiredYear    ${expiredYear}
    Input Text    id=cvv    ${cvv}
    Input Text    id=cardName    ${cardName}
    Element Text Should Be    id=totalPrice    ${totalPrice}
     Click Button    id=Payment

ได้รับการแจ้งเตือน
    [Arguments]    ${notify}
    Element Text Should Be    id=title    ชำระเงินสำเร็จ
    Element Text Should Be    id=notify    ${notify}