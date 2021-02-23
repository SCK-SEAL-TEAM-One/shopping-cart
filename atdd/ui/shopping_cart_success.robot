*** Settings ***
Library    SeleniumLibrary
Library     DatabaseLibrary
Suite Setup     Open Browser   ${url}     gc  #headlesschrome
Suite Teardown  Close Browser

Resource    resource_success.robot

*** Test Cases ***
ซื้อสินค้า 1 ชิ้น เลือกจัดส่งแบบ kerry เลือกตัดเงินผ่าน บัตรเครดิต visa ได้รับ notofication ผ่าน in-app
    ดูรายละเอียดสินค้า
    ตรวจสอบข้อมูลสินค้า     43 Piece dinner Set       12.95 USD       1       CoolKidz
    เพิ่มสินค้าในตะกร้า
    ตรวจสอบที่อยู่ในการจัดส่ง        ณัฐญา ชุติบุตร        405/37 ถ.มหิดล ต.ท่าศาลา อ.เมือง จ.เชียงใหม่ 50000        0970809292
    ตรวจสอบสรุปรายการสั่งซื้อ        12.95 USD         2.00 USD         14.95 USD
    ตรวจสอบตะกร้าสินค้า     43 Piece dinner Set       12.95       1       Coolkidz
    ยืนยันคำสั่งซื้อ
    ชำระค่าสินค้า    4719700591590995    7    20    752    Karnwat Wongudom    121.95 USD
    ได้รับการแจ้งเตือน    วันเวลาที่ชำระเงิน 1/3/2563 13:30:00 หมายเลขคำสั่งซื้อ 8004359103 คุณสามารถติดตามสินค้าผ่านช่องทาง Kerry ด้วยหมายเลข 1785261900
