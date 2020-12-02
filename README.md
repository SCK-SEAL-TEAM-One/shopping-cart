# Shopping cart
### ใช้สำหรับ Workshop ประกอบการสอน
# โจทย์เว็บไซต์ e-commerce ขายของเล่น

# Convention Web UI
## Common Convention
- โค้ดการทำงานหนึ่งบรรทัด ไม่มีการใส่ Semi-colon `;`
```
console.log(this.test)
```

## Function And Parameter Naming Conventions
- ส่วนที่เกี่ยวข้องกับ UI จะใช้รูปแบบการตั้งชื่อแบบ **camelCase ขึ้นต้นต้วยตัวพิมพ์ใหญ่**
- ส่วนที่เป็นการคำนวน จะใช้รูปแบบการตั้งชื่อแบบ **camelCase ขึ้นต้นต้วยตัวพิมพ์เล็ก**
```
func Homepage() //UI
func calculateTotalPrice() //logic
```

## HTML Element ID 
- จะใช้รูปแบบการตั้งชื่อแบบ **camelCase ขึ้นต้นต้วยตัวพิมพ์เล็ก**
```
receiverName
totalAmount
```

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
order
product
```

## File Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชั่นแบบ **camelCase ขึ้นต้นต้วยตัวพิมพ์ใหญ่**
```
Order.js
Shipping.test.js
```

## Test Name
- ใช้รูปแบบการตั้งชื่อให้สื่อความหมาย  เช่น
```
describe('CalculateTotalPrice', () => {
it('Input 5+8 should be 13',()=>{
        const expected = 13
        const productPriceList = ["5","8"]

        const actual = CalculateTotalPrice(productPriceList)

        expect(actual).toEqual(expected)
    })
}
```

## Variable Name
- ชื่อตัวแปรเป็นคำเดียวให้ตั้งชื่อเป็นพิมพ์เล็กทั้งหมด เช่น
```
day, month, year
```

- ชื่อตัวแปรมีความยาวตั้งแต่ 2 คำขึ้นไป ให้คำหลังขึ้นตันด้วยตัวอักษรตัวใหญ่เสมอ ในรูปแบบ **camelCase** เช่น
```
startDay, endMonth
```

- ชื่อตัวแปรเก็บค่าให้เติม "List" ต่อท้ายตัวแปรเสมอ เช่น
```
orderList
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวอักษรพิมพ์ใหญ่ทั้งหมด เช่น
```
HOUR, MINUTE
```

---

# Convention Web Service
## Function And Parameter Naming Conventions
- ใช้รูปแบบการตั้งชื่อฟังก์ชั่นแบบ **camelCase ขึ้นต้นต้วยตัวพิมพ์ใหญ่**
```
func CalculateTotalPrice()
```

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
order
product
```

## File Name
- camelCase ขึ้นต้นด้วยตัวใหญ่ เช่น
```
OrderService.go
ProductRepository.go
OrderService_test.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
order
product
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_CalculateAge_Input_Birth_Date_18042003_Should_be_16
```

## Variable Name
- ชื่อตัวแปรเป็นคำเดียวให้ตั้งชื่อเป็นพิมพ์เล็กทั้งหมด เช่น
```
day, month, year
```

- ชื่อตัวแปรมีความยาวตั้งแต่ 2 คำขึ้นไป ให้คำหลังขึ้นตันด้วยตัวอักษรตัวใหญ่เสมอ ในรูปแบบ **camelCase** เช่น
```
startDay, endMonth
```

- ชื่อตัวแปรเก็บค่าให้เติม "List" ต่อท้ายตัวแปรเสมอ เช่น
```
orderList
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวอักษรพิมพ์ใหญ่ทั้งหมด เช่น
```
HOUR, MINUTE
```

---


## ข้อตกลง Commit Message ร่วมกัน
```
[Created]: สร้างไฟล์ใหม่สำหรับ...

[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code

[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา

[Deleted]: ลบไฟล์ออก 'ชื่อไฟล์' เนื่องจาก...

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน
```

## How to run Acceptance test (API and UI)
### API test
1. ติดตั้ง newman
```
npm install -g newman
```
2. คำสั่งที่ใช้ run newman
```
newman run atdd/api/shopping_cart_success.json -e atdd/api/environment/local_environment.json -d atdd/api/data/shopping_cart_success.json
```
หรือใช้คำสั่งย่อ
```
make run_newman
```

### UI test
1. ติดตั้ง Robot framework
```
download Python2.7.xx (www.python.org) or Python3.x.x 
except 3.7.x (have bug)
```
2. ติดตั้ง Robot framework
```
pip install robotframework
```
3. ติดตั้ง seleniumlibrary
```
pip install robotframework-seleniumlibrary
```
4. ติดตั้ง chrome driver
```
brew --cask install chromedriver
```
5. run robot framework
```
robot atdd/ui/shopping_cart_success.robot
```
หรือใช้คำสั่งย่อ
```
make run_robot
```
