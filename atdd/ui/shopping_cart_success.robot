*** Settings ***
Library    SeleniumLibrary

*** Variable ***
${url}   http://localhost:3000/products/2
${productName}       43 Piece dinner Set
${productPrice}       12.95 USD
${productQuantity}       1
${productBrand}       Coolkidz
${totalProductPrice}        12.95 USD
${totalShippingCharge}         2.00 USD
${totalAmount}         14.95 USD
${recipientName}        ณัฐญา ชุติบุตร
${shippingAddress}        405/37 ถ.มหิดล
${shippingSubDistrict}        ท่าศาลา
${shippingDistrict}        เมือง
${shippingProvince}        เชียงใหม่
${shippingZipCode}        50000
${recipientPhoneNumber}        0970809292
${cardNumber}    4719700591590995
${expiredMonth}    7
${expiredYear}    20
${cvv}    752
${cardName}    Karnwat Wongudom
${totalPrice}    102.00 USD
${notify}    วันเวลาที่ชำระเงิน  1/3/2563 13:30:00  หมายเลขคำสั่งซื้อ 8004359103  คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry ด้วยหมายเลข 1785261900 

*** Test Cases ***
ซื้อสินค้า 1 ชิ้น เลือกจัดส่งแบบ kerry เลือกตัดเงินผ่าน บัตรเครดิต visa ได้รับ notofication ผ่าน in-app
    Open Browser    about:blank    chrome
    Go To           ${url}
    ตรวจสอบข้อมูลสินค้า 
    เพิ่มสินค้าในตะกร้า
    กรอกที่อยู่จัดส่ง
    ตรวจสอบคำสั่งซื้อ
    ตรวจสอบที่อยู่ในการจัดส่ง
    ตรวจสอบสรุปรายการสั่งซื้อ
    ตรวจสอบตะกร้าสินค้า
    ยืนยันคำสั่งซื้อ

*** Keywords ***
ตรวจสอบข้อมูลสินค้า 
    Element Text Should Be       id=productName     ${productName}
    Element Text Should Be       id=productPrice     ${productPrice}
    Element Text Should Be       id=productQuantity     ${productQuantity}
    Element Text Should Be       id=productBrand     ${productBrand}

เพิ่มสินค้าในตะกร้า
    Click Element        id=addCart

กรอกที่อยู่จัดส่ง
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
    Element Text Should Be     id=recipientName        ${recipientName}
    Element Text Should Be     id=shippingAddress        ${shippingAddress}
    Element Text Should Be     id=shippingSubDistrict        ${shippingSubDistrict}
    Element Text Should Be     id=shippingDistrict        ${shippingDistrict}
    Element Text Should Be     id=shippingProvince        ${shippingProvince}
    Element Text Should Be     id=shippingZipCode        ${shippingZipCode}
    Element Text Should Be     id=recipientPhoneNumber        ${recipientPhoneNumber}

ตรวจสอบสรุปรายการสั่งซื้อ
    Element Text Should Be       id=totalProductPrice     ${totalProductPrice}
    Element Text Should Be       id=totalShippingCharge     ${totalShippingCharge}
    Element Text Should Be       id=totalAmount          ${totalAmount}

ตรวจสอบตะกร้าสินค้า
    Element Text Should Be       id=productName-1     ${productName}
    Element Text Should Be       id=productPrice-1     ${productPrice}
    Element Text Should Be       id=productQuantity-1     ${productQuantity}
    Element Text Should Be       id=productBrand-1     ${productBrand}

ยืนยันคำสั่งซื้อ
    Click Element        id=confirmPayment

ชำระค่าสินค้า
    Click Element    id=paymentType
    Input Text    id=cardNumber    ${cardNumber}
    Input Text    id=expiredMonth    ${expiredMonth}
    Input Text    id=expiredYear    ${expiredYear}
    Input Text    id=cvv    ${cvv}
    Input Text    id=cardName    ${cardName}
    Element Text Should Be    id=totalPrice    ${totalPrice}

ได้รับการแจ้งเตือน
    Element Text Should Be    id=notify    ${notify}
    Click Button    id=buttonClose