*** Variable ***
${url}   http://localhost/Product-list

*** Keywords ***
ดูรายละเอียดสินค้า
    Wait Until Element Contains    id=productName-2    43 Piece dinner Set
    Click Element    id=viewMore-1
    
ตรวจสอบข้อมูลสินค้า 
    [Arguments]    ${productName}    ${productPrice}    ${productQuantity}    ${productBrand}
    Wait Until Element Contains    id=productName    ${productName}
    Element Text Should Be       id=productName     ${productName}
    Element Text Should Be       id=productPrice     ${productPrice}
    Input Text      id=productQuantity     ${productQuantity}
    Element Text Should Be       id=productBrand     ${productBrand}

เพิ่มสินค้าในตะกร้า
    Click Element        id=addToCart

กรอกที่อยู่จัดส่ง
    [Arguments]    ${recipientName}    ${shippingAddress}    ${shippingSubDistrict}    ${shippingDistrict}    ${shippingProvince}    ${shippingZipCode}    ${recipientPhoneNumber}
    Set Selenium Speed     0.1
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
    Set Selenium Speed     0.1
    Wait Until Element Contains    id=receiverName    ${recipientName}  
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
    Textfield Value Should Be       id=productQuantity-1     ${productQuantity}

ยืนยันคำสั่งซื้อ
    Click Element        id=confirmPayment

ชำระค่าสินค้า
    [Arguments]    ${cardNumber}    ${expiredMonth}    ${expiredYear}    ${cvv}    ${cardName}    ${totalPrice}
    Wait Until Element Is Visible   id=payment    
    Input Text    id=cardNumber    ${cardNumber} 
    Input Text    id=expiredMonth    ${expiredMonth}
    Input Text    id=expiredYear    ${expiredYear}
    Input Text    id=cvv    ${cvv}
    Input Text    id=cardName    ${cardName}
    Element Text Should Be    id=totalPrice    ${totalPrice}
    Click Button    id=payment

ได้รับการแจ้งเตือน
    Set Selenium Speed     0.5
    [Arguments]    ${notify}
    Element Text Should Be    id=title    ชำระเงินสำเร็จ
    Element Text Should Be    id=notify    ${notify}

คำสั่งซื้อมีจำนวนทั้งหมดเท่ากับ
    [Arguments]    ${expected_rows}
    Connect To Database     pymysql   toy     sealteam    sckshuhari    localhost     3306    
    Row Count Is Equal To X     Select * From orders    ${expected_rows}
    Disconnect from Database